package object

type Environment struct {
	Directory      string
	Kind           Kind
	Name           string
	DeployServices []string
	Variables      map[string]string
}
