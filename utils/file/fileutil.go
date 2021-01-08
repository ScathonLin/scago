package file

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

//ListFiles get all files in the directory specified.
// if the param recursive is true, it will be get all files recusively, else just get files in current dir.
func ListFiles(filePath string, recursive bool) ([]string, error) {
	file, err := os.Open(filePath)
	defer func() { _ = file.Close() }()
	if err != nil {
		return nil, err
	}
	if !recursive {
		return file.Readdirnames(-1)
	}
	return listFilesRecusively(filePath)
}

func listFilesRecusively(filePath string) (files []string, err error) {
	if !IsDir(filePath) {
		// check if the file pointed by filePath is directory, if not, return directly to advoid unmeaning file open operation.
		return
	}
	file, err := os.Open(filePath)
	defer func() { _ = file.Close() }()
	if err != nil {
		return files, err
	}
	subFiles, err := file.Readdir(-1)
	if err != nil {
		return files, err
	}
	for _, subFile := range subFiles {
		wholePath := path.Join(filePath, subFile.Name())
		files = append(files, wholePath)
		// get files in sub dir recursively.
		fls, err := listFilesRecusively(wholePath)
		if err != nil {
			return files, err
		}
		files = append(files, fls...)
	}
	return files, nil
}

func IsDir(filePath string) bool {
	stat, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func FileExists(filePath string) bool {
	filePath = path.Clean(filePath)
	stat, err := os.Lstat(filePath)
	return (err == nil || os.IsExist(err)) && !stat.IsDir()
}

func DirExists(dirPath string) bool {
	dirPath = path.Clean(dirPath)
	lstat, err := os.Lstat(dirPath)
	return (err == nil || os.IsExist(err)) && lstat.IsDir()
}

func IsSafeFilePath(filePath string) bool {
	return strings.Compare(path.Clean(filePath), filePath) == 0
}

func ReadFileToLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("failed to open file : " + filePath)
	}
	reader := bufio.NewReader(file)
	lines := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}
func CopyDir(srcDir, dstDir string, createIfNotExists bool) error {
	if IsDir(srcDir) {
		dirFile, e := os.Open(srcDir)
		if e != nil {
			return e
		}
		subFiles, e := dirFile.Readdir(-1)
		if e != nil {
			return e
		}
		if !DirExists(dstDir) {
			_ = os.MkdirAll(dstDir, os.ModePerm)
		}
		for _, sf := range subFiles {
			e := CopyDir(path.Join(srcDir, sf.Name()), path.Join(dstDir, sf.Name()), createIfNotExists)
			if e != nil {
				panic(e)
			}
		}
	} else {
		// copy file.
		reader, e := os.Open(srcDir)
		defer func() { _ = reader.Close() }()
		if e != nil {
			panic(e)
		}
		writer, e := os.OpenFile(dstDir, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer func() { _ = writer.Close() }()
		if e != nil {
			panic(e)
		}
		_, e = io.Copy(writer, reader)
		if e != nil {
			panic(e)
		}
	}
	return nil
}
