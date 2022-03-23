// Package placeholders implements simple templating format used within Lifecycle's config files.
package placeholders

import (
	"fmt"
	"reflect"
)

// ReplaceFunc is a type of function used for replacing placeholders with
// values. It accepts case-sensitive name of the placeholder and returns value
// and error.
type ReplaceFunc func(name string) (value string, err error)

// ReplaceWithValues replaces placeholders within strings in input data using
// provided values. It supports only the data structures produced by
// yaml.Unmarshal() when unmarshalling to interface{}, which includes: string,
// int, float64, bool, []interface{}, map[interface{}]interface{}.
//
// Replacement values are a string to string map, but for convenience nested
// values may be represented with a nested maps, for example: { "foo.bar":
// "value" } may be also represented as: { "foo": { "bar": "value" } }.
//
// Replacement values keys doesn't include leading ".".
func ReplaceWithValues(input interface{}, values map[string]interface{}) (interface{}, error) {
	collection, err := newValuesCollection(values)
	if err != nil {
		return nil, err
	}

	return Replace(input, collection.Get)
}

// Replace replaces placeholders within strings in input data using provided
// function. It supports only the data structures produced by yaml.Unmarshal()
// when unmarshalling to interface{}, which includes: string, int, float64,
// bool, []interface{}, map[interface{}]interface{}.
//
// Placeholder name provided to replace function is case-sensitive.
func Replace(input interface{}, fn ReplaceFunc) (interface{}, error) {
	res, err := process(reflect.ValueOf(input), func(str string) (string, error) {
		return replaceMarkers(str, fn)
	})
	if err != nil {
		return nil, err
	}

	return res.Interface(), nil
}

// process crawls through provided data structure and generates it's copy with
// all strings replaced using provided function.
func process(v reflect.Value, fn ReplaceFunc) (res reflect.Value, err error) {
	switch v.Kind() {
	case reflect.Map:
		return processMap(v, fn)
	case reflect.Slice:
		return processSlice(v, fn)
	case reflect.String:
		return processString(v, fn)
	case reflect.Float64, reflect.Int, reflect.Bool:
		return v, nil
	case reflect.Interface:
		v = reflect.ValueOf(v.Interface())
		if v.Kind() != reflect.Interface {
			return process(v, fn)
		}
		fallthrough
	default:
		return res, fmt.Errorf("processing placeholders in %q is not supported", v.Kind())
	}
}

func processMap(v reflect.Value, fn ReplaceFunc) (res reflect.Value, err error) {
	m := reflect.MakeMapWithSize(v.Type(), v.Len())
	for iter := v.MapRange(); iter.Next(); {
		mk, err := process(iter.Key(), fn)
		if err != nil {
			return res, err
		}
		mv, err := process(iter.Value(), fn)
		if err != nil {
			return res, err
		}
		m.SetMapIndex(mk, mv)
	}
	return m, nil
}

func processSlice(v reflect.Value, fn ReplaceFunc) (res reflect.Value, err error) {
	s := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
	for i := 0; i < v.Len(); i++ {
		si, err := process(v.Index(i), fn)
		if err != nil {
			return res, err
		}
		s.Index(i).Set(si)
	}
	return s, nil
}

func processString(v reflect.Value, fn ReplaceFunc) (res reflect.Value, err error) {
	str, err := fn(v.String())
	if err == nil {
		res = reflect.ValueOf(str)
	}
	return
}
