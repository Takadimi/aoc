package main

import (
	"flag"
	"fmt"
	"os"

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

	fmt.Println("Part One:", partOne(lines))
	fmt.Println("Part Two:", partTwo(lines))
}

func partOne(lines []string) int {
	rucksacks, err := parseRucksacks(lines)
	if err != nil {
		panic(err)
	}
	prioritySum := 0
	for _, r := range rucksacks {
		p := priorityOfItemPresentInBothCompartments(r)
		prioritySum += p
	}
	return prioritySum
}

func partTwo(lines []string) int {
	rucksacks, err := parseRucksacks(lines)
	if err != nil {
		panic(err)
	}
	groups := groupRucksacks(rucksacks, 3)
	prioritySum := 0
	for _, g := range groups {
		p := priorityOfItemPresentInAllOfGroup(g)
		prioritySum += p
	}
	return prioritySum
}

type Rucksack struct {
	All map[int]bool
	A   map[int]bool
	B   map[int]bool
}

func parseRucksacks(lines []string) ([]Rucksack, error) {
	rucksacks := []Rucksack{}
	for _, line := range lines {
		rucksack := Rucksack{
			All: make(map[int]bool),
			A:   make(map[int]bool),
			B:   make(map[int]bool),
		}
		compartentSize := len(line) / 2
		for i, char := range line {
			priority := charToPriority[char]
			rucksack.All[priority] = true
			if i < compartentSize {
				rucksack.A[priority] = true
			} else {
				rucksack.B[priority] = true
			}
		}
		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks, nil
}

func priorityOfItemPresentInBothCompartments(r Rucksack) int {
	for ap := range r.A {
		for bp := range r.B {
			if ap == bp {
				return bp
			}
		}
	}
	return 0
}

func priorityOfItemPresentInAllOfGroup(group []Rucksack) int {
	firstRucksack := group[0]
	for p := range firstRucksack.All {
		presentInAll := true
		for _, r := range group[1:] {
			if !r.All[p] {
				presentInAll = false
			}
		}

		if presentInAll {
			return p
		}
	}

	return 0
}

func groupRucksacks(rucksacks []Rucksack, groupCount int) [][]Rucksack {
	groups := make([][]Rucksack, len(rucksacks)/groupCount)

	for i, r := range rucksacks {
		groupIndex := i / groupCount
		groups[groupIndex] = append(groups[groupIndex], r)
	}

	return groups
}

var charToPriority = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}
