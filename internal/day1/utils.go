package day1

import (
	"fmt"
	"log"

	"github.com/wincus/adventofcode/internal/common"
)

// Solve returns the solutions of Day1
func Solve() {

	data, err := common.GetData(1)

	if err != nil {
		log.Panic(err)
	}

	intData, err := common.ToInt(data)

	if err != nil {
		log.Panic(err)
	}

	// find combinations with n=2 and n=3
	for _, t := range []int{2, 3} {
		c := tCombinations(t, len(intData))
		for _, i := range c {
			if m := sum(i, intData); m == 2020 {
				fmt.Printf("Solution for n=%v: %v\n", t, multiply(i, intData))
			}
		}
	}
}

func multiply(c, d []int) int {

	r := 1

	for _, i := range c {
		r = r * d[i]
	}

	return r

}

func sum(c, d []int) int {

	r := 0

	for _, i := range c {
		r = r + d[i]
	}

	return r

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
