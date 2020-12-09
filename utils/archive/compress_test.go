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
		"D:\\devtools\\GO\\GOPATH\\src\\scago\\utils\\archive\\compress.go",
		"D:\\devtools\\GO\\GOPATH\\src\\scago\\utils\\archive\\compress_test.go",
		"D:\\devtools\\GO\\GOPATH\\src\\scago\\utils\\archive\\uncompress.go",
		"D:\\devtools\\GO\\GOPATH\\src\\scago\\utils\\archive\\uncompress_test.go",
	}
	tarFilePath := "D:\\devtools\\GO\\GOPATH\\src\\scago\\utils\\archive\\test.tar"
	if e := TarFiles(filePaths, tarFilePath); e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("Done!")
}
