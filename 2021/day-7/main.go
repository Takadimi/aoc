package main

import (
	"flag"
	"fmt"
	"math"
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

	crabPositions, err := parseCrabPositions(lines[0])

	fmt.Println("Part one: ", partOne(crabPositions))
	fmt.Println("Part two: ", partTwo(crabPositions))
}

func partOne(crabPositions []int) int {
	highestPos := highestPosition(crabPositions)
	cheapestFuelCost := totalFuelCostForPositionAtConstantBurn(crabPositions, 0)
	for i := 1; i <= highestPos; i++ {
		fuelCost := totalFuelCostForPositionAtConstantBurn(crabPositions, i)
		if fuelCost < cheapestFuelCost {
			cheapestFuelCost = fuelCost
		}
	}

	return cheapestFuelCost
}

func partTwo(crabPositions []int) int {
	highestPos := highestPosition(crabPositions)
	cheapestFuelCost := totalFuelCostForPositionAtIncrementalBurn(crabPositions, 0)
	for i := 1; i <= highestPos; i++ {
		fuelCost := totalFuelCostForPositionAtIncrementalBurn(crabPositions, i)
		if fuelCost < cheapestFuelCost {
			cheapestFuelCost = fuelCost
		}
	}

	return cheapestFuelCost
}

func highestPosition(crabPositions []int) int {
	highest := 0
	for _, position := range crabPositions {
		if position > highest {
			highest = position
		}
	}
	return highest
}

func totalFuelCostForPositionAtConstantBurn(crabPositions []int, targetPosition int) int {
	totalFuelCost := 0
	for _, position := range crabPositions {
		totalFuelCost += int(math.Abs(float64(position - targetPosition)))
	}

	return totalFuelCost
}

func totalFuelCostForPositionAtIncrementalBurn(crabPositions []int, targetPosition int) int {
	totalFuelCost := 0
	for _, position := range crabPositions {
		distance := int(math.Abs(float64(position - targetPosition)))
		fuelCost := 0
		for i := 1; i <= distance; i++ {
			fuelCost += i
		}
		totalFuelCost += fuelCost
	}

	return totalFuelCost
}

func parseCrabPositions(line string) ([]int, error) {
	parts := strings.Split(line, ",")
	positions := []int{}
	for _, p := range parts {
		position, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}

	return positions, nil
}
