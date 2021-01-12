package day13

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
				"939",
				"7,13,x,x,59,x,31,19",
			},
			p:    common.Part1,
			want: 295,
		},
		{
			input: []string{
				"0",
				"7,13,x,x,59,x,31,19",
			},
			p:    common.Part2,
			want: 1068781,
		},
		{
			input: []string{
				"0",
				"17,x,13,19",
			},
			p:    common.Part2,
			want: 3417,
		},
		{
			input: []string{
				"0",
				"67,7,59,61",
			},
			p:    common.Part2,
			want: 754018,
		},
		{
			input: []string{
				"0",
				"67,x,7,59,61",
			},
			p:    common.Part2,
			want: 779210,
		},
		{
			input: []string{
				"0",
				"67,7,x,59,61",
			},
			p:    common.Part2,
			want: 1261476,
		},
		{
			input: []string{
				"0",
				"1789,37,47,1889",
			},
			p:    common.Part2,
			want: 1202161486,
		},
	}

	for _, test := range tests {

		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
