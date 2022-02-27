package tengoutil

import (
	"errors"
	"fmt"
	"io/fs"
	"math"
	"reflect"
	"testing"

	"github.com/d5/tengo/v2"
	"github.com/stretchr/testify/assert"
)

type testEncoder struct {
	object tengo.Object
	err    error
}

func (t testEncoder) EncodeTengoObject() (tengo.Object, error) {
	return t.object, t.err
}

// TODO: Split it into proper test cases
func TestToObject(t *testing.T) {
	cases := []struct {
		input    any
		expected tengo.Object
		assertFn func(tengo.Object) error
		fails    bool
	}{
		{ // 1
			input:    1,
			expected: &tengo.Int{Value: 1},
		},
		{ // 2
			input:    uint(1),
			expected: &tengo.Int{Value: 1},
		},
		{ // 3
			input:    float64(1),
			expected: &tengo.Float{Value: 1},
		},
		{ // 4
			input:    float32(1),
			expected: &tengo.Float{Value: 1},
		},
		{ // 5
			input: uint64(math.MaxUint64),
			fails: true,
		},
		{ // 6
			input:    "string",
			expected: &tengo.String{Value: "string"},
		},
		{ // 7
			input:    '0',
			expected: &tengo.Int{Value: '0'},
		},
		{ // 8
			input:    true,
			expected: tengo.TrueValue,
		},
		{ // 9
			input:    false,
			expected: tengo.FalseValue,
		},
		{ // 10
			input:    errors.New("abc"),
			expected: &tengo.Error{Value: &tengo.String{Value: "abc"}},
		},
		{ // 11
			input:    false,
			expected: tengo.FalseValue,
		},
		{ // 12
			input:    fs.FileMode(1),
			expected: &tengo.Int{Value: 1},
		},
		{ // 13
			input:    testEncoder{object: &tengo.String{Value: "abc"}, err: nil},
			expected: &tengo.String{Value: "abc"},
		},
		{ // 14
			input: testEncoder{object: nil, err: errors.New("")},
			fails: true,
		},
		{ // 15
			input:    map[string]any{},
			expected: &tengo.Map{Value: map[string]tengo.Object{}},
		},
		{ // 16
			input: map[any]any{},
			fails: true,
		},
		{ // 17
			input: map[string]any{
				"string": "abc",
				"float":  12.3,
				"slice":  []bool{true, false},
			},
			expected: &tengo.Map{Value: map[string]tengo.Object{
				"string": &tengo.String{Value: "abc"},
				"float":  &tengo.Float{Value: 12.3},
				"slice":  &tengo.Array{Value: []tengo.Object{tengo.TrueValue, tengo.FalseValue}},
			}},
		},
		{ // 18
			input: []any{
				"abc",
				12.3,
				[]bool{true, false},
			},
			expected: &tengo.Array{Value: []tengo.Object{
				&tengo.String{Value: "abc"},
				&tengo.Float{Value: 12.3},
				&tengo.Array{Value: []tengo.Object{tengo.TrueValue, tengo.FalseValue}},
			}},
		},
		{ // 19
			input: [...]any{
				"abc",
				12.3,
				[]bool{true, false},
			},
			expected: &tengo.Array{Value: []tengo.Object{
				&tengo.String{Value: "abc"},
				&tengo.Float{Value: 12.3},
				&tengo.Array{Value: []tengo.Object{tengo.TrueValue, tengo.FalseValue}},
			}},
		},
		{ // 20
			input:    func() *int { i := 1; return &i }(),
			expected: &tengo.Int{Value: 1},
		},
		{ // 21
			input: func(x string) string { return x },
			assertFn: func(o tengo.Object) error {
				res, err := o.(*tengo.UserFunction).Call(&tengo.String{Value: "abc"})
				if err != nil {
					return err
				}
				if !res.Equals(&tengo.String{Value: "abc"}) {
					return fmt.Errorf("expected 'abc', got %s", res)
				}
				return nil
			},
		},
		{ // 22
			input: func() any {
				type Inner struct {
					Field string
				}
				type Outer struct {
					Inner
				}
				return Outer{Inner{"string"}}
			}(),
			expected: &tengo.ImmutableMap{Value: map[string]tengo.Object{
				"Inner": &tengo.ImmutableMap{Value: map[string]tengo.Object{
					"Field": &tengo.String{Value: "string"},
				}},
			}},
		},
		{ // 23
			input: struct {
				Public  bool
				private bool
			}{},
			expected: &tengo.ImmutableMap{Value: map[string]tengo.Object{"Public": tengo.FalseValue}},
		},
		{ // 23
			input: struct {
				A bool `tengo:""`
				B bool `tengo:"b"`
				C bool `tengo:"-"`
				D bool `tengo:"-,"`
			}{},
			expected: &tengo.ImmutableMap{Value: map[string]tengo.Object{"A": tengo.FalseValue, "b": tengo.FalseValue, "-": tengo.FalseValue}},
		},
		{ // 23
			input: struct {
				A bool `tengo:",omitempty"`
				B bool `tengo:"b,omitempty"`
				C bool `tengo:"-,omitempty"`
			}{},
			expected: &tengo.ImmutableMap{Value: map[string]tengo.Object{}},
		},
		{ // 24
			input: struct {
				bool `tengo:"bool"` // it's not exported, so tag should be ignored
				string
				int
			}{},
			expected: &tengo.ImmutableMap{Value: map[string]tengo.Object{}},
		},
		{ // 25
			input: struct {
				Map    map[string]any `tengo:",immutable"`
				Array  []bool         `tengo:",immutable"`
				String string         `tengo:",immutable"`
			}{
				Map: map[string]any{
					"a": []bool{},
					"b": map[string]int{},
				},
			},
			expected: &tengo.ImmutableMap{Value: map[string]tengo.Object{
				"Map": &tengo.ImmutableMap{Value: map[string]tengo.Object{
					"a": &tengo.Array{Value: []tengo.Object{}},
					"b": &tengo.Map{Value: map[string]tengo.Object{}},
				}},
				"Array":  &tengo.ImmutableArray{Value: []tengo.Object{}},
				"String": &tengo.String{},
			}},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			o, err := ToObject(c.input)
			if c.fails {
				if err == nil {
					t.Errorf("expected error")
				}
				if o != nil {
					t.Errorf("expected result to be nil, found: %v", o)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
				if c.expected != nil && !reflect.DeepEqual(c.expected, o) {
					t.Errorf("\nexpected: %s\nfound: %s\n", c.expected, o)
				}
				if c.assertFn != nil {
					if err := c.assertFn(o); err != nil {
						t.Error(err)
					}
				}
			}
		})
	}
}

