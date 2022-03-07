package script

import (
	"testing"

	"github.com/g2a-com/cicd/internal/object"
	fakelogger "github.com/g2a-com/cicd/internal/utils/fake_logger"
	"github.com/stretchr/testify/assert"
)

func Test_input_is_passed_down_to_script(t *testing.T) {
	log := fakelogger.New()
	executor := object.Executor{Script: `import("log").print(input)`}
	script := New(executor)
	script.Logger = log

	_, err := script.Run(map[string]any{
		"foo": "bar",
	})

	assert.NoError(t, err)
	assert.Contains(t, log.Messages, fakelogger.Message{Level: "info", Method: "Print", Args: []any{`{foo: "bar"}`}})
}

func Test_returns_results_added_by_script(t *testing.T) {
	executor := object.Executor{Script: `addResult("a", "b"); addResult("c")`}
	script := New(executor)
	script.Logger = fakelogger.New()

	result, err := script.Run(nil)

	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func Test_returns_error_when_script_has_invalid_syntax(t *testing.T) {
	executor := object.Executor{Script: `if`}
	script := New(executor)
	script.Logger = fakelogger.New()

	_, err := script.Run(nil)

	assert.Error(t, err)
}

func Test_returns_error_when_script_is_aborted(t *testing.T) {
	executor := object.Executor{Script: `abort("error")`}
	script := New(executor)
	script.Logger = fakelogger.New()

	_, err := script.Run(nil)

	assert.Error(t, err)
}
