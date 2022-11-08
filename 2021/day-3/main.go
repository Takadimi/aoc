package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

func parseLines(lines []string) ([]int64, int, error) {
	ints := []int64{}
	for _, l := range lines {
		asInt, err := strconv.ParseInt(l, 2, 64)
		if err != nil {
			return ints, 0, err
		}
		ints = append(ints, asInt)
	}

	bitWidth := len(lines[0])

	return ints, bitWidth, nil
}

func pow(base int64, exponent int) int64 {
	var result int64 = 0
	var i int64 = 0
	for ; i < base; i++ {
		result = result * int64(exponent)
	}
	return result
}

func calculateGammaRate(numbers []int64, bitWidth int) int64 {
	var gammaRate int64 = 0
	for i := bitWidth - 1; i >= 0; i-- {
		zeroBitCount := 0
		oneBitCount := 0
		for _, n := range numbers {
			shifted := n >> i
			if shifted&1 == 1 {
				oneBitCount++
			} else {
				zeroBitCount++
			}
		}

		if oneBitCount > zeroBitCount {
			gammaRate = gammaRate | (1 << i)
		}
	}

	return gammaRate
}

func calculateEpsilonRate(gammaRate int64, bitWidth int) int64 {
	var mask int64 = 0
	for i := 0; i < bitWidth; i++ {
		mask |= 1 << i
	}

	return gammaRate ^ mask
}

func calculateOxygenGeneratorRating(numbers []int64, bitWidth int) int64 {
	filteredNumbers := numbers
	for i := bitWidth - 1; i >= 0; i-- {
		numbersWithZeroInBitPosition := []int64{}
		numbersWithOneInBitPosition := []int64{}
		for _, n := range filteredNumbers {
			shifted := n >> i
			if shifted&1 == 1 {
				numbersWithOneInBitPosition = append(numbersWithOneInBitPosition, n)
			} else {
				numbersWithZeroInBitPosition = append(numbersWithZeroInBitPosition, n)
			}
		}

		if len(numbersWithOneInBitPosition) >= len(numbersWithZeroInBitPosition) {
			filteredNumbers = numbersWithOneInBitPosition
		} else {
			filteredNumbers = numbersWithZeroInBitPosition
		}

		if len(filteredNumbers) == 1 {
			return filteredNumbers[0]
		}
	}

	return -1
}

func calculateCarbonDioxideScrubberRating(numbers []int64, bitWidth int) int64 {
	filteredNumbers := numbers
	for i := bitWidth - 1; i >= 0; i-- {
		numbersWithZeroInBitPosition := []int64{}
		numbersWithOneInBitPosition := []int64{}
		for _, n := range filteredNumbers {
			shifted := n >> i
			if shifted&1 == 1 {
				numbersWithOneInBitPosition = append(numbersWithOneInBitPosition, n)
			} else {
				numbersWithZeroInBitPosition = append(numbersWithZeroInBitPosition, n)
			}
		}

		if len(numbersWithZeroInBitPosition) <= len(numbersWithOneInBitPosition) {
			filteredNumbers = numbersWithZeroInBitPosition
		} else {
			filteredNumbers = numbersWithOneInBitPosition
		}

		if len(filteredNumbers) == 1 {
			return filteredNumbers[0]
		}
	}

	return -1
}

func main() {
	flag.Parse()

	lines, err := lines(*inputFile)
	if err != nil {
		panic(err)
	}

	numbers, bitWidth, err := parseLines(lines)
	if err != nil {
		panic(err)
	}

	gammaRate := calculateGammaRate(numbers, bitWidth)
	fmt.Printf("%d (%012b)\n", gammaRate, gammaRate)
	epsilonRate := calculateEpsilonRate(gammaRate, bitWidth)
	fmt.Printf("%d (%012b)\n", epsilonRate, epsilonRate)

	fmt.Printf("Power consumption = %d * %d = %d\n", gammaRate, epsilonRate, gammaRate*epsilonRate)

	oxygenGeneratorRating := calculateOxygenGeneratorRating(numbers, bitWidth)
	fmt.Printf("%d (%012b)\n", oxygenGeneratorRating, oxygenGeneratorRating)

	carbonDioxideScrubberRating := calculateCarbonDioxideScrubberRating(numbers, bitWidth)
	fmt.Printf("%d (%012b)\n", carbonDioxideScrubberRating, carbonDioxideScrubberRating)

	fmt.Printf("Life support rating = %d * %d = %d\n", oxygenGeneratorRating, carbonDioxideScrubberRating, oxygenGeneratorRating*carbonDioxideScrubberRating)
}
