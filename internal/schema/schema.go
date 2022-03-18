package schema

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/icza/dyno"
	"github.com/qri-io/jsonschema"
	"gopkg.in/yaml.v3"
)

//
func Migrate(input []byte) ([]byte, error) {
	var data yaml.Node

	err := yaml.Unmarshal(input, &data)
	if err != nil {
		return []byte{}, err
	}

	for _, n := range data.Content {
		err := migrateDocument(n)
		if err != nil {
			return []byte{}, err
		}
	}

	return yaml.Marshal(&data)
}

func migrateDocument(rootNode *yaml.Node) error {
	versionNode := findMapValue(rootNode, "apiVersion")
	if versionNode == nil {
		return &MigrationError{rootNode, "missing apiVersion field"}
	}

	switch versionNode.Value {
	case "v1beta4":
		err := migrateDocumentFromV1Beta4ToV2(rootNode)
		if err != nil {
			return err
		}
		fallthrough
	case "v2.0":
		break
	default:
		return &MigrationError{versionNode, fmt.Sprintf("unsupported version: %s", versionNode.Value)}
	}

	return nil
}

func Validate(apiVersion string, kind string, content interface{}) error {
	ctx := context.Background()

	schema, ok := SCHEMAS[apiVersion+"/"+kind]
	if !ok {
		return fmt.Errorf(`kind "%s" is not supported by api version "%s"`, kind, apiVersion)
	}

	validator := &jsonschema.Schema{}
	if err := json.Unmarshal(schema, validator); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	result := validator.Validate(ctx, dyno.ConvertMapI2MapS(content))
	if len(*result.Errs) > 0 {
		var err error
		for _, e := range *result.Errs {
			err = multierror.Append(err, errors.New(e.Error()))
		}
		return err
	}

	return nil
}
