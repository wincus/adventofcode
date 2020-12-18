package day5

import (
	"fmt"
	"log"
	"sort"

	"github.com/wincus/adventofcode/internal/common"
)

type row int

type column int

type seat struct {
	r  row
	c  column
	id int
}

// Solve returns the solution for day3 problem
func Solve(s []string, p common.Part) int {

	var max int
	var seats []int

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		seat, err := newSeat(line)

		if err != nil {
			log.Printf("error decoding seat with code: %v", s)
		}

		seats = append(seats, seat.id)

		if seat.id > max {
			max = seat.id
		}
	}

	if p == common.Part1 {
		return max
	}

	if p == common.Part2 {
		sort.Ints(seats)
		for i := seats[1]; i < len(seats)-1; i++ {
			if (seats[i+1] - seats[i]) > 1 {
				return seats[i] + 1
			}
		}
	}

	return 0
}

func newSeat(d string) (seat, error) {

	var s seat

	err := s.r.decode(d)

	if err != nil {
		return s, fmt.Errorf("could not decode row: %v", err)
	}

	err = s.c.decode(d)

	if err != nil {
		return s, fmt.Errorf("could not decode column: %v", err)
	}

	s.id = int(s.r)*8 + int(s.c)

	return s, nil
}

// decode row
func (r *row) decode(d string) error {

	var min int = 0
	var max int = 127
	var delta = max

	if len(d) != 10 {
		return fmt.Errorf("row code should have 10 letters")
	}

	for i := 0; i < 7; i++ {

		delta = delta >> 1

		if string(d[i]) == "F" {
			max = max - delta - 1
		}

		if string(d[i]) == "B" {
			min = min + delta + 1
		}
	}

	if min != max {
		return fmt.Errorf("something went wrong, not sure what: min=%v, max=%v", min, max)
	}

	*r = row(max)

	return nil
}

// decode column
func (c *column) decode(d string) error {

	var min int = 0
	var max int = 7
	var delta = max

	if len(d) != 10 {
		return fmt.Errorf("row code should have 10 letters")
	}

	for i := 7; i < 10; i++ {

		delta = delta >> 1

		if string(d[i]) == "L" {
			max = max - delta - 1
		}

		if string(d[i]) == "R" {
			min = min + delta + 1
		}
	}

	if min != max {
		return fmt.Errorf("something went wrong, not sure what: min=%v, max=%v", min, max)
	}

	*c = column(max)

	return nil
}
