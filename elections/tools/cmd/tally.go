package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"maps"
	"math"
	"os"
	"strconv"
	"strings"

	"sigs.k8s.io/yaml"

	"github.com/spf13/cobra"
)

const (
	DefaultConfigFileName = "config.yaml"

	HeaderRowsDefault       = 1
	FirstRankingColumnIndex = 1
)

//go:embed .git_hash
var gitHash string

type tallyConfig struct {
	// The number of header rows before ranking rows begin.
	HeaderRows int `json:"headerRows",omitempty`

	// The number of columns before rank data starts on each row.
	PrefixColCount int `json:"prefixColCount",omitempty`

	// The number of columns after the rank data on each row.
	SuffixColCount int `json:"suffixColCount",omitempty`

	Candidates []string `json:"candidates"`

	WinnerCount int `json:"winnerCount"`
}

// Returns a list of raw rankings -- nearly identical to the textual input
// Blank cells are represented by 0.
func getRankRows(headerRows, prefixColCount, suffixColCount, candidateCount int, cr *csv.Reader) ([][]int, error) {

	var outRows [][]int
	rowCount := 0
	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		rowCount += 1
		if rowCount <= headerRows {
			continue
		}

		if len(row) != candidateCount+prefixColCount+suffixColCount {
			return nil, fmt.Errorf("row %d has %d columns, expected %d (%d ranking columns + %d prefix columns + %d suffix columns)", rowCount, len(row), candidateCount+prefixColCount+suffixColCount, candidateCount, prefixColCount, suffixColCount)
		}

		rankingRow := row[prefixColCount : len(row)-suffixColCount]

		rankings := map[int]bool{}

		var outRow []int
		for i, cell := range rankingRow {
			if cell == "" {
				// Blank cells indicate no preference and are represented here by a 0.
				outRow = append(outRow, 0)
				continue
			}

			cellInt, err := strconv.Atoi(cell)
			if err != nil {
				return nil, fmt.Errorf("invalid non-integer cell at row %d, col %d: %v", rowCount, i+prefixColCount, err)
			}

			if cellInt < 1 {
				return nil, fmt.Errorf("cell at row %d, col %d has value %d, the lowest allowed ranking is 1", rowCount, 1+prefixColCount, cellInt)
			}

			if cellInt > candidateCount {
				return nil, fmt.Errorf("cell at row %d, col %d has value %d, the highest allowed ranking is %d (the number of candidates)", rowCount, 1+prefixColCount, cellInt, candidateCount)
			}

			if _, ok := rankings[cellInt]; ok {
				return nil, fmt.Errorf("found multiple instances of ranking %d in row %d", cellInt, rowCount)
			}
			rankings[cellInt] = true

			outRow = append(outRow, cellInt)
		}

		outRows = append(outRows, outRow)
	}

	return outRows, nil
}

// TODO: Consider a type alias.
func newMatchupMatrix(candidateCount int) [][]int {
	var m [][]int
	for i := 0; i < candidateCount; i++ {
		var row []int
		for j := 0; j < candidateCount; j++ {
			row = append(row, 0)
		}
		m = append(m, row)
	}

	return m
}

