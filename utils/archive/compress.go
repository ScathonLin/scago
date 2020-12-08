// author: linhuadong(scathonlin)
// creatdate: 2020/12/06
package archive

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"strings"
)

const emptyRtn string = ""

func ZipFiles(filePaths []string, zipFilePath string) error {
	if filePaths == nil || len(filePaths) == 0 {
		return errors.New("no file to zip")
	}

	zipFile, err := os.Create(zipFilePath)
	defer func() { _ = zipFile.Close() }()
	if err != nil {
		return err
	}
	zipWriter := zip.NewWriter(zipFile)
	defer func() { _ = zipWriter.Close() }()
	for _, filePath := range filePaths {
		// parse and get file name.
		entryPath := filePath[strings.LastIndex(filePath, string(os.PathSeparator))+1:]
		// open file.
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		// create entry file handle
		entry, err := zipWriter.Create(entryPath)
		if err != nil {
			return err
		}
		// write file to zip entry.
		if err = transformFileTo(entry, file); err != nil {
			return err
		}
	}
	return nil
}

func Tar(filePath string) (string, error) {
	return "", nil
}

func Gzip(filePath string) (string, error) {
	return "", nil
}

func transformFileTo(writer io.Writer, file *os.File) error {
	defer func() { _ = file.Close() }()
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	return nil
}
