package log

import (
	"github.com/d5/tengo/v2"
	"github.com/g2a-com/cicd/internal/tengoutil"
	logger "github.com/g2a-com/klio-logger-go/v2"
)

type module struct {
	logger logger.Logger
}

func New(logger logger.Logger) *module {
	return &module{
		logger: logger,
	}
}

func (m *module) Import(name string) (any, error) {
	return tengoutil.ToImmutableObject(map[string]any{
		"__module_name__": name,
		"print":           createPrintFunc(m.logger.WithLevel(logger.InfoLevel)),
		"printf":          createPrintfFunc(m.logger.WithLevel(logger.InfoLevel)),
		"fatal":           createPrintFunc(m.logger.WithLevel(logger.FatalLevel)),
		"fatalf":          createPrintfFunc(m.logger.WithLevel(logger.FatalLevel)),
		"err":             createPrintFunc(m.logger.WithLevel(logger.ErrorLevel)),
		"error":           createPrintFunc(m.logger.WithLevel(logger.ErrorLevel)),
		"errorf":          createPrintfFunc(m.logger.WithLevel(logger.ErrorLevel)),
		"warn":            createPrintFunc(m.logger.WithLevel(logger.WarnLevel)),
		"warnf":           createPrintfFunc(m.logger.WithLevel(logger.WarnLevel)),
		"info":            createPrintFunc(m.logger.WithLevel(logger.InfoLevel)),
		"infof":           createPrintfFunc(m.logger.WithLevel(logger.InfoLevel)),
		"verbose":         createPrintFunc(m.logger.WithLevel(logger.VerboseLevel)),
		"verbosef":        createPrintfFunc(m.logger.WithLevel(logger.VerboseLevel)),
		"debug":           createPrintFunc(m.logger.WithLevel(logger.DebugLevel)),
		"debugf":          createPrintfFunc(m.logger.WithLevel(logger.DebugLevel)),
		"spam":            createPrintFunc(m.logger.WithLevel(logger.SpamLevel)),
		"spamf":           createPrintfFunc(m.logger.WithLevel(logger.SpamLevel)),
	})
}

func createPrintFunc(l logger.Logger) func(...tengo.Object) {
	return func(args ...tengo.Object) {
		msg := ""
		for _, arg := range args {
			if arg.TypeName() != "string" && msg != "" {
				msg += " "
			}
			str, _ := tengo.ToString(arg)
			msg += str
		}
		l.Print(msg)
	}
}

func createPrintfFunc(l logger.Logger) func(string, ...tengo.Object) {
	return func(format string, args ...tengo.Object) {
		msg, _ := tengo.Format(format, args...)
		l.Print(msg)
	}
}
