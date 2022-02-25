// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tengoutil

import (
	"reflect"
	"strings"
)

// tagOptions is the string following a comma in a struct field's "tengo"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag extracts "tengo" tag from StructField and splits it into its name and
// comma-separated options. If name is not provided, it uses field's name.
func parseTag(field reflect.StructField) (string, tagOptions) {
	tag := field.Tag.Get("tengo")
	if tag == "-" {
		return "", tagOptions("")
	}
	name, opt, _ := strings.Cut(tag, ",")
	if name == "" {
		name = field.Name
	}
	return name, tagOptions(opt)
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if name == optionName {
			return true
		}
	}
	return false
}
