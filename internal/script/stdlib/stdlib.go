package stdlib

import (
	"fmt"

	"github.com/d5/tengo/v2"
	execModule "github.com/g2a-com/cicd/internal/script/stdlib/exec"
	logModule "github.com/g2a-com/cicd/internal/script/stdlib/log"
	"github.com/g2a-com/cicd/internal/tengoutil"
	logger "github.com/g2a-com/klio-logger-go/v2"
)

type stdlib struct {
	logger   logger.Logger
	builtins map[string]interface{}
}

func New(l logger.Logger) *stdlib {
	abort, _ := tengoutil.ToObject(func(err interface{}) {
		panic(&AbortError{err})
	})

	return &stdlib{
		logger: l,
		builtins: map[string]interface{}{
			"abort": abort,
		},
	}
}

func (s *stdlib) AddBuiltin(name string, value interface{}) (err error) {
	if _, ok := s.builtins[name]; ok {
		return fmt.Errorf("builtin %q is already registered in standard library", name)
	}
	s.builtins[name], err = tengoutil.ToObject(value)
	return err
}

func (s *stdlib) InitializeScript(script *tengo.Script) error {
	// Set imports
	mm := tengo.NewModuleMap()
	mm.Add("exec", execModule.New(s.logger))
	mm.Add("log", logModule.New(s.logger))
	script.SetImports(mm)

	// Set builtins
	for name, obj := range s.builtins {
		err := script.Add(name, obj)
		if err != nil {
			return err
		}
	}

	return nil
}

type AbortError struct {
	value interface{}
}

func (e *AbortError) Error() string {
	return fmt.Sprint(e.value)
}

func (e *AbortError) Is(err error) bool {
	if ae, ok := err.(*AbortError); ok {
		return ae.value == e.value
	}
	return false
}
