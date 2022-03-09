package scheme

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/icza/dyno"
)

func toInternal(obj interface{}) interface{} {
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

func toInternalProject(obj interface{}) interface{} {
	files := []interface{}{}

	for _, f := range getSlice(obj, "files") {
		if isString(f) {
			files = append(files, map[string]interface{}{"glob": f})
		} else {
			if has(f, "git") {
				gitFiles := get(f, "git", "files")
				if isString(gitFiles) {
					gitFiles = []interface{}{gitFiles}
				}
				for _, glob := range getSlice(gitFiles) {
					files = append(files, map[string]interface{}{
						"glob": glob,
						"git": map[string]interface{}{
							"url": getString(f, "git", "url"),
							"rev": getString(f, "git", "rev"),
						},
					})
				}
			}
		}
	}

	return map[string]interface{}{
		"kind":      "Project",
		"files":     files,
		"name":      getString(obj, "name"),
		"variables": getMap(obj, "variables"),
	}
}

func toInternalService(obj interface{}) interface{} {
	toBuild := mapSlice(getSlice(obj, "artifacts"), toInternalEntry)
	tags := mapSlice(getSlice(obj, "tags"), toInternalEntry)
	releases := mapSlice(getSlice(obj, "releases"), toInternalEntry)

	toPush := []interface{}{}
	for i, v := range getSlice(obj, "artifacts") {
		e := toInternalEntry(i, or(get(v, "push"), v))
		if getString(e, "type") != "" {
			toPush = append(toPush, e)
		}
	}

	tasks := map[string]interface{}{}
	for k, v := range getMap(obj, "tasks") {
		tasks[k] = mapSlice(getSlice(v), toInternalEntry)
	}

	return map[string]interface{}{
		"kind": "Service",
		"name": getString(obj, "name"),
		"build": map[string]interface{}{
			"artifacts": map[string]interface{}{
				"toBuild": toBuild,
				"toPush":  toPush,
			},
			"tags": tags,
		},
		"deploy": map[string]interface{}{
			"releases": releases,
		},
		"run": map[string]interface{}{
			"tasks": tasks,
		},
	}
}

func toInternalEnvironment(obj interface{}) interface{} {
	return map[string]interface{}{
		"kind":           "Environment",
		"name":           getString(obj, "name"),
		"deployServices": getSlice(obj, "deployServices"),
		"variables":      getMap(obj, "variables"),
	}
}

func toInternalExecutor(obj interface{}) interface{} {
	schema, err := json.Marshal(getMap(obj, "schema"))
	if err != nil {
		panic(err)
	}
	return map[string]interface{}{
		"kind":   getString(obj, "kind"),
		"name":   getString(obj, "name"),
		"script": getString(obj, "script"),
		"schema": string(schema),
	}
}

func toInternalEntry(i int, obj interface{}) interface{} {
	if obj == false {
		return nil
	}
	if isString(obj) {
		return map[string]interface{}{"index": int64(i), "type": obj, "spec": nil}
	}
	for k, v := range getMap(obj) {
		if k != "push" {
			return map[string]interface{}{"index": int64(i), "type": k, "spec": v}
		}
	}
	return nil
}

func toV2(obj interface{}) interface{} {
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

func toV2Project(obj interface{}) interface{} {
	files := []interface{}{
		map[string]interface{}{
			"git": map[string]interface{}{
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

	return map[string]interface{}{
		"apiVersion": "g2a-cli/v2.0",
		"kind":       "Project",
		"name":       or(get(obj, "name"), "project"),
		"files":      files,
		"tasks":      map[string]interface{}{},
		"variables":  map[string]interface{}{},
	}
}

func toV2Service(obj interface{}) interface{} {
	artifacts := getSlice(obj, "build", "artifacts")

	if hooks := get(obj, "hooks", "pre-build"); hooks != nil {
		artifacts = prepend(artifacts, map[string]interface{}{
			"script": hooksToScript(hooks),
			"push":   false,
		})
	}
	if hooks := get(obj, "hooks", "post-build"); hooks != nil {
		artifacts = append(artifacts, map[string]interface{}{
			"script": hooksToScript(hooks),
			"push":   false,
		})
	}

	releases := getSlice(obj, "deploy", "releases")

	if hooks := get(obj, "hooks", "pre-deploy"); hooks != nil {
		releases = prepend(releases, map[string]interface{}{
			"script": hooksToScript(hooks),
		})
	}
	if hooks := get(obj, "hooks", "post-deploy"); hooks != nil {
		releases = append(releases, map[string]interface{}{
			"script": hooksToScript(hooks),
		})
	}

	tags := []interface{}{}

	for k, v := range getMap(obj, "build", "tagPolicy") {
		tags = append(tags, map[string]interface{}{k: v})
	}

	return map[string]interface{}{
		"apiVersion": "g2a-cli/v2.0",
		"kind":       "Service",
		"name":       getString(obj, "name"),
		"tags":       tags,
		"artifacts":  artifacts,
		"releases":   releases,
		"tasks":      map[string]interface{}{},
	}
}

func toV2Environment(obj interface{}) interface{} {
	return map[string]interface{}{
		"apiVersion":     "g2a-cli/v2.0",
		"kind":           "Environment",
		"name":           getString(obj, "name"),
		"deployServices": getSlice(obj, "deployServices"),
		"variables":      getMap(obj, "variables"),
	}
}

func hooksToScript(hooks interface{}) interface{} {
	sh := "set -e"
	for _, hook := range getSlice(hooks) {
		sh += fmt.Sprintf("\n%s", hook)
	}
	return map[string]interface{}{"sh": sh}
}

func getString(v interface{}, path ...interface{}) string {
	str, err := dyno.GetString(v, path...)
	if err != nil {
		return ""
	}
	return str
}

func isString(v interface{}, path ...interface{}) bool {
	_, err := dyno.GetString(v, path...)
	return err == nil
}

func getSlice(v interface{}, path ...interface{}) []interface{} {
	res, err := dyno.GetSlice(v, path...)
	if err != nil {
		return []interface{}{}
	}
	return res
}

func getMap(v interface{}, path ...interface{}) map[string]interface{} {
	res, err := dyno.GetMapS(v, path...)
	if err != nil {
		return map[string]interface{}{}
	}
	return res
}

func get(v interface{}, path ...interface{}) interface{} {
	val, err := dyno.Get(v, path...)
	if err != nil {
		return nil
	}
	return val
}

func has(v interface{}, path ...interface{}) bool {
	_, err := dyno.Get(v, path...)
	return err == nil
}

func or(v1 interface{}, v2 interface{}) interface{} {
	if v1 != nil {
		return v1
	} else {
		return v2
	}
}

func mapSlice(v []interface{}, fn func(int, interface{}) interface{}) []interface{} {
	res := make([]interface{}, len(v))
	for i := range v {
		res[i] = fn(i, v[i])
	}
	return res
}

func prepend(slice []interface{}, elems ...interface{}) []interface{} {
	return append(elems, slice...)
}
