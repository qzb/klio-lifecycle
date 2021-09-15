// based on https://github.com/barkimedes/go-deepcopy
package placeholders

import (
	"errors"
	"fmt"
	. "reflect"
	"regexp"
	"sort"
	"strings"
)

type copier func(interface{}, map[uintptr]interface{}, map[string]string, bool) (interface{}, error)

const recursionDepthLimit = 10

var placeholderRegexp *regexp.Regexp
var copiers map[Kind]copier

func init() {
	placeholderRegexp = regexp.MustCompile(`{{\s*([A-Za-z0-9_.]+)\s*}}`)
	copiers = map[Kind]copier{
		Bool:       processPrimitive,
		Int:        processPrimitive,
		Int8:       processPrimitive,
		Int16:      processPrimitive,
		Int32:      processPrimitive,
		Int64:      processPrimitive,
		Uint:       processPrimitive,
		Uint8:      processPrimitive,
		Uint16:     processPrimitive,
		Uint32:     processPrimitive,
		Uint64:     processPrimitive,
		Uintptr:    processPrimitive,
		Float32:    processPrimitive,
		Float64:    processPrimitive,
		Complex64:  processPrimitive,
		Complex128: processPrimitive,
		Array:      processArray,
		Map:        processMap,
		Ptr:        processPointer,
		Slice:      processSlice,
		String:     processString,
		Struct:     processStruct,
	}
}

// ProcessStruct recursively replaces placeholders with specified values in every
// string field in the struct.
func ProcessStruct(inputStruct interface{}, values map[string]interface{}) (interface{}, error) {
	ptrs := make(map[uintptr]interface{})
	return processAnything(inputStruct, ptrs, normalizeValues(values), true)
}

