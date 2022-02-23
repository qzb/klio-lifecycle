package main

type DeployerInput struct {
	Spec   any  `tengo:"spec"`
	Force  bool `tengo:"force"`
	DryRun bool `tengo:"dryRun"`
	Wait   int  `tengo:"wait"`
}
