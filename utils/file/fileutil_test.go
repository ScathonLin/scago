package file

import (
	"fmt"
	"testing"
)

func TestListFiles(t *testing.T) {
	filePath := "/Users/scathon/coding/golang/GOPATH/src/scago/utils"
	files, err := ListFiles(filePath, false)
	if err != nil {
		return
	}
	fmt.Println(files)
	files, err = ListFiles(filePath, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, file := range files {
		fmt.Println(file)
	}

}
