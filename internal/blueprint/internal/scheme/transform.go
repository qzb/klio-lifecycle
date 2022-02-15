package scheme

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/icza/dyno"
)

func toInternal(obj any) any {
	if getString(obj, "apiVersion") != "g2a-cli/v2.0" {
		obj = toV2(obj)
	}

	switch kind := getString(obj, "kind"); kind {
	case "Project":
		return toInternalProject(obj)
	case "Service":
		return toInternalService(obj)
	case "Environment":
		return toInternalEnvironment(obj)
	case "Builder", "Deployer", "Pusher", "Tagger":
		return toInternalExecutor(obj)
	default:
		panic(fmt.Errorf("unsupported kind: %s", kind))
	}
}

func toInternalProject(obj any) any {
	files := []any{}

	for _, f := range getSlice(obj, "files") {
		if isString(f) {
			files = append(files, map[string]any{"glob": f})
		} else {
			if has(f, "git") {
				gitFiles := get(f, "git", "files")
				if isString(gitFiles) {
					gitFiles = []any{gitFiles}
				}
				for _, glob := range getSlice(gitFiles) {
					files = append(files, map[string]any{
						"glob": glob,
						"git": map[string]any{
							"url": getString(f, "git", "url"),
							"rev": getString(f, "git", "rev"),
						},
					})
				}
			}
		}
	}

	return map[string]any{
		"kind":      "Project",
		"files":     files,
		"name":      getString(obj, "name"),
		"variables": getMap(obj, "variables"),
	}
}

func toInternalService(obj any) any {
	toBuild := mapSlice(getSlice(obj, "artifacts"), toInternalEntry)
	tags := mapSlice(getSlice(obj, "tags"), toInternalEntry)
	releases := mapSlice(getSlice(obj, "releases"), toInternalEntry)

	toPush := []any{}
	for i, v := range getSlice(obj, "artifacts") {
		e := toInternalEntry(i, or(get(v, "push"), v))
		if getString(e, "type") != "" {
			toPush = append(toPush, e)
		}
	}

	tasks := map[string]any{}
	for k, v := range getMap(obj, "tasks") {
		tasks[k] = mapSlice(getSlice(v), toInternalEntry)
	}

	return map[string]any{
		"kind": "Service",
		"name": getString(obj, "name"),
		"build": map[string]any{
			"artifacts": map[string]any{
				"toBuild": toBuild,
				"toPush":  toPush,
			},
			"tags": tags,
		},
		"deploy": map[string]any{
			"releases": releases,
		},
		"run": map[string]any{
			"tasks": tasks,
		},
	}
}

func toInternalEnvironment(obj any) any {
	return map[string]any{
		"kind":           "Environment",
		"name":           getString(obj, "name"),
		"deployServices": getSlice(obj, "deployServices"),
		"variables":      getMap(obj, "variables"),
	}
}

func toInternalExecutor(obj any) any {
	schema, err := json.Marshal(getMap(obj, "schema"))
	if err != nil {
		panic(err)
	}
	return map[string]any{
		"kind":   getString(obj, "kind"),
		"name":   getString(obj, "name"),
		"js":     getString(obj, "js"),
		"schema": string(schema),
	}
}

func toInternalEntry(i int, obj any) any {
	if obj == false {
		return nil
	}
	if isString(obj) {
		return map[string]any{"index": int64(i), "type": obj, "spec": nil}
	}
	for k, v := range getMap(obj) {
		if k != "push" {
			return map[string]any{"index": int64(i), "type": k, "spec": v}
		}
	}
	return nil
}

func toV2(obj any) any {
	if getString(obj, "apiVersion") != "g2a-cli/v1beta4" {
		panic(fmt.Errorf("unsupported version: %s", get(obj, "apiVersion")))
	}

	switch kind := getString(obj, "kind"); kind {
	case "Project":
		return toV2Project(obj)
	case "Service":
		return toV2Service(obj)
	case "Environment":
		return toV2Environment(obj)
	default:
		panic(fmt.Errorf("unsupported kind: %s", kind))
	}
}

