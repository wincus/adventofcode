package day7

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
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			p:    common.Part1,
			want: 4,
		},
		{
			input: []string{
				"shiny gold bags contain 2 dark red bags.",
				"dark red bags contain 2 dark orange bags.",
				"dark orange bags contain 2 dark yellow bags.",
				"dark yellow bags contain 2 dark green bags.",
				"dark green bags contain 2 dark blue bags.",
				"dark blue bags contain 2 dark violet bags.",
				"dark violet bags contain no other bags.",
			},
			p:    common.Part2,
			want: 126,
		},
	}

	for _, test := range tests {

		graph = make(map[string]map[string]int)

		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
