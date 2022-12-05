package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

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

	sections := splitBySection(lines)
	if len(sections) != 2 {
		panic("expected 2 sections")
	}
	startingStacksSection, procedureSection := sections[0], sections[1]

	procedure := parseProcedure(procedureSection)

	partOneStacks := parseStacks(startingStacksSection)
	partTwoStacks := parseStacks(startingStacksSection)

	fmt.Println("Part one:", partOne(partOneStacks, procedure))
	fmt.Println("Part two:", partTwo(partTwoStacks, procedure))
}

func partOne(stacks [][]string, procedure []Instruction) string {
	for _, instruction := range procedure {
		fromStack := stacks[instruction.From]
		toStack := stacks[instruction.To]

		for i := 0; i < instruction.Count; i++ {
			top := fromStack[len(fromStack)-1]
			toStack = append(toStack, top)
			fromStack = fromStack[:len(fromStack)-1]
		}

		stacks[instruction.From] = fromStack
		stacks[instruction.To] = toStack
	}

	return topItems(stacks)
}

func partTwo(stacks [][]string, procedure []Instruction) string {
	for _, instruction := range procedure {
		fromStack := stacks[instruction.From]
		toStack := stacks[instruction.To]

		topNItems := fromStack[len(fromStack)-instruction.Count:]
		toStack = append(toStack, topNItems...)
		fromStack = fromStack[:len(fromStack)-instruction.Count]

		stacks[instruction.From] = fromStack
		stacks[instruction.To] = toStack
	}

	return topItems(stacks)
}

func topItems(stacks [][]string) string {
	topItems := ""
	for i := 1; i < len(stacks); i++ {
		stack := stacks[i]
		topItems += stack[len(stack)-1]
	}
	return topItems
}

func parseStacks(lines []string) [][]string {
	stackIDsLine := lines[len(lines)-1]
	stackIDIndexes := map[int]int{}
	for i, id := range stackIDsLine {
		if id != ' ' {
			idValue, err := strconv.Atoi(string(id))
			if err != nil {
				panic(err)
			}
			stackIDIndexes[i] = idValue
		}
	}

	stacks := make([][]string, len(stackIDIndexes)+1)

	for _, l := range lines[:len(lines)-1] {
		for i, char := range l {
			if unicode.IsLetter(char) {
				stackIDIndex := stackIDIndexes[i]
				stacks[stackIDIndex] = append([]string{string(char)}, stacks[stackIDIndex]...)
			}
		}
	}

	return stacks
}

type Instruction struct {
	Count int
	From  int
	To    int
}

func parseProcedure(lines []string) []Instruction {
	procedure := []Instruction{}
	for _, l := range lines {
		fields := strings.Fields(l)
		if len(fields) != 6 {
			panic("expected 6 fields in instruction")
		}
		countField, fromField, toField := fields[1], fields[3], fields[5]
		count, err := strconv.Atoi(countField)
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(fromField)
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(toField)
		if err != nil {
			panic(err)
		}

		procedure = append(procedure, Instruction{Count: count, From: from, To: to})
	}

	return procedure
}

func splitBySection(lines []string) [][]string {
	sections := [][]string{}

	currentSection := []string{}
	for _, l := range lines {
		if l == "" {
			sections = append(sections, currentSection)
			currentSection = []string{}
			continue
		}

		currentSection = append(currentSection, l)
	}
	sections = append(sections, currentSection)

	return sections
}
