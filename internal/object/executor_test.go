package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unmarshalling_empty_executor(t *testing.T) {
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0,
		kind: Tagger,
		name: test,
		script: "",
	}`)

	result, err := NewExecutor("dir/file.yaml", input)

	assert.NoError(t, err)
	assert.Equal(t, "dir/file.yaml", result.Metadata().Filename())
	assert.Equal(t, TaggerKind, result.Kind())
	assert.Equal(t, "test", result.Name())
	assert.Equal(t, "dir", result.Directory())
	assert.Equal(t, `tagger "test"`, result.DisplayName())
}

func Test_validating_empty_executor_passes(t *testing.T) {
	collection := testCollection([]Object{})
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0,
		kind: Tagger,
		name: test,
		script: "",
	}`)

	executor, _ := NewExecutor("dir/file.yaml", input)
	err := executor.Validate(collection)

	assert.NoError(t, err)
}
