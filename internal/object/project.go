package object

type Project struct {
	Directory string `placeholders:"disable"`
	Kind      Kind   `placeholders:"disable"`
	Name      string `placeholders:"disable"`
	Files     []struct {
		Glob string
		Git  *struct {
			URL string
			Rev string
		}
	}
	Variables map[string]string
}
