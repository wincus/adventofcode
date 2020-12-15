package day3

import (
	"reflect"
	"testing"

	"github.com/wincus/adventofcode/internal/common"
)

type Test struct {
	input []string
	p     common.Part
	want  int
}

type TestSlope struct {
	input string
	want  slope
}

func TestSolver(t *testing.T) {

	tests := []Test{
		{
			input: []string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			p:    common.Part1,
			want: 7,
		},
		{
			input: []string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			p:    common.Part2,
			want: 336,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}

}

func TestSlopeParser(t *testing.T) {

	tests := []TestSlope{
		{
			input: "r1,d1",
			want:  slope{r: 1, d: 1},
		},
		{
			input: "r1,d1,u1,l1",
			want:  slope{r: 1, d: 1, u: 1, l: 1},
		},
		{
			input: "r10",
			want:  slope{r: 10},
		},
		{
			input: "r10,r10",
			want:  slope{r: 20},
		},
	}

	for _, test := range tests {
		got, err := newSlope(test.input)

		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(got, test.want) {
			t.Error(err)
		}
	}
}
