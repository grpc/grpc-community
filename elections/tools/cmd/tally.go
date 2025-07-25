package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"sigs.k8s.io/yaml"

	"github.com/grpc/grpc-community/elections/tools/pkg/score"

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

		if len(losers) == 0 {
			sb.WriteString("\n*no candidates were eliminated*\n")
		}

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

	losers, winners, tie, matrixSum := score.ScoreRows(rankRows, conf.WinnerCount)

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
