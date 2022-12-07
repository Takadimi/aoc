package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

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

	fileSizesByDir, err := parseFileSizesByDir(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part one:", partOne(fileSizesByDir))
	fmt.Println("Part two:", partTwo(fileSizesByDir))
}

func partOne(fileSizesByDir map[string]int) int {
	sum := 0
	for _, fileSize := range fileSizesByDir {
		if fileSize <= 100_000 {
			sum += fileSize
		}
	}

	return sum
}

func partTwo(fileSizesByDir map[string]int) int {
	totalDiskSpace := 70_000_000
	diskSpaceNeeded := 30_000_000
	diskSpaceUsed := fileSizesByDir["/"]
	additionalDiskSpaceNeeded := diskSpaceNeeded - (totalDiskSpace - diskSpaceUsed)

	fileSizeOfDirToDelete := diskSpaceUsed
	for _, fileSize := range fileSizesByDir {
		if fileSize >= additionalDiskSpaceNeeded && fileSize < fileSizeOfDirToDelete {
			fileSizeOfDirToDelete = fileSize
		}
	}

	return fileSizeOfDirToDelete
}

func parseFileSizesByDir(lines []string) (map[string]int, error) {
	fileSizesByPath := map[string]int{}

	currentPath := ""
	for _, l := range lines {
		fields := strings.Fields(l)

		if dir, isChangeDirCommand := matchChangeDirCommand(fields); isChangeDirCommand {
			currentPath = path.Clean(currentPath + "/" + dir)
			continue
		}

		if isListDirCommand := matchListDirCommand(fields); isListDirCommand {
			continue
		}

		if _, isListedDir := matchListedDir(fields); isListedDir {
			continue
		}

		if fileSize, fileName, isListedFile := matchListedFile(fields); isListedFile {
			fileSizesByPath[path.Clean(currentPath+"/"+fileName)] = fileSize
			continue
		}
	}

	fileSizesByDir := map[string]int{}
	for filePath, fileSize := range fileSizesByPath {
		dir, _ := path.Split(filePath)

		fileSizesByDir[path.Clean(dir)] += fileSize

		for dir != "/" {
			dir = path.Clean(dir + "/..")
			fileSizesByDir[dir] += fileSize
		}
	}

	return fileSizesByDir, nil
}

func matchChangeDirCommand(fields []string) (string, bool) {
	if len(fields) != 3 || fields[0] != "$" || fields[1] != "cd" {
		return "", false
	}

	return fields[2], true
}

func matchListDirCommand(fields []string) bool {
	if len(fields) != 2 || fields[0] != "$" || fields[1] != "ls" {
		return false
	}

	return true
}

func matchListedFile(fields []string) (int, string, bool) {
	if len(fields) != 2 {
		return 0, "", false
	}

	fileSize, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, "", false
	}

	return fileSize, fields[1], true
}

func matchListedDir(fields []string) (string, bool) {
	if len(fields) != 2 || fields[0] != "dir" {
		return "", false
	}

	return fields[1], true
}
