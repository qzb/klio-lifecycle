package stdlib

import (
	"github.com/d5/tengo/v2"

	"github.com/g2a-com/cicd/internal/tengoutil"
	logger "github.com/g2a-com/klio-logger-go"
	"github.com/spf13/afero"
)

type Stdlib struct {
	Fs               afero.Fs
	Logger           *logger.Logger
	WorkingDirectory string
	Builtins         map[string]any
}

func (s *Stdlib) AddToScript(script *tengo.Script) error {
	// Set imports
	modules := map[string]map[string]any{
		"exec": s.createExecModule(),
		"log":  s.createLogModule(),
	}
	ms := tengo.NewModuleMap()
	for name, attrs := range modules {
		m, err := tengoutil.ToObjectsMap(attrs)
		if err != nil {
			return err
		}
		ms.AddBuiltinModule(name, m)
	}
	script.SetImports(ms)

	// Set builtins
	s.Builtins["abort"] = func(err error) {
		panic(err)
	}
	builtins, err := tengoutil.ToObjectsMap(s.Builtins)
	if err != nil {
		return err
	}
	for name, obj := range builtins {
		err := script.Add(name, obj)
		if err != nil {
			return err
		}
	}

	return nil
}