func replacePlaceholders(str string, values map[string]string) (string, error) {
	var err error

	result := []byte(str)

	for depth := 0; placeholderRegexp.Match(result) && err == nil; depth++ {
		if depth > recursionDepthLimit {
			return "", errors.New("exceeded recursion limit while replacing placeholders")
		}

		result = placeholderRegexp.ReplaceAllFunc(result, func(s []byte) []byte {
			placeholderName := strings.ToLower(placeholderRegexp.FindStringSubmatch(string(s))[1])
			value, ok := values[placeholderName]
			if !ok && err == nil {
				placeholders := []string{}
				for k := range values {
					placeholders = append(placeholders, k)
				}
				placeholders = sort.StringSlice(placeholders)
				err = fmt.Errorf("value for {{%s}} placeholder is not specified, available placeholders:\n  %s", placeholderName, strings.Join(placeholders, ", "))
			}

			return []byte(value)
		})
	}

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func normalizeValues(params map[string]interface{}) map[string]string {
	result := map[string]string{}

	for k, v := range params {
		switch val := v.(type) {
		case string:
			result[strings.ToLower("."+k)] = val
		case map[string]string:
			for k2, v2 := range val {
				result[strings.ToLower("."+k+"."+k2)] = v2
			}
		case map[string]interface{}:
			for k2, v2 := range normalizeValues(val) {
				result[strings.ToLower("."+k+k2)] = v2
			}
		default:
			panic("invalid params type")
		}
	}

	return result
}

// Primitive makes a copy of a primitive type...which just means it returns the input value.
// This is wholly uninteresting, but I included it for consistency's sake.
func processPrimitive(x interface{}, ptrs map[uintptr]interface{}, values map[string]string, enabled bool) (interface{}, error) {
	kind := ValueOf(x).Kind()
	if kind == Array || kind == Chan || kind == Func || kind == Interface || kind == Map || kind == Ptr || kind == Slice || kind == Struct || kind == UnsafePointer {
		return nil, fmt.Errorf("unable to copy %v (a %v) as a primitive", x, kind)
	}
	return x, nil
}

func processAnything(x interface{}, ptrs map[uintptr]interface{}, values map[string]string, enabled bool) (interface{}, error) {
	v := ValueOf(x)
	if !v.IsValid() {
		return x, nil
	}
	if c, ok := copiers[v.Kind()]; ok {
		return c(x, ptrs, values, enabled)
	}
	t := TypeOf(x)
	return nil, fmt.Errorf("unable to make a deep copy of %v (type: %v) - kind %v is not supported", x, t, v.Kind())
}

func processSlice(x interface{}, ptrs map[uintptr]interface{}, values map[string]string, enabled bool) (interface{}, error) {
	v := ValueOf(x)
	if v.Kind() != Slice {
		return nil, fmt.Errorf("must pass a value with kind of Slice; got %v", v.Kind())
	}
	// Create a new slice and, for each item in the slice, make a deep copy of it.
	size := v.Len()
	t := TypeOf(x)
	dc := MakeSlice(t, size, size)
	for i := 0; i < size; i++ {
		item, err := processAnything(v.Index(i).Interface(), ptrs, values, enabled)
		if err != nil {
			return nil, err
		}
		iv := ValueOf(item)
		if iv.IsValid() {
			dc.Index(i).Set(iv)
		}
	}
	return dc.Interface(), nil
}

func processMap(x interface{}, ptrs map[uintptr]interface{}, values map[string]string, enabled bool) (interface{}, error) {
	v := ValueOf(x)
	if v.Kind() != Map {
		return nil, fmt.Errorf("must pass a value with kind of Map; got %v", v.Kind())
	}
	t := TypeOf(x)
	dc := MakeMapWithSize(t, v.Len())
	iter := v.MapRange()
	for iter.Next() {
		item, err := processAnything(iter.Value().Interface(), ptrs, values, enabled)
		if err != nil {
			return nil, err
		}
		k, err := processAnything(iter.Key().Interface(), ptrs, values, enabled)
		if err != nil {
			return nil, err
		}
		dc.SetMapIndex(ValueOf(k), ValueOf(item))
	}
	return dc.Interface(), nil
}

func processPointer(x interface{}, ptrs map[uintptr]interface{}, values map[string]string, enabled bool) (interface{}, error) {
	v := ValueOf(x)
	if v.Kind() != Ptr {
		return nil, fmt.Errorf("must pass a value with kind of Ptr; got %v", v.Kind())
	}
	addr := v.Pointer()
	if dc, ok := ptrs[addr]; ok {
		return dc, nil
	}
	t := TypeOf(x)
	dc := New(t.Elem())
	ptrs[addr] = dc.Interface()
	if !v.IsNil() {
		item, err := processAnything(v.Elem().Interface(), ptrs, values, enabled)
		if err != nil {
			return nil, err
		}
		iv := ValueOf(item)
		if iv.IsValid() {
			dc.Elem().Set(ValueOf(item))
		}
	}
	return dc.Interface(), nil
}

func processStruct(x interface{}, ptrs map[uintptr]interface{}, values map[string]string, enabled bool) (interface{}, error) {
	v := ValueOf(x)
	if v.Kind() != Struct {
		return nil, fmt.Errorf("must pass a value with kind of Struct; got %v", v.Kind())
	}
	t := TypeOf(x)
	dc := New(t)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}
		item, err := processAnything(v.Field(i).Interface(), ptrs, values, enabled && f.Tag.Get("placeholders") != "disable")
		if err != nil {
			return nil, err
		}
		dc.Elem().Field(i).Set(ValueOf(item))
	}
	return dc.Elem().Interface(), nil
}

func processArray(x interface{}, ptrs map[uintptr]interface{}, values map[string]string, enabled bool) (interface{}, error) {
	v := ValueOf(x)
	if v.Kind() != Array {
		return nil, fmt.Errorf("must pass a value with kind of Array; got %v", v.Kind())
	}
	t := TypeOf(x)
	size := t.Len()
	dc := New(ArrayOf(size, t.Elem())).Elem()
	for i := 0; i < size; i++ {
		item, err := processAnything(v.Index(i).Interface(), ptrs, values, enabled)
		if err != nil {
			return nil, err
		}
		dc.Index(i).Set(ValueOf(item))
	}
	return dc.Interface(), nil
}

func processString(x interface{}, ptrs map[uintptr]interface{}, values map[string]string, enabled bool) (interface{}, error) {
	if ValueOf(x).Kind() != String {
		return nil, fmt.Errorf("must pass a value with kind of String; got %v", ValueOf(x).Kind())
	}
	if str, ok := x.(string); ok && enabled {
		return replacePlaceholders(str, values)
	}
	return x, nil
}
