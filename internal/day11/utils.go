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

	n := l.countOccupiedNeighbors(s, p)

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

func (l layout) countOccupiedNeighbors(i int, p common.Part) int {

	var counter int
	var n []seat

	if p == common.Part1 {
		n = l.localNeighbours(i)
	}

	if p == common.Part2 {
		n = l.remoteNeighbours(i)
	}

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

// i only care for occupied sets
func (l layout) remoteNeighbours(i int) []seat {

	var n = make([]seat, 0, 8)

	for _, dir := range []direction{up, down, left, right, upleft, upright, downleft, downright} {
		n = append(n, l.directedRemoteNeighour(i, dir))
	}

	return n
}

func (l layout) directedRemoteNeighour(i int, d direction, p common.Part) seat {

	var s seat = floor

	// convert index into a 2D space
	x, y := convertTo2D(i, l.width)

	switch d {
	case up:

		if y == 0 {
			return s
		}

		c := 1

		for {
			yc := y + c

			ic := convertTo1D(x, yc, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if xc == l.width-1 || yc == l.width-1 || p == common.Part1 {
				return s
			}

			c++
		}

		for yc := y - 1; yc >= 0; yc-- {

			ic := convertTo1D(x, yc, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}
		}

	case down:
		if y == l.width {
			return s
		}

		for yc := y + 1; yc < l.width; yc++ {

			ic := convertTo1D(x, yc, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}
		}

	case left:

		if x == 0 {
			return s
		}

		for xc := x - 1; xc >= 0; xc-- {

			ic := convertTo1D(xc, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}
		}

	case right:

		if x == l.width {
			return s
		}

		for xc := x + 1; xc < l.width; xc++ {

			ic := convertTo1D(xc, y, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}
		}

	case upleft:

		if x == 0 || y == 0 {
			return s
		}

		c := 1

		for {

			xc := x - c
			yc := y - c

			ic := convertTo1D(xc, yc, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if xc == 0 || yc == 0 {
				return s
			}

			c++

		}

	case upright:

		if x == l.width-1 || y == 0 {
			return s
		}

		c := 1

		for {

			xc := x + c
			yc := y - c

			ic := convertTo1D(xc, yc, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if xc == l.width-1 || yc == 0 {
				return s
			}

			c++

		}

	case downleft:

		if x == 0 || y == l.width-1 {
			return s
		}

		c := 1

		for {

			xc := x - c
			yc := y + c

			ic := convertTo1D(xc, yc, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if xc == 0 || yc == l.width-1 {
				return s
			}

			c++

		}

	case downright:

		if x == l.width-1 || y == l.width-1 {
			return s
		}

		c := 1

		for {

			xc := x + c
			yc := y + c

			ic := convertTo1D(xc, yc, l.width)

			if l.seats[ic] == empty || l.seats[ic] == occupied {
				return l.seats[ic]
			}

			if xc == l.width-1 || yc == l.width-1 {
				return s
			}

			c++

		}

	}

	return s

}

func (l layout) localNeighbours(i int) []seat {

	var n = make([]seat, 0, 8)

	// convert index into 2D
	x, y := convertTo2D(i, l.width)

	for xc := x - 1; xc <= x+1; xc++ {
		for yc := y - 1; yc <= y+1; yc++ {

			if xc < 0 || xc >= l.width || yc < 0 || yc > l.width {
				continue
			}

			ic := convertTo1D(xc, yc, l.width)

			if ic < 0 || i == ic || ic >= l.width*l.width {
				continue
			}

			n = append(n, l.seats[ic])
		}
	}

	return n

}

func convertTo2D(i, w int) (x, y int) {
	return i % w, i / w
}

func convertTo1D(x, y, w int) int {
	return x + w*y
}
