package tengoutil

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/d5/tengo/v2"
)

type TengoEncoder interface {
	EncodeTengoObject() (tengo.Object, error)
}

// ToObjects - converts input to tengo.Object
func ToObject(value any) (tengo.Object, error) {
	return toObject(value, false)
}

// ToImmutableObject - converts input to tengo.Object, uses immutable maps and arrays instead of regular ones.
func ToImmutableObject(value any) (tengo.Object, error) {
	return toObject(value, true)
}

// ToObjectsMap - runs ToObject on map entries, fixes function names
func ToObjectsMap(attrs map[string]any) (map[string]tengo.Object, error) {
	obj, err := toObject(attrs, false)
	if err != nil {
		return nil, err
	}
	return obj.(*tengo.Map).Value, nil
}

// ToCallableFunc - converts regular function to one compatible with Tengo.
// Returned function validates number and types of arguments.
func ToCallableFunc(fn any) (tengo.CallableFunc, error) {
	// Check if it's already a callable function
	if cf, ok := fn.(tengo.CallableFunc); ok {
		return cf, nil
	}

	// Prepare Type and Value
	t := reflect.TypeOf(fn)
	v := reflect.ValueOf(fn)

	// Only functions can be converted to functions.
	if v.Kind() != reflect.Func {
		panic(fmt.Sprintf("Cannot convert %s to tengo function", v.Kind()))
	}

	// Check what exactly this function returns.
	returnsError := t.NumOut() != 0 && t.Out(t.NumOut()-1).Implements(reflect.TypeOf((*error)(nil)).Elem())
	returnsResult := t.NumOut() == 2 || (t.NumOut() == 1 && !returnsError)

	if t.NumOut() > 2 || (t.NumOut() > 1 && !returnsError) {
		return nil, errors.New("Cannot convert functions returning more than 2 results")
	}

	// TODO: check if result can be converted to tengo

	// Return CallableFunc
	return func(args ...tengo.Object) (obj tengo.Object, err error) {
		// Return error on panics - when error is returned tengo aborts script execution and returns RuntimeError
		defer func() {
			if e, ok := recover().(error); ok && e != nil {
				err = e
			}
		}()

		if (t.IsVariadic() && len(args) < t.NumIn()-1) || (!t.IsVariadic() && len(args) != t.NumIn()) {
			return nil, tengo.ErrWrongNumArguments
		}

		// Convert tengo objects to list of reflect values.
		// If actual and expected types don't match - return runtime error.
		inputs := make([]reflect.Value, len(args))
		for i, arg := range args {
			var argType reflect.Type

			if i >= t.NumIn()-1 && t.IsVariadic() {
				argType = t.In(t.NumIn() - 1).Elem()
			} else {
				argType = t.In(i)
			}

			value := reflect.New(argType)
			err := decodeObject(arg, value)
			if err != nil {
				if err, ok := err.(*DecodingError); ok {
					return nil, tengo.ErrInvalidArgumentType{
						Name:     strings.Join(append([]string{fmt.Sprint(i + 1)}, err.Path...), "."),
						Expected: err.Expected,
						Found:    fmt.Sprintf("%s (%s)", err.Object.TypeName(), err.Object),
					}
				}
				return nil, err
			}

			inputs[i] = value.Elem()
		}

		// Run function
		outputs := v.Call(inputs)

		// Return
		if returnsError {
			if errValue := outputs[len(outputs)-1]; !errValue.IsNil() {
				return toObject(errValue.Interface(), false)
			} else if !returnsResult {
				return tengo.TrueValue, nil
			}
		}

		if returnsResult {
			return toObject(outputs[0].Interface(), false)
		}

		return tengo.UndefinedValue, nil
	}, nil
}

