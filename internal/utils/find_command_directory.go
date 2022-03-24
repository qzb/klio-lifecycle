package utils

import (
	"os"
	"path/filepath"
)

// FindCommandDirectory returns directory of currently run command
func FindCommandDirectory() string {
	filename, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(filename)
}
