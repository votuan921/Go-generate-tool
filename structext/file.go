package structext

import (
	"os"
	"path/filepath"
	"strings"
)

//FileName returns only file name by given file path
func FileName(path string) string {
	return filepath.Base(path)
}

// FileNameWithoutExt returns file name without file extension
func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// FileDir return dir of given file
func FileDir(path string) string {
	return filepath.Dir(path)
}

// DirJoin joins a dir and a path to full path
func DirJoin(dir, filePath string) string {
	return filepath.Join(dir, filePath)
}

// FileExisting checks given file is existing or not
func FileExisting(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil { // unknown error
		return false
	}

	return !info.IsDir()
}

// Mkdirp is equivalent to `mkdir -p` command
func Mkdirp(dirPath string) error {
	return os.MkdirAll(dirPath, os.ModeDir)
}
