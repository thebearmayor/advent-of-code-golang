package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

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
	scanner := bufio.NewScanner(strings.NewReader(input))

	firsts := make([]int, 0, 1000)
	seconds := make([]int, 0, 1000)
	freqs := make(map[int]int)

	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "   ", 2)
		first, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Not a number:", parts[0], err)
		}

		second, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Not a number:", parts[1], err)
		}

		firsts = append(firsts, first)
		seconds = append(seconds, second)
		freqs[second] += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanning error:", err)
		os.Exit(1)
	}

	slices.Sort(firsts)
	slices.Sort(seconds)

	out1 := 0
	for i := range firsts {
		out1 += int(math.Abs(float64(firsts[i] - seconds[i])))
	}

	out2 := 0
	for i := range firsts {
		freq := freqs[firsts[i]]
		out2 += firsts[i] * freq
	}

	if part2 {
		return out2
	}

	return out1
}
