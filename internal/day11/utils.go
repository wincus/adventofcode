package day11

import (
	"github.com/wincus/adventofcode/internal/common"
)

type seat int
type direction int

const (
	floor seat = iota
	empty
	occupied
	unknown
)

const (
	up direction = iota
	upleft
	left
	downleft
	down
	downright
	right
	upright
)

var directions = [...]direction{up, upleft, left, downleft, down, downright, right, upright}

type layout struct {
	seats      []seat
	width      int
	genChanges int
}

// Solve returns the solution for day11 problem
func Solve(s []string, p common.Part) int {

	var l layout

	l = newLayout(s)

	for {
		l.next(p)

		if l.genChanges == 0 {
			return l.count()
		}
	}
}

func (l *layout) next(p common.Part) {

	temp := make([]seat, l.width*l.width)

	l.genChanges = 0

	for i := 0; i < len(l.seats); i++ {
		temp[i] = l.populate(i, p)

		if temp[i] != l.seats[i] {
			l.genChanges++
		}
	}

	l.seats = temp
}

func newLayout(d []string) layout {

	var l layout
	var seats = make([]seat, len(d[0])*len(d))
	var counter int

	for _, line := range d {

		if len(line) == 0 {
			continue
		}

		for _, c := range line {
			seats[counter] = parseSeat(string(c))
			counter++
		}
	}

	l.seats = seats
	l.width = len(d[0])

	return l

}

func parseSeat(c string) seat {

	switch c {
	case ".":
		return floor
	case "L":
		return empty
	case "#":
		return occupied
	default:
		return unknown
	}

}

func (s seat) String() string {
	switch s {
	case floor:
		return "."
	case empty:
		return "L"
	case occupied:
		return "#"
	case unknown:
		return "?"
	default:
		return "???"
	}
}

func (l *layout) populate(s int, p common.Part) seat {

	var maxNeighbors int

	if p == common.Part1 {
		maxNeighbors = 4
	}

	if p == common.Part2 {
		maxNeighbors = 5
	}

	if l.seats[s] == floor {
		return floor
	}

	n := l.countNeighbours(s, p)

	// no neighbours
	if n == 0 {
		return occupied
	}

	// to crowded
	if n >= maxNeighbors {
		return empty
	}

	return l.seats[s]
}

func (l layout) countNeighbours(i int, p common.Part) int {

	var counter int
	var n []seat

	n = l.neighbours(i, p)

	for _, s := range n {
		if s == occupied {
			counter++
		}
	}

	return counter

}

func (l layout) count() int {

	var counter int

	for i := 0; i < len(l.seats); i++ {
		if l.seats[i] == occupied {
			counter++
		}
	}
	return counter
}

func (l layout) neighbours(i int, p common.Part) []seat {

	var n = make([]seat, 0, 8)

	for _, dir := range directions {
		n = append(n, l.getNeighbour(i, dir, p))
	}

	return n
}

func (l layout) getNeighbour(i int, d direction, p common.Part) seat {

	var s seat = floor

	// convert index into a 2D space
	x, y := convertTo2D(i, l.width)

	switch d {
	case up:

		for {

			if y <= 0 {
				return s
			}

			y--

			ic := convertTo1D(x, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if p == common.Part1 {
				return s
			}
		}

	case down:

		for {

			if y >= l.width-1 {
				return s
			}

			y++

			ic := convertTo1D(x, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if p == common.Part1 {
				return s
			}
		}

	case left:

		for {

			if x <= 0 {
				return s
			}

			x--

			ic := convertTo1D(x, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if p == common.Part1 {
				return s
			}
		}

	case right:

		for {

			if x >= l.width-1 {
				return s
			}

			x++

			ic := convertTo1D(x, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if p == common.Part1 {
				return s
			}
		}

	case upleft:

		for {

			if x <= 0 || y <= 0 {
				return s
			}

			x--
			y--

			ic := convertTo1D(x, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if p == common.Part1 {
				return s
			}
		}

	case upright:

		for {

			if x >= l.width-1 || y <= 0 {
				return s
			}

			x++
			y--

			ic := convertTo1D(x, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if p == common.Part1 {
				return s
			}
		}

	case downleft:

		for {

			if x <= 0 || y >= l.width-1 {
				return s
			}

			x--
			y++

			ic := convertTo1D(x, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if p == common.Part1 {
				return s
			}
		}

	case downright:

		for {

			if x >= l.width-1 || y >= l.width-1 {
				return s
			}

			x++
			y++

			ic := convertTo1D(x, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if p == common.Part1 {
				return s
			}
		}

	}

	return s

}

func convertTo2D(i, w int) (x, y int) {
	return i % w, i / w
}

func convertTo1D(x, y, w int) int {
	return x + w*y
}
