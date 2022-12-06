package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Takadimi/aoc/2021/file"
)

var inputFileFlag = flag.String("inputFile", "sample.txt", "Relative file path to use as input.")
var debugFlag = flag.Bool("debug", false, "Output debug logs.")

func main() {
	flag.Parse()
	inputFile := *inputFileFlag

	lines, err := file.Lines(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	seafloorMap := parseMap(lines)

	fmt.Println("Part one:", partOne(seafloorMap))
}

type Position struct {
	X, Y int
}

type Move struct {
	From Position
	To   Position
}

const (
	EastboundCucumber  rune = '>'
	SouthboundCucumber rune = 'v'
	Empty              rune = '.'
)

func partOne(seafloorMap [][]rune) int {
	width := len(seafloorMap[0])
	height := len(seafloorMap)

	stepCount := 0

	printMap(seafloorMap)

	for {
		stepCount++

		eastboundMoves := []Move{}
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				currentSpot := seafloorMap[i][j]
				if currentSpot == EastboundCucumber {
					nextSpotXIndex := (j + 1) % width
					nextSpot := seafloorMap[i][nextSpotXIndex]
					if nextSpot == Empty {
						eastboundMoves = append(eastboundMoves, Move{
							From: Position{X: j, Y: i},
							To:   Position{X: nextSpotXIndex, Y: i},
						})
					}
				}
			}
		}

		for _, m := range eastboundMoves {
			seafloorMap[m.To.Y][m.To.X] = EastboundCucumber
			seafloorMap[m.From.Y][m.From.X] = Empty
		}

		southboundMoves := []Move{}
		for j := 0; j < width; j++ {
			for i := 0; i < height; i++ {
				currentSpot := seafloorMap[i][j]
				if currentSpot == SouthboundCucumber {
					nextSpotYIndex := (i + 1) % height
					nextSpot := seafloorMap[nextSpotYIndex][j]
					if nextSpot == Empty {
						southboundMoves = append(southboundMoves, Move{
							From: Position{X: j, Y: i},
							To:   Position{X: j, Y: nextSpotYIndex},
						})
					}
				}
			}
		}

		for _, m := range southboundMoves {
			seafloorMap[m.To.Y][m.To.X] = SouthboundCucumber
			seafloorMap[m.From.Y][m.From.X] = Empty
		}

		printMap(seafloorMap)

		if len(eastboundMoves) == 0 && len(southboundMoves) == 0 {
			return stepCount
		}
	}
}

func printMap(m [][]rune) {
	if !*debugFlag {
		return
	}
	fmt.Println("------------------------------")
	for _, row := range m {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
	fmt.Println("------------------------------")
}

func parseMap(lines []string) [][]rune {
	seafloorMap := newMap(len(lines[0]), len(lines))

	for i, l := range lines {
		for j, c := range l {
			seafloorMap[i][j] = c
		}
	}

	return seafloorMap
}

func newMap(width, height int) [][]rune {
	m := make([][]rune, height)
	for i := range m {
		m[i] = make([]rune, width)
	}
	return m
}
