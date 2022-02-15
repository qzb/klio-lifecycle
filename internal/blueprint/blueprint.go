package blueprint

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/g2a-com/cicd/internal/blueprint/internal/placeholders"
	"github.com/g2a-com/cicd/internal/blueprint/internal/remotes"
	"github.com/g2a-com/cicd/internal/object"
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

type Opts struct {
	ProjectFile string
	Mode        Mode
	Services    []string
	Params      map[string]string
	Environment string
	Tag         string
}

type Blueprint struct {
	Documents map[string]*document
	project   *document
	opts      Opts
}

func Load(opts Opts) (*Blueprint, error) {
	if opts.Mode == "" {
		panic("mode is not specified")
	}
	if opts.Mode == DeployMode && opts.Environment == "" {
		panic("environment is requited in deploy mode")
	}

	b := &Blueprint{
		Documents: make(map[string]*document),
		opts:      opts,
	}

	// Load all documents from project file
	projectFile, err := filepath.Abs(b.opts.ProjectFile)
	if err != nil {
		return nil, err
	}
	b.opts.ProjectFile = projectFile

	docs, err := readFile(b.opts.ProjectFile, b.opts.Mode)
	if err != nil {
		return nil, err
	}

	err = b.addDocuments(docs...)
	if err != nil {
		return nil, err
	}

	// Project file MUST contain single Project object
	if b.project == nil {
		return nil, fmt.Errorf("there is no project configuration in the file:\n\t%s", b.opts.ProjectFile)
	}

	// Fetch repositories containing remote files
	for _, file := range b.GetProject().Files {
		if file.Git != nil {
			err := remotes.Fetch(b.GetProject().Directory, file.Git.URL, file.Git.Rev)
			if err != nil {
				return nil, fmt.Errorf("failed to fetch revision %q from git repository %q: %s", file.Git.Rev, file.Git.URL, err)
			}
		}
	}

	// Load all extra files specified in the Project object
	processedFiles := map[string]bool{b.opts.ProjectFile: true}
	for _, file := range b.GetProject().Files {
		baseDir := b.GetProject().Directory
		if file.Git != nil {
			baseDir, _ = remotes.GetDir(baseDir, file.Git.URL, file.Git.Rev)
		}

		paths, err := filepath.Glob(filepath.Join(baseDir, filepath.FromSlash(file.Glob)))
		if err != nil {
			return nil, err
		}

		for _, p := range paths {
			if _, ok := processedFiles[p]; ok {
				continue
			} else {
				processedFiles[p] = true
			}

			docs, err := readFile(p, b.opts.Mode)
			if err != nil {
				return nil, fmt.Errorf(`file "%s" contains invalid document: %s`, p, err)
			}

			err = b.addDocuments(docs...)
			if err != nil {
				return nil, err
			}
		}
	}

	// Check conssitency of the blueprint (if there is only one project, if all releases have deployers, etc...)
	err = b.checkConsistency()
	if err != nil {
		return nil, err
	}

	// Fill placeholders
	err = b.fillPlaceholders()
	if err != nil {
		return nil, err
	}

	// If user didn't specified particular services, include all of them
	if len(b.opts.Services) == 0 {
		b.opts.Services = b.getServiceNames()
	}

	// Check artifacts, releases, etc match schemas enforced by executors
	err = b.validateSpecsAgainstExecutorSchemas()
	if err != nil {
		return nil, err
	}

	return b, nil
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
	services := make([]object.Service, 0, len(b.opts.Services))
	for _, name := range b.opts.Services {
		s, _ := b.GetService(name)
		services = append(services, s)
	}
	return services
}

