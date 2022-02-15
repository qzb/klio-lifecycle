//go:generate ../../../../scripts/generate-scheme-module.sh

package scheme

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/icza/dyno"
	"github.com/qri-io/jsonschema"
)

func ToInternal(content interface{}) (interface{}, error) {
	return toInternal(dyno.ConvertMapI2MapS(content)), nil
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
