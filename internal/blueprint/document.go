package blueprint

import (
	"fmt"
	"path/filepath"

	"github.com/g2a-com/cicd/internal/blueprint/internal/scheme"
	"github.com/g2a-com/cicd/internal/object"
	"github.com/icza/dyno"
)

type document struct {
	FilePath    string
	Index       int
	Kind        object.Kind
	Name        string
	APIVersion  string
	DisplayName string
	Object      interface{}
}

func newDocument(filePath string, index int, mode Mode, content interface{}) (*document, error) {
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	apiVersion, err := dyno.GetString(content, "apiVersion")
	if err != nil {
		return nil, err
	}

	kind, err := dyno.GetString(content, "kind")
	if err != nil {
		return nil, err
	}

	name, err := dyno.GetString(content, "name")
	if err != nil && kind != "Project" {
		return nil, err
	}

	normalized, err := scheme.ToInternal(content)
	if err != nil {
		return nil, err
	}

	obj, err := object.NewObject(object.Kind(kind), filepath.Dir(filePath), normalized)
	if err != nil {
		return nil, err
	}

	if kind == string(object.ServiceKind) {
		service := obj.(object.Service)
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
		FilePath:    filePath,
		APIVersion:  apiVersion,
		Kind:        object.Kind(kind),
		Name:        name,
		DisplayName: fmt.Sprintf("%s %q", kind, name),
		Object:      obj,
	}, nil
}
