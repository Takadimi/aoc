package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/Takadimi/aoc/2022/file"
)

var inputFileFlag = flag.String("inputFile", "sample.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	inputFile := *inputFileFlag

	calorieEntries, err := file.Lines(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	caloriesByElf, err := parseCalorieEntries(calorieEntries)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part 1:", partOne(caloriesByElf))
	fmt.Println("Part 2:", partTwo(caloriesByElf))
}

func partOne(caloriesByElf []int) int {
	sortByHighestCalories(caloriesByElf)
	return caloriesByElf[0]
}

func partTwo(caloriesByElf []int) int {
	sortByHighestCalories(caloriesByElf)
	sumOfLastThree := 0
	for _, calories := range caloriesByElf[:3] {
		sumOfLastThree += calories
	}
	return sumOfLastThree
}

func parseCalorieEntries(calorieEntries []string) ([]int, error) {
	caloriesByElf := []int{}
	currentCaloriesForElf := 0
	for _, calorieEntry := range calorieEntries {
		if calorieEntry == "" {
			caloriesByElf = append(caloriesByElf, currentCaloriesForElf)
			currentCaloriesForElf = 0
			continue
		}

		calorieCount, err := strconv.Atoi(calorieEntry)
		if err != nil {
			return nil, err
		}
		currentCaloriesForElf += calorieCount
	}
	caloriesByElf = append(caloriesByElf, currentCaloriesForElf)
	return caloriesByElf, nil
}

func sortByHighestCalories(caloriesByElf []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesByElf)))
}
