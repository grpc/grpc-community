package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
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

// Returns a list of winner tiers and a bool indicating whether or not there was a tie causing a failure of selection.
func scoreSumMatrixInternal(sumMatrix [][]int, winnerCount int, removedCandidates map[int]bool) ([][]int, bool) {
	// Candidates with no losses
	lossCounts := map[int]int{}
	for i, _ := range sumMatrix {
		// Don't evaluate loss count for removed candidates.
		if _, ok := removedCandidates[i]; ok {
			continue
		}

		losses := 0
		for j, _ := range sumMatrix {
			if i == j {
				continue
			}

			// Don't consider removed candidates for losses.
			if _, ok := removedCandidates[j]; ok {
				continue
			}

			if sumMatrix[i][j] < sumMatrix[j][i] {
				losses += 1
			}
		}
		lossCounts[i] = losses
	}

	minLosses := math.MaxInt
	for _, lossCount := range lossCounts {
		if lossCount < minLosses {
			minLosses = lossCount
		}
	}

	tierWinners := []int{}
	for i, losses := range lossCounts {
		if losses == minLosses {
			tierWinners = append(tierWinners, i)
		}
	}

	if len(tierWinners) == winnerCount {
		return [][]int{tierWinners}, false
	} else if len(tierWinners) < winnerCount {
		for _, winner := range tierWinners {
			removedCandidates[winner] = true
		}
		subResult, err := scoreSumMatrixInternal(sumMatrix, winnerCount-len(tierWinners), removedCandidates)
		conjointResult := append([][]int{tierWinners}, subResult...)
		return conjointResult, err
	} else {
		return [][]int{tierWinners}, true
	}
}

// Returns winners in tiers
func scoreSumMatrix(sumMatrix [][]int, winnerCount int) ([][]int, bool) {
	return scoreSumMatrixInternal(sumMatrix, winnerCount, map[int]bool{})
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

func generateResultsMarkdown(candidates []string, winnerTiers [][]int, tie bool, sumMatrix [][]int) string {
	var sb strings.Builder

	// TODO: Implement tied results.
	if tie {
		fmt.Fprintf(os.Stderr, "tie not yet implemented\n")
		os.Exit(1)
	}

	flattenedWinners := []int{}
	for _, tier := range winnerTiers {
		flattenedWinners = append(flattenedWinners, tier...)
	}

	// TODO: Response count.

	sb.WriteString("## Elected Steering Committee\n")

	for _, candidateIndex := range flattenedWinners {
		sb.WriteString(fmt.Sprintf("- %s\n", candidates[candidateIndex]))
	}
	sb.WriteString("\n")

	// TODO: Extend the output tiers to include _everyone_ including those who didn't win.
	sb.WriteString("## Instant Run-Off Tiers\n")
	for i, tier := range winnerTiers {
		if len(tier) > 1 {
			sb.WriteString(fmt.Sprintf("**Tied for Place %d**\n", i+1))
		} else {
			sb.WriteString(fmt.Sprintf("**Place %d**\n", i+1))
		}
		for _, candidateIndex := range tier {
			sb.WriteString(fmt.Sprintf("- %s\n", candidates[candidateIndex]))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

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

	// fmt.Printf("scoring matrix: %v\n", matrixSum)

	// TODO: Test this function.
	// TODO: output the sum matrix to a markdown file.
	winners, tie := scoreSumMatrix(matrixSum, conf.WinnerCount)
	if tie {
		fmt.Fprintf(os.Stderr, "tie: %v\n", winners)
		os.Exit(1)
	}

	md := generateResultsMarkdown(conf.Candidates, winners, tie, matrixSum)
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
