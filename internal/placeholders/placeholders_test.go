package placeholders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_replacing_single_placeholder_in_a_string(t *testing.T) {
	values := map[string]interface{}{
		"placeholder": "egg spam",
	}
	input := "foo {{ .placeholder }} bar"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "foo egg spam bar", output)
}

func Test_replacing_multiple_placeholders_in_a_single_string_is_valid(t *testing.T) {
	values := map[string]interface{}{
		"placeholder1": "egg spam",
		"placeholder2": "potato",
	}
	input := "foo {{ .placeholder1 }} {{ .placeholder2 }} {{ .placeholder1 }} bar"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "foo egg spam potato egg spam bar", output)
}

func Test_replacing_a_string_without_any_placeholders_is_valid(t *testing.T) {
	values := map[string]interface{}{
		"placeholder": "tomato",
	}
	input := "foo bar"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "foo bar", output)
}

func Test_replacing_using_nested_values_is_valid(t *testing.T) {
	values := map[string]interface{}{
		"the": map[string]interface{}{
			"cake": map[string]interface{}{
				"is": map[string]interface{}{
					"a": "lie",
				},
			},
		},
	}
	input := "{{ .the.cake.is.a }}"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "lie", output)
}

func Test_replacing_using_string_to_string_map_is_valid(t *testing.T) {
	values := map[string]interface{}{
		"the": map[string]string{
			"cake": "is.a.lie",
		},
	}
	input := "{{ .the.cake }}"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "is.a.lie", output)
}

func Test_replacing_using_flattened_values_is_valid(t *testing.T) {
	values := map[string]interface{}{
		"the.cake.is": "a.lie",
	}
	input := "{{ .the.cake.is }}"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "a.lie", output)
}

func Test_replacing_using_a_mix_of_flattened_and_nested_values_is_valid(t *testing.T) {
	values := map[string]interface{}{
		"the.cake.is": map[string]interface{}{
			"a": "lie",
		},
	}
	input := "{{ .the.cake.is.a }}"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "lie", output)
}

func Test_replacement_values_are_case_insensitive(t *testing.T) {
	values := map[string]interface{}{
		"TheCakeIsA": "lie",
	}
	input := "{{ .tHeCAKEiSA }}"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "lie", output)
}

func Test_whitespace_around_name_in_a_placeholder_marker_is_ignored(t *testing.T) {
	values := map[string]interface{}{
		"foo": "bar",
	}
	input := "{{.foo}} {{ .foo }} {{  .foo     }} {{\t.foo \t }}"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "bar bar bar bar", output)
}

func Test_placeholders_in_replacement_values_are_replaced(t *testing.T) {
	values := map[string]interface{}{
		"placeholder1": "1",
		"placeholder2": "{{ .placeholder1 }} 2",
		"placeholder3": "{{ .placeholder2 }} 3",
	}
	input := "{{ .placeholder3 }}"

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, "1 2 3", output)
}

func Test_when_there_is_no_value_for_placeholder_replacing_ends_with_error(t *testing.T) {
	values := map[string]interface{}{
		"Foo": "1",
	}
	input := "{{ .Bar }}"

	_, err := Replace(input, values)

	assert.Equal(t, err, &MissingPlaceholderError{MissingName: ".Bar", ValidNames: []string{".Foo"}})
}

func Test_duplicating_placeholder_values_ends_with_error(t *testing.T) {
	values := map[string]interface{}{
		"Foo": map[string]interface{}{
			"Bar": "egg",
		},
		"foo.bar": "spam",
	}
	input := ""

	_, err := Replace(input, values)

	assert.Equal(t, err, &DuplicatedPlaceholderError{Name1: ".foo.bar", Name2: ".Foo.Bar"})
}

func Test_using_placeholder_with_invalid_name_ends_with_error(t *testing.T) {
	names := []string{
		"",
		".",
		"..",
		".a",
		"a.",
		"ą",
		"$",
		"-",
		"a-a",
		"a-9",
	}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			values := map[string]interface{}{name: "value"}
			input := ""

			_, err := Replace(input, values)

			assert.Equal(t, &InvalidPlaceholderNameError{Name: "." + name}, err)
		})
	}
}

func Test_replacing_cyclic_placeholders_ends_with_error(t *testing.T) {
	values := map[string]interface{}{
		"placeholder1": "{{ .placeholder2 }}",
		"placeholder2": "{{ .placeholder1 }}",
	}
	input := "{{ .placeholder1 }}"

	_, err := Replace(input, values)

	assert.Error(t, err)
	assert.Equal(t, &CyclicPlaceholderError{Cycle: []string{".placeholder2", ".placeholder1", ".placeholder2"}}, err)
}

func Test_cyclic_placeholder_errors_preserve_original_letter_casing(t *testing.T) {
	values := map[string]interface{}{
		"placeholder1": "{{ .placehOlder2 }}",
		"placeholder2": "{{ .Placeholder1 }}",
	}
	input := "{{ .placeholder1 }}"

	_, err := Replace(input, values)

	assert.Error(t, err)
	assert.Equal(t, &CyclicPlaceholderError{Cycle: []string{".placehOlder2", ".Placeholder1", ".placehOlder2"}}, err)
}

func Test_replacing_placeholders_in_map_values_is_supported(t *testing.T) {
	values := map[string]interface{}{
		"foo": "bar",
	}
	input := map[interface{}]interface{}{
		"": "{{ .foo }}",
	}

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, map[interface{}]interface{}{"": "bar"}, output)
}

func Test_replacing_placeholders_in_map_keys_is_supported(t *testing.T) {
	values := map[string]interface{}{
		"foo": "bar",
	}
	input := map[interface{}]interface{}{
		"{{ .foo }}": "",
	}

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, map[interface{}]interface{}{"bar": ""}, output)
}

func Test_replacing_placeholders_in_slices_is_supported(t *testing.T) {
	values := map[string]interface{}{
		"foo": "bar",
	}
	input := []interface{}{"{{ .foo }}"}

	output, err := Replace(input, values)

	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"bar"}, output)
}

func Test_replacing_placeholders_in_ints_floats_and_booleans_doesnt_do_anything(t *testing.T) {
	inputs := []interface{}{1.23, 123, true}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%T", input), func(t *testing.T) {
			values := map[string]interface{}{"foo": "bar"}

			output, err := Replace(input, values)

			assert.NoError(t, err)
			assert.Equal(t, input, output)
		})
	}
}

func Test_replacing_placeholders_in_not_supported_types_returns_error(t *testing.T) {
	inputs := []interface{}{
		struct{ Foo string }{"{{ .foo }}"},
		[1]string{"{{ .foo }}"},
		(*interface{})(nil),
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%T", input), func(t *testing.T) {
			values := map[string]interface{}{"foo": "bar"}

			_, err := Replace(input, values)

			assert.EqualError(t, err, fmt.Sprintf("replacing placeholders in %q is not supported", reflect.TypeOf(input).Kind()))
		})
	}
}
