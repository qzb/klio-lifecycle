package exec

import (
	"errors"
	"os/exec"
	"strings"

	log "github.com/g2a-com/klio-logger-go"
)

type Command struct {
	Command             string
	Args                []string
	Dir                 string
	ErrorTextFromStderr bool
	StdoutLogger        *log.Logger
	StderrLogger        *log.Logger
	StdoutText          string
	StderrText          string
	Tags                []string
}

func NewCommand(cmd string, args ...string) *Command {
	return &Command{
		Command:      cmd,
		Args:         args,
		StdoutLogger: log.StandardLogger(),
		StderrLogger: log.ErrorLogger(),
	}
}

func (cmd *Command) Run() error {
	log.Debugf(`running: %s %s`, cmd.Command, joinCommandLineArgs(cmd.Args...))

	externalCmd := exec.Command(cmd.Command, cmd.Args...)
	externalCmd.Dir = cmd.Dir
	externalCmd.Stdout = &writer{Text: &cmd.StdoutText, Logger: cmd.StdoutLogger}
	externalCmd.Stderr = &writer{Text: &cmd.StderrText, Logger: cmd.StderrLogger}

	err := externalCmd.Run()

	if _, ok := err.(*exec.ExitError); cmd.ErrorTextFromStderr && ok {
		return errors.New(cmd.StderrText)
	} else {
		return err
	}
}

func joinCommandLineArgs(args ...string) string {
	var result string
	for i, arg := range args {
		if i != 0 {
			result += " "
		}
		if strings.ContainsAny(arg, " $\t") {
			result += "'" + strings.ReplaceAll(arg, "'", "'\\''") + "'"
		} else {
			result += arg
		}
	}
	return result
}
