package stdlib

import (
	"fmt"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/exec"
	"github.com/g2a-com/cicd/internal/tengoutil"
	logger "github.com/g2a-com/klio-logger-go"
)

func (s *Stdlib) createExecModule() map[string]any {
	return map[string]any{
		"command": func(cmdName string, args ...string) command {
			cmd := exec.NewCommand(cmdName, args...)
			cmd.StdoutLogger = s.Logger.WithLevel(logger.VerboseLevel)
			cmd.StderrLogger = s.Logger.WithLevel(logger.ErrorLevel)
			cmd.Dir = s.WorkingDirectory
			return command{cmd}
		},
	}
}

type command struct {
	*exec.Command
}

func (c command) EncodeTengoObject() (tengo.Object, error) {
	return tengoutil.ToImmutableObject(map[string]any{
		"set_stdout_level": func(levelName string) error {
			level, ok := logger.ParseLevel(levelName)
			if !ok {
				return fmt.Errorf("invalid level: %s", levelName)
			}
			c.StdoutLogger = c.StdoutLogger.WithLevel(level)
			return nil
		},
		"set_stderr_level": func(levelName string) error {
			level, ok := logger.ParseLevel(levelName)
			if !ok {
				return fmt.Errorf("invalid level: %s", levelName)
			}
			c.StderrLogger = c.StderrLogger.WithLevel(level)
			return nil
		},
		"get_stderr_text": func() string {
			return c.StderrText
		},
		"get_stdout_text": func() string {
			return c.StdoutText
		},
		"run": func() error {
			return c.Run()
		},
	})
}
