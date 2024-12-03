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
	scanner := bufio.NewScanner(strings.NewReader(input))

	result := 0
	for scanner.Scan() {
		var ints []int
		for _, s := range strings.Split(scanner.Text(), " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Not a number:", s, err)
				os.Exit(1)
			}
			ints = append(ints, n)
		}

		if !part2 {
			valid := true
			for i := 0; i < len(ints)-1; i++ {
				diff := int(math.Abs(float64(ints[i] - ints[i+1])))
				if diff < 1 || diff > 3 {
					valid = false
				}
			}

			sorted := slices.Sorted(slices.Values(ints))
			reversed := slices.Clone(sorted)
			slices.Reverse(reversed)
			if !(slices.Equal(ints, sorted) || slices.Equal(ints, reversed)) {
				valid = false
			}

			if valid {
				result++
			}
		}

		if part2 {
			valid := true
			for i := 0; i < len(ints)-1; i++ {
				diff := int(math.Abs(float64(ints[i] - ints[i+1])))
				if diff < 1 || diff > 3 {
					valid = false
				}
			}

			sorted := slices.Sorted(slices.Values(ints))
			reversed := slices.Clone(sorted)
			slices.Reverse(reversed)
			if !(slices.Equal(ints, sorted) || slices.Equal(ints, reversed)) {
				valid = false
			}

			if !valid {
				for x := range ints {
					ints2 := slices.Clone(ints)
					ints2 = slices.Delete(ints2, x, x+1)
					for i := 0; i < len(ints2)-1; i++ {
						diff := int(math.Abs(float64(ints2[i] - ints2[i+1])))
						if diff < 1 || diff > 3 {
							valid = false
						}
					}
		
					sorted := slices.Sorted(slices.Values(ints2))
					reversed := slices.Clone(sorted)
					slices.Reverse(reversed)
					if !(slices.Equal(ints2, sorted) || slices.Equal(ints2, reversed)) {
						valid = false
					}
				}
			}

			if valid {
				result++
			}
		}
	}

	return result
}
