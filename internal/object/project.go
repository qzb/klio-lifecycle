package object

type Project struct {
	Directory string
	Kind      Kind
	Name      string
	Files     []string
	Variables map[string]string
}
