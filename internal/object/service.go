package object

type ServiceEntry struct {
	Index int
	Type  string `placeholders:"disable"`
	Spec  interface{}
}

type Service struct {
	Directory string `placeholders:"disable"`
	Kind      Kind   `placeholders:"disable"`
	Name      string `placeholders:"disable"`
	Build     struct {
		Tags      []ServiceEntry
		Artifacts struct {
			ToBuild []ServiceEntry
			ToPush  []ServiceEntry
		}
	}
	Deploy struct {
		Releases []ServiceEntry
	}
	Run struct {
		Tasks map[string][]ServiceEntry
	}
}
