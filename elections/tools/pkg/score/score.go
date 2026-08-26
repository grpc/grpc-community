package score

import (
	"fmt"
	"maps"
	"math"

	"cmp"
)

// A row-major square matrix representing a match-up between two or more candidates.
type MatchupMatrix = [][]int

func newMatchupMatrix(candidateCount int) MatchupMatrix {
	var m MatchupMatrix
	for i := 0; i < candidateCount; i++ {
		var row []int
		for j := 0; j < candidateCount; j++ {
			row = append(row, 0)
		}
		m = append(m, row)
	}

	return m
}

func addMatrices(a, b MatchupMatrix) MatchupMatrix {
	if len(a) != len(b) {
		// We use panic here because this represents a programmer error, not a user error.
		panic(fmt.Sprintf("incompatible matrices with row counts %d and %d", len(a), len(b)))
	}
	c := newMatchupMatrix(len(a))

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			// We use panic here because this represents a programmer error, not a user error.
			panic(fmt.Sprintf("incompatible matrices with column counts %d and %d on row %d", len(a[i]), len(b[i]), i))
		}

		if len(a[i]) != len(a) {
			// We use panic here because this represents a programmer error, not a user error.
			panic("encountered non-square matrix")
		}

		for j := 0; j < len(a); j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}

	return c
}

// Takes a flat ranking row (isomorphic to the input CSV) and outputs a matchup matrix in a row-major format.
func rankRowToMatchupMatrix(rankRow []int, candidateCount int) MatchupMatrix {
	m := newMatchupMatrix(candidateCount)

	for i, ranking := range rankRow {
		if ranking == 0 {
			continue
		}
		for j := 0; j < candidateCount; j++ {
			otherRanking := rankRow[j]
			if i == j {
				continue
			}

			if ranking < otherRanking || otherRanking == 0 {
				// A pairwise win occurs when candidate A's ranking is lower than candidate B or candidate A
				// has received a ranking and candidate B has not.
				m[i][j] = 1
			}
		}
	}

	return m
}

// PlacementMatrix is a row-major matrix where row i represents candidate i and column j represents the number of times candidate i came in rank j.
type PlacementMatrix = [][]int

// Slice of length `candidateCount` of slices of of length `candidateCount`. Values in the slice are tne number of placements of each rank that each candidate has.
func calculatePlacements(rankRows [][]int, candidateCount int) (placements PlacementMatrix) {
	for i := 0; i < candidateCount; i++ {
		candidateRankings := []int{}
		for j := 0; j < candidateCount; j++ {
			candidateRankings = append(candidateRankings, 0)
		}
		placements = append(placements, candidateRankings)
	}

	for _, row := range rankRows {
		for candidateIndex, ranking := range row {
			if ranking == 0 {
				// 0 represents no ranking
				continue
			}
			placements[candidateIndex][ranking-1] += 1
		}
	}

	return placements
}

// Returns the candidates tied for last, in no particular order.
func leastPreferenceInternal(placements PlacementMatrix, candidatesToConsider map[int]bool, place int) []int {
	if len(candidatesToConsider) == 0 {
		panic("programming error")
	}

	if place == len(placements) {
		// There are no more places to break the tie.
		panic("programming error")
	}

	lowestRank := math.MaxInt
	for candidateIndex := 0; candidateIndex < len(placements); candidateIndex++ {
		if _, ok := candidatesToConsider[candidateIndex]; !ok {
			continue
		}

		rankCountForThisCandidate := placements[candidateIndex][place]
		if rankCountForThisCandidate < lowestRank {
			lowestRank = rankCountForThisCandidate
		}
	}

	lowestCandidates := map[int]bool{}
	for candidateIndex := 0; candidateIndex < len(placements); candidateIndex++ {
		if _, ok := candidatesToConsider[candidateIndex]; !ok {
			continue
		}

		if placements[candidateIndex][place] == lowestRank {
			lowestCandidates[candidateIndex] = true
		}
	}

	lowestCandidatesSlice := []int{}
	for c, _ := range lowestCandidates {
		lowestCandidatesSlice = append(lowestCandidatesSlice, c)
	}

	if len(lowestCandidates) == 1 {
		return lowestCandidatesSlice
	} else if len(lowestCandidates) > 1 {
		if place == len(placements)-1 {
			// There's no more place data to use. Return a tie.
			return lowestCandidatesSlice
		} else {
			// Recur and use the next lowest place.
			return leastPreferenceInternal(placements, lowestCandidates, place+1)
		}
	} else {
		panic("programming error")
	}
}

// Returns the least preference candidate and whether or not there's been a tie.
func leastPreference(placements PlacementMatrix, removedCandidates map[int]bool) []int {
	candidatesToConsider := map[int]bool{}
	for candidateIndex := 0; candidateIndex < len(placements); candidateIndex++ {
		if _, ok := removedCandidates[candidateIndex]; !ok {
			candidatesToConsider[candidateIndex] = true
		}
	}
	return leastPreferenceInternal(placements, candidatesToConsider, 0)
}

