package object

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/g2a-com/cicd/internal/object/internal/scheme"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Object interface {
	Metadata() Metadata
	Name() string
	Kind() Kind
	DisplayName() string
	Directory() string
	Validate(ObjectCollection) error
}

type ObjectCollection interface {
	GetObject(kind Kind, name string) Object
}

type GenericObject struct {
	metadata Metadata
	Data     struct {
		Kind Kind
		Name string
	} `mapstructure:",squash" yaml:",inline"`
}

func (o GenericObject) Name() string {
	return o.Data.Name
}

func (o GenericObject) Kind() Kind {
	return o.Data.Kind
}

func (o GenericObject) Metadata() Metadata {
	return o.metadata
}

func (o GenericObject) Directory() string {
	return filepath.Dir(o.metadata.Filename())
}

func (o GenericObject) DisplayName() string {
	return fmt.Sprintf("%s %q", strings.ToLower(string(o.Kind())), o.Name())
}

func (o GenericObject) Validate(ObjectCollection) error {
	return nil
}

func NewObject(filename string, data *yaml.Node) (Object, error) {
	var obj GenericObject
	err := data.Decode(&obj)
	if err != nil {
		return nil, err
	}

	switch obj.Kind() {
	case ProjectKind:
		return NewProject(filename, data)
	case ServiceKind:
		return NewService(filename, data)
	case EnvironmentKind:
		return NewEnvironment(filename, data)
	case BuilderKind, DeployerKind, PusherKind, TaggerKind:
		return NewExecutor(filename, data)
	default:
		return nil, fmt.Errorf("unknown kind %q", obj.Kind())
	}
}

func decode(data *yaml.Node, result interface{}) (err error) {
	var aux interface{}

	err = data.Decode(&aux)
	if err != nil {
		return
	}

	aux, err = scheme.ToInternal(aux)
	if err != nil {
		return
	}

	decoderConfig := &mapstructure.DecoderConfig{
		ErrorUnused: false,
		Squash:      true,
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

	return decoder.Decode(aux)
}
