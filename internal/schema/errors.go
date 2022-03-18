package schema

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type MigrationError struct {
	Node    *yaml.Node
	Message string
}

func (e *MigrationError) Error() string {
	return fmt.Sprintf("error in line %d: %s", e.Node.Line, e.Message)
}
