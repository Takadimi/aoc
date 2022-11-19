package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "sample1.txt", "Relative file path to use as input.")

func textLines(fileName string) ([]string, error) {
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

type line struct {
	A, B point
}

type point struct {
	X, Y int
}

func parsePoint(pointText string) (p point, err error) {
	parts := strings.Split(pointText, ",")
	if len(parts) != 2 {
		return p, fmt.Errorf("invalid coordinate")
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return p, fmt.Errorf("failed to parse x")
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return p, fmt.Errorf("failed to parse y")
	}

	return point{x, y}, nil
}

func parseLines(textLines []string) (lines []line, err error) {
	for i, tl := range textLines {
		fields := strings.Fields(tl)
		if len(fields) != 3 {
			return lines, fmt.Errorf("invalid format on line %d", i)
		}

		a, err := parsePoint(fields[0])
		if err != nil {
			return lines, fmt.Errorf("failed to parse first point on line %d: %w", i, err)
		}
		b, err := parsePoint(fields[2])
		if err != nil {
			return lines, fmt.Errorf("failed to parse second point on line %d: %w", i, err)
		}

		lines = append(lines, line{a, b})
	}

	return lines, nil
}

func filterOnlyStraightLines(lines []line) (filtered []line) {
	for _, l := range lines {
		if l.A.X == l.B.X || l.A.Y == l.B.Y {
			filtered = append(filtered, l)
		}
	}

	return filtered
}

func extrapolatePoints(l line) (points []point) {
	// 0,9 -> 5,9
	// ----------
	// 0,9
	// 1,9
	// 2,9
	// 3,9
	// 4,9
	// 5,9

	cp := l.A
	for {
		points = append(points, cp)

		if cp.X == l.B.X && cp.Y == l.B.Y {
			break
		}

		if cp.X < l.B.X {
			cp.X++
		} else if cp.X > l.B.X {
			cp.X--
		}

		if cp.Y < l.B.Y {
			cp.Y++
		} else if cp.Y > l.B.Y {
			cp.Y--
		}
	}

	return points
}

func visitedPoints(lines []line) map[point]int {
	visited := make(map[point]int, 0)
	for _, l := range lines {
		ep := extrapolatePoints(l)
		for _, p := range ep {
			vp, _ := visited[p]
			vp++
			visited[p] = vp
		}
	}
	return visited
}

func countOfPointsVisitedMultipleTimes(lines []line) int {
	visited := visitedPoints(lines)

	count := 0
	for _, c := range visited {
		if c > 1 {
			count++
		}
	}
	return count
}

func main() {
	flag.Parse()

	l, err := textLines(*inputFile)
	if err != nil {
		fmt.Println("failed to extract lines from input")
		os.Exit(1)
	}

	lines, err := parseLines(l)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Count (Straight Lines Only): %d\n", countOfPointsVisitedMultipleTimes(filterOnlyStraightLines(lines)))
	fmt.Printf("Count (Straight & Diagonal Lines): %d\n", countOfPointsVisitedMultipleTimes(lines))
}
