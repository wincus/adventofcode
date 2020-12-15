package day3

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode/internal/common"
)

type grid struct {
	data []string
	posX int
	posY int
}

type slope struct {
	u, d, l, r int
}

var slopeMap = map[common.Part][]string{
	common.Part1: {
		"r3,d1",
	},
	common.Part2: {
		"r1,d1",
		"r3,d1",
		"r5,d1",
		"r7,d1",
		"r1,d2",
	},
}

type gridObject int

const (
	empty gridObject = iota
	tree
	unknown
)

// Solve returns the solution for day3 problem
func Solve(s []string, p common.Part) int {

	var m grid
	var slopes []*slope

	m.load(s)

	for _, s := range slopeMap[p] {
		sl, err := newSlope(s)

		if err != nil {
			log.Panic(err)
		}

		slopes = append(slopes, sl)
	}

	result := 1

	for _, s := range slopes {
		m.moveTopLeft()

		treeCounter := 0

		for !m.offLimits() {
			if m.whatsHere() == tree {
				treeCounter++
			}
			m.jump(s)
		}

		result = result * treeCounter

	}

	return result
}

func (t *grid) load(s []string) {
	t.data = s
}

func (t *grid) jump(s *slope) {

	t.posX = t.posX - s.l + s.r
	t.posY = t.posY - s.d + s.u

	// "simulates" a map that repeats
	// itself on the x-axis
	if t.posX > len(t.data[0])-1 {
		t.posX = t.posX - len(t.data[0])
	}
}

func (t *grid) whatsHere() gridObject {

	y := len(t.data) - 1 - t.posY

	if len(t.data[y]) == 0 {
		return unknown
	}

	c := string(t.data[y][t.posX])

	switch c {
	case ".":
		return empty
	case "#":
		return tree
	default:
		return unknown
	}
}

func (t *grid) offLimits() bool {
	return t.posY < 0
}

func (t *grid) moveTopLeft() {
	t.posX = 0
	t.posY = len(t.data) - 1
}

func newSlope(m string) (*slope, error) {
	s := new(slope)
	err := s.parse(m)

	if err != nil {
		return s, err
	}

	return s, nil
}

func (s *slope) parse(m string) error {

	t := strings.Split(m, ",")

	for _, c := range t {

		if string(c[0]) == "r" || string(c[0]) == "R" {

			i, err := strconv.Atoi(c[1:])

			if err != nil {
				return fmt.Errorf("could not parse slope")
			}

			s.r = s.r + i
		}

		if string(c[0]) == "d" || string(c[0]) == "D" {

			i, err := strconv.Atoi(c[1:])

			if err != nil {
				return fmt.Errorf("could not parse slope")
			}

			s.d = s.d + i
		}

		if string(c[0]) == "u" || string(c[0]) == "U" {

			i, err := strconv.Atoi(c[1:])

			if err != nil {
				return fmt.Errorf("could not parse slope")
			}

			s.u = s.u + i
		}

		if string(c[0]) == "l" || string(c[0]) == "L" {

			i, err := strconv.Atoi(c[1:])

			if err != nil {
				return fmt.Errorf("could not parse slope")
			}

			s.l = s.l + i
		}
	}

	return nil
}
