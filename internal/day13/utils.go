package day13

import (
	"log"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode/internal/common"
)

// Solve returns the solution for day13 problem
func Solve(s []string, p common.Part) int {

	ts, busIDs, err := parse(s)

	if err != nil {
		log.Printf("could not parse input: %v", err)
		return 0
	}

	if p == common.Part1 {
		for tsi := ts; true; tsi++ {
			for _, id := range busIDs {

				if id == 0 {
					continue
				}

				if tsi%id == 0 {
					return (tsi - ts) * id
				}
			}
		}
	}

	if p == common.Part2 {

		for t := busIDs[0]; true; t += busIDs[0] {

			//log.Printf("checking t=%v", t)

			if check(t, busIDs) {
				return t
			}
		}
	}

	return 0
}

func check(t int, ids []int) bool {

	for i, id := range ids {

		if id == 0 {
			continue
		}

		if (t+i)%id != 0 {
			return false
		}
	}

	return true
}

func parse(s []string) (int, []int, error) {

	var b []int

	ts, err := strconv.Atoi(s[0])

	if err != nil {
		return 0, []int{}, err
	}

	ids := strings.Split(s[1], ",")

	for _, id := range ids {

		if id == "x" {
			b = append(b, 0)
			continue
		}

		x, err := strconv.Atoi(id)

		if err != nil {
			log.Printf("could not parse input: %v", x)
			continue
		}

		b = append(b, x)

	}

	return ts, b, nil

}
