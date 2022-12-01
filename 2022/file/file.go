package file

import (
	"bufio"
	"os"
)

func Lines(fileName string) ([]string, error) {
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
