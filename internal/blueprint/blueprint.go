package blueprint

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/g2a-com/cicd/internal/object"
	"github.com/g2a-com/cicd/internal/placeholders"
	log "github.com/g2a-com/klio-logger-go"
	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"
)

type Mode string

const (
	BuildMode  Mode = "build"
	DeployMode Mode = "deploy"
	RunMode    Mode = "run"
)

type Preprocessor func([]byte) ([]byte, error)

type Blueprint struct {
	Mode           Mode
	Services       []string
	Params         map[string]string
	Environment    string
	Tag            string
	Preprocessors  []Preprocessor
	documents      map[string]*document
	project        *document
	processedFiles map[string]bool
}

func (b *Blueprint) init() error {
	if b.processedFiles == nil {
		b.processedFiles = map[string]bool{}
	}
	if b.documents == nil {
		b.documents = map[string]*document{}
	}

	if b.Mode == "" {
		return errors.New("mode is not specified")
	}
	if b.Mode == DeployMode && b.Environment == "" {
		return errors.New("environment is requited in deploy mode")
	}
	return nil
}

func (b *Blueprint) Validate() (err error) {
	// Check conssitency of the blueprint (if there is only one project, if all releases have deployers, etc...)
	err = b.checkConsistency()
	if err != nil {
		return err
	}

	// Check artifacts, releases, etc match schemas enforced by executors
	err = b.validateSpecsAgainstExecutorSchemas()
	if err != nil {
		return err
	}

	return nil
}

