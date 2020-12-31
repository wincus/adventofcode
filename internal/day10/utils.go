package day10

import (
	"log"
	"sort"
	"strconv"

	"github.com/wincus/adventofcode/internal/common"
)

// Solve returns the solution for day10 problem
func Solve(s []string, p common.Part) int {

	r, err := parse(s)

	if err != nil {
		log.Printf("error parsing input")
		return 0
	}

	// prepend a 0 to account for the charging outlet
	r = append([]int{0}, r...)

	// append last value to be equal to n_max + 3
	// to account for the device
	r = append(r, r[len(r)-1]+3)

	if p == common.Part1 {

		var m = make(map[int]int)

		d := parseDeltas(r)

		for _, v := range d {
			m[v]++
		}

		return m[1] * m[3]

	}

	if p == common.Part2 {
		seg := parseSegments(r)
		return countPaths(seg)
	}

	return 0

}

func tribonacci(n int) int {

	if n == 0 || n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}

func parseSegments(r []int) []int {

	var s []int
	var count int = 1

	for i := 0; i < len(r)-1; i++ {
		if r[i+1]-r[i] == 1 {
			count++
			continue
		}
		s = append(s, count)
		count = 1
	}

	return s
}

func countPaths(s []int) int {

	var mul int = 1

	for i := 0; i < len(s); i++ {
		mul *= tribonacci(s[i] + 1)
	}

	return mul
}

func parse(s []string) ([]int, error) {

	var ints []int

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		i, err := strconv.Atoi(line)

		if err != nil {
			return ints, err
		}

		ints = append(ints, i)
	}

	sort.Ints(ints)
	return ints, nil
}

func parseDeltas(r []int) []int {

	var deltas []int

	for i := 0; i < len(r)-1; i++ {
		deltas = append(deltas, r[i+1]-r[i])
	}

	return deltas
}
