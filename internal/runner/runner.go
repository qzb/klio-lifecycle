package runner

import (
	"github.com/g2a-com/klio-lifecycle/internal/blueprint"
	"github.com/g2a-com/klio-lifecycle/internal/object"
)

type Interpreter func(string, map[string]interface{}) ([]string, error)

type Result struct {
	Service string `json:"service"`
	Entry   int    `json:"entry"`
	Result  string `json:"result"`
}

type BuilderRunner struct {
	Interpreter Interpreter
	Blueprint   *blueprint.Blueprint
	Service     object.Service
	Entry       object.ServiceEntry
	Tags        []string
}

func (r *BuilderRunner) Run() ([]Result, error) {
	return run(r.Blueprint, r.Service, r.Entry, object.BuilderKind, map[string]interface{}{
		"spec": r.Entry.Spec,
		"tags": r.Tags,
	}, r.Interpreter)
}

type TaggerRunner struct {
	Interpreter Interpreter
	Blueprint   *blueprint.Blueprint
	Service     object.Service
	Entry       object.ServiceEntry
}

func (r *TaggerRunner) Run() ([]Result, error) {
	return run(r.Blueprint, r.Service, r.Entry, object.TaggerKind, map[string]interface{}{
		"spec": r.Entry.Spec,
	}, r.Interpreter)
}

type PusherRunner struct {
	Interpreter Interpreter
	Blueprint   *blueprint.Blueprint
	Service     object.Service
	Entry       object.ServiceEntry
	Tags        []string
	Artifacts   []string
}

func (r *PusherRunner) Run() ([]Result, error) {
	return run(r.Blueprint, r.Service, r.Entry, object.PusherKind, map[string]interface{}{
		"spec":      r.Entry.Spec,
		"tags":      r.Tags,
		"artifacts": r.Artifacts,
	}, r.Interpreter)
}