func Test_decoding_to_not_a_pointer_returns_error(t *testing.T) {
	object := tengo.Int{Value: 123}
	result := 0

	err := DecodeObject(&object, result)

	assert.Error(t, err)
}

func Test_decoding_to_nil_returns_error(t *testing.T) {
	object := tengo.Int{Value: 123}
	var result *int

	err := DecodeObject(&object, result)

	assert.Error(t, err)
}

func Test_decoding_to_nested_pointers_is_valid(t *testing.T) {
	object := tengo.Int{Value: 123}
	result := 0
	pointer := &result
	nestedPointer := &pointer

	err := DecodeObject(&object, nestedPointer)

	assert.NoError(t, err)
	assert.Equal(t, 123, result)
}

func Test_decoding_int_within_int_range_to_int_is_valid(t *testing.T) {
	values := []int64{math.MinInt, 0, math.MaxInt}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := int(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, int(value), result)
	}
}

func Test_decoding_int_within_int8_range_to_int8_is_valid(t *testing.T) {
	values := []int64{math.MinInt8, 0, math.MaxInt8}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := int8(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, int8(value), result)
	}
}

func Test_decoding_int_within_int16_range_to_int16_is_valid(t *testing.T) {
	values := []int64{math.MinInt16, 0, math.MaxInt16}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := int16(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, int16(value), result)
	}
}

