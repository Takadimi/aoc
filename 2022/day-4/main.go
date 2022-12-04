package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Takadimi/aoc/2022/file"
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

	assignmentPairs, err := parseAssignmentPairs(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part One:", partOne(assignmentPairs))
	fmt.Println("Part Two:", partTwo(assignmentPairs))
}

func partOne(pairs [][2]Range) int {
	sumOfFullyContainedPairs := 0
	for _, p := range pairs {
		first, second := p[0], p[1]
		longest := first
		other := second
		if second.length() > longest.length() {
			longest = second
			other = first
		}

		if longest.fullyContains(other) {
			sumOfFullyContainedPairs++
		}
	}

	return sumOfFullyContainedPairs
}

func partTwo(pairs [][2]Range) int {
	sumOfFullyIntersectingPairs := 0
	for _, p := range pairs {
		first, second := p[0], p[1]
		if first.intersects(second) {
			sumOfFullyIntersectingPairs++
		}
	}

	return sumOfFullyIntersectingPairs
}

func (r Range) fullyContains(otherRange Range) bool {
	return r.Start <= otherRange.Start && r.End >= otherRange.End
}

func (r Range) intersects(otherRange Range) bool {
	return r.Start <= otherRange.End && r.End >= otherRange.Start
}

func (r Range) length() int {
	return (r.End - r.Start) + 1
}

func parseAssignmentPairs(lines []string) ([][2]Range, error) {
	assignmentPairs := [][2]Range{}
	for _, l := range lines {
		pairParts := strings.Split(l, ",")
		if len(pairParts) != 2 {
			return nil, errors.New("unexpected pair parts length")
		}

		firstRange, err := parseRange(pairParts[0])
		if err != nil {
			return nil, err
		}
		secondRange, err := parseRange(pairParts[1])
		if err != nil {
			return nil, err
		}

		assignmentPairs = append(assignmentPairs, [2]Range{firstRange, secondRange})
	}

	return assignmentPairs, nil
}

func parseRange(rangeString string) (Range, error) {
	rangeParts := strings.Split(rangeString, "-")
	if len(rangeParts) != 2 {
		return Range{}, errors.New("unexpected range parts length")
	}

	rangeStart, err := strconv.Atoi(rangeParts[0])
	if err != nil {
		return Range{}, err
	}
	rangeEnd, err := strconv.Atoi(rangeParts[1])
	if err != nil {
		return Range{}, err
	}

	return Range{rangeStart, rangeEnd}, nil
}

type Range struct {
	Start, End int
}
