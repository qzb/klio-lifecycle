package stdlib

import (
	"os"
)

func (s *Stdlib) createOsModule() map[string]any {
	return map[string]any{
		"getenv":   os.Getenv,
		"getwd":    os.Getwd,
		"hostname": os.Hostname,
		"chdir": func(dir string) error {
			s.Logger.Printf("Changing working directory to %q", dir)
			return os.Chdir(dir)
		},
	}
}
