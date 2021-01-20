package day14

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode/internal/common"
)

const size = 32

type memory map[int]int

type mask struct {
	unset int // and-mask
	set   int // or-mask
}

// Solve returns the solution for day13 problem
func Solve(s []string, p common.Part) int {

	var mem = make(memory)
	var m mask

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "mask") {
			m.parse(line)
		}

		if strings.HasPrefix(line, "mem") {
			mem.parse(line, m)
		}
	}

	var sum int

	for _, v := range mem {
		sum += v
	}

	return sum
}

func (m *mask) parse(s string) {

	// initialize mask
	m.set = 0
	m.unset = 1<<size - 1

	sMask := strings.Split(s, " ")[2]

	for i := len(sMask) - 1; i >= 0; i-- {

		c := string(sMask[i])
		p := len(sMask) - i - 1

		switch c {
		case "X":
			continue
		case "0":
			m.unset -= 1 << p
		case "1":
			m.set += 1 << p
		default:
			log.Printf("could not recognize mask char: %v", c)
		}
	}
}

func (m *mask) apply(i int) int {
	return i&m.unset | m.set
}

func (m *mask) string() string {
	return fmt.Sprintf("and: %b\nor:  %b\n", m.unset, m.set)
}

func (mem memory) parse(s string, m mask) error {

	re := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	d := re.FindStringSubmatch(s)

	if d == nil {
		return fmt.Errorf("could not parse instruction: %v", s)
	}

	addr, err := strconv.Atoi(d[1])

	if err != nil {
		return err
	}

	value, err := strconv.Atoi(d[2])

	if err != nil {
		return err
	}

	mem[addr] = m.apply(value)

	return nil

}
