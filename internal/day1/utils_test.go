package day1

import (
	"reflect"
	"testing"
)

type testT struct {
	t    int
	n    int
	want [][]int
}

func TestTCombinations(t *testing.T) {

	tests := []testT{
		{
			t: 3,
			n: 6,
			want: [][]int{
				{0, 1, 2},
				{0, 1, 3},
				{0, 2, 3},
				{1, 2, 3},
				{0, 1, 4},
				{0, 2, 4},
				{1, 2, 4},
				{0, 3, 4},
				{1, 3, 4},
				{2, 3, 4},
				{0, 1, 5},
				{0, 2, 5},
				{1, 2, 5},
				{0, 3, 5},
				{1, 3, 5},
				{2, 3, 5},
				{0, 4, 5},
				{1, 4, 5},
				{2, 4, 5},
				{3, 4, 5},
			},
		},
	}

	for _, test := range tests {
		got := tCombinations(test.t, test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v expected %v", got, test.want)
		}
	}
}
