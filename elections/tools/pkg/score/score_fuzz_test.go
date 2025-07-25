package score

import (
	"fmt"
	"testing"

	"bytes"
	"encoding/binary"
)

func factorial(x int) int {
	f := 1
	for i := 2; i <= x; i++ {
		f *= i
	}
	return f
}

// Returnst he smallest number of bits required to represent the supplied int.
func log2(x int) int {
	i := 0
	n := 1
	for {
		if n >= x {
			return i
		}
		i++
		n *= 2
	}
}

func sliceRemove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func lehmerCodeToRowInternal(lehmer int, i int, permutation []int, elements []int) []int {
	f := factorial(i)
	d := lehmer / f
	k := lehmer % f
	permutation = append(permutation, elements[d])
	elements = sliceRemove(elements, d)

	if i == 0 {
		return permutation
	}

	return lehmerCodeToRowInternal(k, i-1, permutation, elements)
}

func lehmerCodeToRow(lehmer int, elementCount int) []int {
	elements := []int{}
	for i := 1; i <= elementCount; i++ {
		elements = append(elements, i)
	}
	return lehmerCodeToRowInternal(lehmer, elementCount-1, []int{}, elements)
}

func rowEncodingToRows(rowEncoding []byte, numCandidates, numVoters int) [][]int {
	// We'll draw `numVoters` permutations from the byte slice. If there are no remaining bytes, we will assume the next row is a 0.
	// Each row will be a Lehmer Code represented by ceil(ceil(log_2(numCandidates!)) / 8) bytes.

	// So we will draw no more than numVoters * ceil(ceil(log_2(numCandidates!)) / 8) bytes from the slice.
	encodingIndex := 0
	drawBytes := func(n int) []byte {
		b := []byte{}
		for i := 0; i < n; i++ {
			if encodingIndex < len(rowEncoding) {
				b = append(b, rowEncoding[i])
				encodingIndex++
			} else {
				b = append(b, 0)
			}
		}
		return b
	}

	permutationCount := factorial(numCandidates)

	// TODO: By using more bits than we actually need (the ceiling operation),
	// we're wasting bits by not generating unique configurations from them. For every datum, we should be wasting
	// less than a single bit. This will require sub-byte data tracking.
	byteCount := (log2(permutationCount) + 7) / 8

	readLehmerCode := func() int {
		b := drawBytes(byteCount)
		padBytes := 8 - (len(b) % 8)
		for i := 0; i < padBytes; i++ {
			b = append(b, 0)
		}
		// fmt.Printf("b: %#v\n", b)
		var num uint64
		err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &num)
		if err != nil {
			panic(fmt.Sprintf("error marshalling bytes to int: %v", err))
		}

		return int(num) % permutationCount
	}

	// TODO: Support abstained votes with a bit mask.
	rows := [][]int{}
	for i := 0; i < numVoters; i++ {
		rows = append(rows, lehmerCodeToRow(readLehmerCode(), numCandidates))
	}

	return rows
}

// Returns:
// - winner - the Condorcet winner if `ok` is true
// - ok - whether or not there is a Condorcet winner
func getCondorcetWinner(sumMatrix MatchupMatrix) (int, bool) {
	candidateCount := len(sumMatrix)
	losslessCandidates := []int{}
	for a := 0; a < candidateCount; a++ {
		lossless := true
		for b := 0; b < candidateCount; b++ {
			if a == b {
				continue
			}

			if beats(sumMatrix, a, b) != 1 {
				lossless = false
				break
			}
		}
		if lossless == true {
			losslessCandidates = append(losslessCandidates, a)
		}
	}

	if len(losslessCandidates) == 1 {
		return losslessCandidates[0], true
	}

	return 0, false
}

func FuzzScoreRows(f *testing.F) {
	f.Fuzz(func(t *testing.T, rowEncoding []byte) {
		// TODO: Vary the integer parameters here a bit.
		initialWinnerCount := 7
		voterCount := 7
		candidateCount := 9
		rows := rowEncodingToRows(rowEncoding, candidateCount, voterCount)

		var oldLosers []int = nil

		for winnerCount := initialWinnerCount; winnerCount >= 1; winnerCount-- {
			t.Run(fmt.Sprintf("WinnerCount=%d", winnerCount), func(t *testing.T) {
				losers, winners, tie, sumMatrix := ScoreRows(rows, winnerCount)

				ctxMsg := fmt.Sprintf("winners: %v\nlosers: %v\ntie: %v\nrows: %v\nseed: %v\nsumMatrix: %v", winners, losers, tie, rows, rowEncoding, sumMatrix)

				t.Run("WinnerCount", func(t *testing.T) {
					if len(tie) == 0 && len(winners) != winnerCount {
						t.Errorf("selected wrong number of winners, want: %d, got: %d", winnerCount, len(winners))
					}
				})

				t.Run("ReturnTotal", func(t *testing.T) {
					total := len(winners) + len(losers) + len(tie)
					if total != candidateCount {
						t.Errorf("expected winners, losers, and tie to total %d but got %d\n%s\n", candidateCount, total, ctxMsg)
					}
				})

				t.Run("CondorcetWinner", func(t *testing.T) {
					if cw, ok := getCondorcetWinner(sumMatrix); ok {
						cwInWinners := false
						for _, w := range winners {
							if w == cw {
								cwInWinners = true
							}
						}

						if !cwInWinners {
							t.Errorf("Condorcet winner %d was not in returned winners\n%s", cw, ctxMsg)
						}
					}
				})

				if oldLosers != nil {
					t.Run("LosersMonotonic", func(t *testing.T) {
						if len(losers) < len(oldLosers) {
							t.Errorf("decreased winner count but losers were not a superset of previous losers")
						}

						for i := 0; i < len(oldLosers); i++ {
							if losers[i] != oldLosers[i] {
								t.Errorf("losers at winnerCount=%d (%v) was not a superset of losers at winnerCount=%d (%v)", winnerCount, losers, winnerCount+1, oldLosers)
							}
						}
					})
				}
				oldLosers = losers
			})
		}

		// To test:
		// x The winners, losers, and tie always sum to the candidates
		// - When decreasing the winner count, the new losers are always a superset of the old ones.
		// - The intersection of the Smith set and the losers is always null
	})
}
