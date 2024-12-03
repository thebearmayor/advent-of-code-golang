package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {

	// solve part 1 here
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do(?:n't)*\(\)`)
	insts := re.FindAllStringSubmatch(input, -1)

	result := 0
	do := true
	// return insts
	for _, inst := range insts {
		if inst[0] == "do()" {
			do = true
			continue
		}

		if inst[0] == "don't()" {
			do = false
			continue
		}

		if !part2 || do {
			var n, m int
			var err error
			s := inst[1]
			if n, err = strconv.Atoi(s); err != nil {
				fmt.Printf("%T, %v\n", s, s)
			}
			s = inst[2]
			if m, err = strconv.Atoi(s); err != nil {
				fmt.Printf("%T, %v\n", s, s)
			}
			result += n * m
		}
	}
	return result
}
