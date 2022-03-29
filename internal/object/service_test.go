package object

import (
	"testing"

	"github.com/g2a-com/cicd/internal/schema"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_unmarshalling_empty_service(t *testing.T) {
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0,
		kind: Service,
		name: test,
	}`)

	result, err := NewService("dir/file.yaml", input)

	assert.NoError(t, err)
	assert.Equal(t, "dir/file.yaml", result.Metadata().Filename())
	assert.Equal(t, ServiceKind, result.Kind())
	assert.Equal(t, "test", result.Name())
	assert.Equal(t, "dir", result.Directory())
	assert.Equal(t, `service "test"`, result.DisplayName())
}

func Test_validating_empty_service_passes(t *testing.T) {
	collection := testCollection([]Object{})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0,
		kind: Service,
		name: test,
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.NoError(t, err)
}

func Test_validating_service_using_unknown_tagger_fails(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: TaggerKind, name: "known", schema: "{}"},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		tags: [{ unknown: {} }],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.Error(t, err)
}

func Test_validating_service_using_known_tagger_passes(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: TaggerKind, name: "known", schema: "{}"},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		tags: [{ known: {} }],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.NoError(t, err)
}

func Test_validating_service_using_unknown_builder_fails(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: BuilderKind, name: "known", schema: "{}"},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		artifacts: [{
			unknown: {},
			push: false,
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.Error(t, err)
}

func Test_validating_service_using_known_builder_passes(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: BuilderKind, name: "known", schema: "{}"},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		artifacts: [{
			known: {},
			push: false,
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.NoError(t, err)
}

func Test_validating_service_using_unknown_pusher_fails(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: BuilderKind, name: "known", schema: "{}"},
		fakeObject{kind: PusherKind, name: "known", schema: "{}"},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		artifacts: [{
			known: {},
			push: { unknown: {} }
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.Error(t, err)
}

func Test_validating_service_using_known_pusher_passes(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: BuilderKind, name: "known", schema: "{}"},
		fakeObject{kind: PusherKind, name: "known", schema: "{}"},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		artifacts: [{
			known: {},
			push: { known: {} }
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.NoError(t, err)
}

func Test_validating_service_using_unknown_deployer_fails(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: DeployerKind, name: "known", schema: "{}"},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		releases: [ { unknown: {} } ],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.Error(t, err)
}

func Test_validating_service_using_known_deployer_passes(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: DeployerKind, name: "known", schema: "{}"},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		releases: [ { known: {} } ],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.NoError(t, err)
}

func Test_validating_service_with_tags_entry_not_matching_tagger_schema_fails(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: TaggerKind, name: "type", schema: `{ "required": ["foo"] }`},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		tags: [{
			type: {},
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.Error(t, err)
}

func Test_validating_service_with_tags_entry_matching_tagger_schema_passes(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: TaggerKind, name: "type", schema: `{ "required": ["foo"] }`},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		tags: [{
			type: { foo: true },
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.NoError(t, err)
}

func Test_validating_service_with_artifacts_entry_not_matching_builder_schema_fails(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: BuilderKind, name: "type", schema: `{ "required": ["foo"] }`},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		artifacts: [{
			type: {},
			push: false,
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.Error(t, err)
}

func Test_validating_service_with_artifacts_entry_matching_builder_schema_passes(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: BuilderKind, name: "type", schema: `{ "required": ["foo"] }`},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		artifacts: [{
			type: { foo: true },
			push: false,
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.NoError(t, err)
}

func Test_validating_service_with_artifacts_entry_not_matching_pusher_schema_fails(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: BuilderKind, name: "type", schema: `{}`},
		fakeObject{kind: PusherKind, name: "type", schema: `{ "required": ["foo"] }`},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		artifacts: [{
			type: {},
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.Error(t, err)
}

func Test_validating_service_with_artifacts_entry_matching_pusher_schema_passes(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: BuilderKind, name: "type", schema: `{}`},
		fakeObject{kind: PusherKind, name: "type", schema: `{ "required": ["foo"] }`},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		artifacts: [{
			type: { foo: true },
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.NoError(t, err)
}

func Test_validating_service_with_releases_entry_not_matching_deployer_schema_fails(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: DeployerKind, name: "type", schema: `{ "required": ["foo"] }`},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		releases: [{
			type: {},
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.Error(t, err)
}

func Test_validating_service_with_releases_entry_matching_deployer_schema_passes(t *testing.T) {
	collection := testCollection([]Object{
		fakeObject{kind: DeployerKind, name: "type", schema: `{ "required": ["foo"] }`},
	})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0, kind: Service, name: test,
		releases: [{
			type: { foo: true },
		}],
	}`)

	service, _ := NewService("dir/file.yaml", input)
	err := service.Validate(collection)

	assert.NoError(t, err)
}

// testInput validates input against schema and returns it back. Use only in
// tests.
func prepareTestInput(input string) *yaml.Node {
	_, err := schema.Validate([]byte(input))
	if err != nil {
		panic(err)
	}
	result := &yaml.Node{}
	err = yaml.Unmarshal([]byte(input), result)
	if err != nil {
		panic(err)
	}
	return result
}

type testCollection []Object

func (c testCollection) GetObject(kind Kind, name string) Object {
	for _, o := range c {
		if o.Kind() == kind && o.Name() == name {
			return o
		}
	}
	return nil
}
