package archive

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

// description: tools to compress or uncompress archived file such as gzip, tar, zip.
// author: linhuadong(scathonlin)
// date: 2020-12-2 21:44

func UncompressTarGzFile(tarGzFilePath, uncompressDir string) (*os.File, error) {
	// get the canonical path.
	tarGzFilePath = path.Clean(tarGzFilePath)
	uncompressDir = path.Clean(uncompressDir)
	targzFile, err := os.Open(tarGzFilePath)
	defer func() { _ = targzFile.Close() }()
	if err != nil {
		return nil, errors.New("failed to open targz file with " + tarGzFilePath)
	}
	gzreader, err := gzip.NewReader(targzFile)
	defer func() { _ = gzreader.Close() }()
	if err != nil {
		return nil, errors.New("failed to get gzip reader with file: " + tarGzFilePath)
	}
	tarReader := tar.NewReader(gzreader)
	for entry, err := tarReader.Next(); err != io.EOF; entry, err = tarReader.Next() {
		if entry == nil {
			continue
		}
		// create entry file full path.
		entryFullPath := path.Clean(path.Join(uncompressDir, entry.Name))
		// do check to avoid to be attacked by crossing dir attacks.
		if strings.LastIndex(entryFullPath, uncompressDir) != 0 {
			// attention!! there is a risk of crossing dir attack!
			return nil, errors.New("the targz file is dangerous! system denied to process it")
		}
		if entry.Typeflag == tar.TypeDir {
			// current entry is a dir.
			mkdirAll(entryFullPath, os.ModePerm)
		} else {
			// current entry is a file
			// 1. parse and get dir path of entry file.
			entryFileDir := entryFullPath[:strings.LastIndex(entryFullPath, string(os.PathSeparator))]
			mkdirAll(entryFileDir, os.ModePerm)
			entryFile, err := os.Create(entryFullPath)
			if err != nil {
				panic("failed to create entry file with path: " + entryFullPath)
			}
			transferReaderToFile(entryFile, tarReader)
		}
	}
	return os.Open(uncompressDir)
}

func mkdirAll(dir string, mode os.FileMode) {
	dir = path.Clean(dir)
	_, err := os.Lstat(dir)
	if err == nil || os.IsExist(err) {
		return
	}
	if err := os.MkdirAll(dir, mode); err != nil {
		panic("failed to mkdir dir: " + dir + " with error : " + err.Error())
	}
}

func transferReaderToFile(dstFile *os.File, reader io.Reader) {
	defer func() { _ = dstFile.Close() }()
	if _, err := io.Copy(dstFile, reader); err != nil {
		panic("failed to transfer reader to file : " + dstFile.Name())
	}
}