func (b *Blueprint) addDocuments(documents ...*document) error {
	for _, d := range documents {
		key := string(d.Kind) + "/" + d.Name
		duplicate, ok := b.Documents[key]
		if ok {
			return fmt.Errorf("%s %q is duplicated, it's defined in:\n\t* %s\n\t* %s", strings.ToLower(string(d.Kind)), d.Name, duplicate.FilePath, d.FilePath)
		}
		b.Documents[key] = d

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

	for _, d := range b.Documents {
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

		for _, name := range b.opts.Services {
			_, ok := b.getObject(object.ServiceKind, name)
			if !ok {
				err = multierror.Append(err, fmt.Errorf("service %q does not exist, available services: %s", name, strings.Join(b.getServiceNames(), ", ")))
			}
		}

		if b.opts.Environment != "" {
			_, ok := b.getObject(object.EnvironmentKind, b.opts.Environment)
			if !ok {
				err = multierror.Append(err, fmt.Errorf("environment %q does not exist, available environments: %s", b.opts.Environment, strings.Join(b.getEnvironmentNames(), ", ")))
			}
		}
	}

	return err
}

func (b *Blueprint) fillPlaceholders() error {
	project := b.GetProject()

	for _, d := range b.Documents {
		values := map[string]interface{}{
			"Project": map[string]interface{}{
				"Dir":  project.Directory,
				"Name": project.Name,
				"Vars": project.Variables,
			},
			"Params": b.opts.Params,
		}

		if b.opts.Environment != "" {
			environment, _ := b.GetEnvironment(b.opts.Environment)
			values["Environment"] = map[string]interface{}{
				"Dir":  environment.Directory,
				"Name": environment.Name,
				"Vars": environment.Variables,
			}
		}

		if d.Kind == object.ServiceKind {
			service, _ := b.GetService(d.Name)
			values["Service"] = map[string]interface{}{
				"Dir":  service.Directory,
				"Name": service.Name,
			}
		}

		if b.opts.Tag != "" {
			values["Tag"] = b.opts.Tag
		}

		if d.APIVersion == "g2a-cli/v1beta4" {
			values["Dirs"] = map[string]interface{}{"Project": project.Directory}

			if b.opts.Environment != "" {
				environment, _ := b.GetEnvironment(b.opts.Environment)
				values["Env"] = environment.Variables
				values["Dirs"].(map[string]interface{})["Environment"] = environment.Directory
			}

			if d.Kind == object.ServiceKind {
				service, _ := b.GetService(d.Name)
				values["Dirs"].(map[string]interface{})["Service"] = service.Directory
			}

			if b.opts.Tag != "" {
				values["Opts"] = map[string]interface{}{"Tag": b.opts.Tag}
			}
		}

		obj, err := placeholders.ProcessStruct(d.Object, values)
		if err != nil {
			return err
		}
		d.Object = obj
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

	for _, doc := range b.Documents {
		if doc.Kind != object.ServiceKind {
			continue
		}

		service := doc.Object.(object.Service)

		for _, entry := range service.Build.Tags {
			validate(doc, b.Documents[string(object.TaggerKind)+"/"+entry.Type], entry.Spec)
		}
		for _, entry := range service.Build.Artifacts.ToBuild {
			validate(doc, b.Documents[string(object.BuilderKind)+"/"+entry.Type], entry.Spec)
		}
		for _, entry := range service.Build.Artifacts.ToPush {
			validate(doc, b.Documents[string(object.PusherKind)+"/"+entry.Type], entry.Spec)
		}
		for _, entry := range service.Deploy.Releases {
			validate(doc, b.Documents[string(object.DeployerKind)+"/"+entry.Type], entry.Spec)
		}
	}

	return err
}

func (b *Blueprint) getObject(kind object.Kind, name string) (interface{}, bool) {
	key := string(kind) + "/" + name
	d, ok := b.Documents[key]
	if !ok {
		return nil, false
	}
	return d.Object, true
}

func (b *Blueprint) getServiceNames() (names []string) {
	for _, d := range b.Documents {
		if d.Kind == object.ServiceKind {
			names = append(names, d.Name)
		}
	}
	sort.Strings(names)
	return names
}

func (b *Blueprint) getEnvironmentNames() (names []string) {
	for _, d := range b.Documents {
		if d.Kind == object.EnvironmentKind {
			names = append(names, d.Name)
		}
	}
	sort.Strings(names)
	return names
}

func readFile(filePath string, mode Mode) ([]*document, error) {
	log.Debugf("Loading file: %s", filePath)
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
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
