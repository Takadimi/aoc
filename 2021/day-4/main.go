package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
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

func splitByEmptyLine(lines []string) [][]string {
	byEmptyLine := [][]string{}
	currentLine := []string{}
	for _, l := range lines {
		if len(l) == 0 {
			byEmptyLine = append(byEmptyLine, currentLine)
			currentLine = []string{}
			continue
		}
		currentLine = append(currentLine, l)
	}
	byEmptyLine = append(byEmptyLine, currentLine)

	return byEmptyLine
}

func numbersFromFields(fields []string) ([]int, error) {
	numbers := []int{}
	for _, ns := range fields {
		n, err := strconv.Atoi(ns)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}

	return numbers, nil
}

func parseNumbersToDraw(section []string) ([]int, error) {
	if len(section) != 1 {
		return nil, errors.New("invalid lines submitted")
	}
	numbersLine := section[0]
	fields := strings.Split(numbersLine, ",")

	return numbersFromFields(fields)
}

func parseBoards(sections [][]string) ([]board, error) {
	currentBoard := board{
		Squares: make([][]square, 0),
	}

	boards := []board{}
	for _, section := range sections {
		currentBoard = board{
			Squares: make([][]square, 0),
		}

		for y, line := range section {
			lineNumbers, err := parseBoardLineNumbers(line)
			if err != nil {
				return nil, err
			}

			row := []square{}
			for x, n := range lineNumbers {
				row = append(row, square{
					Number:   n,
					Position: position{X: x, Y: y},
					IsMarked: false,
				})
			}

			currentBoard.Squares = append(currentBoard.Squares, row)
		}
		boards = append(boards, currentBoard)
	}

	return boards, nil
}

func parseBoardLineNumbers(line string) ([]int, error) {
	return numbersFromFields(strings.Fields(line))
}

type board struct {
	Squares          [][]square
	LastMarkedSquare *square
}

type square struct {
	Number   int
	Position position
	IsMarked bool
}

type position struct {
	X, Y int
}

func (b *board) Check(n int) bool {
	var marked *square
	for y := 0; y < len(b.Squares); y++ {
		line := b.Squares[y]
		for x := 0; x < len(line); x++ {
			square := &line[x]
			if square.Number == n {
				square.IsMarked = true
				marked = square
				b.LastMarkedSquare = square
			}
		}
	}

	if marked != nil {
		line := b.Squares[marked.Position.Y]
		for x := 0; x < len(line); x++ {
			if !line[x].IsMarked {
				break
			}
			if x == len(line)-1 {
				return true
			}
		}

		for y := 0; y < len(b.Squares); y++ {
			if !b.Squares[y][marked.Position.X].IsMarked {
				break
			}
			if y == len(b.Squares)-1 {
				return true
			}
		}
	}

	return false
}

func (b *board) Score() int {
	unmarkedSum := 0
	for y := 0; y < len(b.Squares); y++ {
		line := b.Squares[y]
		for x := 0; x < len(line); x++ {
			square := &line[x]
			if !square.IsMarked {
				unmarkedSum += square.Number
			}
		}
	}

	return unmarkedSum * b.LastMarkedSquare.Number
}

func (b *board) Marked() []square {
	marked := []square{}
	for _, row := range b.Squares {
		for _, square := range row {
			if square.IsMarked {
				marked = append(marked, square)
			}
		}
	}
	return marked
}

func (b *board) String() string {
	sb := strings.Builder{}

	s := ""
	for _, line := range b.Squares {
		lineNumbers := []string{}
		for _, n := range line {
			markedText := " "
			if n.IsMarked {
				markedText = "X"
			}
			lineNumbers = append(lineNumbers, fmt.Sprintf("%s %d (%d,%d)", markedText, n.Number, n.Position.X, n.Position.Y))
		}
		s += strings.Join(lineNumbers, "\t|\t") + "\n"
	}

	w := tabwriter.NewWriter(&sb, 1, 0, 1, ' ', 0)
	fmt.Fprint(w, s)
	w.Flush()

	return sb.String()
}

func checkBoards(boards []board, number int) ([]board, []board) {
	winners := []board{}
	remaining := []board{}

	for i := range boards {
		var b *board = &boards[i]
		if b.Check(number) {
			winners = append(winners, *b)
		} else {
			remaining = append(remaining, *b)
		}
	}

	return winners, remaining
}

func main() {
	flag.Parse()

	lines, err := lines(*inputFile)
	if err != nil {
		panic(err)
	}

	sections := splitByEmptyLine(lines)
	if len(sections) < 2 {
		fmt.Println("Expected one section of numbers to draw and at least one section of a board")
		os.Exit(1)
	}

	numbersToDraw, err := parseNumbersToDraw(sections[0])
	if err != nil {
		panic(err)
	}

	boards, err := parseBoards(sections[1:])
	if err != nil {
		panic(err)
	}

	var firstWinner *board
	var lastWinner *board
	for _, n := range numbersToDraw {
		winners, remaining := checkBoards(boards, n)
		if len(winners) > 0 {
			if firstWinner == nil {
				firstWinner = &winners[0]
			}
			if len(remaining) == 0 {
				lastWinner = &winners[0]
			}
		}
		boards = remaining
	}

	fmt.Printf("First Winner (Score: %d):\n%s\n", firstWinner.Score(), firstWinner)
	fmt.Printf("Last Winner (Score: %d):\n%s\n", lastWinner.Score(), lastWinner)
}