func toObject(value any, immutable bool) (_ tengo.Object, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	if value == nil {
		return tengo.UndefinedValue, nil
	}

	switch x := value.(type) {
	case TengoEncoder:
		return x.EncodeTengoObject()
	case tengo.Object:
		return x, nil
	case error:
		return &tengo.Error{Value: &tengo.String{Value: x.Error()}}, nil
	}

	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	switch t.Kind() {
	case reflect.Ptr:
		return toObject(v.Elem().Interface(), false)

	case reflect.String:
		return &tengo.String{Value: v.String()}, nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return &tengo.Int{Value: v.Convert(reflect.TypeOf(int(0))).Int()}, nil

	case reflect.Uint64:
		if v.Uint() > math.MaxInt64 {
			return nil, fmt.Errorf("value %d overflows int64", v.Uint())
		}
		return &tengo.Int{Value: int64(v.Uint())}, nil

	case reflect.Float32, reflect.Float64:
		return &tengo.Float{Value: v.Float()}, nil

	case reflect.Func:
		fn, err := ToCallableFunc(value)
		if err != nil {
			return nil, err
		}
		return &tengo.UserFunction{Value: fn}, nil

	case reflect.Bool:
		if v.Bool() {
			return tengo.TrueValue, nil
		} else {
			return tengo.FalseValue, nil
		}

	case reflect.Slice, reflect.Array:
		objects := make([]tengo.Object, v.Len())
		for i := range objects {
			objects[i], err = toObject(v.Index(i).Interface(), false)
			if err != nil {
				return nil, err
			}
		}
		if immutable {
			return &tengo.ImmutableArray{Value: objects}, nil
		} else {
			return &tengo.Array{Value: objects}, nil
		}

	case reflect.Map:
		if keyKind := t.Key().Kind(); keyKind != reflect.String {
			return nil, fmt.Errorf("only maps with string keys can be converted to tengo object, found %s", keyKind)
		}
		objects := make(map[string]tengo.Object, v.Len())
		iter := v.MapRange()
		for iter.Next() {
			key := iter.Key().String()
			objects[key], err = toObject(iter.Value().Interface(), false)
			if err != nil {
				return nil, err
			}
			if uf, ok := objects[key].(*tengo.UserFunction); ok {
				uf.Name = key
			}
		}
		if immutable {
			return &tengo.ImmutableMap{Value: objects}, nil
		} else {
			return &tengo.Map{Value: objects}, nil
		}

	case reflect.Struct:
		objects := map[string]tengo.Object{}
		for i := 0; i < t.NumField(); i++ {
			ft := t.Field(i)
			fv := v.Field(i)
			name, opts := parseTag(ft)
			if name == "" || !ft.IsExported() || (fv.IsZero() && opts.Contains("omitempty")) {
				continue
			}
			if opts.Contains("immutable") {
				objects[name], err = toObject(fv.Interface(), true)
			} else {
				objects[name], err = toObject(fv.Interface(), false)
			}
			if err != nil {
				return nil, err
			}
			if uf, ok := objects[name].(*tengo.UserFunction); ok {
				uf.Name = name
			}
		}
		return &tengo.ImmutableMap{Value: objects}, nil

	default:
		return nil, fmt.Errorf("%s cannot be converted to tengo object", t.Kind())
	}
}

type DecodingError struct {
	Expected string
	Object   tengo.Object
	Path     []string
}

func (e *DecodingError) Error() string {
	msg := fmt.Sprintf("expected %q, found %q", e.Expected, e.Object.TypeName())
	if len(e.Path) > 0 {
		msg += ", path: " + strings.Join(e.Path, ".")
	}
	return msg
}

func DecodeObject(obj tengo.Object, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("FromObject accepts only non-empty pointers")
	}
	return decodeObject(obj, rv.Elem())
}

