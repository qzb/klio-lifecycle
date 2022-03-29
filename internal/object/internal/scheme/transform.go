package scheme

import (
	"encoding/json"
	"fmt"

	"github.com/icza/dyno"
)

func ToInternal(content interface{}) (interface{}, error) {
	return toInternal(dyno.ConvertMapI2MapS(content)), nil
}

func toInternal(obj interface{}) interface{} {
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
	return map[string]interface{}{
		"kind":      "Project",
		"files":     getSlice(obj, "files"),
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
