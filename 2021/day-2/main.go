package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "sample1.txt", "Relative file path to use as input.")

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

type command struct {
	Direction string
	Amount    int
}

func parseCommands(lines []string) ([]command, error) {
	commands := []command{}
	for _, l := range lines {
		parts := strings.Split(l, " ")
		if len(parts) != 2 {
			return commands, errors.New("invalid command format, expected `{direction} {amount}`")
		}
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			return commands, err
		}
		commands = append(commands, command{parts[0], amount})
	}

	return commands, nil
}

func processCommands(commands []command) (int, int) {
	horizontalPosition := 0
	depth := 0

	for _, command := range commands {
		switch command.Direction {
		case "forward":
			horizontalPosition += command.Amount
		case "down":
			depth += command.Amount // going down increases depth in ocean
		case "up":
			depth -= command.Amount // going up decreases depth in ocean
		default:
			panic("no support for direction " + command.Direction)
		}
	}

	return horizontalPosition, depth
}

func processCommandsWithAim(commands []command) (int, int) {
	aim := 0
	horizontalPosition := 0
	depth := 0

	for _, command := range commands {
		switch command.Direction {
		case "forward":
			horizontalPosition += command.Amount
			depth += aim * command.Amount
		case "down":
			aim += command.Amount // aim goes up to increase depth in ocean
		case "up":
			aim -= command.Amount // aim goes down to decrease depth in ocean
		default:
			panic("no support for direction " + command.Direction)
		}
	}

	return horizontalPosition, depth
}

func main() {
	flag.Parse()

	lines, err := lines(*inputFile)
	if err != nil {
		panic(err)
	}
	commands, err := parseCommands(lines)
	if err != nil {
		panic(err)
	}

	horizontalPosition, depth := processCommands(commands)
	fmt.Printf("Horizontal Position %d * Depth %d = %d\n", horizontalPosition, depth, horizontalPosition*depth)

	horizontalPosition, depth = processCommandsWithAim(commands)
	fmt.Printf("(With Aim) Horizontal Position %d * Depth %d = %d\n", horizontalPosition, depth, horizontalPosition*depth)
}
