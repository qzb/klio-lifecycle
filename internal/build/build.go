package build

import (
	"os"

	"github.com/g2a-com/klio-lifecycle/internal/blueprint"
	"github.com/g2a-com/klio-lifecycle/internal/flags"
	runner "github.com/g2a-com/klio-lifecycle/internal/runner"
	"github.com/g2a-com/klio-lifecycle/internal/utils"
	log "github.com/g2a-com/klio-logger-go"
)

type options struct {
	Push        bool              `flag:"push" alias:"p" help:"Push artifacts to remote registry"`
	Services    []string          `flag:"services" alias:"s" help:"List of services to build (skip to build all services)"`
	Params      map[string]string `flag:"param" help:"Parameters to use in configuration files (key=value pairs)"`
	ProjectFile string            `flag:"project-file" alias:"f" help:"Path to project file"`
	ResultFile  string            `flag:"result-file" help:"Where to write result file"`
}

func Run(runInterpreter runner.Interpreter) {
	// Exit nicely on panics
	defer utils.HandlePanics()

	// Parse options
	opts := options{
		Push:       true,
		ResultFile: "result.json",
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

	// Build
	for _, service := range blueprint.ListServices() {
		l := l.WithTags(service.Name)

		if len(service.Build.Artifacts.ToBuild) == 0 {
			l.WithLevel(log.VerboseLevel).Print("No artifacts to build")
			continue
		}

		// Generate tags
		for _, entry := range service.Build.Tags {
			r := runner.TaggerRunner{
				Interpreter: runInterpreter,
				Blueprint:   blueprint,
				Service:     service,
				Entry:       entry,
			}

			res, err := r.Run()
			if err != nil {
				panic(err)
			}

			result.Tags = append(result.Tags, res...)
		}

		if len(result.getTags(service)) == 0 {
			l.WithLevel(log.WarnLevel).Print("No tags to build")
			continue
		}

		// Build artifacts
		for _, entry := range service.Build.Artifacts.ToBuild {
			r := runner.BuilderRunner{
				Interpreter: runInterpreter,
				Blueprint:   blueprint,
				Service:     service,
				Entry:       entry,
				Tags:        result.getTags(service),
			}

			res, err := r.Run()
			if err != nil {
				panic(err)
			}

			result.Artifacts = append(result.Artifacts, res...)
		}
	}

	// Push artifacts
	if opts.Push {
		for _, service := range blueprint.ListServices() {
			for _, entry := range service.Build.Artifacts.ToPush {
				r := runner.PusherRunner{
					Interpreter: runInterpreter,
					Blueprint:   blueprint,
					Service:     service,
					Entry:       entry,
					Tags:        result.getTags(service),
					Artifacts:   result.getArtifacts(service, entry),
				}

				res, err := r.Run()
				if err != nil {
					panic(err)
				}

				result.PushedArtifacts = append(result.PushedArtifacts, res...)
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
