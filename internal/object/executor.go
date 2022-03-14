package object

import "github.com/qri-io/jsonschema"

type Executor struct {
	Directory string
	Kind      Kind
	Name      string
	Script    string
	Schema    jsonschema.Schema
}
