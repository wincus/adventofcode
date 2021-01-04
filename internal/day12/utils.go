package day12

import (
	"log"
	"strconv"

	"github.com/wincus/adventofcode/internal/common"
)

type opcode int

type inst struct {
	action opcode
	value  int
}

type nav struct {
	direction    opcode
	instructions []inst
	x, y         int
}

const (
	N opcode = iota
	W
	S
	E
	L
	R
	F
)

// Solve returns the solution for day12 problem
func Solve(s []string, p common.Part) int {

	var n nav

	n = parse(s)

	n.navigate()

	return abs(n.x) + abs(n.y)

}

func (n *nav) navigate() {

	for _, i := range n.instructions {

		switch i.action {

		case N, S, E, W, F:
			n.move(i.value, i.action)
		case L:
			n.direction = rotate(i.value, n.direction)
		case R:
			n.direction = rotate(i.value*-1, n.direction)
		}
	}
}

func (n *nav) move(v int, o opcode) {

	switch o {
	case N:
		n.y += v
	case S:
		n.y -= v
	case E:
		n.x += v
	case W:
		n.x -= v
	case F:
		n.move(v, n.direction)
	}

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

func parse(s []string) nav {

	var n nav
	var r []inst

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		r = append(r, newInst(line))

	}

	n.direction = E // initial direction
	n.instructions = r

	return n
}

func newInst(s string) inst {

	var i inst

	switch string(s[0]) {
	case "N":
		i.action = N
	case "S":
		i.action = S
	case "E":
		i.action = E
	case "W":
		i.action = W
	case "L":
		i.action = L
	case "R":
		i.action = R
	case "F":
		i.action = F
	default:
		log.Printf("action not supported: %v", string(s[0]))
		return i
	}

	v, err := strconv.Atoi(s[1:])

	if err != nil {
		log.Printf("could not parse instruction: %v", err)
	}

	i.value = v
	return i
}
