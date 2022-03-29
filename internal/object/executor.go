package object

import (
	"github.com/qri-io/jsonschema"
	"gopkg.in/yaml.v3"
)

type Executor struct {
	GenericObject

	Script string
	schema jsonschema.Schema
}

var _ Object = Executor{}

func NewExecutor(filename string, data *yaml.Node) (executor Executor, err error) {
	executor.GenericObject.metadata = NewMetadata(filename, data)
	err = decode(data, &executor)
	return
}

func (e Executor) Schema() *jsonschema.Schema {
	return &e.schema
}