func Test_decoding_int_within_int32_range_to_int32_is_valid(t *testing.T) {
	values := []int64{math.MinInt32, 0, math.MaxInt32}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := int32(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, int32(value), result)
	}
}

func Test_decoding_int_within_int64_range_to_int64_is_valid(t *testing.T) {
	values := []int64{math.MinInt64, 0, math.MaxInt64}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := int64(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, int64(value), result)
	}
}

func Test_decoding_int_within_uint_range_to_uint_is_valid(t *testing.T) {
	values := []int64{0, math.MaxInt64}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, uint(value), result)
	}
}

func Test_decoding_int_within_uint8_range_to_uint8_is_valid(t *testing.T) {
	values := []int64{0, math.MaxUint8}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint8(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, uint8(value), result)
	}
}

func Test_decoding_int_within_uint16_range_to_uint16_is_valid(t *testing.T) {
	values := []int64{0, math.MaxUint16}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint16(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, uint16(value), result)
	}
}

func Test_decoding_int_within_uint32_range_to_uint32_is_valid(t *testing.T) {
	values := []int64{0, math.MaxUint32}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint32(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, uint32(value), result)
	}
}

func Test_decoding_int_within_uint64_range_to_uint64_is_valid(t *testing.T) {
	values := []int64{0, math.MaxInt64}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint64(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, uint64(value), result)
	}
}

func Test_decoding_int_to_type_compatible_with_int_is_valid(t *testing.T) {
	type custom int
	object := tengo.Int{Value: 123}
	result := custom(0)

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, custom(123), result)
}

func Test_decoding_int_to_any_is_valid(t *testing.T) {
	object := tengo.Int{Value: 123}
	result := any(0)

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, any(123), result)
}

func Test_decoding_int_to_float64_returns_error(t *testing.T) {
	object := tengo.Int{Value: 0}
	result := float64(0)

	err := DecodeObject(&object, &result)

	assert.Equal(t, &DecodingError{Object: &object, Expected: "float"}, err)
}

func Test_decoding_int_to_bool_returns_error(t *testing.T) {
	object := tengo.Int{Value: 0}
	result := false

	err := DecodeObject(&object, &result)

	assert.Equal(t, &DecodingError{Object: &object, Expected: "bool"}, err)
}

func Test_decoding_int_overflowing_int8_to_int8_returns_error(t *testing.T) {
	values := []int64{math.MinInt8 - 1, math.MaxInt8 + 1}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := int8(0)

		err := DecodeObject(&object, &result)

		assert.Equal(t, &DecodingError{Object: &object, Expected: "int (within int8 value range)"}, err)
	}
}

func Test_decoding_int_overflowing_int16_to_int16_returns_error(t *testing.T) {
	values := []int64{math.MinInt16 - 1, math.MaxInt16 + 1}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := int16(0)

		err := DecodeObject(&object, &result)

		assert.Equal(t, &DecodingError{Object: &object, Expected: "int (within int16 value range)"}, err)
	}
}

func Test_decoding_int_overflowing_int32_to_int32_returns_error(t *testing.T) {
	values := []int64{math.MinInt32 - 1, math.MaxInt32 + 1}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := int32(0)

		err := DecodeObject(&object, &result)

		assert.Equal(t, &DecodingError{Object: &object, Expected: "int (within int32 value range)"}, err)
	}
}

func Test_decoding_int_overflowing_uint_to_uint_returns_error(t *testing.T) {
	values := []int64{-1}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint(0)

		err := DecodeObject(&object, &result)

		assert.Equal(t, &DecodingError{Object: &object, Expected: "int (within uint value range)"}, err)
	}
}

func Test_decoding_int_overflowing_uint8_to_uint8_returns_error(t *testing.T) {
	values := []int64{-1, math.MaxUint8 + 1}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint8(0)

		err := DecodeObject(&object, &result)

		assert.Equal(t, &DecodingError{Object: &object, Expected: "int (within uint8 value range)"}, err)
	}
}

