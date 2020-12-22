package day7

import (
	"log"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode/internal/common"
)

var graph = make(map[string]map[string]int)

// Solve returns the solution for day6 problem
func Solve(s []string, p common.Part) int {

	var counter int

	if len(graph) == 0 {
		// register all known bagTypes
		for _, line := range s {
			err := register(line)

			if err != nil {
				log.Printf("could not process line: %v", err)
			}
		}
	}

	if p == common.Part1 {
		for bag := range graph {
			if canHold(bag, "shiny gold") {
				counter++
			}
		}
	}

	if p == common.Part2 {
		counter = holdSum("shiny gold") - 1
	}

	return counter

}

// Given a bag type s returns how many
// bags are on it
func holdSum(s string) int {

	var counter int = 1

	for bag := range graph[s] {
		counter += graph[s][bag] * holdSum(bag)
	}

	return counter
}

func canHold(s, p string) bool {

	if s == p {
		return false // a bag cannot hold itself!
	}

	for bag := range graph[s] {
		if bag == p {
			return true
		}

		if canHold(bag, p) {
			return true
		}
	}

	return false
}

func register(s string) error {

	if len(s) == 0 {
		return nil
	}

	tokens := strings.Split(s, " ")

	from := strings.Join(tokens[:2], " ")

	graph[from] = make(map[string]int)

	if strings.HasSuffix(s, "contain no other bags.") {
		return nil
	}

	deps := strings.Split(strings.Join(tokens[4:], " "), ",")

	for _, line := range deps {

		tokens := strings.Split(line, " ")

		if len(tokens[0]) == 0 {
			tokens = tokens[1:]
		}

		q, err := strconv.Atoi(strings.TrimSpace(tokens[0]))

		if err != nil {
			log.Printf("could not parse line: %v", err)
		}

		to := strings.Join(tokens[1:3], " ")

		graph[from][to] = q

	}

	return nil

}
