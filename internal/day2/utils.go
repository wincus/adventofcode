package day2

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/wincus/adventofcode/internal/common"
)

type password struct {
	min int
	max int
	r   rune
	p   string
}

// Solve returns the solution for day2 problem
func Solve(s []string, p common.Part) int {

	var counter int

	for _, line := range s {

		pass, err := newPassword(line)

		if err != nil {
			log.Print(err)
			continue
		}

		if p == common.Part1 && pass.isValid() {
			counter++
		}

		if p == common.Part2 && pass.arePositionsValid() {
			counter++
		}
	}

	return counter
}

func newPassword(s string) (password, error) {

	re := regexp.MustCompile(`^(\d+)-(\d+)\s(\w):\s(\w+)$`)

	d := re.FindStringSubmatch(s)

	if d == nil {
		return password{}, fmt.Errorf("could not parse password")
	}

	min, err := strconv.Atoi(d[1])

	if err != nil {
		return password{}, fmt.Errorf("could not parse min: %v", err)
	}

	max, err := strconv.Atoi(d[2])

	if err != nil {
		return password{}, fmt.Errorf("could not parse max: %v", err)
	}

	if len(d[3]) != 1 {
		return password{}, fmt.Errorf("more than one char was found")
	}

	return password{
		min: min,
		max: max,
		r:   rune(d[3][0]),
		p:   d[4],
	}, nil
}

func (p *password) isValid() bool {

	var counter int

	// an empty password would pass
	// validation checks!
	if len(p.p) == 0 {
		return false
	}

	for _, r := range p.p {
		if r == p.r {
			counter++
		}
	}

	return counter >= p.min && counter <= p.max

}

func (p *password) arePositionsValid() bool {

	var counter int

	// an empty password would pass
	// validation checks!
	if len(p.p) == 0 {
		return false
	}

	if rune(p.p[p.min-1]) == p.r {
		counter++
	}

	if rune(p.p[p.max-1]) == p.r {
		counter++
	}

	return counter == 1

}
