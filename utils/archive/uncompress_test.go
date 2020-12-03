package archive

import (
	"fmt"
	"os"
	"testing"
)

const (
	testTarGzFile      string = "/Users/scathon/projects/golang/tmp/test.tar.gz"
	testTarGzUnComPath string = "/Users/scathon/projects/golang/tmp/untargz"
	testZipFile        string = "/Users/scathon/projects/golang/tmp/test.zip"
	testZipUnComPath   string = "/Users/scathon/projects/golang/tmp/unzip"
	testGzipFile       string = "/Users/scathon/projects/golang/tmp/gzip_input.txt.gz"
	testUnGipFilePath  string = "/Users/scathon/projects/golang/tmp/gunzip.txt"
)

func TestUncompressTarGzFile(t *testing.T) {
	err := UncompressTarGzFile(testTarGzFile, testTarGzUnComPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Done!")
}

func TestUncompressZipFile(t *testing.T) {
	err := UncompressZipFile(testZipFile, testZipUnComPath, 10000, 10<<20)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Done!")
}

func TestUncompressGzipFile(t *testing.T) {
	err := UncompressGzipFile(testGzipFile, testUnGipFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file, _ := os.Open(testUnGipFilePath)
	defer func() { _ = file.Close() }()
	fmt.Println("Done!")
}
