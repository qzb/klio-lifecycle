package blueprint

import (
	"bytes"
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
	"gopkg.in/yaml.v3"
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
	for _, d := range b.documents {
		e := d.Object.Validate(b)
		if e != nil {
			err = multierror.Append(err, e)
		}
	}

	if b.Environment != "" {
		if b.GetObject(object.EnvironmentKind, b.Environment) == nil {
			err = multierror.Append(err, fmt.Errorf("environment %q does not exist, available environments: %s", b.Environment, strings.Join(b.getEnvironmentNames(), ", ")))
		}
	}

	for _, name := range b.getServiceNames() {
		if b.GetObject(object.ServiceKind, name) == nil {
			err = multierror.Append(err, fmt.Errorf("service %q does not exist, available services: %s", name, strings.Join(b.getServiceNames(), ", ")))
		}
	}

	return err
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
						globs = append(globs, path.Join(project.Directory(), entry))
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
	o := b.GetObject(kind, name)
	e, ok := o.(object.Executor)
	return e, o != nil && ok
}

// GetEnvironment gets Service object by the name
func (b *Blueprint) GetService(name string) (object.Service, bool) {
	o := b.GetObject(object.ServiceKind, name)
	return o.(object.Service), o != nil
}

// GetEnvironment gets Environment object by the name
func (b *Blueprint) GetEnvironment(name string) (object.Environment, bool) {
	o := b.GetObject(object.EnvironmentKind, name)
	return o.(object.Environment), o != nil
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

func (b *Blueprint) ExpandPlaceholders() (err error) {
	project := b.GetProject()

	for _, d := range b.documents {
		service, ok := d.Object.(object.Service)
		if !ok {
			continue
		}

		values := map[string]interface{}{
			"Project": map[string]interface{}{
				"Dir":  project.Directory(),
				"Name": project.Name(),
				"Vars": project.Variables,
			},
			"Service": map[string]interface{}{
				"Dir":  service.Directory(),
				"Name": service.Name(),
			},
			"Params": b.Params,
		}

		if b.Environment != "" {
			environment, _ := b.GetEnvironment(b.Environment)
			values["Environment"] = map[string]interface{}{
				"Dir":  environment.Directory(),
				"Name": environment.Name(),
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

func (b *Blueprint) GetObject(kind object.Kind, name string) object.Object {
	key := string(kind) + "/" + name
	d, ok := b.documents[key]
	if !ok {
		return nil
	}
	return d.Object
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
		var content yaml.Node

		err := decoder.Decode(&content)
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf(`file "%s" contains invalid document: %s`, filePath, err)
			}
			break
		}

		doc, err := newDocument(filePath, i, mode, &content)
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
