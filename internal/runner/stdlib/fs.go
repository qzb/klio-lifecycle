package stdlib

import (
	"os"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/tengoutil"
	"github.com/spf13/afero"
)

func (s *Stdlib) createFsModule() map[string]any {
	fs := s.Fs
	l := s.Logger

	return map[string]any{
		"chmod": func(name string, perm os.FileMode) {
			l.Printf("Creating permissions of %s to %03o", name, perm)
			fs.Chmod(name, perm)
		},
		"mkdir": func(name string, perm os.FileMode) {
			l.Printf("Creating directory %q with mode %03o", name, perm)
			fs.Mkdir(name, perm)
		},
		"mkdir_all": func(name string, perm os.FileMode) {
			l.Printf("Creating directory %q with permissions %03o", name, perm)
			fs.MkdirAll(name, perm)
		},
		"remove": func(name string) {
			l.Printf("Removing %s", name)
			fs.Remove(name)
		},
		"remove_all": func(name string) {
			l.Printf("Removing %s", name)
			fs.RemoveAll(name)
		},
		"rename": func(oldname, newname string) {
			l.Printf("Renaming %q to %q", oldname, newname)
			fs.Rename(oldname, newname)
		},
		"stat": func(name string) (fileInfo, error) {
			info, err := fs.Stat(name)
			return fileInfo{info}, err
		},
		"dir_exists": func(path string) (bool, error) {
			return afero.DirExists(fs, path)
		},
		"exists": func(path string) (bool, error) {
			return afero.Exists(fs, path)
		},
		"is_dir": func(path string) (bool, error) {
			return afero.IsDir(fs, path)
		},
		"is_empty": func(path string) (bool, error) {
			return afero.IsEmpty(fs, path)
		},
		"read_dir": func(dirname string) (result []fileInfo, err error) {
			infos, err := afero.ReadDir(fs, dirname)
			for _, info := range infos {
				result = append(result, fileInfo{info})
			}
			return result, err
		},
		"read_file": func(filename string) ([]byte, error) {
			return afero.ReadFile(fs, filename)
		},
		"temp_dir": func() (name string, err error) {
			return afero.TempDir(fs, "", "") // TODO: set dir and prefix
		},
		"write_file": func(filename string, data []byte, perm os.FileMode) error {
			l.Printf("Writing to file %s", filename)
			return afero.WriteFile(fs, filename, data, perm)
		},
	}
}

type fileInfo struct {
	os.FileInfo
}

func (f fileInfo) EncodeTengoObject() (tengo.Object, error) {
	return tengoutil.ToImmutableObject(map[string]any{
		"name":     f.Name,
		"size":     f.Size,
		"mode":     f.Mode,
		"mod_time": f.ModTime,
		"is_dir":   f.IsDir,
	})
}
