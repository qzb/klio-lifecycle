package scheme

import (
	"testing"

	"github.com/go-test/deep"
	"gopkg.in/yaml.v3"
)

func TestToInternal(t *testing.T) {
	cases := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
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
					"glob",
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
			actual, _ := ToInternal(c.input)

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
