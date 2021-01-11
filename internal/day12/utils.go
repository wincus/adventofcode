package day12

import (
	"log"
	"strconv"

	"github.com/wincus/adventofcode/internal/common"
)

type pos struct {
	dir opcode
	m   map[opcode]int
}

type opcode int

const (
	E opcode = iota
	N
	W
	S
)

// Solve returns the solution for day12 problem
func Solve(s []string, p common.Part) int {

	var boat pos
	var wayPoint pos

	boat.m = make(map[opcode]int)
	wayPoint.m = make(map[opcode]int)

	// initial WayPoint conditions
	wayPoint.m[E] = 10
	wayPoint.m[N] = 1

	// for each line in input
	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		if p == common.Part1 {
			err := boat.move(line)

			if err != nil {
				log.Printf("Warning, could not parse instruction: %v", line)
			}
		}

		if p == common.Part2 {
			err := boat.moveWithWayPoint(line, wayPoint)

			if err != nil {
				log.Printf("Warning, could not parse instruction: %v", line)
			}

		}

	}

	return abs(boat.m[E]-boat.m[W]) + abs(boat.m[N]-boat.m[S])

}

func (p *pos) move(s string) error {

	v, err := strconv.Atoi(s[1:])

	if err != nil {
		return err
	}

	switch string(s[0]) {
	case "N":
		p.m[N] += v
	case "S":
		p.m[S] += v
	case "E":
		p.m[E] += v
	case "W":
		p.m[W] += v
	case "L":
		p.dir = rotate(v, p.dir)
	case "R":
		p.dir = rotate(-1*v, p.dir)
	case "F":
		p.m[p.dir] += v
	}

	return nil

}

func (p *pos) moveWithWayPoint(s string, w pos) error {

	v, err := strconv.Atoi(s[1:])

	if err != nil {
		return err
	}

	switch string(s[0]) {
	case "N":
		w.move(s)
	case "S":
		w.move(s)
	case "E":
		w.move(s)
	case "W":
		w.move(s)
	case "L":
		w.m[rotate(v, E)], w.m[rotate(v, N)], w.m[rotate(v, W)], w.m[rotate(v, S)] = w.m[E], w.m[N], w.m[W], w.m[S]
	case "R":
		nv := -1 * v
		w.m[rotate(nv, E)], w.m[rotate(nv, N)], w.m[rotate(nv, W)], w.m[rotate(nv, S)] = w.m[E], w.m[N], w.m[W], w.m[S]
	case "F":
		p.m[E] += w.m[E] * v
		p.m[N] += w.m[N] * v
		p.m[W] += w.m[W] * v
		p.m[S] += w.m[S] * v
	}

	return nil
}

func rotate(n int, o opcode) opcode {

	n = n % 360

	if n < 0 {
		n = 360 + n
	}

	return opcode((n/90 + int(o)) % 4)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