func Test_decoding_int_overflowing_uint16_to_uint16_returns_error(t *testing.T) {
	values := []int64{-1, math.MaxUint16 + 1}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint16(0)

		err := DecodeObject(&object, &result)

		assert.Equal(t, &DecodingError{Object: &object, Expected: "int (within uint16 value range)"}, err)
	}
}

func Test_decoding_int_overflowing_uint32_to_uint32_returns_error(t *testing.T) {
	values := []int64{-1, math.MaxUint32 + 1}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint32(0)

		err := DecodeObject(&object, &result)

		assert.Equal(t, &DecodingError{Object: &object, Expected: "int (within uint32 value range)"}, err)
	}
}

func Test_decoding_int_overflowing_uint64_to_uint64_returns_error(t *testing.T) {
	values := []int64{-1}
	for _, value := range values {
		object := tengo.Int{Value: value}
		result := uint64(0)

		err := DecodeObject(&object, &result)

		assert.Equal(t, &DecodingError{Object: &object, Expected: "int (within uint64 value range)"}, err)
	}
}

func Test_decoding_float_within_float32_range_to_float32_is_valid(t *testing.T) {
	values := []float64{math.SmallestNonzeroFloat32, 0, math.MaxFloat32, math.Inf(1)}
	for _, value := range values {
		object := tengo.Float{Value: value}
		result := float32(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, float32(value), result)
	}
}

func Test_decoding_float_within_float64_range_to_float64_is_valid(t *testing.T) {
	values := []float64{math.SmallestNonzeroFloat32, 0, math.MaxFloat64, math.Inf(1)}
	for _, value := range values {
		object := tengo.Float{Value: value}
		result := float64(0)

		err := DecodeObject(&object, &result)

		assert.NoError(t, err)
		assert.Equal(t, float64(value), result)
	}
}

func Test_decoding_float_to_type_compatible_with_float64_is_valid(t *testing.T) {
	type custom float64
	object := tengo.Float{Value: 123}
	result := custom(0)

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, custom(123), result)
}

func Test_decoding_float_to_any_is_valid(t *testing.T) {
	object := tengo.Float{Value: 123}
	result := any(0)

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, any(float64(123)), result)
}

func Test_decoding_float_overflowing_float32_to_float32_returns_error(t *testing.T) {
	object := tengo.Float{Value: math.MaxFloat64}
	result := float32(0)

	err := DecodeObject(&object, &result)

	assert.Equal(t, &DecodingError{Object: &object, Expected: "float (within float32 value range)"}, err)
}

func Test_decoding_float_to_int_returns_error(t *testing.T) {
	object := tengo.Float{Value: 0}
	result := int(0)

	err := DecodeObject(&object, &result)

	assert.Equal(t, &DecodingError{Object: &object, Expected: "int"}, err)
}

func Test_decoding_string_to_string_is_valid(t *testing.T) {
	object := tengo.String{Value: "test"}
	result := ""

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, "test", result)
}

func Test_decoding_string_to_type_compatible_with_string_is_valid(t *testing.T) {
	type custom string
	object := tengo.String{Value: "test"}
	result := custom("")

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, custom("test"), result)
}

func Test_decoding_string_to_any_is_valid(t *testing.T) {
	object := tengo.String{Value: "test"}
	result := any("")

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, any("test"), result)
}

func Test_decoding_string_to_bool_returns_error(t *testing.T) {
	object := tengo.String{Value: ""}
	result := false

	err := DecodeObject(&object, &result)

	assert.Equal(t, &DecodingError{Object: &object, Expected: "bool"}, err)
}

func Test_decoding_string_to_slice_of_runes_returns_error(t *testing.T) {
	object := tengo.String{Value: "test"}
	result := []rune{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, &DecodingError{Object: &object, Expected: "array"}, err)
}

