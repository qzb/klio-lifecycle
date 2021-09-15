package utils

import (
	"os"
	"os/user"
	"path/filepath"
)

// FindProjectFile returns root directory (directory containing g2a.yaml file)
// for a current project
func FindProjectFile() string {
	dir, err := os.Getwd()

	if err != nil {
		return ""
	}

	dir, err = filepath.EvalSymlinks(dir)

	if err != nil {
		return ""
	}

	userHomeDir, _ := userHomeDir()

	for true {
		// Home directory cannot be a project directory
		if dir == userHomeDir {
			return ""
		}

		// Root directory of the filesystem cannot be a project directory
		if dir == filepath.Dir(dir) {
			return ""
		}

		// Check if project file exist
		for _, filename := range []string{"project.yaml", "project.yml", "project.json", "g2a.yaml"} {
			path := filepath.Join(dir, filename)
			if file, err := os.Stat(path); err == nil && !file.IsDir() {
				cwd, err := os.Getwd()
				if err != nil {
					return path
				}
				relPath, _ := filepath.Rel(cwd, path)
				if err != nil {
					return path
				}
				return relPath
			}
		}

		dir = filepath.Dir(dir)
	}

	return ""
}

func userHomeDir() (string, bool) {
	currentUser, err := user.Current()

	if err != nil {
		return "", false
	}

	homeDir, err := filepath.EvalSymlinks(currentUser.HomeDir)

	if err != nil {
		return "", false
	}

	return homeDir, true
}
