package build

import (
	"github.com/g2a-com/klio-lifecycle/internal/object"
	runner "github.com/g2a-com/klio-lifecycle/internal/runner"
)

type Result struct {
	Tags            []runner.Result `json:"tags"`
	Artifacts       []runner.Result `json:"artifacts"`
	PushedArtifacts []runner.Result `json:"pushedArtifacts"`
}

func (r *Result) getTags(service object.Service) (tags []string) {
	for _, r := range r.Tags {
		if r.Service == service.Name {
			tags = append(tags, r.Result)
		}
	}
	return tags
}

func (r *Result) getArtifacts(service object.Service, entry object.ServiceEntry) (artifacts []string) {
	for _, r := range r.Artifacts {
		if r.Service == service.Name && r.Entry == entry.Index {
			artifacts = append(artifacts, r.Result)
		}
	}
	return artifacts
}
