package main

import (
	"log"

	"github.com/wincus/adventofcode2021/internal/common"
	"github.com/wincus/adventofcode2021/internal/day14"
)

func main() {

	d, err := common.GetData(14)

	if err != nil {
		log.Panicf("no data, no game ... sorry!")
	}

	for _, p := range []common.Part{common.Part1, common.Part2} {
		log.Printf("Solution for Part %v: %v", p, day14.Solve(d, p))
	}
}
