package archive

import (
	"fmt"
	"os"
	"testing"
)

const (
	testTarGzFile      string = "/Users/scathon/projects/golang/tmp/test.tar.gz"
	testTarGzUnComPath string = "/Users/scathon/projects/golang/tmp/untargz"
	testZipFile        string = "/Users/scathon/tmp/test1.zip"
	testZipUnComPath   string = "/Users/scathon/tmp/unzipdir"
	testGzipFile       string = "/Users/scathon/projects/golang/tmp/gzip_input.txt.gz"
	testUnGipFilePath  string = "/Users/scathon/projects/golang/tmp/gunzip.txt"
	tarFilePath        string = "/Users/scathon/tmp/test2.tar"
	unTarDirPath       string = "/Users/scathon/tmp/test2-untar"
)

func TestUncompressTarGzFile(t *testing.T) {
	err := UnTarGZ(testTarGzFile, testTarGzUnComPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Done!")
}

func TestUncompressZipFile(t *testing.T) {
	err := UnZip(testZipFile, testZipUnComPath, 10000, 10<<20)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Done!")
}

func TestUncompressGzipFile(t *testing.T) {
	err := GUnzip(testGzipFile, testUnGipFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file, _ := os.Open(testUnGipFilePath)
	defer func() { _ = file.Close() }()
	fmt.Println("Done!")
}

func TestUnTar(t *testing.T) {
	err := UnTar(tarFilePath, unTarDirPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Done!")
}
