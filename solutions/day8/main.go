package main

import (
	"log"

	"github.com/wincus/adventofcode/internal/common"
	"github.com/wincus/adventofcode/internal/day8"
)

func main() {

	d, err := common.GetData(8)

	if err != nil {
		log.Panicf("no data, no game ... sorry: %v", err)
	}

	for _, p := range []common.Part{common.Part1, common.Part2} {
		log.Printf("Solution for Part %v: %v", p, day8.Solve(d, p))
	}
}
