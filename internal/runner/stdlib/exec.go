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
			cmd.Dir = s.WorkingDirectory
			return command{cmd, s.Logger}
		},
	}
}

type command struct {
	*exec.Command
	*logger.Logger
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
		"run":         c.Run,
		"must_run":    c.mustRun,
		"output":      c.mustOutput,
		"must_output": c.output,
	})
}

func (c command) run() error {
	if c.StdoutLogger == nil {
		c.Command.StdoutLogger = c.Logger.WithLevel(logger.InfoLevel)
	}
	if c.StderrLogger == nil {
		c.Command.StderrLogger = c.Logger.WithLevel(logger.ErrorLevel)
	}
	return c.Command.Run()
}

func (c command) mustRun() {
	err := c.run()
	if err != nil {
		panic(err)
	}
}

func (c command) output() (string, error) {
	if c.Command.StdoutLogger == nil {
		c.Command.StdoutLogger = c.Logger.WithLevel(logger.VerboseLevel)
	}
	err := c.run()
	return c.Command.StdoutText, err
}

func (c command) mustOutput() string {
	out, err := c.output()
	if err != nil {
		panic(err)
	}
	return out
}
