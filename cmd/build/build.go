package main

import (
	"fmt"
	"os"

	"github.com/g2a-com/cicd/internal/blueprint"
	"github.com/g2a-com/cicd/internal/flags"
	"github.com/g2a-com/cicd/internal/object"
	"github.com/g2a-com/cicd/internal/script"
	"github.com/g2a-com/cicd/internal/utils"
	log "github.com/g2a-com/klio-logger-go/v2"
)

type options struct {
	Push        bool              `flag:"push" alias:"p" help:"Push artifacts to remote registry"`
	Services    []string          `flag:"services" alias:"s" help:"List of services to build (skip to build all services)"`
	Params      map[string]string `flag:"param" help:"Parameters to use in configuration files (key=value pairs)"`
	ProjectFile string            `flag:"project-file" alias:"f" help:"Path to project file"`
	ResultFile  string            `flag:"result-file" help:"Where to write result file"`
}

func main() {
	// Exit nicely on panics
	defer utils.HandlePanics()

	// Parse options
	opts := options{
		ResultFile:  "build-result.json",
		ProjectFile: utils.FindProjectFile(),
	}
	flags.ParseArgs(&opts, os.Args)

	// Prepare logger
	l := log.StandardLogger()

	// Handle results
	result := &Result{}
	defer utils.SaveResult(opts.ResultFile, result)

	// Check if project file exists
	if !utils.FileExists(opts.ProjectFile) {
		panic("cannot find project.yaml")
	}

	// Load blueprint
	blueprint, err := blueprint.Load(blueprint.Opts{
		Mode:        blueprint.BuildMode,
		ProjectFile: opts.ProjectFile,
		Params:      opts.Params,
		Services:    opts.Services,
	})
	if err != nil {
		panic(err)
	}

	// Change working directory
	os.Chdir(blueprint.GetProject().Directory)

	// Helper for getting executors
	getExecutor := func(kind object.Kind, name string) object.Executor {
		e, ok := blueprint.GetExecutor(kind, name)
		if !ok {
			panic(fmt.Errorf("%s %q does not exist", kind, name))
		}
		return e
	}

	// Build
	for _, service := range blueprint.ListServices() {
		l := l.WithTags(service.Name)

		if len(service.Build.Artifacts.ToBuild) == 0 {
			l.WithLevel(log.VerboseLevel).Print("No artifacts to build")
			continue
		}

		// Generate tags
		for _, entry := range service.Build.Tags {
			s := script.New(getExecutor(object.TaggerKind, entry.Type))
			s.Logger = l

			res, err := s.Run(TaggerInput{
				Spec: entry.Spec,
				Dirs: Dirs{
					Project: blueprint.GetProject().Directory,
					Service: service.Directory,
				},
			})
			if err != nil {
				panic(err)
			}

			result.addTags(service, entry, res)
		}

		if len(result.getTags(service)) == 0 {
			l.WithLevel(log.WarnLevel).Print("No tags to build")
			continue
		}

		// Build artifacts
		for _, entry := range service.Build.Artifacts.ToBuild {
			s := script.New(getExecutor(object.BuilderKind, entry.Type))
			s.Logger = l

			res, err := s.Run(BuilderInput{
				Spec: entry.Spec,
				Tags: result.getTags(service),
				Dirs: Dirs{
					Project: blueprint.GetProject().Directory,
					Service: service.Directory,
				},
			})
			if err != nil {
				panic(err)
			}

			result.addArtifacts(service, entry, res)
		}
	}

	// Push artifacts
	if opts.Push {
		for _, service := range blueprint.ListServices() {
			l := l.WithTags("push", service.Name)

			for _, entry := range service.Build.Artifacts.ToPush {
				s := script.New(getExecutor(object.PusherKind, entry.Type))
				s.Logger = l

				res, err := s.Run(PusherInput{
					Spec:      entry.Spec,
					Tags:      result.getTags(service),
					Artifacts: result.getArtifacts(service, entry),
					Dirs: Dirs{
						Project: blueprint.GetProject().Directory,
						Service: service.Directory,
					},
				})
				if err != nil {
					panic(err)
				}

				result.addPushedArtifacts(service, entry, res)
			}
		}
	}

	// Print success message
	switch count := len(blueprint.ListServices()); count {
	case 0:
		l.Print("There was nothing to build")
	case 1:
		l.Print("Successfully built 1 service")
	default:
		l.Printf("Successfully built %v services", count)
	}
}