func Test_decoding_string_to_slice_of_bytes_returns_error(t *testing.T) {
	object := tengo.String{Value: "test"}
	result := []byte{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, &DecodingError{Object: &object, Expected: "bytes"}, err)
}

func Test_decoding_true_to_bool_is_valid(t *testing.T) {
	object := tengo.TrueValue
	result := false

	err := DecodeObject(object, &result)

	assert.NoError(t, err)
	assert.Equal(t, true, result)
}

func Test_decoding_false_to_bool_is_valid(t *testing.T) {
	object := tengo.FalseValue
	result := false

	err := DecodeObject(object, &result)

	assert.NoError(t, err)
	assert.Equal(t, false, result)
}

func Test_decoding_bool_to_type_compatible_with_bool_is_valid(t *testing.T) {
	type custom bool
	object := tengo.TrueValue
	result := custom(false)

	err := DecodeObject(object, &result)

	assert.NoError(t, err)
	assert.Equal(t, custom(true), result)
}

func Test_decoding_bytes_to_slice_of_bytes_is_valid(t *testing.T) {
	object := tengo.Bytes{Value: []byte("abc")}
	result := []byte{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, []byte("abc"), result)
}

func Test_decoding_bytes_to_any_is_valid(t *testing.T) {
	object := tengo.Bytes{Value: []byte("abc")}
	result := any([]byte{})

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, any([]byte("abc")), result)
}

func Test_decoding_bytes_to_type_compatible_with_slice_of_bytes_is_valid(t *testing.T) {
	type custom []byte
	object := tengo.Bytes{Value: []byte("abc")}
	result := custom{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, custom("abc"), result)
}

func Test_decoding_bytes_makes_copy_of_underlying_value(t *testing.T) {
	bytes := []byte("abc")
	object := tengo.Bytes{Value: bytes}
	result := []byte{}

	err := DecodeObject(&object, &result)
	bytes[0] = '_'

	assert.NoError(t, err)
	assert.Equal(t, []byte("abc"), result)
}

func Test_decoding_bytes_to_slice_of_bytes_overrides_its_existing_content(t *testing.T) {
	bytes := []byte("abc")
	object := tengo.Bytes{Value: bytes}
	result := []byte("test")

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, []byte("abc"), result)
}

func Test_decoding_bytes_to_string_returns_error(t *testing.T) {
	object := tengo.Bytes{Value: []byte("test")}
	result := ""

	err := DecodeObject(&object, &result)

	assert.Equal(t, &DecodingError{Object: &object, Expected: "string"}, err)
}

func Test_decoding_array_to_slice_is_valid(t *testing.T) {
	object := tengo.Array{Value: []tengo.Object{&tengo.Int{Value: 1}, &tengo.Int{Value: 2}, &tengo.Int{Value: 3}}}
	result := []int{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, result)
}

func Test_decoding_array_to_any_is_valid(t *testing.T) {
	object := tengo.ImmutableArray{Value: []tengo.Object{&tengo.Int{Value: 1}, &tengo.String{Value: "test"}, tengo.TrueValue}}
	result := any([]any{})

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, any([]any{1, "test", true}), result)
}

func Test_decoding_array_to_slice_of_bytes_is_valid(t *testing.T) {
	type custom bool
	object := tengo.Array{Value: []tengo.Object{}}
	result := []byte{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, []byte{}, result)
}

func Test_decoding_array_of_mixed_type_elements_to_slice_of_any_is_valid(t *testing.T) {
	object := tengo.ImmutableArray{Value: []tengo.Object{&tengo.Int{Value: 1}, &tengo.String{Value: "test"}, tengo.TrueValue}}
	result := []any{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, []any{1, "test", true}, result)
}

func Test_decoding_array_to_slice_fails_when_element_types_are_incompatible(t *testing.T) {
	invalidItem := tengo.Float{Value: 3}
	object := tengo.ImmutableArray{Value: []tengo.Object{&tengo.Int{Value: 1}, &tengo.Int{Value: 2}, &invalidItem}}
	result := []int{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, &DecodingError{Object: &invalidItem, Expected: "int", Path: []string{"2"}}, err)
}

