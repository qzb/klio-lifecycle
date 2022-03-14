// Package placeholders implements simple templating format used within Lifecycle's config files.
package placeholders

import (
	"fmt"
	"reflect"
	"regexp"
)

var placeholderRegexp *regexp.Regexp = regexp.MustCompile(`{{\s*([A-Za-z0-9_.]+)\s*}}`)

// Replace replaces placeholders within strings in input data using provided
// values. It supports only the data structures produced by yaml.Unmarshal()
// when unmarshalling to interface{}, which includes: string, int, float64,
// bool, []interface{}, map[interface{}]interface{}.
//
// Replacement values are a string to string map, but for convenience nested
// values may be represented with a nested maps, for example: { "foo.bar":
// "value" } may be also represented as: { "foo": { "bar": "value" } }.
//
// Replacement values keys doesn't include leading ".".
func Replace(input interface{}, values map[string]interface{}) (interface{}, error) {
	v, err := newValuesCollection(values)
	if err != nil {
		return nil, err
	}
	res, err := process(reflect.ValueOf(input), v)
	if err != nil {
		return nil, err
	}
	return res.Interface(), nil
}

func process(v reflect.Value, values *valuesCollection) (res reflect.Value, err error) {
	switch v.Kind() {
	case reflect.Map:
		return processMap(v, values)
	case reflect.Slice:
		return processSlice(v, values)
	case reflect.String:
		return processString(v, values)
	case reflect.Float64, reflect.Int, reflect.Bool:
		return v, nil
	case reflect.Interface:
		v = reflect.ValueOf(v.Interface())
		if v.Kind() != reflect.Interface {
			return process(v, values)
		}
		fallthrough
	default:
		return res, fmt.Errorf("replacing placeholders in %q is not supported", v.Kind())
	}
}

func processMap(v reflect.Value, values *valuesCollection) (res reflect.Value, err error) {
	m := reflect.MakeMapWithSize(v.Type(), v.Len())
	for iter := v.MapRange(); iter.Next(); {
		mk, err := process(iter.Key(), values)
		if err != nil {
			return res, err
		}
		mv, err := process(iter.Value(), values)
		if err != nil {
			return res, err
		}
		m.SetMapIndex(mk, mv)
	}
	return m, nil
}

func processSlice(v reflect.Value, values *valuesCollection) (res reflect.Value, err error) {
	s := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
	for i := 0; i < v.Len(); i++ {
		si, err := process(v.Index(i), values)
		if err != nil {
			return res, err
		}
		s.Index(i).Set(si)
	}
	return s, nil
}

func processString(v reflect.Value, values *valuesCollection) (res reflect.Value, err error) {
	str := placeholderRegexp.ReplaceAllStringFunc(v.String(), func(s string) (value string) {
		value, err = values.Get(placeholderRegexp.FindStringSubmatch(s)[1])
		return
	})

	if err == nil {
		res = reflect.ValueOf(str)
	}

	return
}
