package script

import (
	"fmt"
	"strings"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/object"
	"github.com/g2a-com/cicd/internal/script/stdlib"
	logger "github.com/g2a-com/klio-logger-go/v2"
)

type Script struct {
	executor object.Executor
	Logger   logger.Logger
}

func New(executor object.Executor) *Script {
	script := &Script{}
	script.executor = executor
	script.Logger = logger.StandardLogger()
	return script
}

func (s *Script) Run(input interface{}) (results []string, err error) {
	displayName := strings.ToLower(fmt.Sprintf("%s %q", s.executor.Kind, s.executor.Name))

	s.Logger.WithLevel(logger.SpamLevel).Printf("Running %s", displayName)

	// Create a new tengo script instance
	script := tengo.NewScript([]byte(s.executor.Script))

	// Prepare function for updating results
	addResult := func(rs ...string) {
		results = append(results, rs...)
	}

	// Set imports & builtins
	std := stdlib.New(s.Logger)
	err = std.AddBuiltin("input", input)
	if err != nil {
		return results, fmt.Errorf("Cannot initialize standard library for %s:\n\t%s", displayName, err)
	}
	err = std.AddBuiltin("addResult", addResult)
	if err != nil {
		return results, fmt.Errorf("Cannot initialize standard library for %s:\n\t%s", displayName, err)
	}
	err = std.InitializeScript(script)
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
