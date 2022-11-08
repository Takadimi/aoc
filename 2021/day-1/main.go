package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "sample1.txt", "Relative file path to use as input.")

func parseMeasurements(inputFile string) []int64 {
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	fields := strings.Fields(string(b))

	measurements := []int64{}
	for _, field := range fields {
		measurement, err := strconv.ParseInt(field, 10, 64)
		if err != nil {
			panic(err)
		}
		measurements = append(measurements, measurement)
	}

	return measurements
}

func simpleMeasurementIncreaseCount(inputFile string) int {
	measurements := parseMeasurements(inputFile)
	measurementIncreaseCount := 0

	for i, measurement := range measurements {
		if i == 0 {
			continue
		}
		previousMeasurement := measurements[i-1]
		if measurement > previousMeasurement {
			measurementIncreaseCount++
		}
	}

	return measurementIncreaseCount
}

func measurementWindowSum(window []int64) int64 {
	var sum int64 = 0
	for _, measurement := range window {
		sum += measurement
	}
	return sum
}

func slidingWindowMeasurementIncreaseCount(inputFile string) int {
	measurements := parseMeasurements(inputFile)
	measurementIncreaseCount := 0

	for i := 3; i < len(measurements); i++ {
		previousWindow := measurements[i-3 : i]
		currentWindow := measurements[i-2 : i+1]

		if measurementWindowSum(currentWindow) > measurementWindowSum(previousWindow) {
			measurementIncreaseCount++
		}
	}

	return measurementIncreaseCount
}

func main() {
	flag.Parse()

	fmt.Println(simpleMeasurementIncreaseCount(*inputFile))
	fmt.Println(slidingWindowMeasurementIncreaseCount(*inputFile))
}
