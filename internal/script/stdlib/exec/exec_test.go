package exec

import (
	"errors"
	"testing"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/tengoutil"
	fakelogger "github.com/g2a-com/cicd/internal/utils/fake_logger"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/exec"
	testingexec "k8s.io/utils/exec/testing"
)

func Test_run_runs_command(t *testing.T) {
	cmd := prepareFakeCmd("stdout", "stderr", 0)
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	result, err := run(mod, `exec.run("echo", "abc", "123")`)

	assert.NoError(t, err)
	assert.Equal(t, []string{"echo", "abc", "123"}, cmd.Argv)
	assert.Equal(t, map[string]interface{}{
		"stdout_text": "stdout",
		"stderr_text": "stderr",
		"error":       nil,
		"exit_code":   0,
	}, result)
}

func Test_run_prints_logs_at_info_and_error_levels(t *testing.T) {
	log := fakelogger.New()
	cmd := prepareFakeCmd("stdout", "stderr", 0)
	mod := New(log)
	mod.exec = prepareFakeExec(cmd)

	_, _ = run(mod, `exec.run("cmd")`)

	assert.Equal(t, []fakelogger.Message{
		{Level: "info", Method: "Write", Args: []interface{}{[]byte("stdout")}},
		{Level: "error", Method: "Write", Args: []interface{}{[]byte("stderr")}},
	}, log.Messages)
}

func Test_run_aborts_script_on_error(t *testing.T) {
	cmd := prepareFakeCmd("stdout", "stderr", 1)
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	_, err := run(mod, `exec.run("cmd")`)

	assert.Error(t, err)
}

func Test_run_silently_runs_command(t *testing.T) {
	cmd := prepareFakeCmd("stdout", "stderr", 0)
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	result, err := run(mod, `exec.run_silently("echo", "abc", "123")`)

	assert.NoError(t, err)
	assert.Equal(t, []string{"echo", "abc", "123"}, cmd.Argv)
	assert.Equal(t, map[string]interface{}{
		"stdout_text": "stdout",
		"stderr_text": "stderr",
		"error":       nil,
		"exit_code":   0,
	}, result)
}

func Test_run_silently_prints_logs_at_debug_and_error_levels(t *testing.T) {
	log := fakelogger.New()
	cmd := prepareFakeCmd("stdout", "stderr", 0)
	mod := New(log)
	mod.exec = prepareFakeExec(cmd)

	_, _ = run(mod, `exec.run_silently("cmd")`)

	assert.Equal(t, []fakelogger.Message{
		{Level: "debug", Method: "Write", Args: []interface{}{[]byte("stdout")}},
		{Level: "error", Method: "Write", Args: []interface{}{[]byte("stderr")}},
	}, log.Messages)
}

func Test_run_silently_aborts_script_on_error(t *testing.T) {
	cmd := prepareFakeCmd("stdout", "stderr", 1)
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	_, err := run(mod, `exec.run_silently("cmd")`)

	assert.Error(t, err)
}

func Test_command_run_runs_command(t *testing.T) {
	cmd := prepareFakeCmd("stdout", "stderr", 0)
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	result, err := run(mod, `exec.command({ name: "echo", args: [ "abc", "123" ] }).run()`)

	assert.NoError(t, err)
	assert.Equal(t, []string{"echo", "abc", "123"}, cmd.Argv)
	assert.Equal(t, map[string]interface{}{
		"stdout_text": "stdout",
		"stderr_text": "stderr",
		"error":       nil,
		"exit_code":   0,
	}, result)
}

func Test_command_run_by_default_prints_logs_at_info_and_error_levels(t *testing.T) {
	log := fakelogger.New()
	cmd := prepareFakeCmd("stdout", "stderr", 0)
	mod := New(log)
	mod.exec = prepareFakeExec(cmd)

	_, _ = run(mod, `exec.command({ name: "cmd" }).run()`)

	assert.Equal(t, []fakelogger.Message{
		{Level: "info", Method: "Write", Args: []interface{}{[]byte("stdout")}},
		{Level: "error", Method: "Write", Args: []interface{}{[]byte("stderr")}},
	}, log.Messages)
}

