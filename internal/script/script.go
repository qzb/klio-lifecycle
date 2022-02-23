package script

import (
	"fmt"
	"strings"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/object"
	"github.com/g2a-com/cicd/internal/script/stdlib"
	logger "github.com/g2a-com/klio-logger-go"
	"github.com/spf13/afero"
)

type Script struct {
	executor object.Executor
	Fs       afero.Fs
	Logger   *logger.Logger
	Dir      string
}

func New(executor object.Executor) *Script {
	script := &Script{}
	script.executor = executor
	script.Logger = logger.StandardLogger()
	script.Fs = afero.NewOsFs()
	return script
}

func (s *Script) Run(input any) (results []string, err error) {
	displayName := strings.ToLower(fmt.Sprintf("%s %q", s.executor.Kind, s.executor.Name))

	s.Logger.WithLevel(logger.SpamLevel).Printf("Running %s", displayName)

	// Create a new tengo script instance
	script := tengo.NewScript([]byte(s.executor.Script))

	// Prepare function for updating results
	addResult := func(rs ...string) {
		results = append(results, rs...)
	}

	// Set imports & builtins
	std := stdlib.Stdlib{
		Fs:               afero.OsFs{},
		Logger:           s.Logger,
		WorkingDirectory: s.Dir,
		Builtins: map[string]any{
			"input":     input,
			"addResult": addResult,
		},
	}
	err = std.AddToScript(script)
	if err != nil {
		return results, fmt.Errorf("Cannot initialize standard library for %s:\n\t%s", displayName, err)
	}

	// Run the script
	_, err = script.Run()
	if err != nil {
		return results, fmt.Errorf("Error occurred while running %s:\n\t%s", displayName, err)
	}

	return results, nil
}
