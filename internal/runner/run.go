package runner

import (
	"fmt"

	"github.com/g2a-com/klio-lifecycle/internal/blueprint"
	"github.com/g2a-com/klio-lifecycle/internal/object"
)

func run(b *blueprint.Blueprint, service object.Service, entry object.ServiceEntry, kind object.Kind, input map[string]interface{}, runInterpreter Interpreter) (results []Result, err error) {
	executor, ok := b.GetExecutor(kind, entry.Type)
	if !ok {
		return results, fmt.Errorf("builder %q does not exist", entry.Type)
	}

	output, err := runInterpreter(executor.Script, input)
	if err != nil {
		return results, err
	}

	for _, o := range output {
		results = append(results, Result{
			Service: service.Name,
			Entry:   entry.Index,
			Result:  o,
		})
	}

	return results, nil
}
