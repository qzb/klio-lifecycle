package schema

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/hashicorp/go-multierror"
	"github.com/icza/dyno"
	"github.com/qri-io/jsonschema"
	"gopkg.in/yaml.v2"
)

func Validate(input []byte) ([]byte, error) {
	reader := bytes.NewReader(input)
	decoder := yaml.NewDecoder(reader)

	for {
		var content interface{}

		err := decoder.Decode(&content)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		apiVersion, err := dyno.GetString(content, "apiVersion")
		if err != nil {
			return nil, errors.New("missing apiVersion")
		}
		kind, err := dyno.GetString(content, "kind")
		if err != nil {
			return nil, errors.New("missing kind")
		}

		ctx := context.Background()

		schema, ok := SCHEMAS[apiVersion+"/"+kind]
		if !ok {
			return nil, fmt.Errorf("kind %q is not supported by api version %q", kind, apiVersion)
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
			return nil, err
		}
	}

	return input, nil
}
