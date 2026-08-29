package file

import (
	"os"
	"path/filepath"
)

func ValidDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func Joinwd(elem ...string) string {
	wd, _ := os.Getwd()
	dirs := []string{wd}
	dirs = append(dirs, elem...)
	return filepath.Join(dirs...)
}
