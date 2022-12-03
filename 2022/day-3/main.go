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
			p := priorityByItemType[char]
			rucksack.All[p] = true
			if i < compartentSize {
				rucksack.A[p] = true
			} else {
				rucksack.B[p] = true
			}
		}
		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks, nil
}

func mapItemTypeToPriority() map[rune]int {
	priorityMap := map[rune]int{}
	priority := 0

	for itemType := 'a'; itemType <= 'z'; itemType++ {
		priority++
		priorityMap[itemType] = priority
	}
	for itemType := 'A'; itemType <= 'Z'; itemType++ {
		priority++
		priorityMap[itemType] = priority
	}

	return priorityMap
}

var priorityByItemType = mapItemTypeToPriority()
