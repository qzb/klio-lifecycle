package stdlib

import (
	"fmt"
	"testing"

	"github.com/d5/tengo/v2"
	fakelogger "github.com/g2a-com/cicd/internal/utils/fake_logger"
	"github.com/stretchr/testify/assert"
)

func Test_builtin_abort_function_aborts_script_with_specified_error(t *testing.T) {
	stdlib := New(fakelogger.New())

	script := tengo.NewScript([]byte(`abort("error message")`))
	stdlib.InitializeScript(script)
	_, err := script.Run()

	assert.ErrorIs(t, err, &AbortError{"error message"})
}

func Test_builtin_abort_function_accepts_tengo_errors(t *testing.T) {
	stdlib := New(fakelogger.New())

	script := tengo.NewScript([]byte(`abort(error(7))`))
	stdlib.InitializeScript(script)
	_, err := script.Run()

	assert.ErrorIs(t, err, &AbortError{7})
}

func Test_all_modules_can_be_imported(t *testing.T) {
	modules := []string{"log", "exec"}
	for _, module := range modules {
		t.Run(module, func(t *testing.T) {
			stdlib := New(fakelogger.New())

			script := tengo.NewScript([]byte(fmt.Sprintf(`import("%s")`, module)))
			stdlib.InitializeScript(script)
			_, err := script.Run()

			assert.NoError(t, err)
		})
	}
}

func Test_additional_builtins_are_available_in_script(t *testing.T) {
	script := tengo.NewScript([]byte{})
	stdlib := New(fakelogger.New())

	errAdd := stdlib.AddBuiltin("test", 7)
	stdlib.InitializeScript(script)
	compiled, errRun := script.Run()

	assert.NoError(t, errAdd)
	assert.NoError(t, errRun)
	assert.Equal(t, 7, compiled.Get("test").Int())
}

func Test_logger_is_passed_down_to_modules(t *testing.T) {
	log := fakelogger.New()
	stdlib := New(log)

	script := tengo.NewScript([]byte(`import("log").print("test")`))
	stdlib.InitializeScript(script)
	_, err := script.Run()

	assert.NoError(t, err)
	assert.Len(t, log.Messages, 1)
}
