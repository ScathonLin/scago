// author: linhuadong(scathonlin)
// creatdate: 2020/12/06
package archive

import (
	"fmt"
	"testing"
)

func TestZipFiles(t *testing.T) {
	filePaths := []string{
		"/Users/scathon/coding/golang/GOPATH/src/scago/utils/archive/compress.go",
		"/Users/scathon/coding/golang/GOPATH/src/scago/utils/archive/compress_test.go",
		"/Users/scathon/coding/golang/GOPATH/src/scago/utils/archive/uncompress.go",
		"/Users/scathon/coding/golang/GOPATH/src/scago/utils/archive/uncompress_test.go",
		"/Users/scathon/coding/golang/GOPATH/src/scago/README.md",
	}
	zipFilePath := "/Users/scathon/tmp/test.zip"
	if err := ZipFiles(filePaths, zipFilePath); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Done")
}

func TestTar(t *testing.T) {
	filePaths := []string{
		"/Users/scathon/coding/golang/GOPATH/src/scago/utils/archive/compress.go",
		"/Users/scathon/coding/golang/GOPATH/src/scago/utils/archive/compress_test.go",
		"/Users/scathon/coding/golang/GOPATH/src/scago/utils/archive/uncompress.go",
		"/Users/scathon/coding/golang/GOPATH/src/scago/utils/archive/uncompress_test.go",
	}
	tarFilePath := "/Users/scathon/tmp/test.tar"
	if e := TarFiles(filePaths, tarFilePath); e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("Done!")
}

func TestZipDir(t *testing.T) {
	dirToZip := "/Users/scathon/coding/golang/GOPATH/src/scago"
	zipFilePath := "/Users/scathon/tmp/test1.zip"
	err := ZipDir(dirToZip, zipFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Done!")
}

func TestGzip(t *testing.T) {
	filePath := "/Users/scathon/tmp/test.tar"
	gzipFilePath := "/Users/scathon/tmp/test.tar.gz"
	err := Gzip(filePath, gzipFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Done!")
}

func TestTarDir(t *testing.T) {
	dirToTar := "/Users/scathon/coding/golang/GOPATH/src/scago"
	tarFilePath := "/Users/scathon/tmp/test2.tar"
	err := TarDir(dirToTar, tarFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Done!")
}