func Test_decoding_array_to_slice_overrides_its_existing_content(t *testing.T) {
	object := tengo.Array{Value: []tengo.Object{&tengo.Int{Value: 0}}}
	result := []int{1, 2, 3}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, []int{0}, result)
}

func Test_decoding_immutable_array_to_slice_is_valid(t *testing.T) {
	object := tengo.ImmutableArray{Value: []tengo.Object{&tengo.Int{Value: 1}, &tengo.Int{Value: 2}, &tengo.Int{Value: 3}}}
	result := []int{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, result)
}

func Test_decoding_map_to_struct_is_valid(t *testing.T) {
	type S struct{ Foo bool }
	object := tengo.Map{Value: map[string]tengo.Object{"Foo": tengo.TrueValue}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, S{true}, result)
}

func Test_decoding_map_to_struct_overrides_only_fields_with_matching_map_entries(t *testing.T) {
	type S struct {
		Foo string
		Bar string
		Egg string
	}
	object := tengo.Map{Value: map[string]tengo.Object{"Foo": &tengo.String{Value: "overriden"}}}
	result := S{Foo: "original", Bar: "original"}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, S{"overriden", "original", ""}, result)
}

func Test_decoding_map_to_struct_ignores_unexported_fields(t *testing.T) {
	type S struct {
		foo bool
	}
	object := tengo.Map{Value: map[string]tengo.Object{"foo": tengo.TrueValue}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, err, &DecodingError{Object: object.Value["foo"], Path: []string{"foo"}, Expected: "undefined"})
}

func Test_decoding_map_to_struct_writes_to_promoted_fields_of_anonymous_struct_members(t *testing.T) {
	type S1 struct {
		Foo bool
		Bar bool
	}
	type S2 struct {
		S1
		Bar bool
	}
	object := tengo.Map{Value: map[string]tengo.Object{"Foo": tengo.TrueValue, "Bar": tengo.TrueValue}}
	result := S2{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, S2{S1{true, false}, true}, result)
}

func Test_decoding_map_to_struct_ignores_implicit_names_of_anonymous_struct_members(t *testing.T) {
	type S struct {
		bool
	}
	object := tengo.Map{Value: map[string]tengo.Object{"bool": tengo.TrueValue}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, err, &DecodingError{Object: tengo.TrueValue, Path: []string{"bool"}, Expected: "undefined"})
}

func Test_decoding_map_to_struct_ignores_names_set_with_tag_on_anonymous_struct_members(t *testing.T) {
	type S struct {
		bool `tengo:"name"`
	}
	object := tengo.Map{Value: map[string]tengo.Object{"name": tengo.TrueValue}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, err, &DecodingError{Object: tengo.TrueValue, Path: []string{"name"}, Expected: "undefined"})
}

func Test_decoding_map_to_struct_uses_names_from_tags(t *testing.T) {
	type S struct {
		Foo string `tengo:"bar"`
	}
	object := tengo.Map{Value: map[string]tengo.Object{"bar": &tengo.String{Value: "text"}}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, S{"text"}, result)
}

func Test_decoding_map_to_struct_ignores_fields_disabled_with_tags(t *testing.T) {
	type S struct {
		Foo string `tengo:"-"`
	}
	object := tengo.Map{Value: map[string]tengo.Object{"Foo": &tengo.String{Value: "text"}}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, err, &DecodingError{Object: object.Value["Foo"], Path: []string{"Foo"}, Expected: "undefined"})
}

func Test_decoding_map_to_struct_fails_when_element_types_are_incompatible(t *testing.T) {
	type S struct {
		Foo string
	}
	object := tengo.Map{Value: map[string]tengo.Object{"Foo": tengo.TrueValue}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, err, &DecodingError{Object: object.Value["Foo"], Path: []string{"Foo"}, Expected: "string"})
}

