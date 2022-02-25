package stdlib

import (
	logger "github.com/g2a-com/klio-logger-go"
)

func (s *Stdlib) createLogModule() map[string]any {
	return map[string]any{
		"print":    createPrintFunc(s.Logger.WithLevel(logger.InfoLevel)),
		"printf":   createPrintfFunc(s.Logger.WithLevel(logger.InfoLevel)),
		"fatal":    createPrintFunc(s.Logger.WithLevel(logger.FatalLevel)),
		"fatalf":   createPrintfFunc(s.Logger.WithLevel(logger.FatalLevel)),
		"err":      createPrintFunc(s.Logger.WithLevel(logger.ErrorLevel)),
		"error":    createPrintFunc(s.Logger.WithLevel(logger.ErrorLevel)),
		"errorf":   createPrintfFunc(s.Logger.WithLevel(logger.ErrorLevel)),
		"warn":     createPrintFunc(s.Logger.WithLevel(logger.WarnLevel)),
		"warnf":    createPrintfFunc(s.Logger.WithLevel(logger.WarnLevel)),
		"info":     createPrintFunc(s.Logger.WithLevel(logger.InfoLevel)),
		"infof":    createPrintfFunc(s.Logger.WithLevel(logger.InfoLevel)),
		"verbose":  createPrintFunc(s.Logger.WithLevel(logger.VerboseLevel)),
		"verbosef": createPrintfFunc(s.Logger.WithLevel(logger.VerboseLevel)),
		"debug":    createPrintFunc(s.Logger.WithLevel(logger.DebugLevel)),
		"debugf":   createPrintfFunc(s.Logger.WithLevel(logger.DebugLevel)),
		"spam":     createPrintFunc(s.Logger.WithLevel(logger.SpamLevel)),
		"spamf":    createPrintfFunc(s.Logger.WithLevel(logger.SpamLevel)),
	}
}

func createPrintFunc(l *logger.Logger) func(...any) {
	return func(args ...any) {
		l.Print(args...)
	}
}

func createPrintfFunc(l *logger.Logger) func(string, ...any) {
	return func(format string, args ...any) {
		l.Printf(format, args...)
	}
}
