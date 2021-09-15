package placeholders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplacePlaceholdersSingleVar(t *testing.T) {
	replacements := map[string]string{
		".Envs.ResourcesNamespace": "test1",
	}

	newString, err := replacePlaceholders("../../scripts/helm/create-mongodb-secrets -rn {{ .Envs.ResourcesNamespace }}", replacements)

	assert.NoError(t, err)
	assert.Equal(t, newString, "../../scripts/helm/create-mongodb-secrets -rn test1")
}

func TestReplacePlaceholdersMultipleVars(t *testing.T) {
	replacements := map[string]string{
		".Envs.ResourcesNamespace": "test1",
		".Envs.Namespace":          "test2",
	}

	newString, err := replacePlaceholders("../../scripts/helm/create-mongodb-secrets -rn {{ .Envs.ResourcesNamespace }} -n {{ .Envs.Namespace }}", replacements)

	assert.NoError(t, err)
	assert.Equal(t, newString, "../../scripts/helm/create-mongodb-secrets -rn test1 -n test2")
}

func TestReplacePlaceholdersShouldReturnErrorWhenVariableIsMissing(t *testing.T) {
	replacements := map[string]string{
		".Envs.ResourcesNamespace": "test1",
	}

	newString, err := replacePlaceholders("../../scripts/helm/create-mongodb-secrets -rn {{ .Envs.ResourcesNamespace }} -n {{ .Envs.Namespace }}", replacements)

	assert.Equal(t, "", newString)
	assert.Error(t, assert.AnError, err)
}

func TestReplacePlaceholdersShouldHandleRecursivePlaceholders(t *testing.T) {
	replacements := map[string]string{
		".Test1": "{{ .Test2 }}",
		".Test2": "test2",
	}

	newString, err := replacePlaceholders("{{ .Test1 }}, {{ .Test2 }}", replacements)

	assert.Equal(t, "test2, test2", newString)
	assert.NoError(t, err)
}

func TestReplacePlaceholdersShouldReturnErrrorForCyclicPlaceholders(t *testing.T) {
	replacements := map[string]string{
		".Test1": "{{ .Test2 }}",
		".Test2": "{{ .Test1 }}",
	}

	newString, err := replacePlaceholders("{{ .Test1 }}, {{ .Test2 }}", replacements)

	assert.Equal(t, "", newString)
	assert.Error(t, assert.AnError, err)
}
