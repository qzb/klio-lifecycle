package log

import (
	"testing"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/tengoutil"
	fakelogger "github.com/g2a-com/cicd/internal/utils/fake_logger"
	"github.com/stretchr/testify/assert"
)

func Test_log_print_writes_formatted_message_to_logger_at_info_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.print("foo", "bar", 2000, true, { "foo": "bar" }, [ 1, 2, 3 ])`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "info", Method: "Print", Args: []interface{}{`foobar 2000 true {foo: "bar"} [1, 2, 3]`}}}, l.Messages)
}

func Test_log_printf_writes_formatted_message_to_logger_at_info_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.printf("%o", 8)`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "info", Method: "Print", Args: []interface{}{`10`}}}, l.Messages)
}

func Test_log_spam_writes_formatted_message_to_logger_at_spam_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.spam("foo", "bar", 2000, true, { "foo": "bar" }, [ 1, 2, 3 ])`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "spam", Method: "Print", Args: []interface{}{`foobar 2000 true {foo: "bar"} [1, 2, 3]`}}}, l.Messages)
}

func Test_log_spamf_writes_formatted_message_to_logger_at_spam_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.spamf("%o", 8)`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "spam", Method: "Print", Args: []interface{}{`10`}}}, l.Messages)
}

func Test_log_debug_writes_formatted_message_to_logger_at_debug_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.debug("foo", "bar", 2000, true, { "foo": "bar" }, [ 1, 2, 3 ])`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "debug", Method: "Print", Args: []interface{}{`foobar 2000 true {foo: "bar"} [1, 2, 3]`}}}, l.Messages)
}

func Test_log_debugf_writes_formatted_message_to_logger_at_debug_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.debugf("%o", 8)`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "debug", Method: "Print", Args: []interface{}{`10`}}}, l.Messages)
}

func Test_log_verbose_writes_formatted_message_to_logger_at_verbose_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.verbose("foo", "bar", 2000, true, { "foo": "bar" }, [ 1, 2, 3 ])`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "verbose", Method: "Print", Args: []interface{}{`foobar 2000 true {foo: "bar"} [1, 2, 3]`}}}, l.Messages)
}

func Test_log_verbosef_writes_formatted_message_to_logger_at_verbose_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.verbosef("%o", 8)`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "verbose", Method: "Print", Args: []interface{}{`10`}}}, l.Messages)
}

func Test_log_info_writes_formatted_message_to_logger_at_info_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.info("foo", "bar", 2000, true, { "foo": "bar" }, [ 1, 2, 3 ])`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "info", Method: "Print", Args: []interface{}{`foobar 2000 true {foo: "bar"} [1, 2, 3]`}}}, l.Messages)
}

func Test_log_infof_writes_formatted_message_to_logger_at_info_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.infof("%o", 8)`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "info", Method: "Print", Args: []interface{}{`10`}}}, l.Messages)
}

func Test_log_warn_writes_formatted_message_to_logger_at_warn_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.warn("foo", "bar", 2000, true, { "foo": "bar" }, [ 1, 2, 3 ])`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "warn", Method: "Print", Args: []interface{}{`foobar 2000 true {foo: "bar"} [1, 2, 3]`}}}, l.Messages)
}

func Test_log_warnf_writes_formatted_message_to_logger_at_warn_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.warnf("%o", 8)`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "warn", Method: "Print", Args: []interface{}{`10`}}}, l.Messages)
}

func Test_log_err_writes_formatted_message_to_logger_at_error_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.err("foo", "bar", 2000, true, { "foo": "bar" }, [ 1, 2, 3 ])`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "error", Method: "Print", Args: []interface{}{`foobar 2000 true {foo: "bar"} [1, 2, 3]`}}}, l.Messages)
}

func Test_log_error_writes_formatted_message_to_logger_at_error_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log["error"]("foo", "bar", 2000, true, { "foo": "bar" }, [ 1, 2, 3 ])`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "error", Method: "Print", Args: []interface{}{`foobar 2000 true {foo: "bar"} [1, 2, 3]`}}}, l.Messages)
}

func Test_log_errorf_writes_formatted_message_to_logger_at_error_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.errorf("%o", 8)`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "error", Method: "Print", Args: []interface{}{`10`}}}, l.Messages)
}

func Test_log_fatal_writes_formatted_message_to_logger_at_fatal_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.fatal("foo", "bar", 2000, true, { "foo": "bar" }, [ 1, 2, 3 ])`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "fatal", Method: "Print", Args: []interface{}{`foobar 2000 true {foo: "bar"} [1, 2, 3]`}}}, l.Messages)
}

func Test_log_fatalf_writes_formatted_message_to_logger_at_fatal_level(t *testing.T) {
	l := fakelogger.New()
	m := New(l)

	_, err := run(m, `log.fatalf("%o", 8)`)

	assert.NoError(t, err)
	assert.Equal(t, []fakelogger.Message{{Level: "fatal", Method: "Print", Args: []interface{}{`10`}}}, l.Messages)
}

func run(m *module, code string) (result interface{}, err error) {
	modules := tengo.NewModuleMap()
	modules.Add("log", m)
	script := tengo.NewScript([]byte(`log := import("log"); result := ` + code))
	script.SetImports(modules)
	compiled, err := script.Run()
	if err == nil {
		err = tengoutil.DecodeObject(compiled.Get("result").Object(), &result)
	}
	return
}
