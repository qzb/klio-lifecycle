package main

type DeployerInput struct {
	Spec   interface{} `tengo:"spec"`
	Force  bool        `tengo:"force"`
	DryRun bool        `tengo:"dryRun"`
	Wait   int         `tengo:"wait"`
	Dirs   Dirs        `tengo:"dirs"`
}

type Dirs struct {
	Project     string `tengo:"project"`
	Environment string `tengo:"environment"`
	Service     string `tengo:"service"`
}
