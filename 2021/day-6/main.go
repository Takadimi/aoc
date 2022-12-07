package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Takadimi/aoc/2021/file"
)

var inputFileFlag = flag.String("inputFile", "sample.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	inputFile := *inputFileFlag

	lines, err := file.Lines(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	initialNumbers, err := parseInitialNumbers(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part one:", partOne(initialNumbers))
	fmt.Println("Part two:", partTwo(initialNumbers))
}

func partOne(initialNumbers []int) int {
	return fishOverDays(initialNumbers, 80)
}

func partTwo(initialNumbers []int) int {
	return fishOverDays(initialNumbers, 256)
}

func fishOverDays(initialNumbers []int, days int) int {
	clock := [9]int{}
	for _, n := range initialNumbers {
		clock[n] = clock[n] + 1
	}

	for i := 0; i < days; i++ {
		oldClock := clock
		clock = [9]int{}
		for day := range oldClock {
			if day == 0 {
				clock[6] += oldClock[day]
				clock[8] += oldClock[day]
				continue
			}

			clock[day-1] += oldClock[day]
		}
	}

	totalFish := 0
	for _, daySum := range clock {
		totalFish += daySum
	}

	return totalFish
}

func parseInitialNumbers(lines []string) ([]int, error) {
	if len(lines) != 1 {
		return nil, errors.New("expected one line of input")
	}
	parts := strings.Split(lines[0], ",")
	numbers := []int{}
	for _, numberString := range parts {
		n, err := strconv.Atoi(numberString)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, int(n))
	}
	return numbers, nil
}

/*
   Starting point: 3
   After 18d, this would generate 3 new fish

   18 - 3 = 15d -- +1 fish

   15 / 6 = 2.5 = 2 -- +2 fish

   Starting point: 4
   After 18d, this would generate 3 new fish

   18 - 4 = 14d  -- +1 fish
   14 / 6 = 2.3 = 2 -- +2 fish

   Starting point: 3
   After 18d, this would generate 3 new fish

   18 - 3 = 15d -- +1 fish

   15 / 6 = 2.5 = 2 -- +2 fish

   Starting point: 1
   After 18d, this would generate 3 new fish

   18 - 1 = 17d -- +1 fish

   17 / 6 = 2.83 = 2 -- +2 fish

   Starting point: 2
   After 18d, this would generate 3 new fish

   18 - 2 = 16d -- +1 fish

   16 / 6 = 2.6 = 2 -- +2 fish

   ---

   What's the freakin' formulaaaa?

   1 fish starting at 3
   18 - 3 = 15 -- +1 fish
   15 - 6 = 9 -- +1 fish
   9 - 6 = 3 -- +1 fish
   3 - 6 = -6 -- +0 fish

   floor(15 / 6) = 2
   2 * (15/8) = 4
*/
