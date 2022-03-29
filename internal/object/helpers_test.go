package object

import "github.com/qri-io/jsonschema"

type fakeObject struct {
	metadata    metadata
	kind        Kind
	name        string
	directory   string
	displayName string
	schema      string
}

func (o fakeObject) Name() string {
	return o.name
}

func (o fakeObject) Kind() Kind {
	return o.kind
}

func (o fakeObject) Metadata() Metadata {
	return o.metadata
}

func (o fakeObject) Directory() string {
	return o.directory
}

func (o fakeObject) DisplayName() string {
	return o.displayName
}

func (o fakeObject) Validate(ObjectCollection) error {
	return nil
}

func (o fakeObject) Schema() *jsonschema.Schema {
	return jsonschema.Must(o.schema)
}
