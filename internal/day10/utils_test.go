package day10

import (
	"testing"

	"github.com/wincus/adventofcode/internal/common"
)

type Test struct {
	input []string
	p     common.Part
	want  int
}

func TestSolver(t *testing.T) {

	tests := []Test{
		{
			input: []string{
				"28",
				"33",
				"18",
				"42",
				"31",
				"14",
				"46",
				"20",
				"48",
				"47",
				"24",
				"23",
				"49",
				"45",
				"19",
				"38",
				"39",
				"11",
				"1",
				"32",
				"25",
				"35",
				"8",
				"17",
				"7",
				"9",
				"4",
				"2",
				"34",
				"10",
				"3",
			},
			p:    common.Part1,
			want: 220,
		},
		{
			input: []string{
				"16",
				"10",
				"15",
				"5",
				"1",
				"11",
				"7",
				"19",
				"6",
				"12",
				"4",
			},
			p:    common.Part2,
			want: 8,
		},
		{
			input: []string{
				"28",
				"33",
				"18",
				"42",
				"31",
				"14",
				"46",
				"20",
				"48",
				"47",
				"24",
				"23",
				"49",
				"45",
				"19",
				"38",
				"39",
				"11",
				"1",
				"32",
				"25",
				"35",
				"8",
				"17",
				"7",
				"9",
				"4",
				"2",
				"34",
				"10",
				"3",
			},
			p:    common.Part2,
			want: 19208,
		},
	}

	for _, test := range tests {

		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
