package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_creating_new_metadata_instance(t *testing.T) {
	input := prepareTestInput(`{
		apiVersion: g2a-cli/v2.0,
		kind: Project,
		name: test,
	}`)

	result := NewMetadata("dir/file.yaml", input)

	assert.Equal(t, "dir/file.yaml", result.Filename())
	assert.Equal(t, 1, result.Line())
}
