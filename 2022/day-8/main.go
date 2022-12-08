package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

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

	treeMap, err := parseTreeMap(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part one:", sumOfVisibleTrees(treeMap))
	fmt.Println("Part two:", highestScenicScore(treeMap))
}

func sumOfVisibleTrees(treeMap [][]int) int {
	// initialize sum with number of edge trees since they're always visible
	sum := ((len(treeMap) - 1) * 2) + ((len(treeMap[0]) - 1) * 2)

	// evaluate interior trees
	for y := 1; y < len(treeMap)-1; y++ {
		for x := 1; x < len(treeMap[y])-1; x++ {
			tree := treeMap[y][x]

			// North
			visibleNorth := true
			for yy := 0; yy < y; yy++ {
				if treeMap[yy][x] >= tree {
					visibleNorth = false
					break
				}
			}
			if visibleNorth {
				sum++
				continue
			}

			// South
			visibleSouth := true
			for yy := y + 1; yy < len(treeMap); yy++ {
				if treeMap[yy][x] >= tree {
					visibleSouth = false
					break
				}
			}
			if visibleSouth {
				sum++
				continue
			}

			// East
			visibleEast := true
			for xx := x + 1; xx < len(treeMap[y]); xx++ {
				if treeMap[y][xx] >= tree {
					visibleEast = false
					break
				}
			}
			if visibleEast {
				sum++
				continue
			}

			// West
			visibleWest := true
			for xx := 0; xx < x; xx++ {
				if treeMap[y][xx] >= tree {
					visibleWest = false
					break
				}
			}
			if visibleWest {
				sum++
				continue
			}
		}
	}

	return sum
}

func highestScenicScore(treeMap [][]int) int {
	highestScenicScore := 0

	for y := 0; y < len(treeMap); y++ {
		for x := 0; x < len(treeMap[y]); x++ {
			tree := treeMap[y][x]

			// North
			northScore := 0
			for yy := y - 1; yy >= 0; yy-- {
				northScore++
				if treeMap[yy][x] >= tree {
					break
				}
			}

			// South
			southScore := 0
			for yy := y + 1; yy < len(treeMap); yy++ {
				southScore++
				if treeMap[yy][x] >= tree {
					break
				}
			}

			// East
			eastScore := 0
			for xx := x + 1; xx < len(treeMap[y]); xx++ {
				eastScore++
				if treeMap[y][xx] >= tree {
					break
				}
			}

			// West
			westScore := 0
			for xx := x - 1; xx >= 0; xx-- {
				westScore++
				if treeMap[y][xx] >= tree {
					break
				}
			}

			totalScore := northScore * southScore * eastScore * westScore
			if totalScore > highestScenicScore {
				highestScenicScore = totalScore
			}
		}
	}

	return highestScenicScore
}

func parseTreeMap(lines []string) ([][]int, error) {
	treeMap := [][]int{}
	for y := 0; y < len(lines); y++ {
		treeMap = append(treeMap, []int{})
		for x := 0; x < len(lines[y]); x++ {
			v, err := strconv.Atoi(string(lines[y][x]))
			if err != nil {
				return nil, err
			}
			treeMap[y] = append(treeMap[y], v)
		}
	}

	return treeMap, nil
}
