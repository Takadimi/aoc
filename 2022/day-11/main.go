package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
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

	sections := splitBySection(lines)

	monkeys, err := parseMonkeySections(sections)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part one:", partOne(monkeys))
}

func partOne(monkeys []Monkey) int {
	for round := 0; round < 20; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkeys[i].InspectionCount++
				newWorryLevel := monkey.Operation(item)
				newWorryLevel /= 3
				if monkey.Test(newWorryLevel) {
					monkeys[monkey.MonkeyToThrowToIfTrue].Items = append(monkeys[monkey.MonkeyToThrowToIfTrue].Items, newWorryLevel)
				} else {
					monkeys[monkey.MonkeyToThrowToIfFalse].Items = append(monkeys[monkey.MonkeyToThrowToIfFalse].Items, newWorryLevel)
				}
				monkeys[i].Items = monkeys[i].Items[1:]
			}
		}
	}

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].InspectionCount > monkeys[j].InspectionCount
	})

	monkeyBusiness := monkeys[0].InspectionCount * monkeys[1].InspectionCount

	return monkeyBusiness
}

type Monkey struct {
	Items                  []int
	Operation              func(int) int
	Test                   func(int) bool
	MonkeyToThrowToIfTrue  int
	MonkeyToThrowToIfFalse int
	InspectionCount        int
}

func parseMonkeySections(sections [][]string) ([]Monkey, error) {
	monkeys := make([]Monkey, len(sections))

	for _, section := range sections {
		if len(section) != 6 {
			return nil, errors.New("not enough lines in monkey section")
		}

		identifierLine := section[0]
		identifierLineFields := strings.Fields(identifierLine)
		if len(identifierLineFields) != 2 {
			return nil, errors.New("malformed identifer line")
		}
		identifier, err := strconv.Atoi(strings.TrimRight(identifierLineFields[1], ":"))
		if err != nil {
			return nil, err
		}

		monkey := monkeys[identifier]

		startingItemsLine := section[1]
		parts := strings.Split(startingItemsLine, ":")
		if len(parts) != 2 {
			return nil, errors.New("malformed starting items line")
		}
		itemsPartSplitByComma := strings.Split(parts[1], ",")
		startingItems := []int{}
		for _, itemStr := range itemsPartSplitByComma {
			item, err := strconv.Atoi(strings.TrimSpace(itemStr))
			if err != nil {
				return nil, err
			}
			startingItems = append(startingItems, item)
		}
		monkey.Items = startingItems

		operationLine := section[2]
		parts = strings.Split(operationLine, ":")
		if len(parts) != 2 {
			return nil, errors.New("malformed operations line")
		}
		operationStr := strings.TrimSpace(parts[1])
		operationFields := strings.Fields(operationStr)
		if len(operationFields) != 5 {
			return nil, errors.New("malformed operation, operation definition not 5 fields")
		}
		if operationFields[0] != "new" || operationFields[1] != "=" {
			return nil, errors.New("malformed operation, does not start with assignment")
		}
		operation := func(oldWorry int) int {
			leftOperandStr := operationFields[2]
			operator := operationFields[3]
			rightOperandStr := operationFields[4]

			var leftOperand, rightOperand int
			if leftOperandStr == "old" {
				leftOperand = oldWorry
			} else {
				value, convErr := strconv.Atoi(leftOperandStr)
				if convErr != nil {
					panic(err)
				}
				leftOperand = value
			}
			if rightOperandStr == "old" {
				rightOperand = oldWorry
			} else {
				value, convErr := strconv.Atoi(rightOperandStr)
				if convErr != nil {
					panic(err)
				}
				rightOperand = value
			}

			if operator == "*" {
				return leftOperand * rightOperand
			}
			if operator == "+" {
				return leftOperand + rightOperand
			}

			panic("unsupported operator for operation")
		}
		monkey.Operation = operation

		testLine := section[3]
		testFields := strings.Fields(testLine)
		divisibleByValue, err := strconv.Atoi(testFields[len(testFields)-1])
		if err != nil {
			return nil, err
		}
		test := func(newWorry int) bool {
			return newWorry%divisibleByValue == 0
		}
		monkey.Test = test

		trueLine := section[4]
		trueLineFields := strings.Fields(trueLine)
		trueMonkeyIdentifier, err := strconv.Atoi(trueLineFields[len(trueLineFields)-1])
		if err != nil {
			return nil, err
		}
		monkey.MonkeyToThrowToIfTrue = trueMonkeyIdentifier

		falseLine := section[5]
		falseLineFields := strings.Fields(falseLine)
		falseMonkeyIdentifier, err := strconv.Atoi(falseLineFields[len(falseLineFields)-1])
		if err != nil {
			return nil, err
		}
		monkey.MonkeyToThrowToIfFalse = falseMonkeyIdentifier

		monkeys[identifier] = monkey
	}

	return monkeys, nil
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
