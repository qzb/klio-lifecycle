package object

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

func NewObject(kind Kind, dir string, content interface{}) (interface{}, error) {
	switch kind {
	case ProjectKind:
		obj := Project{Directory: dir}
		err := decode(content, &obj)
		return obj, err
	case ServiceKind:
		obj := Service{Directory: dir}
		err := decode(content, &obj)
		return obj, err
	case EnvironmentKind:
		obj := Environment{Directory: dir}
		err := decode(content, &obj)
		return obj, err
	case BuilderKind, DeployerKind, PusherKind, TaggerKind:
		obj := Executor{Directory: dir}
		err := decode(content, &obj)
		return obj, err
	default:
		return nil, fmt.Errorf("unknown kind %s", kind)
	}
}

func decode(content interface{}, result interface{}) error {
	decoderConfig := &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      result,
		DecodeHook: func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
			if f.Kind() != reflect.String {
				return data, nil
			}
			result := reflect.New(t).Interface()
			unmarshaller, ok := result.(json.Unmarshaler)
			if !ok {
				return data, nil
			}
			if err := unmarshaller.UnmarshalJSON([]byte(data.(string))); err != nil {
				return nil, err
			}
			return result, nil
		},
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return err
	}

	return decoder.Decode(content)
}
