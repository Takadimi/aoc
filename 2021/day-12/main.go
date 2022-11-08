package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type cave struct {
	name    string
	toCaves []*cave
	isBig   bool
}

type path struct {
	caves []*cave
}

func (p *path) hasVisited(c *cave) bool {
	for _, cave := range p.caves {
		if cave == c {
			return true
		}
	}
	return false
}

func (p *path) visitCount(c *cave) int {
	visits := 0
	for _, cave := range p.caves {
		if cave == c {
			visits++
		}
	}
	return visits
}

func (p *path) singleSmallCaveVisitedTwice() bool {
	visits := map[string]int{}
	for _, cave := range p.caves {
		if !cave.isBig {
			v, _ := visits[cave.name]
			v++
			visits[cave.name] = v
		}
	}
	for _, v := range visits {
		if v > 1 {
			return true
		}
	}
	return false
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fields := strings.Fields(string(data))

	caveMap := make(map[string]*cave)

	for _, field := range fields {
		parts := strings.Split(field, "-")
		caveAName := parts[0]
		caveBName := parts[1]

		caveA, hasCaveA := caveMap[caveAName]
		if !hasCaveA {
			caveA = new(cave)
			caveA.name = caveAName
			caveA.toCaves = make([]*cave, 0)
			if caveAName == strings.ToUpper(caveAName) {
				caveA.isBig = true
			}
		}
		caveB, hasCaveB := caveMap[caveBName]
		if !hasCaveB {
			caveB = new(cave)
			caveB.name = caveBName
			caveB.toCaves = make([]*cave, 0)
			if caveBName == strings.ToUpper(caveBName) {
				caveB.isBig = true
			}
		}
		caveMap[caveAName] = caveA
		caveMap[caveBName] = caveB

		caveA.toCaves = append(caveA.toCaves, caveB)
		caveB.toCaves = append(caveB.toCaves, caveA)
	}

	// printMap(caveMap)
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")

	currentCave := caveMap["start"]
	path := path{caves: []*cave{}}
	visit(currentCave, path)
}

func visit(c *cave, path path) {
	path.caves = append(path.caves, c)

	if c.name == "end" {
		printPath(path)
		return
	}
	if len(c.toCaves) == 0 {
		return
	}

	var endCave *cave
	for _, cave := range c.toCaves {
		if cave.name == "start" {
			continue
		}
		if cave.name == "end" {
			endCave = cave
			continue
		}
		if !cave.isBig && path.singleSmallCaveVisitedTwice() && path.visitCount(cave) == 1 {
			continue
		}
		if !cave.isBig && path.visitCount(cave) > 1 {
			continue
		}
		visit(cave, path)
	}
	if endCave != nil {
		visit(endCave, path)
	}
}

func printPath(path path) {
	for i, c := range path.caves {
		if i < len(path.caves)-1 {
			fmt.Print(c.name + ",")
		} else {
			fmt.Println(c.name)
		}
	}
}

func printMap(m map[string]*cave) {
	for _, c := range m {
		fmt.Printf("%s - (", c.name)
		for _, tc := range c.toCaves {
			fmt.Printf("%s, ", tc.name)
		}
		fmt.Println(")")
	}
}
