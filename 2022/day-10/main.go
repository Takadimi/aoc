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

	instructions, err := parseInstructions(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	registerX := 1
	cycle := 1
	signalStrengthSum := 0

	for _, instruction := range instructions {
		for c := 0; c < instruction.CycleCount; c++ {
			// every 40th cycle, starting at 20, evaluate signal strength
			if (cycle-20)%40 == 0 {
				signalStrengthSum += (registerX * cycle)
			}
			cycle++
		}

		registerX += instruction.Increment
	}

	fmt.Println(signalStrengthSum)
}

type Instruction struct {
	CycleCount int
	Increment  int
}

func parseInstructions(lines []string) ([]Instruction, error) {
	instructions := []Instruction{}

	for _, l := range lines {
		fields := strings.Fields(l)
		if len(fields) == 0 || len(fields) > 2 {
			return nil, errors.New("unknown instruction format")
		}

		switch fields[0] {
		case "noop":
			instructions = append(instructions, Instruction{CycleCount: 1})
		case "addx":
			increment, err := strconv.Atoi(fields[1])
			if err != nil {
				return nil, err
			}
			instructions = append(instructions, Instruction{CycleCount: 2, Increment: increment})
		default:
			return nil, errors.New("unknown instruction type")
		}
	}

	return instructions, nil
}
