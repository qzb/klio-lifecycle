package fakelogger

import (
	"io"

	logger "github.com/g2a-com/klio-logger-go/v2"
)

type Message struct {
	Level  logger.Level
	Tags   []string
	Method string
	Args   []interface{}
}

type FakeLogger struct {
	level    logger.Level
	tags     []string
	output   io.Writer
	root     *FakeLogger
	Messages []Message
}

func New() *FakeLogger {
	l := &FakeLogger{}
	l.root = l
	return l
}

func (l *FakeLogger) Print(v ...interface{}) logger.Logger {
	l.root.Messages = append(l.root.Messages, Message{
		Level:  l.level,
		Tags:   l.tags,
		Method: "Print",
		Args:   v,
	})
	return l
}

func (l *FakeLogger) Printf(format string, v ...interface{}) logger.Logger {
	l.root.Messages = append(l.root.Messages, Message{
		Level:  l.level,
		Tags:   l.tags,
		Method: "Printf",
		Args:   append([]interface{}{format}, v...),
	})
	return l
}

func (l *FakeLogger) Write(data []byte) (int, error) {
	l.root.Messages = append(l.root.Messages, Message{
		Level:  l.level,
		Tags:   l.tags,
		Method: "Write",
		Args:   []interface{}{data},
	})
	return len(data), nil
}

func (l *FakeLogger) WithLevel(level logger.Level) logger.Logger {
	n := *l
	n.level = level
	return &n
}

func (l *FakeLogger) WithTags(tags ...string) logger.Logger {
	n := *l
	n.tags = tags
	return &n
}

func (l *FakeLogger) WithOutput(output io.Writer) logger.Logger {
	n := *l
	n.output = output
	return &n
}

func (l *FakeLogger) Level() logger.Level {
	return l.level
}

func (l *FakeLogger) Tags() []string {
	return l.tags
}

func (l *FakeLogger) Output() io.Writer {
	return l.output
}
