package archive

import (
	"fmt"
	"testing"
)

const testTarGzFile string = "/Users/scathon/projects/golang/tmp/test.tar.gz"
const testTarGzUnComPath string = "/Users/scathon/projects/golang/tmp/untargz"

func TestUncompressTarGzFile(t *testing.T) {
	_, err := UncompressTarGzFile(testTarGzFile, testTarGzUnComPath)
	if err != nil {
		fmt.Errorf("Failed to uncompress tar.gz file\n")
	}
	fmt.Println("Done!")
}
