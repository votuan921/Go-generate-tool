package struct_extend

import (
	"path/filepath"
	"strings"
)

func FileName(path string) string {
	return filepath.Base(path)
}

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func FileDir(path string) string {
	return filepath.Dir(path)
}

func DirJoin(dir, filePath string) string {
	return filepath.Join(dir, filePath)
}
