package main

type TaggerInput struct {
	Spec any  `tengo:"spec"`
	Dirs Dirs `tengo:"dirs"`
}

type BuilderInput struct {
	Spec any      `tengo:"spec"`
	Tags []string `tengo:"tags"`
	Dirs Dirs     `tengo:"dirs"`
}

type PusherInput struct {
	Spec      any      `tengo:"spec"`
	Tags      []string `tengo:"tags"`
	Artifacts []string `tengo:"artifacts"`
	Dirs      Dirs     `tengo:"dirs"`
}

type Dirs struct {
	Project string `tengo:"project"`
	Service string `tengo:"service"`
}
