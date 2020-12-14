package day2

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
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			p:    common.Part1,
			want: 2,
		},
		{
			input: []string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			p:    common.Part2,
			want: 1,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
