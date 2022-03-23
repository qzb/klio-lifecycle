package schema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validating_valid_document(t *testing.T) {
	input := []byte(`{ apiVersion: g2a-cli/v2.0, kind: Project, name: test }`)

	result, err := Validate(input)

	assert.NoError(t, err)
	assert.Equal(t, input, result)
}

func Test_validating_multiple_valid_documents(t *testing.T) {
	content := `{ apiVersion: g2a-cli/v2.0, kind: Project, name: test }`
	content += "\n---\n"
	content += `{ apiVersion: g2a-cli/v2.0, kind: Service, name: test }`
	input := []byte(content)

	result, err := Validate(input)

	assert.NoError(t, err)
	assert.Equal(t, input, result)
}

func Test_validating_invalid_document_returns_error(t *testing.T) {
	input := []byte(`{ apiVersion: g2a-cli/v2.0, kind: Project, name: 7 }`)

	_, err := Validate(input)

	assert.Error(t, err)
}

func Test_validating_document_without_apiVersion_returns_error(t *testing.T) {
	input := []byte(`{ kind: Project, name: test }`)

	_, err := Validate(input)

	assert.Error(t, err)
}

func Test_validating_document_without_kind_returns_error(t *testing.T) {
	input := []byte(`{ apiVersion: g2a-cli/v2.0, name: test }`)

	_, err := Validate(input)

	assert.Error(t, err)
}

func Test_validating_document_with_invalid_apiVersion_and_kind_combination_returns_error(t *testing.T) {
	cases := [][2]string{
		{`g2a-cli/v1beta4`, `Builder`},
		{`g2a-cli/v1beta4`, `Invalid`},
		{`g2a-cli/v2.0`, `Invalid`},
	}
	for _, c := range cases {
		t.Run(c[0]+"/"+c[1], func(t *testing.T) {
			input := []byte(fmt.Sprintf(`{ apiVersion: %s, kind: %s, name: test }`, c[0], c[1]))

			_, err := Validate(input)

			assert.Error(t, err)
		})
	}
}
