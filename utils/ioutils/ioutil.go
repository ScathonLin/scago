package ioutils

import (
	"bufio"
	"errors"
	"os"
)

func ReadFileToLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("failed to open file : " + filePath)
	}
	reader := bufio.NewReader(file)
	lines := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}
