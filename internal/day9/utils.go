package day9

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/wincus/adventofcode/internal/common"
)

// Solve returns the solution for day9 problem
func Solve(s []string, p common.Part, size int) int {

	r := parse(s)

	if p == common.Part1 {
		for i := 0; i < len(r)-size; i++ {

			n := r[i+size]

			if !isValid(r[i:i+size], n) {
				return n
			}
		}
	}

	if p == common.Part2 {
		// solve part 1 again
		n := Solve(s, common.Part1, size)

		// get lower and upper index for
		// contiguous ints that sum n
		l, u, err := findContiguousSum(r, n)

		if err != nil {
			log.Print(err)
			return 0
		}

		t := make([]int, u-l)

		copy(t, r[l:u])

		sort.Ints(t)

		return t[0] + t[len(t)-1]
	}

	return 0

}

func findContiguousSum(r []int, n int) (int, int, error) {

	for i := 0; i < len(r); i++ {

		count := 0

		for e := 0; (i + e) < len(r); e++ {
			count += r[i+e]

			if count > n {
				break
			}

			if count == n {
				return i, i + e + 1, nil
			}
		}
	}

	return 0, 0, fmt.Errorf("could not found contiguous elements")
}

func isValid(r []int, n int) bool {

	tc := tCombinations(2, len(r))

	for _, p := range tc {
		if r[p[0]]+r[p[1]] == n {
			return true
		}
	}

	return false
}

// returns all t-combinations ct....c2c1 of the n numbers {0,1,...,n-1}
// given n >= t >= 0 in lexicographic order.
// Based on the Algorithm published by Donald Knuth on The Art of Computer
// Programing section 7.2.1.3
func tCombinations(t, n int) [][]int {

	var c []int
	var a []int
	var r [][]int

	// validate n >= t >= 0
	if t > n || t < 0 {
		return nil
	}

	c = make([]int, t+3)

	// L1 Initialize
	for j := 1; j <= t; j++ {
		c[j] = j - 1
	}

	// sentinel values
	c[t+1] = n
	c[t+2] = 0

	for {

		a = make([]int, t) // we need a new allocation
		copy(a, c[1:t+1])

		r = append(r, a)

		j := 1

		for c[j]+1 == c[j+1] {
			c[j] = j - 1
			j++
		}

		if j > t {
			return r
		}

		c[j] = c[j] + 1
	}
}

func parse(s []string) []int {

	var r []int

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		i, err := strconv.Atoi(line)

		if err != nil {
			log.Printf("could not parse int: %v", err)
			continue
		}

		r = append(r, i)

	}

	return r

}
