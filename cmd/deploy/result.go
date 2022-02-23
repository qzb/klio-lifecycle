package main

import (
	"github.com/g2a-com/cicd/internal/object"
)

type ResultEntry struct {
	Service string `json:"service"`
	Entry   int    `json:"entry"`
	Result  string `json:"result"`
}

type Result struct {
	Releases []ResultEntry `json:"releases"`
}

func (r *Result) addReleases(service object.Service, entry object.ServiceEntry, releases []string) {
	for _, release := range releases {
		r.Releases = append(r.Releases, ResultEntry{service.Name, entry.Index, release})
	}
}
