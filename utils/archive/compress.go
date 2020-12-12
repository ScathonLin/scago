// author: linhuadong(scathonlin)
// creatdate: 2020/12/06
package archive

import (
	"archive/tar"
	"archive/zip"
	"errors"
	"github.com/klauspost/compress/gzip"
	"io"
	"os"
	"scago/utils/file"
)

var filePathIllegalErr = errors.New("filePath is illegal")

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
		// open file.
		fileItem, err := os.Open(filePath)
		if err != nil {
			return err
		}
		var fileInfo os.FileInfo
		if fileInfo, err = os.Stat(filePath); err != nil {
			return err
		}
		var fileHeader *zip.FileHeader
		if fileHeader, err = zip.FileInfoHeader(fileInfo); err != nil {
			return err
		}
		zipWriter, err := zipWriter.CreateHeader(fileHeader)
		if err != nil {
			return nil
		}
		// write file to zip entry.
		if err = transformFileTo(zipWriter, fileItem); err != nil {
			return err
		}
	}
	_ = zipWriter.Flush()
	return nil
}

func ZipDir(dirToZip, zipFilePath string) error {
	if !file.IsSafeFilePath(dirToZip) || !file.IsSafeFilePath(zipFilePath) {
		return errors.New("file or dir path params are illegal")
	}
	allFiles, err := file.ListFiles(dirToZip, true)
	if err != nil {
		return err
	}
	zipFile, err := os.Create(zipFilePath)
	defer func() { _ = zipFile.Close() }()
	if err != nil {
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
		var fileHeader *zip.FileHeader
		if fileHeader, e = zip.FileInfoHeader(fileInfo); e != nil {
			return e
		}
		fileHeader.Name = fileItem[len(dirToZip)+1:]
		var entryWriter io.Writer
		if entryWriter, e = zipWriter.CreateHeader(fileHeader); e != nil {
			return e
		}
		var entryFile *os.File
		if entryFile, e = os.Open(fileItem); e != nil {
			return e
		}
		if e = transformFileTo(entryWriter, entryFile); e != nil {
			return e
		}
	}
	return nil
}

func TarFiles(filePaths []string, tarFilePath string) error {
	if filePaths == nil || len(filePaths) == 0 {
		return nil
	}
	err := errors.New("filePath is illegal, denied to process")
	for _, filePath := range filePaths {
		if !file.IsSafeFilePath(filePath) {
			return err
		}
	}
	if !file.IsSafeFilePath(tarFilePath) {
		return err
	}
	tarFile, err := os.Create(tarFilePath)
	defer func() { _ = tarFile.Close() }()
	if err != nil {
		return err
	}
	tarWriter := tar.NewWriter(tarFile)
	for _, filePath := range filePaths {
		fileItem, err := os.Open(filePath)
		if err != nil {
			return err
		}
		var fileInfo os.FileInfo
		if fileInfo, err = os.Stat(filePath); err != nil {
			return err
		}
		var header *tar.Header
		if header, err = tar.FileInfoHeader(fileInfo, ""); err != nil {
			return err
		}
		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}
		if err = transformFileTo(tarWriter, fileItem); err != nil {
			return err
		}
		_ = tarWriter.Flush()
	}
	return nil
}

func Gzip(filePath, gzipFilePath string) error {
	if !file.IsSafeFilePath(filePath) {
		return filePathIllegalErr
	}
	var err error
	fileToGzip, err := os.Open(filePath)
	defer func() { _ = fileToGzip.Close() }()
	if err != nil {
		return err
	}
	gzipFile, err := os.Create(gzipFilePath)
	defer func() { _ = gzipFile.Close() }()
	if err != nil {
		return err
	}
	gzipWriter := gzip.NewWriter(gzipFile)
	defer func() { _ = gzipWriter.Close() }()
	var stat os.FileInfo
	if stat, err = os.Stat(filePath); err != nil {
		return err
	}
	gzipWriter.Name = stat.Name()
	gzipWriter.ModTime = stat.ModTime()
	if err = transformFileTo(gzipWriter, fileToGzip); err != nil {
		return err
	}
	_ = gzipWriter.Flush()
	return nil
}

func transformFileTo(writer io.Writer, file *os.File) error {
	defer func() { _ = file.Close() }()
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	return nil
}
