package object

type Environment struct {
	Directory      string   `placeholders:"disable"`
	Kind           Kind     `placeholders:"disable"`
	Name           string   `placeholders:"disable"`
	DeployServices []string `placeholders:"disable"`
	Variables      map[string]string
}