func Test_command_run_prints_logs_at_specified_levels(t *testing.T) {
	log := fakelogger.New()
	cmd := prepareFakeCmd("stdout", "stderr", 0)
	mod := New(log)
	mod.exec = prepareFakeExec(cmd)

	_, _ = run(mod, `exec.command({ name: "cmd", stdout_level: "verbose", stderr_level: "warn" }).run()`)

	assert.Equal(t, []fakelogger.Message{
		{Level: "verbose", Method: "Write", Args: []interface{}{[]byte("stdout")}},
		{Level: "warn", Method: "Write", Args: []interface{}{[]byte("stderr")}},
	}, log.Messages)
}

func Test_command_run_by_default_aborts_script_on_error(t *testing.T) {
	cmd := prepareFakeCmd("stdout", "stderr", 1)
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	_, err := run(mod, `exec.command({ name: "cmd" }).run()`)

	assert.Error(t, err)
}

func Test_command_run_returns_result_with_error_when_ignore_errors_option_is_enabled(t *testing.T) {
	cmd := prepareFakeCmd("stdout", "stderr", 1)
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	result, err := run(mod, `exec.command({ name: "cmd", ignore_errors: true }).run()`)

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"stdout_text": "stdout",
		"stderr_text": "stderr",
		"error":       "exit code != 0", // This message is set by prepareFakeCmd function
		"exit_code":   1,
	}, result)
}

func Test_command_run_returns_result_with_minus_one_exit_code_when_command_cannot_be_runned(t *testing.T) {
	cmd := &testingexec.FakeCmd{RunScript: []testingexec.FakeAction{
		func() ([]byte, []byte, error) {
			return []byte{}, []byte{}, errors.New("command not found")
		},
	}}
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	result, err := run(mod, `exec.command({ name: "cmd", ignore_errors: true }).run()`)

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"stdout_text": "",
		"stderr_text": "",
		"error":       "command not found",
		"exit_code":   -1,
	}, result)
}

func Test_command_run_runs_command_in_specified_directory(t *testing.T) {
	cmd := prepareFakeCmd("stdout", "stderr", 0)
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	_, err := run(mod, `exec.command({ name: "cmd", dir: "/some/dir" }).run()`)

	assert.NoError(t, err)
	assert.Equal(t, []string{"/some/dir"}, cmd.Dirs)
}

func Test_command_run_runs_command_with_specified_env_variables(t *testing.T) {
	cmd := prepareFakeCmd("stdout", "stderr", 0)
	mod := New(fakelogger.New())
	mod.exec = prepareFakeExec(cmd)

	_, err := run(mod, `exec.command({ name: "cmd", env: ["FOO=bar", "egg=spam"] }).run()`)

	assert.NoError(t, err)
	assert.Equal(t, []string{"FOO=bar", "egg=spam"}, cmd.Env)
}

func run(m *module, code string) (result interface{}, err error) {
	modules := tengo.NewModuleMap()
	modules.Add("exec", m)
	script := tengo.NewScript([]byte(`exec := import("exec"); result := ` + code))
	script.SetImports(modules)
	compiled, err := script.Run()
	if err == nil {
		err = tengoutil.DecodeObject(compiled.Get("result").Object(), &result)
	}
	return
}

func prepareFakeExec(commands ...*testingexec.FakeCmd) *testingexec.FakeExec {
	actions := make([]testingexec.FakeCommandAction, len(commands))
	for i := range actions {
		actions[i] = func(cmd string, args ...string) exec.Cmd {
			return testingexec.InitFakeCmd(commands[i], cmd, args...)
		}
	}
	return &testingexec.FakeExec{CommandScript: actions}
}

func prepareFakeCmd(stdout string, stderr string, exitCode int) *testingexec.FakeCmd {
	return &testingexec.FakeCmd{
		RunScript: []testingexec.FakeAction{
			func() ([]byte, []byte, error) {
				var err error
				if exitCode != 0 {
					err = exec.CodeExitError{Code: exitCode, Err: errors.New("exit code != 0")}
				}
				return []byte(stdout), []byte(stderr), err
			},
		},
	}
}
