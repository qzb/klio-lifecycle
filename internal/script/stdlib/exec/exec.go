package exec

import (
	"bytes"
	"fmt"
	"io"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/tengoutil"
	logger "github.com/g2a-com/klio-logger-go/v2"
	"k8s.io/utils/exec"
)

type module struct {
	exec   exec.Interface
	logger logger.Logger
}

func New(logger logger.Logger) *module {
	return &module{
		exec:   exec.New(),
		logger: logger,
	}
}

func (m *module) Import(name string) (interface{}, error) {
	return tengoutil.ToImmutableObject(map[string]interface{}{
		"__module_name__": name,
		"command":         m.command,
		"run":             m.run,
		"run_silently":    m.runSilently,
	})
}

func (m *module) command(opts cmdOpts) *cmd {
	if opts.StdoutLevel == "" {
		opts.StdoutLevel = logger.InfoLevel
	}
	if opts.StderrLevel == "" {
		opts.StderrLevel = logger.ErrorLevel
	}

	return &cmd{&opts, m.logger, m.exec}
}

func (m *module) run(cmdName string, args ...string) cmdResult {
	opts := &cmdOpts{
		Name:        cmdName,
		Args:        args,
		StdoutLevel: logger.InfoLevel,
		StderrLevel: logger.ErrorLevel,
	}
	return (&cmd{opts, m.logger, m.exec}).run()
}

func (m *module) runSilently(cmdName string, args ...string) cmdResult {
	opts := &cmdOpts{
		Name:        cmdName,
		Args:        args,
		StdoutLevel: logger.DebugLevel,
		StderrLevel: logger.ErrorLevel,
	}
	return (&cmd{opts, m.logger, m.exec}).run()
}

type cmd struct {
	opts *cmdOpts
	log  logger.Logger
	exec exec.Interface
}

func (c *cmd) EncodeTengoObject() (tengo.Object, error) {
	return tengoutil.ToImmutableObject(map[string]interface{}{
		"run": c.run,
	})
}

func (c *cmd) run() cmdResult {
	// Prepare command to run
	execCmd := c.exec.Command(c.opts.Name, c.opts.Args...)
	execCmd.SetDir(c.opts.Dir)
	execCmd.SetEnv(c.opts.Env)

	var stdoutBuffer, stderrBuffer bytes.Buffer

	if c.opts.StdoutLevel == "disable" {
		execCmd.SetStdout(&stdoutBuffer)
	} else {
		level := parseLevel(string(c.opts.StdoutLevel))
		execCmd.SetStdout(io.MultiWriter(&stdoutBuffer, c.log.WithLevel(level)))
	}
	if c.opts.StderrLevel == "disable" {
		execCmd.SetStderr(&stderrBuffer)
	} else {
		level := parseLevel(string(c.opts.StderrLevel))
		execCmd.SetStderr(io.MultiWriter(&stderrBuffer, c.log.WithLevel(level)))
	}

	// Run and prepare result
	result := cmdResult{}
	result.Error = execCmd.Run()
	if result.Error != nil {
		if !c.opts.IgnoreErrors {
			panic(result.Error)
		}
		if e, ok := result.Error.(exec.ExitError); ok {
			result.ExitCode = e.ExitStatus()
		} else {
			result.ExitCode = -1
		}
	}

	result.StdoutText = stdoutBuffer.String()
	result.StderrText = stderrBuffer.String()

	return result
}

type cmdOpts struct {
	Name         string       `tengo:"name"`
	Args         []string     `tengo:"args"`
	Dir          string       `tengo:"dir"`
	Env          []string     `tengo:"env"`
	StdoutLevel  logger.Level `tengo:"stdout_level"`
	StderrLevel  logger.Level `tengo:"stderr_level"`
	IgnoreErrors bool         `tengo:"ignore_errors"`
}

type cmdResult struct {
	ExitCode   int    `tengo:"exit_code"`
	Error      error  `tengo:"error"`
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
