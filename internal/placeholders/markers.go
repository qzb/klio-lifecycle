package placeholders

import "regexp"

var placeholderRegexp *regexp.Regexp = regexp.MustCompile(`{{\s*([A-Za-z0-9_.]+)\s*}}`)

// replaceMarkers replaces all occurrences of placeholder markers (e.g.: {{
// .someting }}) using provided function.
func replaceMarkers(str string, fn func(string) (string, error)) (res string, err error) {
	res = placeholderRegexp.ReplaceAllStringFunc(str, func(s string) (r string) {
		if err != nil {
			return s
		}
		name := placeholderRegexp.FindStringSubmatch(s)[1]
		r, err = fn(name)
		return r
	})
	return
}

// containsMarkers checks if provided string contains any placeholder markers.
func containsMarkers(str string) bool {
	return placeholderRegexp.MatchString(str)
}
