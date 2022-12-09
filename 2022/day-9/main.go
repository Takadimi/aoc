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

	headMotionSeries, err := parseHeadMotionSeries(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	headPosition := Position{0, 0}
	tailPosition := Position{0, 0}
	visitedTailPositions := map[Position]bool{
		tailPosition: true,
	}

	for _, motion := range headMotionSeries {
		for step := 0; step < motion.Steps; step++ {
			switch motion.Dir {
			case Direction_Up:
				headPosition.Y++
			case Direction_Down:
				headPosition.Y--
			case Direction_Left:
				headPosition.X--
			case Direction_Right:
				headPosition.X++
			default:
				panic("unknown direction")
			}

			tpxo := headPosition.X - tailPosition.X
			tpyo := headPosition.Y - tailPosition.Y

			if (tpxo > 0 && tpyo > 1) || (tpxo > 1 && tpyo > 0) {
				// NE
				tailPosition.X++
				tailPosition.Y++
			} else if (tpxo < 0 && tpyo < -1) || (tpxo < -1 && tpyo < 0) {
				// SW
				tailPosition.X--
				tailPosition.Y--
			} else if (tpxo > 0 && tpyo < -1) || (tpxo > 1 && tpyo < 0) {
				// SE
				tailPosition.X++
				tailPosition.Y--
			} else if (tpxo < 0 && tpyo > 1) || (tpxo < -1 && tpyo > 0) {
				// NW
				tailPosition.X--
				tailPosition.Y++
			} else if tpxo > 1 {
				// E
				tailPosition.X++
			} else if tpxo < -1 {
				// W
				tailPosition.X--
			} else if tpyo > 1 {
				// N
				tailPosition.Y++
			} else if tpyo < -1 {
				tailPosition.Y--
			}

			visitedTailPositions[tailPosition] = true
		}
	}

	fmt.Println(len(visitedTailPositions))
}

type Position struct {
	X, Y int
}

type Direction int

const (
	Direction_Unknown Direction = iota
	Direction_Up
	Direction_Down
	Direction_Left
	Direction_Right
)

type Motion struct {
	Dir   Direction
	Steps int
}

func parseHeadMotionSeries(lines []string) ([]Motion, error) {
	motionSeries := []Motion{}

	for _, l := range lines {
		fields := strings.Fields(l)
		if len(fields) != 2 {
			return nil, errors.New("expected 2 fields per line for motion series")
		}
		var dir Direction
		switch fields[0] {
		case "U":
			dir = Direction_Up
		case "D":
			dir = Direction_Down
		case "L":
			dir = Direction_Left
		case "R":
			dir = Direction_Right
		default:
			return nil, errors.New("no matching direction identifier")
		}

		steps, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}

		motionSeries = append(motionSeries, Motion{Dir: dir, Steps: steps})
	}

	return motionSeries, nil
}
