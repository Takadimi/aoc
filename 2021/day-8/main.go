package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Takadimi/aoc/2021/file"
)

var inputFileFlag = flag.String("inputFile", "sample.txt", "Relative file path to use as input.")

var uniqueDigitCount = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

func main() {
	flag.Parse()
	inputFile := *inputFileFlag

	lines, err := file.Lines(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum := 0
	for _, l := range lines {
		sum += instancesOfUniqueDigits(l)
	}

	fmt.Println("Part one:", sum)
}

func instancesOfUniqueDigits(l string) int {
	count := 0
	parts := strings.Split(l, "|")
	fields := strings.Fields(parts[1])
	for _, f := range fields {
		if _, isUnique := uniqueDigitCount[len(f)]; isUnique {
			count++
		}
	}
	return count
}
