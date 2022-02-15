package scheme

import (
	"testing"

	"github.com/go-test/deep"
	"gopkg.in/yaml.v2"
)

func TestToInternal(t *testing.T) {
	cases := []struct {
		name     string
		input    map[string]any
		expected map[string]any
	}{
		{
			name: "v1beta4/Service/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Service",
				"name":       "test",
				"hooks": map[string]any{
					"pre-build": []any{
						"echo pre-build",
						"cat file >&2",
					},
					"post-build": []any{
						"echo post-build",
						"cat file >&2",
					},
					"pre-deploy": []any{
						"echo pre-deploy",
						"cat file >&2",
					},
					"post-deploy": []any{
						"echo post-deploy",
						"cat file >&2",
					},
				},
				"build": map[string]any{
					"tagPolicy": map[string]any{
						"gitSha": map[string]any{"length": 7},
					},
					"artifacts": []any{
						map[string]any{
							"docker": map[string]any{},
						},
						map[string]any{
							"docker": map[string]any{"image": "example.com/test/image"},
						},
					},
				},
				"deploy": map[string]any{
					"releases": []any{
						map[string]any{
							"helm": map[string]any{"chartPath": "bitnami/redis"},
						},
					},
				},
				"extra": true,
			},
			expected: map[string]any{
				"kind": "Service",
				"name": "test",
				"build": map[string]any{
					"tags": []any{
						map[string]any{
							"index": int64(0),
							"spec":  map[string]any{"length": 7},
							"type":  "gitSha",
						},
					},
					"artifacts": map[string]any{
						"toBuild": []any{
							map[string]any{
								"index": int64(0),
								"spec": map[string]any{
									"sh": "set -e\necho pre-build\ncat file >&2",
								},
								"type": "script",
							},
							map[string]any{
								"index": int64(1),
								"spec":  map[string]any{},
								"type":  "docker",
							},
							map[string]any{
								"index": int64(2),
								"spec":  map[string]any{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]any{
								"index": int64(3),
								"spec": map[string]any{
									"sh": "set -e\necho post-build\ncat file >&2",
								},
								"type": "script",
							},
						},
						"toPush": []any{
							map[string]any{
								"index": int64(1),
								"spec":  map[string]any{},
								"type":  "docker",
							},
							map[string]any{
								"index": int64(2),
								"spec":  map[string]any{"image": "example.com/test/image"},
								"type":  "docker",
							},
						},
					},
				},
				"deploy": map[string]any{
					"releases": []any{
						map[string]any{
							"index": int64(0),
							"spec": map[string]any{
								"sh": "set -e\necho pre-deploy\ncat file >&2",
							},
							"type": "script",
						},
						map[string]any{
							"index": int64(1),
							"spec":  map[string]any{"chartPath": "bitnami/redis"},
							"type":  "helm",
						},
						map[string]any{
							"index": int64(2),
							"spec": map[string]any{
								"sh": "set -e\necho post-deploy\ncat file >&2",
							},
							"type": "script",
						},
					},
				},
				"run": map[string]any{
					"tasks": map[string]any{},
				},
			},
		},
		{
			name: "v1beta4/Service/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Service",
				"name":       "test",
			},
			expected: map[string]any{
				"kind": "Service",
				"name": "test",
				"build": map[string]any{
					"artifacts": map[string]any{
						"toBuild": []any{},
						"toPush":  []any{},
					},
					"tags": []any{},
				},
				"deploy": map[string]any{
					"releases": []any{},
				},
				"run": map[string]any{
					"tasks": map[string]any{},
				},
			},
		},
		{
			name: "v1beta4/Project/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Project",
				"services": []any{
					".",
					"dir",
					"dirs/*",
				},
				"environments": []any{
					"envs/*",
				},
				"variables": map[string]any{
					"test": "value",
				},
				"extra": true,
			},
			expected: map[string]any{
				"kind": "Project",
				"name": "project",
				"files": []any{
					map[string]any{
						"git": map[string]any{
							"url": "git@github.com:g2a-com/klio-lifecycle.git",
							"rev": "main",
						},
						"glob": "assets/executors/*/*.yaml",
					},
					map[string]any{
						"glob": "./service.yaml",
					},
					map[string]any{
						"glob": "dir/service.yaml",
					},
					map[string]any{
						"glob": "dirs/*/service.yaml",
					},
					map[string]any{
						"glob": "envs/*/environment.yaml",
					},
				},
				"variables": map[string]any{},
			},
		},
		{
			name: "v1beta4/Project/min_plus_defaults",
			input: map[string]any{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Project",
				"services": []any{
					"services/*",
				},
				"environments": []any{
					"environments/*",
				},
			},
			expected: map[string]any{
				"kind": "Project",
				"name": "project",
				"files": []any{
					map[string]any{
						"git": map[string]any{
							"url": "git@github.com:g2a-com/klio-lifecycle.git",
							"rev": "main",
						},
						"glob": "assets/executors/*/*.yaml",
					},
					map[string]any{
						"glob": "services/*/service.yaml",
					},
					map[string]any{
						"glob": "environments/*/environment.yaml",
					},
				},
				"variables": map[string]any{},
			},
		},
		{
			name: "v1beta4/Project/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Project",
			},
			expected: map[string]any{
				"kind": "Project",
				"name": "project",
				"files": []any{
					map[string]any{
						"git": map[string]any{
							"url": "git@github.com:g2a-com/klio-lifecycle.git",
							"rev": "main",
						},
						"glob": "assets/executors/*/*.yaml",
					},
					map[string]any{
						"glob": "services/*/service.yaml",
					},
					map[string]any{
						"glob": "environments/*/environment.yaml",
					},
				},
				"variables": map[string]any{},
			},
		},
		{
			name: "v1beta4/Environment/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Environment",
				"name":       "test",
				"deployServices": []any{
					"serviceA",
					"serviceB",
				},
				"variables": map[string]any{
					"varA": "value",
					"varB": "value",
				},
			},
			expected: map[string]any{
				"kind": "Environment",
				"name": "test",
				"deployServices": []any{
					"serviceA",
					"serviceB",
				},
				"variables": map[string]any{
					"varA": "value",
					"varB": "value",
				},
			},
		},
		{
			name: "v1beta4/Environment/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Environment",
				"name":       "test",
			},
			expected: map[string]any{
				"kind":           "Environment",
				"name":           "test",
				"deployServices": []any{},
				"variables":      map[string]any{},
			},
		},
		{
			name: "v2.0/Service/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Service",
				"name":       "test",
				"tags": []any{
					"gitTag",
					map[string]any{
						"custom": "latest",
					},
					map[string]any{
						"gitSha": map[string]any{"length": 7},
					},
				},
				"artifacts": []any{
					"docker",
					map[string]any{
						"docker": map[string]any{},
					},
					map[string]any{
						"docker": map[string]any{"image": "example.com/test/image"},
					},
					map[string]any{
						"docker": map[string]any{"image": "example.com/test/image"},
						"push":   false,
					},
					map[string]any{
						"docker": map[string]any{"image": "example.com/test/image"},
						"push":   "script",
					},
					map[string]any{
						"docker": map[string]any{"image": "example.com/test/image"},
						"push":   map[string]any{"script": "script.sh"},
					},
				},
				"releases": []any{
					"npm",
					map[string]any{
						"helm": "bitnami/redis",
					},
					map[string]any{
						"helm": map[string]any{"chartPath": "bitnami/redis"},
					},
				},
				"tasks": map[string]any{
					"prepare": []any{
						map[string]any{"npm": "install"},
					},
					"test": []any{
						map[string]any{"npm": "test"},
						map[string]any{"go": "test ./..."},
					},
					"lint": []any{
						"prettier",
					},
					"task-name": []any{
						map[string]any{
							"script": map[string]any{
								"sh": "./task.sh",
							},
						},
					},
				},
				"extra": true,
			},
			expected: map[string]any{
				"kind": "Service",
				"name": "test",
				"build": map[string]any{
					"tags": []any{
						map[string]any{
							"index": int64(0),
							"spec":  nil,
							"type":  "gitTag",
						},
						map[string]any{
							"index": int64(1),
							"spec":  "latest",
							"type":  "custom",
						},
						map[string]any{
							"index": int64(2),
							"spec":  map[string]any{"length": 7},
							"type":  "gitSha",
						},
					},
					"artifacts": map[string]any{
						"toBuild": []any{
							map[string]any{
								"index": int64(0),
								"spec":  nil,
								"type":  "docker",
							},
							map[string]any{
								"index": int64(1),
								"spec":  map[string]any{},
								"type":  "docker",
							},
							map[string]any{
								"index": int64(2),
								"spec":  map[string]any{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]any{
								"index": int64(3),
								"spec":  map[string]any{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]any{
								"index": int64(4),
								"spec":  map[string]any{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]any{
								"index": int64(5),
								"spec":  map[string]any{"image": "example.com/test/image"},
								"type":  "docker",
							},
						},
						"toPush": []any{
							map[string]any{
								"index": int64(0),
								"spec":  nil,
								"type":  "docker",
							},
							map[string]any{
								"index": int64(1),
								"spec":  map[string]any{},
								"type":  "docker",
							},
							map[string]any{
								"index": int64(2),
								"spec":  map[string]any{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]any{
								"index": int64(4),
								"spec":  nil,
								"type":  "script",
							},
							map[string]any{
								"index": int64(5),
								"spec":  "script.sh",
								"type":  "script",
							},
						},
					},
				},
				"deploy": map[string]any{
					"releases": []any{
						map[string]any{
							"index": int64(0),
							"spec":  nil,
							"type":  "npm",
						},
						map[string]any{
							"index": int64(1),
							"spec":  "bitnami/redis",
							"type":  "helm",
						},
						map[string]any{
							"index": int64(2),
							"spec":  map[string]any{"chartPath": "bitnami/redis"},
							"type":  "helm",
						},
					},
				},
				"run": map[string]any{
					"tasks": map[string]any{
						"prepare": []any{
							map[string]any{
								"index": int64(0),
								"spec":  "install",
								"type":  "npm",
							},
						},
						"test": []any{
							map[string]any{
								"index": int64(0),
								"spec":  "test",
								"type":  "npm",
							},
							map[string]any{
								"index": int64(1),
								"spec":  "test ./...",
								"type":  "go",
							},
						},
						"lint": []any{
							map[string]any{
								"index": int64(0),
								"spec":  nil,
								"type":  "prettier",
							},
						},
						"task-name": []any{
							map[string]any{
								"index": int64(0),
								"spec": map[string]any{
									"sh": "./task.sh",
								},
								"type": "script",
							},
						},
					},
				},
			},
		},
		{
			name: "v2.0/Service/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Service",
				"name":       "test",
			},
			expected: map[string]any{
				"kind": "Service",
				"name": "test",
				"build": map[string]any{
					"artifacts": map[string]any{
						"toBuild": []any{},
						"toPush":  []any{},
					},
					"tags": []any{},
				},
				"deploy": map[string]any{
					"releases": []any{},
				},
				"run": map[string]any{
					"tasks": map[string]any{},
				},
			},
		},
		{
			name: "v2.0/Project/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Project",
				"name":       "test",
				"files": []any{
					"glob",
					map[string]any{"git": map[string]any{
						"url":   "http://github.com/g2a-com/klio-lifecycle",
						"rev":   "master",
						"files": "assets/executors/*/*",
					}},
				},
				"variables": map[string]any{
					"name": "value",
				},
				"extra": true,
			},
			expected: map[string]any{
				"kind": "Project",
				"name": "test",
				"files": []any{
					map[string]any{
						"glob": "glob",
					},
					map[string]any{
						"git": map[string]any{
							"url": "http://github.com/g2a-com/klio-lifecycle",
							"rev": "master",
						},
						"glob": "assets/executors/*/*",
					},
				},
				"variables": map[string]any{
					"name": "value",
				},
			},
		},
		{
			name: "v2.0/Project/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Project",
				"name":       "test",
			},
			expected: map[string]any{
				"kind":      "Project",
				"name":      "test",
				"files":     []any{},
				"variables": map[string]any{},
			},
		},
		{
			name: "v2.0/Environment/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Environment",
				"name":       "test",
				"deployServices": []any{
					"serviceA",
					"serviceB",
				},
				"variables": map[string]any{
					"varA": "value",
					"varB": "value",
				},
			},
			expected: map[string]any{
				"kind": "Environment",
				"name": "test",
				"deployServices": []any{
					"serviceA",
					"serviceB",
				},
				"variables": map[string]any{
					"varA": "value",
					"varB": "value",
				},
			},
		},
		{
			name: "v2.0/Environment/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Environment",
				"name":       "test",
			},
			expected: map[string]any{
				"kind":           "Environment",
				"name":           "test",
				"deployServices": []any{},
				"variables":      map[string]any{},
			},
		},
		{
			name: "v2.0/Builder/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Builder",
				"name":       "test",
				"schema":     map[string]any{},
				"js":         "",
			},
			expected: map[string]any{
				"kind":   "Builder",
				"name":   "test",
				"schema": "{}",
				"js":     "",
			},
		},
		{
			name: "v2.0/Builder/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Builder",
				"name":       "test",
				"js":         "",
			},
			expected: map[string]any{
				"kind":   "Builder",
				"name":   "test",
				"schema": "{}",
				"js":     "",
			},
		},
		{
			name: "v2.0/Deployer/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Deployer",
				"name":       "test",
				"schema":     map[string]any{},
				"js":         "",
			},
			expected: map[string]any{
				"kind":   "Deployer",
				"name":   "test",
				"schema": "{}",
				"js":     "",
			},
		},
		{
			name: "v2.0/Deployer/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Deployer",
				"name":       "test",
				"js":         "",
			},
			expected: map[string]any{
				"kind":   "Deployer",
				"name":   "test",
				"schema": "{}",
				"js":     "",
			},
		},
		{
			name: "v2.0/Pusher/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Pusher",
				"name":       "test",
				"schema":     map[string]any{},
				"js":         "",
			},
			expected: map[string]any{
				"kind":   "Pusher",
				"name":   "test",
				"schema": "{}",
				"js":     "",
			},
		},
		{
			name: "v2.0/Pusher/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Pusher",
				"name":       "test",
				"js":         "",
			},
			expected: map[string]any{
				"kind":   "Pusher",
				"name":   "test",
				"schema": "{}",
				"js":     "",
			},
		},
		{
			name: "v2.0/Tagger/full",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Tagger",
				"name":       "test",
				"schema":     map[string]any{},
				"js":         "",
			},
			expected: map[string]any{
				"kind":   "Tagger",
				"name":   "test",
				"schema": "{}",
				"js":     "",
			},
		},
		{
			name: "v2.0/Tagger/min",
			input: map[string]any{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Tagger",
				"name":       "test",
				"js":         "",
			},
			expected: map[string]any{
				"kind":   "Tagger",
				"name":   "test",
				"schema": "{}",
				"js":     "",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Validate(c.input["apiVersion"].(string), c.input["kind"].(string), c.input)
			if err != nil {
				t.Error(err)
				return
			}

			actual, _ := ToInternal(c.input)

			err = Validate("internal", "Object", c.expected)
			if err != nil {
				t.Errorf("expected value doesn't match schema:\n%s", err)
				return
			}

			if diff := deep.Equal(c.expected, actual); diff != nil {
				e, _ := yaml.Marshal(map[string]any{"EXPECTED": c.expected})
				a, _ := yaml.Marshal(map[string]any{"ACTUAL": actual})
				t.Errorf("\n%s\n%s\n", e, a)
				for _, d := range diff {
					t.Error(d)
					return
				}
			}
		})
	}
}
