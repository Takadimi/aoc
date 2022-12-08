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

			checkVisibility := func(wTree int) bool {
				return wTree < tree
			}
			visibleNorth := walkTreeMapNorth(treeMap, x, y, checkVisibility)
			visibleSouth := walkTreeMapSouth(treeMap, x, y, checkVisibility)
			visibleEast := walkTreeMapEast(treeMap, x, y, checkVisibility)
			visibleWest := walkTreeMapWest(treeMap, x, y, checkVisibility)

			if visibleNorth || visibleSouth || visibleEast || visibleWest {
				sum++
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

			northScore := 0
			southScore := 0
			eastScore := 0
			westScore := 0

			walkTreeMapNorth(treeMap, x, y, func(wTree int) bool {
				northScore++
				return wTree < tree
			})
			walkTreeMapSouth(treeMap, x, y, func(wTree int) bool {
				southScore++
				return wTree < tree
			})
			walkTreeMapEast(treeMap, x, y, func(wTree int) bool {
				eastScore++
				return wTree < tree
			})
			walkTreeMapWest(treeMap, x, y, func(wTree int) bool {
				westScore++
				return wTree < tree
			})

			totalScore := northScore * southScore * eastScore * westScore
			if totalScore > highestScenicScore {
				highestScenicScore = totalScore
			}
		}
	}

	return highestScenicScore
}

type treeWalkFunc func(int) bool

func walkTreeMapNorth(treeMap [][]int, startingX, startingY int, walkFunc treeWalkFunc) bool {
	for cy := startingY - 1; cy >= 0; cy-- {
		if keepWalking := walkFunc(treeMap[cy][startingX]); !keepWalking {
			return false
		}
	}
	return true
}

func walkTreeMapSouth(treeMap [][]int, startingX, startingY int, walkFunc treeWalkFunc) bool {
	for cy := startingY + 1; cy < len(treeMap); cy++ {
		if keepWalking := walkFunc(treeMap[cy][startingX]); !keepWalking {
			return false
		}
	}
	return true
}

func walkTreeMapEast(treeMap [][]int, startingX, startingY int, walkFunc treeWalkFunc) bool {
	for cx := startingX + 1; cx < len(treeMap[startingY]); cx++ {
		if keepWalking := walkFunc(treeMap[startingY][cx]); !keepWalking {
			return false
		}
	}
	return true
}

func walkTreeMapWest(treeMap [][]int, startingX, startingY int, walkFunc treeWalkFunc) bool {
	for cx := startingX - 1; cx >= 0; cx-- {
		if keepWalking := walkFunc(treeMap[startingY][cx]); !keepWalking {
			return false
		}
	}
	return true
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