func toV2Project(obj any) any {
	files := []any{
		map[string]any{
			"git": map[string]any{
				"url":   "git@github.com:g2a-com/klio-lifecycle.git",
				"rev":   "main",
				"files": "assets/executors/*/*.yaml",
			},
		},
	}

	if has(obj, "services") {
		for _, s := range getSlice(obj, "services") {
			files = append(files, path.Join(s.(string), "service.yaml"))
		}
	} else {
		files = append(files, "services/*/service.yaml")
	}

	if has(obj, "environments") {
		for _, e := range getSlice(obj, "environments") {
			files = append(files, path.Join(e.(string), "environment.yaml"))
		}
	} else {
		files = append(files, "environments/*/environment.yaml")
	}

	return map[string]any{
		"apiVersion": "g2a-cli/v2.0",
		"kind":       "Project",
		"name":       or(get(obj, "name"), "project"),
		"files":      files,
		"tasks":      map[string]any{},
		"variables":  map[string]any{},
	}
}

func toV2Service(obj any) any {
	artifacts := getSlice(obj, "build", "artifacts")

	if hooks := get(obj, "hooks", "pre-build"); hooks != nil {
		artifacts = prepend(artifacts, map[string]any{
			"script": hooksToScript(hooks),
			"push":   false,
		})
	}
	if hooks := get(obj, "hooks", "post-build"); hooks != nil {
		artifacts = append(artifacts, map[string]any{
			"script": hooksToScript(hooks),
			"push":   false,
		})
	}

	releases := getSlice(obj, "deploy", "releases")

	if hooks := get(obj, "hooks", "pre-deploy"); hooks != nil {
		releases = prepend(releases, map[string]any{
			"script": hooksToScript(hooks),
		})
	}
	if hooks := get(obj, "hooks", "post-deploy"); hooks != nil {
		releases = append(releases, map[string]any{
			"script": hooksToScript(hooks),
		})
	}

	tags := []any{}

	for k, v := range getMap(obj, "build", "tagPolicy") {
		tags = append(tags, map[string]any{k: v})
	}

	return map[string]any{
		"apiVersion": "g2a-cli/v2.0",
		"kind":       "Service",
		"name":       getString(obj, "name"),
		"tags":       tags,
		"artifacts":  artifacts,
		"releases":   releases,
		"tasks":      map[string]any{},
	}
}

func toV2Environment(obj any) any {
	return map[string]any{
		"apiVersion":     "g2a-cli/v2.0",
		"kind":           "Environment",
		"name":           getString(obj, "name"),
		"deployServices": getSlice(obj, "deployServices"),
		"variables":      getMap(obj, "variables"),
	}
}

func hooksToScript(hooks any) any {
	sh := "set -e"
	for _, hook := range getSlice(hooks) {
		sh += fmt.Sprintf("\n%s", hook)
	}
	return map[string]any{"sh": sh}
}

func getString(v any, path ...any) string {
	str, err := dyno.GetString(v, path...)
	if err != nil {
		return ""
	}
	return str
}

func isString(v any, path ...any) bool {
	_, err := dyno.GetString(v, path...)
	return err == nil
}

func getSlice(v any, path ...any) []any {
	res, err := dyno.GetSlice(v, path...)
	if err != nil {
		return []any{}
	}
	return res
}

func getMap(v any, path ...any) map[string]any {
	res, err := dyno.GetMapS(v, path...)
	if err != nil {
		return map[string]any{}
	}
	return res
}

func get(v any, path ...any) any {
	val, err := dyno.Get(v, path...)
	if err != nil {
		return nil
	}
	return val
}

func has(v any, path ...any) bool {
	_, err := dyno.Get(v, path...)
	return err == nil
}

func or(v1 any, v2 any) any {
	if v1 != nil {
		return v1
	} else {
		return v2
	}
}

func mapSlice(v []any, fn func(int, any) any) []any {
	res := make([]any, len(v), len(v))
	for i := range v {
		res[i] = fn(i, v[i])
	}
	return res
}

func prepend(slice []any, elems ...any) []any {
	return append(elems, slice...)
}
