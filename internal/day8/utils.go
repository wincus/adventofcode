package day8

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode/internal/common"
)

type operation int

const (
	acc operation = iota
	jmp
	nop
	err
)

type instruction struct {
	op  operation
	arg int
}

type program struct {
	instructions   []instruction
	counter        int
	programCounter int
}

// errInfiniteLoop is returned when the program
// was detected to enter an infinite loop
var errInfiniteLoop = errors.New("infinite loop detected")

// map to track if a given instruction has been
// corrected
var hasBeenCorrected = make(map[int]bool)

// Solve returns the solution for day6 problem
func Solve(s []string, p common.Part) int {

	if p == common.Part1 {

		prog, err := load(s)

		if err != nil {
			log.Printf("could not parse program: %v", err)
		}

		c, err := run(prog)

		if err != nil {
			log.Printf("error running program: %v", err)
		}

		return c

	}

	if p == common.Part2 {

		for t := 0; t < len(s); t++ {

			prog, err := loadAndCorrect(s, t)

			if err != nil {
				log.Printf("could not parse program: %v", err)
			}

			c, err := run(prog)

			if err == errInfiniteLoop {
				continue
			}

			log.Printf("changing instruction %v the program finished", t)

			return c

		}
	}

	return 0
}

func run(p program) (int, error) {

	var seen = make(map[int]bool)

	for {

		// the instruction pointed by the program Counter
		// was already executed
		if seen[p.programCounter] {
			return p.counter, errInfiniteLoop
		}

		// reached end of the program
		if p.programCounter == len(p.instructions) {
			log.Printf("reached the end of the program")
			return p.counter, nil
		}

		seen[p.programCounter] = true

		i := p.instructions[p.programCounter]

		switch i.op {
		case nop:
			p.programCounter++
		case acc:
			p.counter = p.counter + i.arg
			p.programCounter++
		case jmp:
			p.programCounter = p.programCounter + i.arg
		}
	}
}

func load(s []string) (program, error) {

	var p program

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		i, err := parse(line)

		if err != nil {
			return p, err
		}

		p.instructions = append(p.instructions, i)
	}

	return p, nil
}

func loadAndCorrect(s []string, threshold int) (program, error) {

	var p program

	for c, line := range s {

		if len(line) == 0 {
			continue
		}

		i, err := parse(line)

		if err != nil {
			return p, err
		}

		// correct instruction, one at a time
		if c < threshold && !hasBeenCorrected[c] {

			hasBeenCorrected[c] = true

			if i.op == jmp {
				i.op = nop
			} else if i.op == nop {
				i.op = jmp
			}
		}

		p.instructions = append(p.instructions, i)
	}

	return p, nil
}

func parse(s string) (instruction, error) {

	var i instruction

	tokens := strings.Split(s, " ")

	if len(tokens) != 2 {
		return i, fmt.Errorf("didn't get 2 tokens as expected")
	}

	op, err := parseOperation(tokens[0])

	if err != nil {
		return i, err
	}

	arg, err := parseArg(tokens[1])

	if err != nil {
		return i, err
	}

	return instruction{op, arg}, nil
}

func parseOperation(s string) (operation, error) {

	switch s {
	case "nop":
		return nop, nil
	case "acc":
		return acc, nil
	case "jmp":
		return jmp, nil
	default:
		return err, fmt.Errorf("unknown op: %v", s)
	}
}

func parseArg(s string) (int, error) {
	return strconv.Atoi(s)
}
