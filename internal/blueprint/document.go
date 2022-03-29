package blueprint

import (
	"path/filepath"

	"github.com/g2a-com/cicd/internal/object"
	"gopkg.in/yaml.v3"
)

type document struct {
	FilePath    string
	Index       int
	Kind        object.Kind
	Name        string
	APIVersion  string
	DisplayName string
	Object      object.Object
}

func newDocument(filename string, index int, mode Mode, data *yaml.Node) (*document, error) {
	filename, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	obj, err := object.NewObject(filename, data)
	if err != nil {
		return nil, err
	}

	if service, ok := obj.(object.Service); ok {
		if mode != BuildMode {
			service.Build.Artifacts.ToBuild = []object.ServiceEntry{}
			service.Build.Artifacts.ToPush = []object.ServiceEntry{}
			service.Build.Tags = []object.ServiceEntry{}
		}
		if mode != DeployMode {
			service.Deploy.Releases = []object.ServiceEntry{}
		}
		obj = service
	}

	return &document{
		FilePath:    filename,
		Kind:        obj.Kind(),
		Name:        obj.Name(),
		DisplayName: obj.DisplayName(),
		Object:      obj,
	}, nil
}
