package ioutils

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFileToLines(t *testing.T) {
	runtimeDir, _ := os.Getwd()
	path := runtimeDir + string(os.PathSeparator) + "ioutil_test.go"
	lines, err := ReadFileToLines(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
