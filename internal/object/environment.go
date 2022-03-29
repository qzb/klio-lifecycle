package object

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v3"
)

type Environment struct {
	GenericObject

	DeployServices []string
	Variables      map[string]string
}

var _ Object = Environment{}

func NewEnvironment(filename string, data *yaml.Node) (environment Environment, err error) {
	environment.GenericObject.metadata = NewMetadata(filename, data)
	err = decode(data, &environment)
	return
}

func (e Environment) Validate(c ObjectCollection) (err error) {
	for _, name := range e.DeployServices {
		if c.GetObject(ServiceKind, name) == nil {
			err = multierror.Append(err, fmt.Errorf("missing service %q deployed to environment %q defined in the file:\n\t  %s", name, e.Name(), e.Metadata().Filename()))
		}
	}
	return
}

func (e Environment) Entries(string) []ServiceEntry {
	return []ServiceEntry{}
}
