package object

type ServiceEntry struct {
	Index int
	Type  string
	Spec  interface{}
}

type Service struct {
	Directory string
	Kind      Kind
	Name      string
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
