package scheme

import (
	"testing"

	"github.com/go-test/deep"
	"gopkg.in/yaml.v2"
)

func TestToInternal(t *testing.T) {
	cases := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "v1beta4/Service/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Service",
				"name":       "test",
				"hooks": map[string]interface{}{
					"pre-build": []interface{}{
						"echo pre-build",
						"cat file >&2",
					},
					"post-build": []interface{}{
						"echo post-build",
						"cat file >&2",
					},
					"pre-deploy": []interface{}{
						"echo pre-deploy",
						"cat file >&2",
					},
					"post-deploy": []interface{}{
						"echo post-deploy",
						"cat file >&2",
					},
				},
				"build": map[string]interface{}{
					"tagPolicy": map[string]interface{}{
						"gitSha": map[string]interface{}{"length": 7},
					},
					"artifacts": []interface{}{
						map[string]interface{}{
							"docker": map[string]interface{}{},
						},
						map[string]interface{}{
							"docker": map[string]interface{}{"image": "example.com/test/image"},
						},
					},
				},
				"deploy": map[string]interface{}{
					"releases": []interface{}{
						map[string]interface{}{
							"helm": map[string]interface{}{"chartPath": "bitnami/redis"},
						},
					},
				},
				"extra": true,
			},
			expected: map[string]interface{}{
				"kind": "Service",
				"name": "test",
				"build": map[string]interface{}{
					"tags": []interface{}{
						map[string]interface{}{
							"index": int64(0),
							"spec":  map[string]interface{}{"length": 7},
							"type":  "gitSha",
						},
					},
					"artifacts": map[string]interface{}{
						"toBuild": []interface{}{
							map[string]interface{}{
								"index": int64(0),
								"spec": map[string]interface{}{
									"sh": "set -e\necho pre-build\ncat file >&2",
								},
								"type": "script",
							},
							map[string]interface{}{
								"index": int64(1),
								"spec":  map[string]interface{}{},
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(2),
								"spec":  map[string]interface{}{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(3),
								"spec": map[string]interface{}{
									"sh": "set -e\necho post-build\ncat file >&2",
								},
								"type": "script",
							},
						},
						"toPush": []interface{}{
							map[string]interface{}{
								"index": int64(1),
								"spec":  map[string]interface{}{},
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(2),
								"spec":  map[string]interface{}{"image": "example.com/test/image"},
								"type":  "docker",
							},
						},
					},
				},
				"deploy": map[string]interface{}{
					"releases": []interface{}{
						map[string]interface{}{
							"index": int64(0),
							"spec": map[string]interface{}{
								"sh": "set -e\necho pre-deploy\ncat file >&2",
							},
							"type": "script",
						},
						map[string]interface{}{
							"index": int64(1),
							"spec":  map[string]interface{}{"chartPath": "bitnami/redis"},
							"type":  "helm",
						},
						map[string]interface{}{
							"index": int64(2),
							"spec": map[string]interface{}{
								"sh": "set -e\necho post-deploy\ncat file >&2",
							},
							"type": "script",
						},
					},
				},
				"run": map[string]interface{}{
					"tasks": map[string]interface{}{},
				},
			},
		},
		{
			name: "v1beta4/Service/min",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Service",
				"name":       "test",
			},
			expected: map[string]interface{}{
				"kind": "Service",
				"name": "test",
				"build": map[string]interface{}{
					"artifacts": map[string]interface{}{
						"toBuild": []interface{}{},
						"toPush":  []interface{}{},
					},
					"tags": []interface{}{},
				},
				"deploy": map[string]interface{}{
					"releases": []interface{}{},
				},
				"run": map[string]interface{}{
					"tasks": map[string]interface{}{},
				},
			},
		},
		{
			name: "v1beta4/Project/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Project",
				"services": []interface{}{
					".",
					"dir",
					"dirs/*/",
				},
				"environments": []interface{}{
					"envs/*",
				},
				"variables": map[string]interface{}{
					"test": "value",
				},
				"extra": true,
			},
			expected: map[string]interface{}{
				"kind": "Project",
				"name": "project",
				"files": []interface{}{
					map[string]interface{}{
						"git": map[string]interface{}{
							"url": "git@github.com:g2a-com/klio-lifecycle.git",
							"rev": "main",
						},
						"glob": "assets/executors/*/*.yaml",
					},
					map[string]interface{}{
						"glob": "service.yaml",
					},
					map[string]interface{}{
						"glob": "dir/service.yaml",
					},
					map[string]interface{}{
						"glob": "dirs/*/service.yaml",
					},
					map[string]interface{}{
						"glob": "envs/*/environment.yaml",
					},
				},
				"variables": map[string]interface{}{},
			},
		},
		{
			name: "v1beta4/Project/min_plus_defaults",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Project",
				"services": []interface{}{
					"services/*",
				},
				"environments": []interface{}{
					"environments/*",
				},
			},
			expected: map[string]interface{}{
				"kind": "Project",
				"name": "project",
				"files": []interface{}{
					map[string]interface{}{
						"git": map[string]interface{}{
							"url": "git@github.com:g2a-com/klio-lifecycle.git",
							"rev": "main",
						},
						"glob": "assets/executors/*/*.yaml",
					},
					map[string]interface{}{
						"glob": "services/*/service.yaml",
					},
					map[string]interface{}{
						"glob": "environments/*/environment.yaml",
					},
				},
				"variables": map[string]interface{}{},
			},
		},
		{
			name: "v1beta4/Project/min",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Project",
			},
			expected: map[string]interface{}{
				"kind": "Project",
				"name": "project",
				"files": []interface{}{
					map[string]interface{}{
						"git": map[string]interface{}{
							"url": "git@github.com:g2a-com/klio-lifecycle.git",
							"rev": "main",
						},
						"glob": "assets/executors/*/*.yaml",
					},
					map[string]interface{}{
						"glob": "services/*/service.yaml",
					},
					map[string]interface{}{
						"glob": "environments/*/environment.yaml",
					},
				},
				"variables": map[string]interface{}{},
			},
		},
		{
			name: "v1beta4/Environment/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Environment",
				"name":       "test",
				"deployServices": []interface{}{
					"serviceA",
					"serviceB",
				},
				"variables": map[string]interface{}{
					"varA": "value",
					"varB": "value",
				},
			},
			expected: map[string]interface{}{
				"kind": "Environment",
				"name": "test",
				"deployServices": []interface{}{
					"serviceA",
					"serviceB",
				},
				"variables": map[string]interface{}{
					"varA": "value",
					"varB": "value",
				},
			},
		},
		{
			name: "v1beta4/Environment/min",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v1beta4",
				"kind":       "Environment",
				"name":       "test",
			},
			expected: map[string]interface{}{
				"kind":           "Environment",
				"name":           "test",
				"deployServices": []interface{}{},
				"variables":      map[string]interface{}{},
			},
		},
		{
			name: "v2.0/Service/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Service",
				"name":       "test",
				"tags": []interface{}{
					"gitTag",
					map[string]interface{}{
						"custom": "latest",
					},
					map[string]interface{}{
						"gitSha": map[string]interface{}{"length": 7},
					},
				},
				"artifacts": []interface{}{
					"docker",
					map[string]interface{}{
						"docker": map[string]interface{}{},
					},
					map[string]interface{}{
						"docker": map[string]interface{}{"image": "example.com/test/image"},
					},
					map[string]interface{}{
						"docker": map[string]interface{}{"image": "example.com/test/image"},
						"push":   false,
					},
					map[string]interface{}{
						"docker": map[string]interface{}{"image": "example.com/test/image"},
						"push":   "script",
					},
					map[string]interface{}{
						"docker": map[string]interface{}{"image": "example.com/test/image"},
						"push":   map[string]interface{}{"script": "script.sh"},
					},
				},
				"releases": []interface{}{
					"npm",
					map[string]interface{}{
						"helm": "bitnami/redis",
					},
					map[string]interface{}{
						"helm": map[string]interface{}{"chartPath": "bitnami/redis"},
					},
				},
				"tasks": map[string]interface{}{
					"prepare": []interface{}{
						map[string]interface{}{"npm": "install"},
					},
					"test": []interface{}{
						map[string]interface{}{"npm": "test"},
						map[string]interface{}{"go": "test ./..."},
					},
					"lint": []interface{}{
						"prettier",
					},
					"task-name": []interface{}{
						map[string]interface{}{
							"script": map[string]interface{}{
								"sh": "./task.sh",
							},
						},
					},
				},
				"extra": true,
			},
			expected: map[string]interface{}{
				"kind": "Service",
				"name": "test",
				"build": map[string]interface{}{
					"tags": []interface{}{
						map[string]interface{}{
							"index": int64(0),
							"spec":  nil,
							"type":  "gitTag",
						},
						map[string]interface{}{
							"index": int64(1),
							"spec":  "latest",
							"type":  "custom",
						},
						map[string]interface{}{
							"index": int64(2),
							"spec":  map[string]interface{}{"length": 7},
							"type":  "gitSha",
						},
					},
					"artifacts": map[string]interface{}{
						"toBuild": []interface{}{
							map[string]interface{}{
								"index": int64(0),
								"spec":  nil,
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(1),
								"spec":  map[string]interface{}{},
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(2),
								"spec":  map[string]interface{}{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(3),
								"spec":  map[string]interface{}{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(4),
								"spec":  map[string]interface{}{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(5),
								"spec":  map[string]interface{}{"image": "example.com/test/image"},
								"type":  "docker",
							},
						},
						"toPush": []interface{}{
							map[string]interface{}{
								"index": int64(0),
								"spec":  nil,
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(1),
								"spec":  map[string]interface{}{},
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(2),
								"spec":  map[string]interface{}{"image": "example.com/test/image"},
								"type":  "docker",
							},
							map[string]interface{}{
								"index": int64(4),
								"spec":  nil,
								"type":  "script",
							},
							map[string]interface{}{
								"index": int64(5),
								"spec":  "script.sh",
								"type":  "script",
							},
						},
					},
				},
				"deploy": map[string]interface{}{
					"releases": []interface{}{
						map[string]interface{}{
							"index": int64(0),
							"spec":  nil,
							"type":  "npm",
						},
						map[string]interface{}{
							"index": int64(1),
							"spec":  "bitnami/redis",
							"type":  "helm",
						},
						map[string]interface{}{
							"index": int64(2),
							"spec":  map[string]interface{}{"chartPath": "bitnami/redis"},
							"type":  "helm",
						},
					},
				},
				"run": map[string]interface{}{
					"tasks": map[string]interface{}{
						"prepare": []interface{}{
							map[string]interface{}{
								"index": int64(0),
								"spec":  "install",
								"type":  "npm",
							},
						},
						"test": []interface{}{
							map[string]interface{}{
								"index": int64(0),
								"spec":  "test",
								"type":  "npm",
							},
							map[string]interface{}{
								"index": int64(1),
								"spec":  "test ./...",
								"type":  "go",
							},
						},
						"lint": []interface{}{
							map[string]interface{}{
								"index": int64(0),
								"spec":  nil,
								"type":  "prettier",
							},
						},
						"task-name": []interface{}{
							map[string]interface{}{
								"index": int64(0),
								"spec": map[string]interface{}{
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
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Service",
				"name":       "test",
			},
			expected: map[string]interface{}{
				"kind": "Service",
				"name": "test",
				"build": map[string]interface{}{
					"artifacts": map[string]interface{}{
						"toBuild": []interface{}{},
						"toPush":  []interface{}{},
					},
					"tags": []interface{}{},
				},
				"deploy": map[string]interface{}{
					"releases": []interface{}{},
				},
				"run": map[string]interface{}{
					"tasks": map[string]interface{}{},
				},
			},
		},
		{
			name: "v2.0/Project/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Project",
				"name":       "test",
				"files": []interface{}{
					"glob",
					map[string]interface{}{"git": map[string]interface{}{
						"url":   "http://github.com/g2a-com/klio-lifecycle",
						"rev":   "master",
						"files": "assets/executors/*/*",
					}},
				},
				"variables": map[string]interface{}{
					"name": "value",
				},
				"extra": true,
			},
			expected: map[string]interface{}{
				"kind": "Project",
				"name": "test",
				"files": []interface{}{
					map[string]interface{}{
						"glob": "glob",
					},
					map[string]interface{}{
						"git": map[string]interface{}{
							"url": "http://github.com/g2a-com/klio-lifecycle",
							"rev": "master",
						},
						"glob": "assets/executors/*/*",
					},
				},
				"variables": map[string]interface{}{
					"name": "value",
				},
			},
		},
		{
			name: "v2.0/Project/min",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Project",
				"name":       "test",
			},
			expected: map[string]interface{}{
				"kind":      "Project",
				"name":      "test",
				"files":     []interface{}{},
				"variables": map[string]interface{}{},
			},
		},
		{
			name: "v2.0/Environment/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Environment",
				"name":       "test",
				"deployServices": []interface{}{
					"serviceA",
					"serviceB",
				},
				"variables": map[string]interface{}{
					"varA": "value",
					"varB": "value",
				},
			},
			expected: map[string]interface{}{
				"kind": "Environment",
				"name": "test",
				"deployServices": []interface{}{
					"serviceA",
					"serviceB",
				},
				"variables": map[string]interface{}{
					"varA": "value",
					"varB": "value",
				},
			},
		},
		{
			name: "v2.0/Environment/min",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Environment",
				"name":       "test",
			},
			expected: map[string]interface{}{
				"kind":           "Environment",
				"name":           "test",
				"deployServices": []interface{}{},
				"variables":      map[string]interface{}{},
			},
		},
		{
			name: "v2.0/Builder/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Builder",
				"name":       "test",
				"schema":     map[string]interface{}{},
				"script":     "",
			},
			expected: map[string]interface{}{
				"kind":   "Builder",
				"name":   "test",
				"schema": "{}",
				"script": "",
			},
		},
		{
			name: "v2.0/Builder/min",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Builder",
				"name":       "test",
				"script":     "",
			},
			expected: map[string]interface{}{
				"kind":   "Builder",
				"name":   "test",
				"schema": "{}",
				"script": "",
			},
		},
		{
			name: "v2.0/Deployer/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Deployer",
				"name":       "test",
				"schema":     map[string]interface{}{},
				"script":     "",
			},
			expected: map[string]interface{}{
				"kind":   "Deployer",
				"name":   "test",
				"schema": "{}",
				"script": "",
			},
		},
		{
			name: "v2.0/Deployer/min",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Deployer",
				"name":       "test",
				"script":     "",
			},
			expected: map[string]interface{}{
				"kind":   "Deployer",
				"name":   "test",
				"schema": "{}",
				"script": "",
			},
		},
		{
			name: "v2.0/Pusher/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Pusher",
				"name":       "test",
				"schema":     map[string]interface{}{},
				"script":     "",
			},
			expected: map[string]interface{}{
				"kind":   "Pusher",
				"name":   "test",
				"schema": "{}",
				"script": "",
			},
		},
		{
			name: "v2.0/Pusher/min",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Pusher",
				"name":       "test",
				"script":     "",
			},
			expected: map[string]interface{}{
				"kind":   "Pusher",
				"name":   "test",
				"schema": "{}",
				"script": "",
			},
		},
		{
			name: "v2.0/Tagger/full",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Tagger",
				"name":       "test",
				"schema":     map[string]interface{}{},
				"script":     "",
			},
			expected: map[string]interface{}{
				"kind":   "Tagger",
				"name":   "test",
				"schema": "{}",
				"script": "",
			},
		},
		{
			name: "v2.0/Tagger/min",
			input: map[string]interface{}{
				"apiVersion": "g2a-cli/v2.0",
				"kind":       "Tagger",
				"name":       "test",
				"script":     "",
			},
			expected: map[string]interface{}{
				"kind":   "Tagger",
				"name":   "test",
				"schema": "{}",
				"script": "",
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
				e, _ := yaml.Marshal(map[string]interface{}{"EXPECTED": c.expected})
				a, _ := yaml.Marshal(map[string]interface{}{"ACTUAL": actual})
				t.Errorf("\n%s\n%s\n", e, a)
				for _, d := range diff {
					t.Error(d)
					return
				}
			}
		})
	}
}
