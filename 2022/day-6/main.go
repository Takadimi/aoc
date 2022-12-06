package main

import (
	"flag"
	"fmt"
	"os"

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

	for i, l := range lines {
		fmt.Printf("Part one (line %d): %d\n", i+1, partOne(l))
		fmt.Printf("Part two (line %d): %d\n", i+1, partTwo(l))
	}
}

func partOne(line string) int {
	return indexAfterNUniqueCharacters(line, 4)
}

func partTwo(line string) int {
	return indexAfterNUniqueCharacters(line, 14)
}

func indexAfterNUniqueCharacters(line string, n int) int {
	for i := 0; i < len(line)-n-1; i++ {
		occurenceMap := map[rune]int{}
		segment := line[i : i+n]
		for _, char := range segment {
			occurenceMap[char] = occurenceMap[char] + 1
		}

		hasDuplicates := false
		for _, occurences := range occurenceMap {
			if occurences > 1 {
				hasDuplicates = true
				break
			}
		}
		if !hasDuplicates {
			return i + n
		}
	}

	return 0
}
