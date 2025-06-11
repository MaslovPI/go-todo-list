package filemanagement

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func GetCSVPath() (string, error) {
	baseDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(baseDir, ".todo")
	os.MkdirAll(dir, 0755)

	return filepath.Join(dir, "data.csv"), nil
}

func LoadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func CloseFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}
