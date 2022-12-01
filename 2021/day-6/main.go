package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
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

func parseAges(line string) (ages []uint8, err error) {
	parts := strings.Split(line, ",")
	for _, p := range parts {
		age, err := strconv.Atoi(p)
		if err != nil {
			return ages, err
		}
		ages = append(ages, uint8(age))
	}

	return ages, nil
}

func evaluateOneDay(ages []uint8) (newAges []uint8) {
	var wg sync.WaitGroup
	workerCount := 15

	var newFishCount uint64
	now := time.Now()

	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		w := w
		go func() {
			defer wg.Done()
			fmt.Println("Starting at ", w)
			var nfc uint64
			for i := w; i < len(ages); i += workerCount {
				a := ages[i]
				newAge := a - 1
				if a == 0 {
					newAge = 6
					nfc++
				}

				ages[i] = newAge
			}
			atomic.AddUint64(&newFishCount, nfc)
		}()
	}
	wg.Wait()

	newFish := []uint8{}
	for i := 0; i < int(newFishCount); i++ {
		newFish = append(newFish, 8)
	}

	// for i := range ages {
	// 	newAge := ages[i] - 1
	// 	if newAge < 0 {
	// 		newAge = 6
	// 		newFish = append(newFish, 8)
	// 	}
	//
	// 	ages[i] = newAge
	// }
	postFishEval := time.Now()
	fmt.Printf("fish age eval: %s\n", postFishEval.Sub(now))

	allFishAges := append(ages, newFish...)
	fmt.Printf("append: %s\n", time.Now().Sub(postFishEval))
	return allFishAges
}

func fishCountAfterDay(day int, ages []uint8) int {
	for i := 0; i < day; i++ {
		fmt.Printf("Day %d\n", i+1)
		ages = evaluateOneDay(ages)
	}
	return len(ages)
}

func main() {
	flag.Parse()

	lines, err := lines(*inputFile)
	if err != nil {
		fmt.Printf("failed to get lines for input file: %s\n", err)
		os.Exit(1)
	}

	if len(lines) != 1 {
		fmt.Printf("expect only on line in input file: %s\n", err)
		os.Exit(1)
	}

	ages, err := parseAges(lines[0])
	if err != nil {
		fmt.Printf("failed to parse ages: %s\n", err)
		os.Exit(1)
	}

	// fmt.Printf("Fish Count (After Day 18): %d\n", fishCountAfterDay(18, ages))
	// fmt.Printf("Fish Count (After Day 80): %d\n", fishCountAfterDay(80, ages))
	fmt.Printf("Fish Count (After Day 256): %d\n", fishCountAfterDay(256, ages))
}
