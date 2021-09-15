package main

import (
	"github.com/g2a-com/cicd/internal/runner"
)

type Result struct {
	Releases []runner.Result `json:"releases"`
}
