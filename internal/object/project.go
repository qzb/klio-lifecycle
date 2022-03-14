package object

type Project struct {
	Directory string
	Kind      Kind
	Name      string
	Files     []struct {
		Glob string
		Git  *struct {
			URL string
			Rev string
		}
	}
	Variables map[string]string
}