func addMatrices(a, b [][]int) [][]int {
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
func rankRowToMatchupMatrix(rankRow []int, candidateCount int) [][]int {
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

// Slice of length `candidateCount` of slices of of length `candidateCount`. Values in the slice are tne number of placements of each rank that each candidate has.
func calculatePlacements(rankRows [][]int, candidateCount int) (placements [][]int) {
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
func leastPreferenceInternal(placements [][]int, candidatesToConsider map[int]bool, place int) []int {
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
func leastPreference(placements [][]int, removedCandidates map[int]bool) []int {
	candidatesToConsider := map[int]bool{}
	for candidateIndex := 0; candidateIndex < len(placements); candidateIndex++ {
		if _, ok := removedCandidates[candidateIndex]; !ok {
			candidatesToConsider[candidateIndex] = true
		}
	}
	return leastPreferenceInternal(placements, candidatesToConsider, 0)
}

// //  Returns:
// //  - Only when the second return value is false: if a has won true, otherwise false.
// //  - true if there was a tie
// func headToHead(sumMatrix [][]int, a, b int) (bool, bool) {
//
// }

// Returns
// - the candidate to eliminate, if tie is not true
// - whether or not there has been a tie
func findEliminatee(sumMatrix [][]int, candidates []int) (int, bool) {
	wins := map[int]int{}
	for _, c := range candidates {
		wins[c] = 0
		for _, o := range candidates {
			if c == 0 {
				continue
			}
			if sumMatrix[c][o] > sumMatrix[o][c] {
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
func scoreSumMatrixInternal(sumMatrix [][]int, placements [][]int, winnerCount int, removedCandidates map[int]bool, losers []int) ([]int, []int, []int) {
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
			for c, _ := range least {
				losers = append(losers, c)
				removedCandidates[c] = true
			}
			return scoreSumMatrixInternal(sumMatrix, placements, winnerCount, removedCandidates, losers)
		} else {
			// This is a tie.
			winners := []int{}
			for c := 0; c < len(sumMatrix); c++ {
				if _, ok := removedCandidates[c]; !ok {
					winners = append(winners, c)
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
func scoreSumMatrix(sumMatrix [][]int, placements [][]int, winnerCount int) (losers []int, winners []int, tied []int) {
	return scoreSumMatrixInternal(sumMatrix, placements, winnerCount, map[int]bool{}, []int{})
}

var rootCommand = &cobra.Command{
	Use:   "tally [csv-file]",
	Short: "Tallies results from a gRPC Steering Committee Election using the Condorcet IRV method.",
	Long: `
Arguments:
  csv-file: The path to the input CSV file.
`,
	Args: cobra.ExactArgs(1),
	RunE: run,
}

var outputMarkdownPath string
var configFilePath string

func generateResultsMarkdown(candidates []string, responseCount int, winners, losers, tie []int, sumMatrix [][]int) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("*tallied results for %d total ballots*\n\n", responseCount))

	if len(tie) > 0 {
		sb.WriteString("This election has resulted in a **tie**. A revote must be held with the eliminated candidates excluded.\n\n")

		sb.WriteString("## Candidates Eligible for Re-vote\n")
		for _, c := range winners {
			sb.WriteString(fmt.Sprintf("- %s\n", candidates[c]))
		}
		for _, c := range tie {
			sb.WriteString(fmt.Sprintf("- %s\n", candidates[c]))
		}
		sb.WriteString("\n")

		sb.WriteString("## Eliminated Candidates\n")
		for _, loser := range losers {
			sb.WriteString(fmt.Sprintf("- %s\n", candidates[loser]))
		}
		sb.WriteString("\n")
	} else {

		sb.WriteString("## Elected Steering Committee\n")

		for _, candidateIndex := range winners {
			sb.WriteString(fmt.Sprintf("- %s\n", candidates[candidateIndex]))
		}
		sb.WriteString("\n")

		sb.WriteString("## Instant Run-Off Elimination\n")
		sb.WriteString("Candidates were eliminated according to the Condorcet IRV method in the following order:\n")
		for _, loser := range losers {
			sb.WriteString(fmt.Sprintf("- %s\n", candidates[loser]))
		}

		sb.WriteString("\n")
	}

	sb.WriteString("## Sum Matrix\n")
	sb.WriteString("*([definition on Wikipedia](https://en.wikipedia.org/wiki/Condorcet_method#Basic_procedure))*\n")

	// Column Headers
	sb.WriteString("| |")
	for _, candidate := range candidates {
		sb.WriteString(fmt.Sprintf(" %s |", candidate))
	}
	sb.WriteString("\n")

	// Header divider row
	sb.WriteString("| -- |")
	for i := 0; i < len(candidates); i++ {
		sb.WriteString(" -- |")
	}
	sb.WriteString("\n")

	for candidateAIndex, candidateA := range candidates {
		sb.WriteString(fmt.Sprintf("| **%s** |", candidateA))
		for candidateBIndex, _ := range candidates {
			sb.WriteString(fmt.Sprintf(" %d |", sumMatrix[candidateAIndex][candidateBIndex]))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	sb.WriteString("---\n")

	sb.WriteString(fmt.Sprintf("*Results generated by [tally version %s](https://github.com/grpc/grpc-community/blob/%s/elections/tools)*\n", gitHash, gitHash))

	return sb.String()
}

func run(cmd *cobra.Command, args []string) error {
	csvFilename := os.Args[1]

	configStr, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("unable to read %s: %v\n", configFilePath, err)
	}

	var conf tallyConfig
	err = yaml.UnmarshalStrict(configStr, &conf)

	if err != nil {
		return fmt.Errorf("config error: %v\n", err)
	}

	csvFile, err := os.Open(csvFilename)
	if err != nil {
		return fmt.Errorf("unable to open %s: %v\n", csvFilename, err)
	}

	cr := csv.NewReader(csvFile)
	rankRows, err := getRankRows(conf.HeaderRows, conf.PrefixColCount, conf.SuffixColCount, len(conf.Candidates), cr)
	if err != nil {
		return fmt.Errorf("non-compliant csv in %s: %v\n", csvFilename, err)
	}

	matrixSum := newMatchupMatrix(len(conf.Candidates))
	for _, row := range rankRows {
		m := rankRowToMatchupMatrix(row, len(conf.Candidates))
		matrixSum = addMatrices(matrixSum, m)
	}

	placements := calculatePlacements(rankRows, len(conf.Candidates))

	// TODO: Test this function.
	losers, winners, tie := scoreSumMatrix(matrixSum, placements, conf.WinnerCount)

	fmt.Printf("Winners: %v\n", winners)
	fmt.Printf("Losers: %v\n", losers)
	fmt.Printf("Tie: %v\n", tie)

	md := generateResultsMarkdown(conf.Candidates, len(rankRows), winners, losers, tie, matrixSum)
	os.WriteFile(outputMarkdownPath, []byte(md), 0666)
	fmt.Printf("Wrote %s\n", outputMarkdownPath)

	return nil
}

func main() {
	rootCommand.PersistentFlags().StringVarP(&outputMarkdownPath, "output", "o", "results.md", "The path of the output markdown file")
	rootCommand.PersistentFlags().StringVarP(&configFilePath, "config", "c", DefaultConfigFileName, "The path of the config file")

	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
