package object

import "github.com/qri-io/jsonschema"

type Executor struct {
	Directory string            `placeholders:"disable"`
	Kind      Kind              `placeholders:"disable"`
	Name      string            `placeholders:"disable"`
	Script    string            `placeholders:"disable"`
	Schema    jsonschema.Schema `placeholders:"disable"`
}
