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
