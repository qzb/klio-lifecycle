package tengoutil

import (
	"errors"
	"fmt"
	"math"
	"reflect"

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
		if len(args) < t.NumIn() || (len(args) > t.NumIn() && !t.IsVariadic()) {
			return nil, tengo.ErrWrongNumArguments
		}

		// Convert tengo objects to list of reflect values.
		// If actual and expected types don't match - return runtime error.
		inputs := make([]reflect.Value, len(args))
		for i, arg := range args {
			var expected string
			var value reflect.Value
			var argType reflect.Type

			if i >= t.NumIn()-1 && t.IsVariadic() {
				argType = t.In(t.NumIn() - 1).Elem()
			} else {
				argType = t.In(i)
			}

			switch argType.Kind() {
			case reflect.String:
				expected = "string"
				if o, ok := arg.(*tengo.String); ok {
					value = reflect.ValueOf(o.Value)
				}
			case reflect.Int:
				expected = "int"
				if o, ok := arg.(*tengo.Int); ok {
					value = reflect.ValueOf(int(o.Value))
				}
			case reflect.Int64:
				expected = "int"
				if o, ok := arg.(*tengo.Int); ok {
					value = reflect.ValueOf(o.Value)
				}
			case reflect.Uint, reflect.Uint32, reflect.Uint64:
				expected = "int (>= 0)"
				if o, ok := arg.(*tengo.Int); ok && o.Value >= 0 {
					value = reflect.ValueOf(o.Value).Convert(argType)
				}
			case reflect.Bool:
				expected = "bool"
				if o, ok := arg.(*tengo.Bool); ok {
					value = reflect.ValueOf(!o.IsFalsy())
				}
			case reflect.Interface:
				expected = "any"
				if arg.TypeName() == "bool" {
					value = reflect.ValueOf(!arg.IsFalsy())
				} else {
					value = reflect.ValueOf(arg).Elem().FieldByName("Value")
				}
			default:
				expected = argType.Kind().String()
			}

			if !value.IsValid() {
				return nil, tengo.ErrInvalidArgumentType{
					Name:     fmt.Sprint(i + 1),
					Expected: expected,
					Found:    arg.TypeName(),
				}
			}

			inputs[i] = value
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
		return &tengo.UserFunction{Value: fn}, err

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
			tag := ft.Tag.Get("tengo")
			name, opts := parseTag(tag)
			if name == "" {
				name = ft.Name
			}
			if tag == "-" || name == "" || !ft.IsExported() || (fv.IsZero() && opts.Contains("omitempty")) {
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
