package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var inputFileFlag = flag.String("inputFile", "sample.txt", "Relative file path to use as input.")

func lines(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result, nil
}

func partOne(caloriesByElf []int) int {
	maxCalories := 0
	for _, calories := range caloriesByElf {
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	return maxCalories
}

func partTwo(caloriesByElf []int) int {
	caloriesSorted := caloriesByElf
	sort.Ints(caloriesSorted)
	lastThree := caloriesSorted[len(caloriesSorted)-3:]
	sumOfLastThree := 0
	for _, calories := range lastThree {
		sumOfLastThree += calories
	}
	return sumOfLastThree
}

func main() {
	flag.Parse()
	inputFile := *inputFileFlag

	calorieEntries, err := lines(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
			fmt.Println(err)
			os.Exit(1)
		}
		currentCaloriesForElf += calorieCount
	}
	caloriesByElf = append(caloriesByElf, currentCaloriesForElf)

	fmt.Println("Part 1:", partOne(caloriesByElf))
	fmt.Println("Part 2:", partTwo(caloriesByElf))
}
