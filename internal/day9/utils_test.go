package day9

import (
	"testing"

	"github.com/wincus/adventofcode/internal/common"
)

type Test struct {
	input        []string
	p            common.Part
	preambleSize int
	want         int
}

func TestSolver(t *testing.T) {

	tests := []Test{
		{
			input: []string{
				"35",
				"20",
				"15",
				"25",
				"47",
				"40",
				"62",
				"55",
				"65",
				"95",
				"102",
				"117",
				"150",
				"182",
				"127",
				"219",
				"299",
				"277",
				"309",
				"576",
			},
			preambleSize: 5,
			p:            common.Part1,
			want:         127,
		},
		{
			input: []string{
				"35",
				"20",
				"15",
				"25",
				"47",
				"40",
				"62",
				"55",
				"65",
				"95",
				"102",
				"117",
				"150",
				"182",
				"127",
				"219",
				"299",
				"277",
				"309",
				"576",
			},
			preambleSize: 5,
			p:            common.Part2,
			want:         62,
		},
	}

	for _, test := range tests {

		got := Solve(test.input, test.p, test.preambleSize)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
