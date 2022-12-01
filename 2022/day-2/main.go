package main

import (
	"flag"
	"fmt"
)

var inputFileFlag = flag.String("inputFile", "sample.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	inputFile := *inputFileFlag

    fmt.Println(inputFile)
}
