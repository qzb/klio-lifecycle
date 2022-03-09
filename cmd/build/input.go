package main

type TaggerInput struct {
	Spec interface{} `tengo:"spec"`
	Dirs Dirs        `tengo:"dirs"`
}

type BuilderInput struct {
	Spec interface{} `tengo:"spec"`
	Tags []string    `tengo:"tags"`
	Dirs Dirs        `tengo:"dirs"`
}

type PusherInput struct {
	Spec      interface{} `tengo:"spec"`
	Tags      []string    `tengo:"tags"`
	Artifacts []string    `tengo:"artifacts"`
	Dirs      Dirs        `tengo:"dirs"`
}

type Dirs struct {
	Project string `tengo:"project"`
	Service string `tengo:"service"`
}
