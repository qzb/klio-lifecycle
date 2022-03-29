package object

import (
	"gopkg.in/yaml.v3"
)

type Project struct {
	GenericObject

	Files     []string
	Variables map[string]string
}

var _ Object = Project{}

func NewProject(filename string, data *yaml.Node) (project Project, err error) {
	project.GenericObject.metadata = NewMetadata(filename, data)
	err = decode(data, &project)
	return
}
