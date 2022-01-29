package runner

import (
	"fmt"
	"strings"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/blueprint"
	"github.com/g2a-com/cicd/internal/object"
	"github.com/g2a-com/cicd/internal/runner/stdlib"
	logger "github.com/g2a-com/klio-logger-go"
	"github.com/spf13/afero"
)

type Result struct {
	Service string `json:"service"`
	Entry   int    `json:"entry"`
	Result  string `json:"result"`
}

type BuilderRunner struct {
	Blueprint *blueprint.Blueprint
	Service   object.Service
	Entry     object.ServiceEntry
	Tags      []string
}

func (r *BuilderRunner) Run() ([]Result, error) {
	return run(r.Blueprint, r.Service, r.Entry, object.BuilderKind, map[string]interface{}{
		"spec": r.Entry.Spec,
		"tags": r.Tags,
	})
}

type TaggerRunner struct {
	Blueprint *blueprint.Blueprint
	Service   object.Service
	Entry     object.ServiceEntry
}

func (r *TaggerRunner) Run() ([]Result, error) {
	return run(r.Blueprint, r.Service, r.Entry, object.TaggerKind, map[string]interface{}{
		"spec": r.Entry.Spec,
	})
}

type PusherRunner struct {
	Blueprint *blueprint.Blueprint
	Service   object.Service
	Entry     object.ServiceEntry
	Tags      []string
	Artifacts []string
}

func (r *PusherRunner) Run() ([]Result, error) {
	return run(r.Blueprint, r.Service, r.Entry, object.PusherKind, map[string]interface{}{
		"spec":      r.Entry.Spec,
		"tags":      r.Tags,
		"artifacts": r.Artifacts,
	})
}

type DeployerRunner struct {
	Blueprint *blueprint.Blueprint
	Service   object.Service
	Entry     object.ServiceEntry
	Force     bool
	DryRun    bool
	Wait      int
}

func (r *DeployerRunner) Run() ([]Result, error) {
	return run(r.Blueprint, r.Service, r.Entry, object.DeployerKind, map[string]interface{}{
		"spec":   r.Entry.Spec,
		"force":  r.Force,
		"dryRun": r.DryRun,
		"wait":   r.Wait,
	})
}

func run(b *blueprint.Blueprint, service object.Service, entry object.ServiceEntry, kind object.Kind, input map[string]interface{}) (results []Result, err error) {
	displayName := fmt.Sprintf("%s %q", strings.ToLower(string(kind)), entry.Type)

	// Try to get an executor object
	executor, ok := b.GetExecutor(kind, entry.Type)
	if !ok {
		return results, fmt.Errorf("%s does not exist", displayName)
	}

	// Create a new tengo script instance
	script := tengo.NewScript([]byte(executor.Script))

	// Prepare function for updating results
	addResult := func(rs ...string) {
		for _, r := range rs {
			results = append(results, Result{Service: service.Name, Entry: entry.Index, Result: r})
		}
	}

	// Set imports & builtins
	std := stdlib.Stdlib{
		Fs:               afero.OsFs{},
		Logger:           logger.StandardLogger().WithTags(service.Name, entry.Type),
		WorkingDirectory: service.Directory,
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
