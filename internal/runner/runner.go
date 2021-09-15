package runner

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/g2a-com/cicd/internal/blueprint"
	"github.com/g2a-com/cicd/internal/exec"
	"github.com/g2a-com/cicd/internal/object"
	log "github.com/g2a-com/klio-logger-go"
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
	l := log.StandardLogger().WithTags(service.Name, entry.Type)

	executor, ok := b.GetExecutor(kind, entry.Type)
	if !ok {
		return results, fmt.Errorf("builder %q does not exist", entry.Type)
	}

	vm := goja.New()

	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	vm.Set("input", input)
	vm.Set("output", []interface{}{})

	vm.Set("exec", func(command string, args []string) string {
		cmd := exec.NewCommand(command, args...)
		cmd.StdoutLogger = l
		cmd.StderrLogger = l.WithLevel(log.ErrorLevel)
		cmd.Dir = service.Directory
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		return cmd.StdoutText
	})

	vm.Set("read", readFile)
	vm.Set("write", writeFile)
	vm.Set("request", request)

	_, err = vm.RunString(executor.JS)
	if err != nil {
		// if jserr, ok := err.(*goja.Exception); ok {
		// 	msg := jserr.Value().Export()
		// 	return "", fmt.Errorf("%s %q failed: %s", strings.ToLower(string(executor.Kind)), executor.Name, msg)
		// }
		return results, err
	}

	var output []string
	err = vm.ExportTo(vm.Get("output"), &output)
	if err != nil {
		return results, err
	}

	for _, o := range output {
		results = append(results, Result{
			Service: service.Name,
			Entry:   entry.Index,
			Result:  o,
		})
	}

	return results, nil
}
