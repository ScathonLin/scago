package archive

import (
	"fmt"
	"testing"
)

const testTarGzFile string = "/Users/scathon/projects/golang/tmp/test.tar.gz"
const testTarGzUnComPath string = "/Users/scathon/projects/golang/tmp/untargz"
const testZipFile string = "/Users/scathon/projects/golang/tmp/test.zip"
const testZipUnComPath string = "/Users/scathon/projects/golang/tmp/unzip"

func TestUncompressTarGzFile(t *testing.T) {
	_, err := UncompressTarGzFile(testTarGzFile, testTarGzUnComPath)
	if err != nil {
		fmt.Errorf("Failed to uncompress tar.gz file\n")
	}
	fmt.Println("Done!")
}

func TestUncompressZipFile(t *testing.T) {
	_, err := UncompressZipFile(testZipFile, testZipUnComPath)
	if err != nil {
		fmt.Errorf("Failed to uncompress .zip file\n")
	}
	fmt.Println("Done!")
}