func decodeObject(obj tengo.Object, v reflect.Value) (err error) {
	t := v.Type()

	switch v.Kind() {
	case reflect.Pointer:
		if _, ok := obj.(*tengo.Undefined); ok {
			return
		}
		return decodeObject(obj, v.Elem())

	case reflect.String:
		if o, ok := obj.(*tengo.String); ok {
			v.Set(reflect.ValueOf(o.Value).Convert(t))
			return
		} else {
			return &DecodingError{Object: obj, Expected: "string"}
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if o, ok := obj.(*tengo.Int); ok {
			if v.OverflowInt(o.Value) {
				return &DecodingError{Object: obj, Expected: fmt.Sprintf("int (within %s value range)", t.Kind())}
			}
			v.Set(reflect.ValueOf(o.Value).Convert(t))
			return
		} else {
			return &DecodingError{Object: obj, Expected: "int"}
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if o, ok := obj.(*tengo.Int); ok {
			if o.Value < 0 || v.OverflowUint(uint64(o.Value)) {
				return &DecodingError{Object: obj, Expected: fmt.Sprintf("int (within %s value range)", t.Kind())}
			}
			v.Set(reflect.ValueOf(o.Value).Convert(t))
		} else {
			return &DecodingError{Object: obj, Expected: "int (>= 0)"}
		}

	case reflect.Float32, reflect.Float64:
		if o, ok := obj.(*tengo.Float); ok {
			if v.OverflowFloat(o.Value) {
				return &DecodingError{Object: obj, Expected: fmt.Sprintf("float (within %s value range)", t.Kind())}
			}
			v.Set(reflect.ValueOf(o.Value).Convert(t))
			return
		} else {
			return &DecodingError{Object: obj, Expected: "float"}
		}

	case reflect.Bool:
		if o, ok := obj.(*tengo.Bool); ok {
			v.Set(reflect.ValueOf(!o.IsFalsy()).Convert(t))
			return
		} else {
			return &DecodingError{Object: obj, Expected: "bool"}
		}

	case reflect.Slice:
		expected := "array"
		if t.Elem().Kind() == reflect.Uint8 {
			expected = "bytes"
			if o, ok := obj.(*tengo.Bytes); ok {
				v.Set(reflect.ValueOf(append([]byte{}, o.Value...)))
				return
			}
		}
		var items []tengo.Object
		switch o := obj.(type) {
		case *tengo.Array:
			items = o.Value
		case *tengo.ImmutableArray:
			items = o.Value
		default:
			return &DecodingError{Object: obj, Expected: expected}
		}
		s := reflect.MakeSlice(t, len(items), len(items))
		for i, o := range items {
			err := decodeObject(o, s.Index(i))
			if err != nil {
				if e, ok := err.(*DecodingError); ok {
					e.Path = append([]string{fmt.Sprint(i)}, e.Path...)
				}
				return err
			}
		}
		v.Set(s)

	// TODO: error when map has some extra keys
	case reflect.Struct:
		var entries map[string]tengo.Object
		switch o := obj.(type) {
		case *tengo.Map:
			entries = o.Value
		case *tengo.ImmutableMap:
			entries = o.Value
		default:
			return &DecodingError{Object: obj, Expected: "map"}
		}
		mapping := map[string][]int{}
		for _, f := range reflect.VisibleFields(t) {
			name, _ := parseTag(f)
			if name == "" || !f.IsExported() || f.Anonymous {
				continue
			}
			mapping[name] = f.Index
		}
		for key, val := range entries {
			if _, ok := val.(*tengo.Undefined); ok {
				continue
			}
			i, ok := mapping[key]
			if !ok {
				return &DecodingError{Object: val, Path: []string{key}, Expected: "undefined"}
			}
			err := decodeObject(val, v.FieldByIndex(i))
			if err != nil {
				if err, ok := err.(*DecodingError); ok {
					err.Path = append([]string{key}, err.Path...)
				}
				return err
			}
		}

	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			return errors.New("decoding to maps is supported only when they have string keys")
		}
		var entries map[string]tengo.Object
		switch o := obj.(type) {
		case *tengo.Map:
			entries = o.Value
		case *tengo.ImmutableMap:
			entries = o.Value
		default:
			return &DecodingError{Object: obj, Expected: "map"}
		}
		for key, val := range entries {
			x := reflect.New(t.Elem()).Elem()
			err := decodeObject(val, x)
			if err != nil {
				if err, ok := err.(*DecodingError); ok {
					err.Path = append([]string{key}, err.Path...)
				}
				return err
			}
			v.SetMapIndex(reflect.ValueOf(key), x)
		}

	case reflect.Interface:
		if t.NumMethod() != 0 {
			if t.Implements(reflect.TypeOf((*tengo.Object)(nil)).Elem()) {
				v.Set(reflect.ValueOf(obj))
				return
			}
			return fmt.Errorf("decoding to not empty interfaces is not supported")
		}
		var x reflect.Value
		switch o := obj.(type) {
		case *tengo.Undefined:
			v.Set(reflect.New(reflect.TypeOf((*any)(nil)).Elem()).Elem())
			return nil
		case *tengo.Int:
			x = reflect.New(reflect.TypeOf(0))
		case *tengo.Float:
			x = reflect.New(reflect.TypeOf(.0))
		case *tengo.Bool:
			x = reflect.New(reflect.TypeOf(false))
		case *tengo.Char:
			x = reflect.New(reflect.TypeOf(rune(0)))
		case *tengo.String:
			x = reflect.New(reflect.TypeOf(""))
		case *tengo.Bytes:
			x = reflect.New(reflect.SliceOf(reflect.TypeOf(byte(0))))
		case *tengo.Array, *tengo.ImmutableArray:
			x = reflect.New(reflect.SliceOf(reflect.TypeOf((*any)(nil)).Elem()))
		case *tengo.Map, *tengo.ImmutableMap:
			m := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf((*any)(nil)).Elem()))
			x = reflect.New(m.Type())
			x.Elem().Set(m)
		case *tengo.Error:
			return decodeObject(o.Value, v)
		default:
			return fmt.Errorf("unsupported conversion from interface for type: %s", obj.TypeName())
		}
		err := decodeObject(obj, x.Elem())
		if err != nil {
			return err
		}
		v.Set(x.Elem())

	default:
		return fmt.Errorf("unsupported kind: %s", t.Kind())
	}

	return
}
