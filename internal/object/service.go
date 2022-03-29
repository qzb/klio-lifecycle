package object

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/qri-io/jsonschema"
	"gopkg.in/yaml.v3"
)

type ServiceEntry struct {
	Index int
	Type  string
	Spec  interface{}
}

type Service struct {
	GenericObject

	Build struct {
		Tags      []ServiceEntry
		Artifacts struct {
			ToBuild []ServiceEntry
			ToPush  []ServiceEntry
		}
	}
	Deploy struct {
		Releases []ServiceEntry
	}
	Run struct {
		Tasks map[string][]ServiceEntry
	}
}

var _ Object = Service{}

func NewService(filename string, data *yaml.Node) (service Service, err error) {
	service.GenericObject.metadata = NewMetadata(filename, data)
	err = decode(data, &service)
	return
}

func (s Service) Validate(c ObjectCollection) (err error) {
	for _, t := range s.Build.Tags {
		if c.GetObject(TaggerKind, t.Type) == nil {
			err = multierror.Append(err, fmt.Errorf("missing tagger %q used by service %q defined in the file:\n\t  %s", t.Type, s.Name(), s.Metadata().Filename()))
		}
	}
	for _, a := range s.Build.Artifacts.ToBuild {
		if c.GetObject(BuilderKind, a.Type) == nil {
			err = multierror.Append(err, fmt.Errorf("missing builder %q used by service %q defined in the file:\n\t  %s", a.Type, s.Name(), s.Metadata().Filename()))
		}
	}
	for _, a := range s.Build.Artifacts.ToPush {
		if c.GetObject(PusherKind, a.Type) == nil {
			err = multierror.Append(err, fmt.Errorf("missing pusher %q used by service %q defined in the file:\n\t  %s", a.Type, s.Name(), s.Metadata().Filename()))
		}
	}
	for _, r := range s.Deploy.Releases {
		if c.GetObject(DeployerKind, r.Type) == nil {
			err = multierror.Append(err, fmt.Errorf("missing deployer %q used by service %q defined in the file:\n\t  %s", r.Type, s.Name(), s.Metadata().Filename()))
		}
	}

	if err != nil {
		return err
	}

	ctx := context.Background()

	validate := func(obj Object, spec interface{}) {
		executor := obj.(interface {
			Object
			Schema() *jsonschema.Schema
		})

		schema := executor.Schema()
		result := schema.Validate(ctx, spec)

		if len(*result.Errs) > 0 {
			for _, e := range *result.Errs {
				err = multierror.Append(err, fmt.Errorf(
					"%s contains invalid configuration for %s:\n\t  %s\n\t  Definition files:\n\t    %s\n\t    %s",
					s.DisplayName(), executor.DisplayName(), e, s.Metadata().Filename(), executor.Metadata().Filename(),
				))
			}
		}
	}

	for _, entry := range s.Build.Tags {
		validate(c.GetObject(TaggerKind, entry.Type), entry.Spec)
	}
	for _, entry := range s.Build.Artifacts.ToBuild {
		validate(c.GetObject(BuilderKind, entry.Type), entry.Spec)
	}
	for _, entry := range s.Build.Artifacts.ToPush {
		validate(c.GetObject(PusherKind, entry.Type), entry.Spec)
	}
	for _, entry := range s.Deploy.Releases {
		validate(c.GetObject(DeployerKind, entry.Type), entry.Spec)
	}

	return
}
