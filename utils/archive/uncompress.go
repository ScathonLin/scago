// description: tools to compress or uncompress archived file such as gzip, tar, zip.
// author: linhuadong(scathonlin)
// date: 2020-12-2 21:44
package archive

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path"
	"scago/utils/file"
	"strings"
)

//GUnzip can be used to uncompress gzip file with given gzip file path and target uncompress file.
func GUnzip(gzipFilePath, gUnzipDir string) error {
	gzipFilePath = path.Clean(gzipFilePath)
	gUnzipDir = path.Clean(gUnzipDir)
	gzipFile, err := os.Open(gzipFilePath)
	defer func() { _ = gzipFile.Close() }()
	if err != nil {
		return err
	}
	gzipReader, err := gzip.NewReader(gzipFile)
	defer func() { _ = gzipReader.Close() }()
	if err != nil {
		return err
	}
	bytes, buf := make([]byte, 0), make([]byte, 1<<10)
	uncompressFile, err := os.Create(gUnzipDir)
	if err != nil {
		return err
	}
	for {
		readlen, err := gzipReader.Read(buf)
		if err == io.EOF {
			_, _ = uncompressFile.Write(buf[:readlen])
			break
		}
		bytes = append(bytes, buf[:readlen]...)
		_, _ = uncompressFile.Write(bytes)
	}
	return nil
}

//UnZip can be used to uncompress zip file.
// 1. maxEntries is be used to symbol the max entry num in the zip file which is counted during unzip procedure,
// 2. limitSize is be used to symbol the total size of unzip file which is calculated during unzip procedure.
// the purpose of limitSize and maxEntries is to protect the program from high compression radio file and zip bomb attacks.
func UnZip(zipFilePath, unZipDir string, maxEntries, limitSize int) error {
	zipFilePath = path.Clean(zipFilePath)
	unZipDir = path.Clean(unZipDir)
	zipFile, err := os.Open(zipFilePath)
	defer func() { _ = zipFile.Close() }()
	if err != nil {
		return err
	}
	zipReader, err := zip.OpenReader(zipFilePath)
	defer func() { _ = zipReader.Close() }()
	if err != nil {
		return err
	}
	entryCounter, totalSize := 0, int64(0)
	for _, zipEntry := range zipReader.File {
		if zipEntry == nil {
			continue
		}
		entryCounter++
		if entryCounter > maxEntries {
			return errors.New("too many entries in zip file, maybe there is a zip bomb attacks")
		}
		entryFullPath := path.Clean(path.Join(unZipDir, zipEntry.Name))
		if strings.LastIndex(entryFullPath, unZipDir) != 0 {
			return errors.New("your zip file may can cause crossing dir attacks,system denied to process it")
		}
		fileInfo := zipEntry.FileInfo()
		if fileInfo.IsDir() {
			mkdirAll(entryFullPath, os.ModePerm)
		} else {
			// create file dir.
			entryFileDir := entryFullPath[:strings.LastIndex(entryFullPath, string(os.PathSeparator))]
			mkdirAll(entryFileDir, os.ModePerm)
			entryFile, err := os.Create(entryFullPath)
			if err != nil {
				return err
			}
			entryReader, err := zipEntry.Open()
			if err != nil {
				return err
			}
			// accumulate the total size.
			totalSize += transferReadCloserToFile(entryFile, entryReader)
			// check if the total size over limit the given max size.
			if totalSize > int64(limitSize) {
				return errors.New("the zip file may have high compression radio, it seems like zip bomb attacks")
			}
		}
	}
	return nil
}

// UnTarGZ can uncompress the *.tar.gz file to specified dir.
func UnTarGZ(tarGzFilePath, unTarGzDir string) error {
	// get the canonical path.
	tarGzFilePath = path.Clean(tarGzFilePath)
	unTarGzDir = path.Clean(unTarGzDir)
	targzFile, err := os.Open(tarGzFilePath)
	defer func() { _ = targzFile.Close() }()
	if err != nil {
		return err
	}
	gzreader, err := gzip.NewReader(targzFile)
	defer func() { _ = gzreader.Close() }()
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(gzreader)
	for entry, err := tarReader.Next(); err != io.EOF; entry, err = tarReader.Next() {
		if entry == nil {
			continue
		}
		// create entry file full path.
		entryFullPath := path.Clean(path.Join(unTarGzDir, entry.Name))
		// do check to avoid to be attacked by crossing dir attacks.
		if strings.LastIndex(entryFullPath, unTarGzDir) != 0 {
			// attention!! there is a risk of crossing dir attack!
			return errors.New("the targz file is dangerous! system denied to process it")
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
	return nil
}

//UnTar uncompress the tar file to the target dir.
func UnTar(tarFilePath, unTarDir string) error {
	if !file.IsSafeFilePath(tarFilePath) || !file.IsSafeFilePath(unTarDir) {
		// file path is illegal.
		return filePathIllegalErr
	}
	unTarDir = strings.TrimSuffix(unTarDir, string(os.PathSeparator))
	// if untar dir is not existed, then create it!
	if !file.DirExists(unTarDir) && os.MkdirAll(unTarDir, os.ModePerm) != nil {
		return errors.New("failed to create dir " + unTarDir)
	}
	tarFile, err := os.Open(tarFilePath)
	defer func() { _ = tarFile.Close() }()
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(tarFile)
	for fileItem, err := tarReader.Next(); err != io.EOF; fileItem, err = tarReader.Next() {
		if fileItem == nil {
			continue
		}
		// we should use fileItem.Name instead of fileItem.FileInfo().Name()
		// because fileItem.Name is same to the relative path which specified during tar files, it save the dir structure,
		// but fileInfo.Name() just storage the file name.
		fileItemName := strings.TrimPrefix(fileItem.Name, string(os.PathSeparator))
		fileItemFullPath := path.Join(unTarDir, fileItemName)
		if fileItem.Typeflag == tar.TypeDir {
			err := os.MkdirAll(fileItemFullPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			// process dir.
			finalFile, err := os.Create(fileItemFullPath)
			if err != nil {
				return err
			}
			transferReaderToFile(finalFile, tarReader)
		}
	}
	return nil
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
func transferReadCloserToFile(dstFile *os.File, reader io.ReadCloser) (transferLen int64) {
	defer func() { _ = dstFile.Close() }()
	defer func() { _ = reader.Close() }()
	var err error
	if transferLen, err = io.Copy(dstFile, reader); err != nil {
		panic("failed to transfer reader to file : " + dstFile.Name())
	}
	return transferLen
}