// Returns:
//
//	-1 if a loses to b
//	1 if a wins against b
//	0 if there is a tie
func beats(sumMatrix MatchupMatrix, a, b int) int {
	return cmp.Compare(sumMatrix[a][b], sumMatrix[b][a])
}

// Returns
// - the candidate to eliminate, if tie is not true
// - whether or not there has been a tie
func findEliminatee(sumMatrix MatchupMatrix, candidates []int) (int, bool) {
	wins := map[int]int{}
	for _, c := range candidates {
		wins[c] = 0
		for _, o := range candidates {
			if c == o {
				continue
			}
			if beats(sumMatrix, c, o) == 1 {
				wins[c] += 1
			}
		}
	}

	potentialElims := []int{}
	for _, c := range candidates {
		if wins[c] == 0 {
			potentialElims = append(potentialElims, c)
		}
	}

	if len(potentialElims) == 1 {
		return potentialElims[0], false
	} else {
		return 0, true
	}
}

// Returns:
// - losers in order from first elimination to last elimination
// - winners in no particular order
// - candidates who tied and who have not conclusively won or lost, in no particular order.
func scoreSumMatrixInternal(sumMatrix MatchupMatrix, placements PlacementMatrix, winnerCount int, removedCandidates map[int]bool, losers []int) ([]int, []int, []int) {
	remainderCount := len(sumMatrix) - len(removedCandidates)
	if remainderCount == winnerCount {
		winners := []int{}
		for i := 0; i < len(sumMatrix); i++ {
			if _, ok := removedCandidates[i]; !ok {
				winners = append(winners, i)
			}
		}
		return losers, winners, []int{}
	} else if remainderCount < winnerCount {
		panic("Eliminated too many candidates")
	}

	least := leastPreference(placements, removedCandidates)
	if len(least) < 2 {
		additionalRemoved := map[int]bool{}
		maps.Copy(additionalRemoved, removedCandidates)
		for _, loser := range least {
			additionalRemoved[loser] = true
		}
		least = append(least, leastPreference(placements, additionalRemoved)...)
	}

	// `least` now has at least two candidates in consideration
	// eliminate the one that loses in a head-to-head match-up
	// unless there's a cycle, in which case, there's a tie

	e, tie := findEliminatee(sumMatrix, least)

	if tie {
		if len(sumMatrix)-len(removedCandidates)-len(least) >= winnerCount {
			// If the choice of which one of these to eliminate doesn't affect the overall result, we can eliminate them all.
			for _, c := range least {
				losers = append(losers, c)
				removedCandidates[c] = true
			}
			return scoreSumMatrixInternal(sumMatrix, placements, winnerCount, removedCandidates, losers)
		} else {
			// TODO: This section is probably too deeply nested now.
			// This is a tie.
			winners := []int{}
			for c := 0; c < len(sumMatrix); c++ {
				if _, ok := removedCandidates[c]; !ok {
					isTie := false
					for _, l := range least {
						if c == l {
							isTie = true
						}
					}
					if !isTie {
						winners = append(winners, c)
					}
				}
			}
			tied := []int{}
			for c, _ := range least {
				tied = append(tied, c)
			}
			return losers, winners, tied
		}
	} else {
		// We have a single definitive candidate to eliminate.
		removedCandidates[e] = true
		losers = append(losers, e)
		return scoreSumMatrixInternal(sumMatrix, placements, winnerCount, removedCandidates, losers)
	}
}

// Returns:
// - losers in order from first elimination to last elimination
// - winners in no particular order
// - candidates who tied and who have not conclusively won or lost, in no particular order.
func scoreSumMatrix(sumMatrix MatchupMatrix, placements PlacementMatrix, winnerCount int) (losers []int, winners []int, tied []int) {
	return scoreSumMatrixInternal(sumMatrix, placements, winnerCount, map[int]bool{}, []int{})
}

// TODO: Test this function.

// Returns:
// - losers in order from first elimination to last elimination
// - winners in no particular order
// - candidates who tied and who have not conclusively won or lost, in no particular order.
func ScoreRows(rankRows [][]int, winnerCount int) ([]int, []int, []int, MatchupMatrix) {
	candidateCount := len(rankRows[0])
	matrixSum := newMatchupMatrix(candidateCount)
	for _, row := range rankRows {
		m := rankRowToMatchupMatrix(row, candidateCount)
		matrixSum = addMatrices(matrixSum, m)
	}

	placements := calculatePlacements(rankRows, candidateCount)

	losers, winners, tied := scoreSumMatrix(matrixSum, placements, winnerCount)
	return losers, winners, tied, matrixSum
}
