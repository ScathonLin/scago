// author: linhuadong(scathonlin)
// creatdate: 2020/12/06
package archive

import (
	"archive/tar"
	"archive/zip"
	"errors"
	"io"
	"os"
	"scago/utils/file"
	"strings"
)

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
		fileItem, err := os.Open(filePath)
		if err != nil {
			return err
		}
		// create entry file handle
		entry, err := zipWriter.Create(entryPath)
		if err != nil {
			return err
		}
		// write file to zip entry.
		if err = transformFileTo(entry, fileItem); err != nil {
			return err
		}
	}
	return nil
}

func ZipDir(dirToZip, zipFilePath string) error {
	if !file.IsSafeFilePath(dirToZip) || !file.IsSafeFilePath(zipFilePath) {
		return errors.New("file or dir path params are illegal")
	}
	allFiles, e := file.ListFiles(dirToZip, true)
	if e != nil {
		return e
	}
	zipFile, e := os.Create(zipFilePath)
	defer func() { _ = zipFile.Close() }()
	if e != nil {
		return nil
	}
	zipWriter := zip.NewWriter(zipFile)
	defer func() { _ = zipWriter.Close() }()
	for _, fileItem := range allFiles {
		fileInfo, e := os.Stat(fileItem)
		if e != nil {
			return e
		}
		if fileInfo.IsDir() {
			continue
		}
		fileHeader, e := zip.FileInfoHeader(fileInfo)
		if e != nil {
			return e
		}
		fileHeader.Name = fileItem[len(dirToZip)+1:]
		entryWriter, e := zipWriter.CreateHeader(fileHeader)
		if e != nil {
			return e
		}
		entryFile, e := os.Open(fileItem)
		if e != nil {
			return e
		}
		e = transformFileTo(entryWriter, entryFile)
		if e != nil {
			return e
		}
	}
	return nil
}

func TarFiles(filePaths []string, tarFilePath string) error {
	if filePaths == nil || len(filePaths) == 0 {
		return nil
	}
	e := errors.New("filePath is illegal, denied to process")
	for _, filePath := range filePaths {
		if !file.IsSafeFilePath(filePath) {
			return e
		}
	}
	if !file.IsSafeFilePath(tarFilePath) {
		return e
	}
	tarFile, e := os.Create(tarFilePath)
	defer func() { _ = tarFile.Close() }()
	if e != nil {
		return e
	}
	tarWriter := tar.NewWriter(tarFile)
	for _, filePath := range filePaths {
		fileItem, e := os.Open(filePath)
		if e != nil {
			return e
		}
		var fileInfo os.FileInfo
		if fileInfo, e = os.Stat(filePath); e != nil {
			return e
		}
		var header *tar.Header
		if header, e = tar.FileInfoHeader(fileInfo, ""); e != nil {
			return e
		}
		if e = tarWriter.WriteHeader(header); e != nil {
			return e
		}
		if e = transformFileTo(tarWriter, fileItem); e != nil {
			return e
		}
	}
	return nil
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
