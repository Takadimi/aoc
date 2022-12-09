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

	fmt.Println("Part one:", len(simulate(headMotionSeries, 1)[0].VisitedPositions))
	fmt.Println("Part two:", len(simulate(headMotionSeries, 10)[8].VisitedPositions))
}

type Tail struct {
	Position         Position
	VisitedPositions map[Position]bool
}

func simulate(headMotionSeries []Motion, tailCount int) []Tail {
	headPosition := Position{0, 0}

	tails := make([]Tail, tailCount)

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

			for i, tail := range tails {
				tpxo := headPosition.X - tail.Position.X
				tpyo := headPosition.Y - tail.Position.Y

				if i > 0 {
					tpxo = tails[i-1].Position.X - tail.Position.X
					tpyo = tails[i-1].Position.Y - tail.Position.Y
				}

				if (tpxo > 0 && tpyo > 1) || (tpxo > 1 && tpyo > 0) {
					// NE
					tail.Position.X++
					tail.Position.Y++
				} else if (tpxo < 0 && tpyo < -1) || (tpxo < -1 && tpyo < 0) {
					// SW
					tail.Position.X--
					tail.Position.Y--
				} else if (tpxo > 0 && tpyo < -1) || (tpxo > 1 && tpyo < 0) {
					// SE
					tail.Position.X++
					tail.Position.Y--
				} else if (tpxo < 0 && tpyo > 1) || (tpxo < -1 && tpyo > 0) {
					// NW
					tail.Position.X--
					tail.Position.Y++
				} else if tpxo > 1 {
					// E
					tail.Position.X++
				} else if tpxo < -1 {
					// W
					tail.Position.X--
				} else if tpyo > 1 {
					// N
					tail.Position.Y++
				} else if tpyo < -1 {
					tail.Position.Y--
				}

				if tail.VisitedPositions == nil {
					tail.VisitedPositions = make(map[Position]bool)
				}
				tail.VisitedPositions[tail.Position] = true
				tails[i] = tail
			}
		}
	}

	return tails
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