func (b *Blueprint) Load(glob string) error {
	err := b.init()
	if err != nil {
		return err
	}

	glob, err = filepath.Abs(glob)
	if err != nil {
		return err
	}

	globs := []string{glob}

	for i := 0; i < len(globs); i++ {
		glob := globs[i]

		paths, err := filepath.Glob(glob)
		if err != nil {
			return err
		}

		for _, p := range paths {
			if _, ok := b.processedFiles[p]; ok {
				continue
			} else {
				b.processedFiles[p] = true
			}

			docs, err := b.readFile(p, b.Mode)
			if err != nil {
				return fmt.Errorf(`file "%s" contains invalid document: %s`, p, err)
			}

			for _, doc := range docs {
				project, ok := doc.Object.(object.Project)
				if ok {
					for _, entry := range project.Files {
						globs = append(globs, path.Join(project.Directory, entry))
					}
				}
			}

			err = b.addDocuments(docs...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetProject get Project object
func (b *Blueprint) GetProject() object.Project {
	return b.project.Object.(object.Project)
}

// GetEnvironment gets Executor object by kind and name
func (b *Blueprint) GetExecutor(kind object.Kind, name string) (object.Executor, bool) {
	d, ok1 := b.getObject(kind, name)
	e, ok2 := d.(object.Executor)
	return e, ok1 && ok2
}

// GetEnvironment gets Service object by the name
func (b *Blueprint) GetService(name string) (object.Service, bool) {
	d, ok := b.getObject(object.ServiceKind, name)
	return d.(object.Service), ok
}

// GetEnvironment gets Environment object by the name
func (b *Blueprint) GetEnvironment(name string) (object.Environment, bool) {
	d, ok := b.getObject(object.EnvironmentKind, name)
	return d.(object.Environment), ok
}

// ListServices returns all service objects in the blueprint
func (b *Blueprint) ListServices() []object.Service {
	names := b.getServiceNames()
	services := make([]object.Service, 0, len(names))
	for _, name := range b.getServiceNames() {
		s, _ := b.GetService(name)
		services = append(services, s)
	}
	return services
}

func (b *Blueprint) addDocuments(documents ...*document) error {
	for _, d := range documents {
		key := string(d.Kind) + "/" + d.Name
		duplicate, ok := b.documents[key]
		if ok {
			return fmt.Errorf("%s %q is duplicated, it's defined in:\n\t* %s\n\t* %s", strings.ToLower(string(d.Kind)), d.Name, duplicate.FilePath, d.FilePath)
		}
		b.documents[key] = d

		if d.Kind == object.ProjectKind {
			if b.project != nil {
				return fmt.Errorf("project is duplicated, it's defined in:\n\t* %s\n\t* %s", b.project.FilePath, d.FilePath)
			}
			b.project = d
		}
	}

	return nil
}

func (b *Blueprint) checkConsistency() error {
	var err error

	for _, d := range b.documents {
		switch d.Kind {
		case object.ServiceKind:
			// check taggers
			for _, t := range d.Object.(object.Service).Build.Tags {
				_, ok := b.getObject(object.TaggerKind, t.Type)
				if !ok {
					err = multierror.Append(err, fmt.Errorf("missing tagger %q used by service %q defined in the file:\n\t  %s", t.Type, d.Name, d.FilePath))
				}
			}
			// check builders
			for _, a := range d.Object.(object.Service).Build.Artifacts.ToBuild {
				_, ok := b.getObject(object.BuilderKind, a.Type)
				if !ok {
					err = multierror.Append(err, fmt.Errorf("missing builder %q used by service %q defined in the file:\n\t  %s", a.Type, d.Name, d.FilePath))
				}
			}
			// check pushers
			for _, a := range d.Object.(object.Service).Build.Artifacts.ToPush {
				_, ok := b.getObject(object.PusherKind, a.Type)
				if !ok {
					err = multierror.Append(err, fmt.Errorf("missing pusher %q used by service %q defined in the file:\n\t  %s", a.Type, d.Name, d.FilePath))
				}
			}
			// check deployers
			for _, r := range d.Object.(object.Service).Deploy.Releases {
				_, ok := b.getObject(object.DeployerKind, r.Type)
				if !ok {
					err = multierror.Append(err, fmt.Errorf("missing deployer %q used by service %q defined in the file:\n\t  %s", r.Type, d.Name, d.FilePath))
				}
			}
		case object.EnvironmentKind:
			// check services
			for _, name := range d.Object.(object.Environment).DeployServices {
				_, ok := b.getObject(object.ServiceKind, name)
				if !ok {
					err = multierror.Append(err, fmt.Errorf("missing service %q deployed to environment %q defined in the file:\n\t  %s", name, d.Name, d.FilePath))
				}
			}
		}

		for _, name := range b.getServiceNames() {
			_, ok := b.getObject(object.ServiceKind, name)
			if !ok {
				err = multierror.Append(err, fmt.Errorf("service %q does not exist, available services: %s", name, strings.Join(b.getServiceNames(), ", ")))
			}
		}

		if b.Environment != "" {
			_, ok := b.getObject(object.EnvironmentKind, b.Environment)
			if !ok {
				err = multierror.Append(err, fmt.Errorf("environment %q does not exist, available environments: %s", b.Environment, strings.Join(b.getEnvironmentNames(), ", ")))
			}
		}
	}

	return err
}

func (b *Blueprint) ExpandPlaceholders() (err error) {
	project := b.GetProject()

	for _, d := range b.documents {
		service, ok := d.Object.(object.Service)
		if !ok {
			continue
		}

		values := map[string]interface{}{
			"Project": map[string]interface{}{
				"Dir":  project.Directory,
				"Name": project.Name,
				"Vars": project.Variables,
			},
			"Service": map[string]interface{}{
				"Dir":  service.Directory,
				"Name": service.Name,
			},
			"Params": b.Params,
		}

		if b.Environment != "" {
			environment, _ := b.GetEnvironment(b.Environment)
			values["Environment"] = map[string]interface{}{
				"Dir":  environment.Directory,
				"Name": environment.Name,
				"Vars": environment.Variables,
			}
		}

		if b.Tag != "" {
			values["Tag"] = b.Tag
		}

		for i := range service.Build.Artifacts.ToBuild {
			service.Build.Artifacts.ToBuild[i].Spec, err = placeholders.ReplaceWithValues(service.Build.Artifacts.ToBuild[i].Spec, values)
			if err != nil {
				return err
			}
		}
		for i := range service.Build.Artifacts.ToPush {
			service.Build.Artifacts.ToPush[i].Spec, err = placeholders.ReplaceWithValues(service.Build.Artifacts.ToPush[i].Spec, values)
			if err != nil {
				return err
			}
		}
		for i := range service.Deploy.Releases {
			service.Deploy.Releases[i].Spec, err = placeholders.ReplaceWithValues(service.Deploy.Releases[i].Spec, values)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Blueprint) validateSpecsAgainstExecutorSchemas() error {
	ctx := context.Background()

	var err error

	validate := func(service *document, executor *document, spec interface{}) {
		schema := executor.Object.(object.Executor).Schema
		result := schema.Validate(ctx, spec)

		if len(*result.Errs) > 0 {
			for _, e := range *result.Errs {
				err = multierror.Append(err, fmt.Errorf(
					"%s contains invalid configuration for %s:\n\t  %s\n\t  Definition files:\n\t    %s\n\t    %s",
					service.DisplayName, executor.DisplayName, e, service.FilePath, executor.FilePath,
				))
			}
		}
	}

	for _, doc := range b.documents {
		if doc.Kind != object.ServiceKind {
			continue
		}

		service := doc.Object.(object.Service)

		for _, entry := range service.Build.Tags {
			validate(doc, b.documents[string(object.TaggerKind)+"/"+entry.Type], entry.Spec)
		}
		for _, entry := range service.Build.Artifacts.ToBuild {
			validate(doc, b.documents[string(object.BuilderKind)+"/"+entry.Type], entry.Spec)
		}
		for _, entry := range service.Build.Artifacts.ToPush {
			validate(doc, b.documents[string(object.PusherKind)+"/"+entry.Type], entry.Spec)
		}
		for _, entry := range service.Deploy.Releases {
			validate(doc, b.documents[string(object.DeployerKind)+"/"+entry.Type], entry.Spec)
		}
	}

	return err
}

func (b *Blueprint) getObject(kind object.Kind, name string) (interface{}, bool) {
	key := string(kind) + "/" + name
	d, ok := b.documents[key]
	if !ok {
		return nil, false
	}
	return d.Object, true
}

func (b *Blueprint) getServiceNames() (names []string) {
	if len(b.Services) > 0 {
		return b.Services
	}
	for _, d := range b.documents {
		if d.Kind == object.ServiceKind {
			names = append(names, d.Name)
		}
	}
	sort.Strings(names)
	return names
}

func (b *Blueprint) getEnvironmentNames() (names []string) {
	for _, d := range b.documents {
		if d.Kind == object.EnvironmentKind {
			names = append(names, d.Name)
		}
	}
	sort.Strings(names)
	return names
}

func (b *Blueprint) readFile(filePath string, mode Mode) ([]*document, error) {
	log.Debugf("Loading file: %s", filePath)
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	for _, preprocessor := range b.Preprocessors {
		buf, err = preprocessor(buf)
		if err != nil {
			return nil, fmt.Errorf(`file "%s" contains invalid document: %s`, filePath, err)
		}
	}

	var documents []*document

	reader := bytes.NewReader(buf)
	decoder := yaml.NewDecoder(reader)

	for i := 0; true; i++ {
		var content interface{}

		err := decoder.Decode(&content)
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf(`file "%s" contains invalid document: %s`, filePath, err)
			}
			break
		}

		doc, err := newDocument(filePath, i, mode, content)
		if err != nil {
			return nil, fmt.Errorf(`file "%s" contains invalid document: %s`, filePath, err)
		}

		if mode != DeployMode && doc.Kind == object.EnvironmentKind {
			continue
		}

		documents = append(documents, doc)
	}

	return documents, nil
}
