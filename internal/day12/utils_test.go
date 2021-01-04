package day12

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
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			p:    common.Part1,
			want: 25,
		},
		{
			input: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			p:    common.Part2,
			want: 286,
		},
	}

	for _, test := range tests {

		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
