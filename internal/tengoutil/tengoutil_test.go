package tengoutil

import (
	"errors"
	"fmt"
	"io/fs"
	"math"
	"reflect"
	"testing"

	"github.com/d5/tengo/v2"
)

type testEncoder struct {
	object tengo.Object
	err    error
}

func (t testEncoder) EncodeTengoObject() (tengo.Object, error) {
	return t.object, t.err
}

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