func Test_decoding_map_to_struct_fails_when_some_entries_are_not_mapped_to_fields(t *testing.T) {
	type S struct {
		Foo bool
	}
	object := tengo.Map{Value: map[string]tengo.Object{"Foo": tengo.TrueValue, "Bar": tengo.TrueValue}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, err, &DecodingError{Object: object.Value["Bar"], Path: []string{"Bar"}, Expected: "undefined"})
}

func Test_decoding_map_to_map_is_valid(t *testing.T) {
	object := tengo.Map{Value: map[string]tengo.Object{"foo": tengo.TrueValue, "bar": tengo.FalseValue}}
	result := map[string]bool{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, map[string]bool{"foo": true, "bar": false}, result)
}

func Test_decoding_map_to_map_fails_when_element_types_are_incompatible(t *testing.T) {
	object := tengo.Map{Value: map[string]tengo.Object{"foo": tengo.TrueValue, "invalid": &tengo.String{}, "bar": tengo.FalseValue}}
	result := map[string]bool{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, err, &DecodingError{Object: object.Value["invalid"], Path: []string{"invalid"}, Expected: "bool"})
}

func Test_decoding_map_to_map_fails_when_target_map_keys_are_not_strings(t *testing.T) {
	object := tengo.Map{Value: map[string]tengo.Object{"foo": tengo.TrueValue, "bar": tengo.FalseValue}}
	result := map[any]bool{}

	err := DecodeObject(&object, &result)

	assert.Error(t, err)
}

func Test_decoding_map_of_mixed_type_entries_to_map_of_any_is_valid(t *testing.T) {
	object := tengo.Map{Value: map[string]tengo.Object{"bool": tengo.TrueValue, "string": &tengo.String{}, "int": &tengo.Int{}, "map": &tengo.Map{}}}
	result := map[string]any{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, map[string]any{"bool": true, "string": "", "int": 0, "map": map[string]any{}}, result)
}

func Test_decoding_map_to_any_is_valid(t *testing.T) {
	object := tengo.Map{Value: map[string]tengo.Object{"bool": tengo.TrueValue, "string": &tengo.String{}, "int": &tengo.Int{}, "map": &tengo.Map{}}}
	result := any(0)

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, any(map[string]any{"bool": true, "string": "", "int": 0, "map": map[string]any{}}), result)
}

func Test_decoding_map_to_map_overrides_only_matching_entries(t *testing.T) {
	object := tengo.Map{Value: map[string]tengo.Object{"foo": tengo.TrueValue}}
	result := map[string]bool{"foo": false, "bar": false}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, map[string]bool{"foo": true, "bar": false}, result)
}

func Test_decoding_errors_contain_correct_path(t *testing.T) {
	type S struct {
		Field map[string][]struct{}
	}
	object := tengo.Map{Value: map[string]tengo.Object{
		"Field": &tengo.Map{Value: map[string]tengo.Object{
			"Key": &tengo.Array{Value: []tengo.Object{
				&tengo.Map{Value: map[string]tengo.Object{
					"Key": tengo.TrueValue}}}}}}}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.Equal(t, err, &DecodingError{Object: tengo.TrueValue, Path: []string{"Field", "Key", "0", "Key"}, Expected: "undefined"})
}

func Test_decoding_immutable_map_to_struct_is_valid(t *testing.T) {
	type S struct{ Foo bool }
	object := tengo.ImmutableMap{Value: map[string]tengo.Object{"Foo": tengo.TrueValue}}
	result := S{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, S{true}, result)
}

func Test_decoding_immutable_map_to_map_is_valid(t *testing.T) {
	object := tengo.ImmutableMap{Value: map[string]tengo.Object{"foo": tengo.TrueValue, "bar": tengo.FalseValue}}
	result := map[string]bool{}

	err := DecodeObject(&object, &result)

	assert.NoError(t, err)
	assert.Equal(t, map[string]bool{"foo": true, "bar": false}, result)
}
