package score

import (
	"sort"
	"testing"

	"reflect"
)

func TestScoreRowsBasic(t *testing.T) {
	tcs := []struct {
		Name        string
		Rows        [][]int
		WinnerCount int
		WantWinners []int
		WantLosers  []int
		WantTie     []int
	}{
		{
			Name: "Basic",
			Rows: [][]int{
				{1, 2, 3, 4, 0, 6, 7, 8, 9},
				{2, 3, 1, 4, 0, 6, 9, 7, 8},
				{2, 3, 1, 4, 0, 6, 7, 9, 8},
			},
			WinnerCount: 7,
			WantWinners: []int{0, 1, 2, 3, 5, 6, 7},
			WantLosers:  []int{4, 8},
			WantTie:     []int{},
		},
		{
			Name: "Basic tie",
			Rows: [][]int{
				{1, 2, 3, 4, 0, 6, 7, 8, 9},
				{2, 3, 1, 4, 0, 6, 9, 7, 8},
				{2, 3, 1, 4, 0, 6, 7, 9, 8},
			},
			WinnerCount: 2,
			WantWinners: []int{0, 1, 2},
			WantLosers:  []int{4, 8, 7, 6, 5, 3},
			WantTie:     []int{0, 1},
		},
		{
			Name: "Blocks 1",
			Rows: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{9, 8, 7, 6, 5, 4, 3, 2, 1},
				{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			WinnerCount: 7,
			WantWinners: []int{0, 1, 2, 3, 6, 7, 8},
			WantLosers:  []int{5, 4},
			WantTie:     []int{},
		},
		{
			Name: "Blocks 2",
			Rows: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{2, 1, 3, 4, 5, 6, 7, 8, 9},
				{4, 2, 1, 3, 5, 6, 7, 8, 9},
				{7, 5, 3, 1, 2, 4, 6, 8, 9},
				{7, 6, 5, 3, 1, 2, 4, 8, 9},
				{8, 9, 7, 6, 5, 4, 3, 2, 1},
				{8, 9, 7, 6, 5, 4, 3, 2, 1},
			},
			WinnerCount: 7,
			WantWinners: []int{0, 1, 2, 3, 4, 5, 8},
			WantLosers:  []int{6, 7},
			WantTie:     []int{},
		},
		{
			Name: "Blocks 3",
			Rows: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{2, 1, 3, 4, 5, 6, 7, 8, 9},
				{4, 2, 1, 3, 5, 6, 7, 8, 9},
				{7, 5, 3, 1, 2, 4, 6, 8, 9},
				{7, 6, 5, 3, 1, 2, 4, 8, 9},
				{7, 6, 5, 4, 2, 1, 3, 8, 9},
				{8, 9, 7, 6, 5, 4, 3, 2, 1},
			},
			WinnerCount: 7,
			WantWinners: []int{0, 1, 2, 3, 4, 5, 6},
			WantLosers:  []int{7, 8},
			WantTie:     []int{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			losers, winners, tie, _ := ScoreRows(tc.Rows, tc.WinnerCount)

			// Winner order does not matter.
			sort.Ints(winners)
			sort.Ints(tc.WantWinners)
			if !reflect.DeepEqual(winners, tc.WantWinners) {
				t.Errorf("yielded unexpected winners: got: %v\nwant: %v\n", winners, tc.WantWinners)
			}

			if !reflect.DeepEqual(losers, tc.WantLosers) {
				t.Errorf("yielded unexpected losers: got: %v\nwant: %v\n", losers, tc.WantLosers)
			}

			// Tie order does not matter
			sort.Ints(tie)
			sort.Ints(tc.WantTie)
			if !reflect.DeepEqual(tie, tc.WantTie) {
				t.Errorf("yielded unexpected ties: got: %v\nwant: %v\n", tie, tc.WantTie)
			}

		})
	}
}
