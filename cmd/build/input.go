package main

type TaggerInput struct {
	Spec any `tengo:"spec"`
}

type BuilderInput struct {
	Spec any      `tengo:"spec"`
	Tags []string `tengo:"tags"`
}

type PusherInput struct {
	Spec      any      `tengo:"spec"`
	Tags      []string `tengo:"tags"`
	Artifacts []string `tengo:"artifacts"`
}
