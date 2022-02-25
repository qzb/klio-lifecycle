package stdlib

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/tengoutil"
	logger "github.com/g2a-com/klio-logger-go"
)

func (s *Stdlib) createExecModule() map[string]any {
	l := s.Logger

	return map[string]any{
		"command": func(opts cmdOpts) *cmd {
			if opts.Dir == "" {
				opts.Dir = s.WorkingDirectory
			}
			if opts.StdoutLevel == "" {
				opts.StdoutLevel = logger.InfoLevel
			}
			if opts.StderrLevel == "" {
				opts.StderrLevel = logger.ErrorLevel
			}

			return &cmd{&opts, l}
		},
		"run": func(cmdName string, args ...string) cmdResult {
			opts := &cmdOpts{
				Name:        cmdName,
				Args:        args,
				Dir:         s.WorkingDirectory,
				StdoutLevel: logger.InfoLevel,
				StderrLevel: logger.ErrorLevel,
			}
			return cmd{opts, l}.run()
		},
		"run_silently": func(cmdName string, args ...string) cmdResult {
			opts := &cmdOpts{
				Name:        cmdName,
				Args:        args,
				Dir:         s.WorkingDirectory,
				StdoutLevel: logger.DebugLevel,
				StderrLevel: logger.ErrorLevel,
			}
			return cmd{opts, l}.run()
		},
	}
}

type cmd struct {
	opts *cmdOpts
	log  *logger.Logger
}

func (c *cmd) EncodeTengoObject() (tengo.Object, error) {
	return tengoutil.ToImmutableObject(map[string]any{
		"run": c.run,
	})
}

func (c cmd) run() cmdResult {
	var stdoutBuffer, stderrBuffer bytes.Buffer

	execCmd := exec.Command(c.opts.Name, c.opts.Args...)
	execCmd.Dir = c.opts.Dir
	execCmd.Env = c.opts.Env
	execCmd.Stdout = &stdoutBuffer
	execCmd.Stderr = &stderrBuffer

	if c.opts.StdoutLevel != "disable" {
		level := parseLevel(string(c.opts.StdoutLevel))
		execCmd.Stdout = io.MultiWriter(execCmd.Stdout, c.log.WithLevel(level))
	}
	if c.opts.StderrLevel != "disable" {
		level := parseLevel(string(c.opts.StderrLevel))
		execCmd.Stderr = io.MultiWriter(execCmd.Stderr, c.log.WithLevel(level))
	}

	err := execCmd.Run()
	if _, ok := err.(*exec.ExitError); err != nil && (!ok || !c.opts.IgnoreExitErrors) {
		panic(fmt.Errorf("Failed to run %s:\n\t%s", execCmd, err))
	}

	return cmdResult{
		ExitCode:   execCmd.ProcessState.ExitCode(),
		StdoutText: stdoutBuffer.String(),
		StderrText: stderrBuffer.String(),
	}
}

type cmdOpts struct {
	Name             string       `tengo:"name"`
	Args             []string     `tengo:"args"`
	Dir              string       `tengo:"dir"`
	Env              []string     `tengo:"env"`
	StdoutLevel      logger.Level `tengo:"stdout_level"`
	StderrLevel      logger.Level `tengo:"stderr_level"`
	IgnoreExitErrors bool         `tengo:"ignore_exit_errors"`
}

type cmdResult struct {
	ExitCode   int    `tengo:"exit_code"`
	StdoutText string `tengo:"stdout_text"`
	StderrText string `tengo:"stderr_text"`
}

func parseLevel(levelName string) logger.Level {
	parsed, ok := logger.ParseLevel(levelName)
	if !ok {
		panic(fmt.Errorf(
			"Unknown level %q, use one of: %s, %s, %s, %s, %s, %s, %s, disable",
			levelName, logger.SpamLevel, logger.DebugLevel, logger.VerboseLevel, logger.InfoLevel, logger.WarnLevel, logger.ErrorLevel, logger.FatalLevel,
		))
	}
	return parsed
}
