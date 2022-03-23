//go:generate ../../scripts/generate-schema-module.sh
package schema

// THIS FILE WAS GENERATED USING SCRIPT, DO NOT CHANGE IT BY HAND!
// RUN "go generate ./internal/schema" TO UPDATE.

var SCHEMAS = map[string][]byte{
	"g2a-cli/v1beta4/Environment": []byte(`
		{
		  "title": "Environment",
		  "type": "object",
		  "additionalProperties": false,
		  "required": [
		    "apiVersion",
		    "kind",
		    "name"
		  ],
		  "properties": {
		    "apiVersion": {
		      "const": "g2a-cli/v1beta4"
		    },
		    "kind": {
		      "const": "Environment"
		    },
		    "name": {
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "deployServices": {
		      "type": "array",
		      "items": {
		        "type": "string"
		      }
		    },
		    "variables": {
		      "type": "object",
		      "patternProperties": {
		        "^[a-zA-Z][a-zA-Z0-9]*$": {
		          "type": "string"
		        }
		      }
		    }
		  }
		}
	`),
	"g2a-cli/v1beta4/Object": []byte(`
		{
		  "title": "Object",
		  "oneOf": [
		    {
		      "title": "Project",
		      "type": "object",
		      "required": [
		        "apiVersion",
		        "kind"
		      ],
		      "properties": {
		        "apiVersion": {
		          "const": "g2a-cli/v1beta4"
		        },
		        "kind": {
		          "const": "Project"
		        },
		        "name": {
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$"
		        },
		        "services": {
		          "type": "array",
		          "items": {
		            "type": "string"
		          }
		        },
		        "environments": {
		          "type": "array",
		          "items": {
		            "type": "string"
		          }
		        }
		      },
		      "$defs": {
		        "globs": {
		          "type": "array",
		          "items": {
		            "type": "string"
		          }
		        }
		      }
		    },
		    {
		      "title": "Service",
		      "type": "object",
		      "required": [
		        "apiVersion",
		        "kind",
		        "name"
		      ],
		      "properties": {
		        "apiVersion": {
		          "const": "g2a-cli/v1beta4"
		        },
		        "kind": {
		          "const": "Service"
		        },
		        "name": {
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$"
		        },
		        "hooks": {
		          "type": "object",
		          "additionalProperties": true,
		          "properties": {
		            "pre-build": {
		              "type": "array",
		              "items": {
		                "type": "string",
		                "minLength": 1
		              }
		            },
		            "pre-deploy": {
		              "type": "array",
		              "items": {
		                "type": "string",
		                "minLength": 1
		              }
		            },
		            "post-build": {
		              "type": "array",
		              "items": {
		                "type": "string",
		                "minLength": 1
		              }
		            },
		            "post-deploy": {
		              "type": "array",
		              "items": {
		                "type": "string",
		                "minLength": 1
		              }
		            }
		          }
		        },
		        "build": {
		          "type": "object",
		          "additionalProperties": false,
		          "required": [
		            "artifacts",
		            "tagPolicy"
		          ],
		          "properties": {
		            "artifacts": {
		              "type": "array",
		              "items": {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": {
		                  "type": "object"
		                }
		              }
		            },
		            "tagPolicy": {
		              "type": "object",
		              "minProperties": 1,
		              "additionalProperties": {
		                "type": "object"
		              }
		            }
		          }
		        },
		        "deploy": {
		          "type": "object",
		          "additionalProperties": false,
		          "required": [
		            "releases"
		          ],
		          "properties": {
		            "releases": {
		              "type": "array",
		              "items": {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": {
		                  "type": "object"
		                }
		              }
		            }
		          }
		        }
		      },
		      "$defs": {
		        "entries": {
		          "type": "array",
		          "items": {
		            "type": "object",
		            "minProperties": 1,
		            "maxProperties": 1,
		            "additionalProperties": {
		              "type": "object"
		            }
		          }
		        },
		        "hooks": {
		          "type": "array",
		          "items": {
		            "type": "string",
		            "minLength": 1
		          }
		        }
		      }
		    },
		    {
		      "title": "Environment",
		      "type": "object",
		      "additionalProperties": false,
		      "required": [
		        "apiVersion",
		        "kind",
		        "name"
		      ],
		      "properties": {
		        "apiVersion": {
		          "const": "g2a-cli/v1beta4"
		        },
		        "kind": {
		          "const": "Environment"
		        },
		        "name": {
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$"
		        },
		        "deployServices": {
		          "type": "array",
		          "items": {
		            "type": "string"
		          }
		        },
		        "variables": {
		          "type": "object",
		          "patternProperties": {
		            "^[a-zA-Z][a-zA-Z0-9]*$": {
		              "type": "string"
		            }
		          }
		        }
		      }
		    }
		  ]
		}
	`),
	"g2a-cli/v1beta4/Project": []byte(`
		{
		  "title": "Project",
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind"
		  ],
		  "properties": {
		    "apiVersion": {
		      "const": "g2a-cli/v1beta4"
		    },
		    "kind": {
		      "const": "Project"
		    },
		    "name": {
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "services": {
		      "type": "array",
		      "items": {
		        "type": "string"
		      }
		    },
		    "environments": {
		      "type": "array",
		      "items": {
		        "type": "string"
		      }
		    }
		  }
		}
	`),
	"g2a-cli/v1beta4/Service": []byte(`
		{
		  "title": "Service",
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind",
		    "name"
		  ],
		  "properties": {
		    "apiVersion": {
		      "const": "g2a-cli/v1beta4"
		    },
		    "kind": {
		      "const": "Service"
		    },
		    "name": {
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "hooks": {
		      "type": "object",
		      "additionalProperties": true,
		      "properties": {
		        "pre-build": {
		          "type": "array",
		          "items": {
		            "type": "string",
		            "minLength": 1
		          }
		        },
		        "pre-deploy": {
		          "type": "array",
		          "items": {
		            "type": "string",
		            "minLength": 1
		          }
		        },
		        "post-build": {
		          "type": "array",
		          "items": {
		            "type": "string",
		            "minLength": 1
		          }
		        },
		        "post-deploy": {
		          "type": "array",
		          "items": {
		            "type": "string",
		            "minLength": 1
		          }
		        }
		      }
		    },
		    "build": {
		      "type": "object",
		      "additionalProperties": false,
		      "required": [
		        "artifacts",
		        "tagPolicy"
		      ],
		      "properties": {
		        "artifacts": {
		          "type": "array",
		          "items": {
		            "type": "object",
		            "minProperties": 1,
		            "maxProperties": 1,
		            "additionalProperties": {
		              "type": "object"
		            }
		          }
		        },
		        "tagPolicy": {
		          "type": "object",
		          "minProperties": 1,
		          "additionalProperties": {
		            "type": "object"
		          }
		        }
		      }
		    },
		    "deploy": {
		      "type": "object",
		      "additionalProperties": false,
		      "required": [
		        "releases"
		      ],
		      "properties": {
		        "releases": {
		          "type": "array",
		          "items": {
		            "type": "object",
		            "minProperties": 1,
		            "maxProperties": 1,
		            "additionalProperties": {
		              "type": "object"
		            }
		          }
		        }
		      }
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Build-result": []byte(`
		{
		  "type": "object",
		  "additionalProperties": true,
		  "required": [
		    "artifacts",
		    "pushed"
		  ],
		  "properties": {
		    "tags": {
		      "type": "array",
		      "items": {
		        "examples": [
		          {
		            "service": "generic-service",
		            "entry": 0,
		            "result": "some result"
		          }
		        ],
		        "type": "object",
		        "additionalProperties": true,
		        "properties": {
		          "service": {
		            "description": "Name of the object, unique within the kind.",
		            "type": "string",
		            "minLength": 1,
		            "pattern": "^[a-z][A-Za-z0-9_-]*$"
		          },
		          "entry": {
		            "type": "integer",
		            "min": 0
		          },
		          "result": {
		            "type": "string"
		          }
		        }
		      }
		    },
		    "artifacts": {
		      "type": "array",
		      "items": {
		        "examples": [
		          {
		            "service": "generic-service",
		            "entry": 0,
		            "result": "some result"
		          }
		        ],
		        "type": "object",
		        "additionalProperties": true,
		        "properties": {
		          "service": {
		            "description": "Name of the object, unique within the kind.",
		            "type": "string",
		            "minLength": 1,
		            "pattern": "^[a-z][A-Za-z0-9_-]*$"
		          },
		          "entry": {
		            "type": "integer",
		            "min": 0
		          },
		          "result": {
		            "type": "string"
		          }
		        }
		      }
		    },
		    "pushedArtifacts": {
		      "type": "array",
		      "items": {
		        "examples": [
		          {
		            "service": "generic-service",
		            "entry": 0,
		            "result": "some result"
		          }
		        ],
		        "type": "object",
		        "additionalProperties": true,
		        "properties": {
		          "service": {
		            "description": "Name of the object, unique within the kind.",
		            "type": "string",
		            "minLength": 1,
		            "pattern": "^[a-z][A-Za-z0-9_-]*$"
		          },
		          "entry": {
		            "type": "integer",
		            "min": 0
		          },
		          "result": {
		            "type": "string"
		          }
		        }
		      }
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Builder": []byte(`
		{
		  "title": "Builder",
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind",
		    "name",
		    "script"
		  ],
		  "additionalProperties": false,
		  "properties": {
		    "apiVersion": {
		      "description": "Version of the configuration format.",
		      "const": "g2a-cli/v2.0"
		    },
		    "kind": {
		      "const": "Builder"
		    },
		    "name": {
		      "description": "Name of the object, unique within the kind.",
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "schema": {
		      "$schema": "https://json-schema.org/draft/2019-09/schema",
		      "$id": "https://json-schema.org/draft/2019-09/schema",
		      "$vocabulary": {
		        "https://json-schema.org/draft/2019-09/vocab/core": true,
		        "https://json-schema.org/draft/2019-09/vocab/applicator": true,
		        "https://json-schema.org/draft/2019-09/vocab/validation": true,
		        "https://json-schema.org/draft/2019-09/vocab/meta-data": true,
		        "https://json-schema.org/draft/2019-09/vocab/format": false,
		        "https://json-schema.org/draft/2019-09/vocab/content": true
		      },
		      "$recursiveAnchor": true,
		      "title": "Core and Validation specifications meta-schema",
		      "allOf": [
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/core",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/core": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Core vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "$id": {
		              "type": "string",
		              "format": "uri-reference",
		              "$comment": "Non-empty fragments not allowed.",
		              "pattern": "^[^#]*#?$"
		            },
		            "$schema": {
		              "type": "string",
		              "format": "uri"
		            },
		            "$anchor": {
		              "type": "string",
		              "pattern": "^[A-Za-z][-A-Za-z0-9.:_]*$"
		            },
		            "$ref": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveRef": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveAnchor": {
		              "type": "boolean",
		              "default": false
		            },
		            "$vocabulary": {
		              "type": "object",
		              "propertyNames": {
		                "type": "string",
		                "format": "uri"
		              },
		              "additionalProperties": {
		                "type": "boolean"
		              }
		            },
		            "$comment": {
		              "type": "string"
		            },
		            "$defs": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/applicator",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/applicator": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Applicator vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "additionalItems": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedItems": {
		              "$recursiveRef": "#"
		            },
		            "items": {
		              "anyOf": [
		                {
		                  "$recursiveRef": "#"
		                },
		                {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                }
		              ]
		            },
		            "contains": {
		              "$recursiveRef": "#"
		            },
		            "additionalProperties": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedProperties": {
		              "$recursiveRef": "#"
		            },
		            "properties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            },
		            "patternProperties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "propertyNames": {
		                "format": "regex"
		              },
		              "default": {}
		            },
		            "dependentSchemas": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              }
		            },
		            "propertyNames": {
		              "$recursiveRef": "#"
		            },
		            "if": {
		              "$recursiveRef": "#"
		            },
		            "then": {
		              "$recursiveRef": "#"
		            },
		            "else": {
		              "$recursiveRef": "#"
		            },
		            "allOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "anyOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "oneOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "not": {
		              "$recursiveRef": "#"
		            }
		          },
		          "$defs": {
		            "schemaArray": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/validation",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/validation": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Validation vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "multipleOf": {
		              "type": "number",
		              "exclusiveMinimum": 0
		            },
		            "maximum": {
		              "type": "number"
		            },
		            "exclusiveMaximum": {
		              "type": "number"
		            },
		            "minimum": {
		              "type": "number"
		            },
		            "exclusiveMinimum": {
		              "type": "number"
		            },
		            "maxLength": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minLength": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "pattern": {
		              "type": "string",
		              "format": "regex"
		            },
		            "maxItems": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minItems": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "uniqueItems": {
		              "type": "boolean",
		              "default": false
		            },
		            "maxContains": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minContains": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 1
		            },
		            "maxProperties": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minProperties": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "required": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            },
		            "dependentRequired": {
		              "type": "object",
		              "additionalProperties": {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            },
		            "const": true,
		            "enum": {
		              "type": "array",
		              "items": true
		            },
		            "type": {
		              "anyOf": [
		                {
		                  "enum": [
		                    "array",
		                    "boolean",
		                    "integer",
		                    "null",
		                    "number",
		                    "object",
		                    "string"
		                  ]
		                },
		                {
		                  "type": "array",
		                  "items": {
		                    "enum": [
		                      "array",
		                      "boolean",
		                      "integer",
		                      "null",
		                      "number",
		                      "object",
		                      "string"
		                    ]
		                  },
		                  "minItems": 1,
		                  "uniqueItems": true
		                }
		              ]
		            }
		          },
		          "$defs": {
		            "nonNegativeInteger": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "nonNegativeIntegerDefault0": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 0
		            },
		            "simpleTypes": {
		              "enum": [
		                "array",
		                "boolean",
		                "integer",
		                "null",
		                "number",
		                "object",
		                "string"
		              ]
		            },
		            "stringArray": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/meta-data",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/meta-data": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Meta-data vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "title": {
		              "type": "string"
		            },
		            "description": {
		              "type": "string"
		            },
		            "default": true,
		            "deprecated": {
		              "type": "boolean",
		              "default": false
		            },
		            "readOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "writeOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "examples": {
		              "type": "array",
		              "items": true
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/format",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/format": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Format vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "format": {
		              "type": "string"
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/content",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/content": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Content vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "contentMediaType": {
		              "type": "string"
		            },
		            "contentEncoding": {
		              "type": "string"
		            },
		            "contentSchema": {
		              "$recursiveRef": "#"
		            }
		          }
		        }
		      ],
		      "type": [
		        "object",
		        "boolean"
		      ],
		      "properties": {
		        "definitions": {
		          "$comment": "While no longer an official keyword as it is replaced by $defs, this keyword is retained in the meta-schema to prevent incompatible extensions as it remains in common use.",
		          "type": "object",
		          "additionalProperties": {
		            "$recursiveRef": "#"
		          },
		          "default": {}
		        },
		        "dependencies": {
		          "$comment": "\"dependencies\" is no longer a keyword, but schema authors should avoid redefining it to facilitate a smooth transition to \"dependentSchemas\" and \"dependentRequired\"",
		          "type": "object",
		          "additionalProperties": {
		            "anyOf": [
		              {
		                "$recursiveRef": "#"
		              },
		              {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            ]
		          }
		        }
		      }
		    },
		    "script": {
		      "type": "string"
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Deploy-result": []byte(`
		{
		  "type": "object",
		  "additionalProperties": true,
		  "required": [
		    "releases"
		  ],
		  "properties": {
		    "releases": {
		      "type": "array",
		      "items": {
		        "examples": [
		          {
		            "service": "generic-service",
		            "entry": 0,
		            "result": "some result"
		          }
		        ],
		        "type": "object",
		        "additionalProperties": true,
		        "properties": {
		          "service": {
		            "description": "Name of the object, unique within the kind.",
		            "type": "string",
		            "minLength": 1,
		            "pattern": "^[a-z][A-Za-z0-9_-]*$"
		          },
		          "entry": {
		            "type": "integer",
		            "min": 0
		          },
		          "result": {
		            "type": "string"
		          }
		        }
		      }
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Deployer": []byte(`
		{
		  "title": "Deployer",
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind",
		    "name",
		    "script"
		  ],
		  "additionalProperties": false,
		  "properties": {
		    "apiVersion": {
		      "description": "Version of the configuration format.",
		      "const": "g2a-cli/v2.0"
		    },
		    "kind": {
		      "const": "Deployer"
		    },
		    "name": {
		      "description": "Name of the object, unique within the kind.",
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "schema": {
		      "$schema": "https://json-schema.org/draft/2019-09/schema",
		      "$id": "https://json-schema.org/draft/2019-09/schema",
		      "$vocabulary": {
		        "https://json-schema.org/draft/2019-09/vocab/core": true,
		        "https://json-schema.org/draft/2019-09/vocab/applicator": true,
		        "https://json-schema.org/draft/2019-09/vocab/validation": true,
		        "https://json-schema.org/draft/2019-09/vocab/meta-data": true,
		        "https://json-schema.org/draft/2019-09/vocab/format": false,
		        "https://json-schema.org/draft/2019-09/vocab/content": true
		      },
		      "$recursiveAnchor": true,
		      "title": "Core and Validation specifications meta-schema",
		      "allOf": [
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/core",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/core": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Core vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "$id": {
		              "type": "string",
		              "format": "uri-reference",
		              "$comment": "Non-empty fragments not allowed.",
		              "pattern": "^[^#]*#?$"
		            },
		            "$schema": {
		              "type": "string",
		              "format": "uri"
		            },
		            "$anchor": {
		              "type": "string",
		              "pattern": "^[A-Za-z][-A-Za-z0-9.:_]*$"
		            },
		            "$ref": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveRef": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveAnchor": {
		              "type": "boolean",
		              "default": false
		            },
		            "$vocabulary": {
		              "type": "object",
		              "propertyNames": {
		                "type": "string",
		                "format": "uri"
		              },
		              "additionalProperties": {
		                "type": "boolean"
		              }
		            },
		            "$comment": {
		              "type": "string"
		            },
		            "$defs": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/applicator",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/applicator": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Applicator vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "additionalItems": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedItems": {
		              "$recursiveRef": "#"
		            },
		            "items": {
		              "anyOf": [
		                {
		                  "$recursiveRef": "#"
		                },
		                {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                }
		              ]
		            },
		            "contains": {
		              "$recursiveRef": "#"
		            },
		            "additionalProperties": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedProperties": {
		              "$recursiveRef": "#"
		            },
		            "properties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            },
		            "patternProperties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "propertyNames": {
		                "format": "regex"
		              },
		              "default": {}
		            },
		            "dependentSchemas": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              }
		            },
		            "propertyNames": {
		              "$recursiveRef": "#"
		            },
		            "if": {
		              "$recursiveRef": "#"
		            },
		            "then": {
		              "$recursiveRef": "#"
		            },
		            "else": {
		              "$recursiveRef": "#"
		            },
		            "allOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "anyOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "oneOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "not": {
		              "$recursiveRef": "#"
		            }
		          },
		          "$defs": {
		            "schemaArray": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/validation",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/validation": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Validation vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "multipleOf": {
		              "type": "number",
		              "exclusiveMinimum": 0
		            },
		            "maximum": {
		              "type": "number"
		            },
		            "exclusiveMaximum": {
		              "type": "number"
		            },
		            "minimum": {
		              "type": "number"
		            },
		            "exclusiveMinimum": {
		              "type": "number"
		            },
		            "maxLength": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minLength": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "pattern": {
		              "type": "string",
		              "format": "regex"
		            },
		            "maxItems": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minItems": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "uniqueItems": {
		              "type": "boolean",
		              "default": false
		            },
		            "maxContains": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minContains": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 1
		            },
		            "maxProperties": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minProperties": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "required": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            },
		            "dependentRequired": {
		              "type": "object",
		              "additionalProperties": {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            },
		            "const": true,
		            "enum": {
		              "type": "array",
		              "items": true
		            },
		            "type": {
		              "anyOf": [
		                {
		                  "enum": [
		                    "array",
		                    "boolean",
		                    "integer",
		                    "null",
		                    "number",
		                    "object",
		                    "string"
		                  ]
		                },
		                {
		                  "type": "array",
		                  "items": {
		                    "enum": [
		                      "array",
		                      "boolean",
		                      "integer",
		                      "null",
		                      "number",
		                      "object",
		                      "string"
		                    ]
		                  },
		                  "minItems": 1,
		                  "uniqueItems": true
		                }
		              ]
		            }
		          },
		          "$defs": {
		            "nonNegativeInteger": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "nonNegativeIntegerDefault0": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 0
		            },
		            "simpleTypes": {
		              "enum": [
		                "array",
		                "boolean",
		                "integer",
		                "null",
		                "number",
		                "object",
		                "string"
		              ]
		            },
		            "stringArray": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/meta-data",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/meta-data": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Meta-data vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "title": {
		              "type": "string"
		            },
		            "description": {
		              "type": "string"
		            },
		            "default": true,
		            "deprecated": {
		              "type": "boolean",
		              "default": false
		            },
		            "readOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "writeOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "examples": {
		              "type": "array",
		              "items": true
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/format",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/format": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Format vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "format": {
		              "type": "string"
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/content",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/content": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Content vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "contentMediaType": {
		              "type": "string"
		            },
		            "contentEncoding": {
		              "type": "string"
		            },
		            "contentSchema": {
		              "$recursiveRef": "#"
		            }
		          }
		        }
		      ],
		      "type": [
		        "object",
		        "boolean"
		      ],
		      "properties": {
		        "definitions": {
		          "$comment": "While no longer an official keyword as it is replaced by $defs, this keyword is retained in the meta-schema to prevent incompatible extensions as it remains in common use.",
		          "type": "object",
		          "additionalProperties": {
		            "$recursiveRef": "#"
		          },
		          "default": {}
		        },
		        "dependencies": {
		          "$comment": "\"dependencies\" is no longer a keyword, but schema authors should avoid redefining it to facilitate a smooth transition to \"dependentSchemas\" and \"dependentRequired\"",
		          "type": "object",
		          "additionalProperties": {
		            "anyOf": [
		              {
		                "$recursiveRef": "#"
		              },
		              {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            ]
		          }
		        }
		      }
		    },
		    "script": {
		      "type": "string"
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Environment": []byte(`
		{
		  "title": "Environment",
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind",
		    "name"
		  ],
		  "properties": {
		    "apiVersion": {
		      "description": "Version of the configuration format.",
		      "const": "g2a-cli/v2.0"
		    },
		    "kind": {
		      "description": "Determines type of the document.",
		      "const": "Environment"
		    },
		    "name": {
		      "description": "Unique name used to identify environment.",
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "deployServices": {
		      "description": "Default list of the services to deploy to this environment. It may be modified by using \"--services\" option.\n",
		      "type": "array",
		      "items": {
		        "description": "Name of the object, unique within the kind.",
		        "type": "string",
		        "minLength": 1,
		        "pattern": "^[a-z][A-Za-z0-9_-]*$"
		      }
		    },
		    "variables": {
		      "description": "Definitions of the variables to use in the configuration files. Names are case-insensitive.\n",
		      "examples": [
		        {
		          "name": "value"
		        }
		      ],
		      "type": "object",
		      "patternProperties": {
		        "^[a-zA-Z][a-zA-Z0-9]*$": {
		          "type": "string"
		        }
		      }
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Executor": []byte(`
		{
		  "title": "Executor",
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind",
		    "name",
		    "script"
		  ],
		  "additionalProperties": false,
		  "properties": {
		    "apiVersion": {
		      "description": "Version of the configuration format.",
		      "const": "g2a-cli/v2.0"
		    },
		    "kind": {
		      "description": "Determines type of the document.",
		      "examples": [
		        "Builder"
		      ],
		      "enum": [
		        "Builder",
		        "Deployer",
		        "Tagger",
		        "Pusher"
		      ]
		    },
		    "name": {
		      "description": "Name used to identify executor. Unique together with kind.",
		      "examples": [
		        "docker"
		      ],
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "schema": {
		      "description": "JSON Schema defining format of the configuration for the executor.",
		      "examples": [
		        {
		          "type": "object",
		          "required": [
		            "image"
		          ],
		          "properties": {
		            "image": {
		              "type": "string"
		            },
		            "context": {
		              "type": "string"
		            }
		          }
		        }
		      ],
		      "$schema": "https://json-schema.org/draft/2019-09/schema",
		      "$id": "https://json-schema.org/draft/2019-09/schema",
		      "$vocabulary": {
		        "https://json-schema.org/draft/2019-09/vocab/core": true,
		        "https://json-schema.org/draft/2019-09/vocab/applicator": true,
		        "https://json-schema.org/draft/2019-09/vocab/validation": true,
		        "https://json-schema.org/draft/2019-09/vocab/meta-data": true,
		        "https://json-schema.org/draft/2019-09/vocab/format": false,
		        "https://json-schema.org/draft/2019-09/vocab/content": true
		      },
		      "$recursiveAnchor": true,
		      "title": "Core and Validation specifications meta-schema",
		      "allOf": [
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/core",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/core": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Core vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "$id": {
		              "type": "string",
		              "format": "uri-reference",
		              "$comment": "Non-empty fragments not allowed.",
		              "pattern": "^[^#]*#?$"
		            },
		            "$schema": {
		              "type": "string",
		              "format": "uri"
		            },
		            "$anchor": {
		              "type": "string",
		              "pattern": "^[A-Za-z][-A-Za-z0-9.:_]*$"
		            },
		            "$ref": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveRef": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveAnchor": {
		              "type": "boolean",
		              "default": false
		            },
		            "$vocabulary": {
		              "type": "object",
		              "propertyNames": {
		                "type": "string",
		                "format": "uri"
		              },
		              "additionalProperties": {
		                "type": "boolean"
		              }
		            },
		            "$comment": {
		              "type": "string"
		            },
		            "$defs": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/applicator",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/applicator": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Applicator vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "additionalItems": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedItems": {
		              "$recursiveRef": "#"
		            },
		            "items": {
		              "anyOf": [
		                {
		                  "$recursiveRef": "#"
		                },
		                {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                }
		              ]
		            },
		            "contains": {
		              "$recursiveRef": "#"
		            },
		            "additionalProperties": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedProperties": {
		              "$recursiveRef": "#"
		            },
		            "properties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            },
		            "patternProperties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "propertyNames": {
		                "format": "regex"
		              },
		              "default": {}
		            },
		            "dependentSchemas": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              }
		            },
		            "propertyNames": {
		              "$recursiveRef": "#"
		            },
		            "if": {
		              "$recursiveRef": "#"
		            },
		            "then": {
		              "$recursiveRef": "#"
		            },
		            "else": {
		              "$recursiveRef": "#"
		            },
		            "allOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "anyOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "oneOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "not": {
		              "$recursiveRef": "#"
		            }
		          },
		          "$defs": {
		            "schemaArray": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/validation",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/validation": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Validation vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "multipleOf": {
		              "type": "number",
		              "exclusiveMinimum": 0
		            },
		            "maximum": {
		              "type": "number"
		            },
		            "exclusiveMaximum": {
		              "type": "number"
		            },
		            "minimum": {
		              "type": "number"
		            },
		            "exclusiveMinimum": {
		              "type": "number"
		            },
		            "maxLength": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minLength": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "pattern": {
		              "type": "string",
		              "format": "regex"
		            },
		            "maxItems": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minItems": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "uniqueItems": {
		              "type": "boolean",
		              "default": false
		            },
		            "maxContains": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minContains": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 1
		            },
		            "maxProperties": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minProperties": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "required": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            },
		            "dependentRequired": {
		              "type": "object",
		              "additionalProperties": {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            },
		            "const": true,
		            "enum": {
		              "type": "array",
		              "items": true
		            },
		            "type": {
		              "anyOf": [
		                {
		                  "enum": [
		                    "array",
		                    "boolean",
		                    "integer",
		                    "null",
		                    "number",
		                    "object",
		                    "string"
		                  ]
		                },
		                {
		                  "type": "array",
		                  "items": {
		                    "enum": [
		                      "array",
		                      "boolean",
		                      "integer",
		                      "null",
		                      "number",
		                      "object",
		                      "string"
		                    ]
		                  },
		                  "minItems": 1,
		                  "uniqueItems": true
		                }
		              ]
		            }
		          },
		          "$defs": {
		            "nonNegativeInteger": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "nonNegativeIntegerDefault0": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 0
		            },
		            "simpleTypes": {
		              "enum": [
		                "array",
		                "boolean",
		                "integer",
		                "null",
		                "number",
		                "object",
		                "string"
		              ]
		            },
		            "stringArray": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/meta-data",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/meta-data": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Meta-data vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "title": {
		              "type": "string"
		            },
		            "description": {
		              "type": "string"
		            },
		            "default": true,
		            "deprecated": {
		              "type": "boolean",
		              "default": false
		            },
		            "readOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "writeOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "examples": {
		              "type": "array",
		              "items": true
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/format",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/format": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Format vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "format": {
		              "type": "string"
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/content",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/content": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Content vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "contentMediaType": {
		              "type": "string"
		            },
		            "contentEncoding": {
		              "type": "string"
		            },
		            "contentSchema": {
		              "$recursiveRef": "#"
		            }
		          }
		        }
		      ],
		      "type": [
		        "object",
		        "boolean"
		      ],
		      "properties": {
		        "definitions": {
		          "$comment": "While no longer an official keyword as it is replaced by $defs, this keyword is retained in the meta-schema to prevent incompatible extensions as it remains in common use.",
		          "type": "object",
		          "additionalProperties": {
		            "$recursiveRef": "#"
		          },
		          "default": {}
		        },
		        "dependencies": {
		          "$comment": "\"dependencies\" is no longer a keyword, but schema authors should avoid redefining it to facilitate a smooth transition to \"dependentSchemas\" and \"dependentRequired\"",
		          "type": "object",
		          "additionalProperties": {
		            "anyOf": [
		              {
		                "$recursiveRef": "#"
		              },
		              {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            ]
		          }
		        }
		      }
		    },
		    "script": {
		      "description": "Implementation of the executor.",
		      "examples": [
		        "for (let tag of input.tags) {\n  let image = input.spec.image + ':' + tag\n  let context = input.spec.context || \".\"\n\n  exec(\"docker\", [ 'build', context, '-t', image ])\n\n  output = image\n}\n"
		      ],
		      "type": "string"
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Object": []byte(`
		{
		  "title": "Object",
		  "oneOf": [
		    {
		      "title": "Builder",
		      "type": "object",
		      "required": [
		        "apiVersion",
		        "kind",
		        "name",
		        "script"
		      ],
		      "additionalProperties": false,
		      "properties": {
		        "apiVersion": {
		          "description": "Version of the configuration format.",
		          "const": "g2a-cli/v2.0"
		        },
		        "kind": {
		          "const": "Builder"
		        },
		        "name": {
		          "description": "Name of the object, unique within the kind.",
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$"
		        },
		        "schema": {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/schema",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/core": true,
		            "https://json-schema.org/draft/2019-09/vocab/applicator": true,
		            "https://json-schema.org/draft/2019-09/vocab/validation": true,
		            "https://json-schema.org/draft/2019-09/vocab/meta-data": true,
		            "https://json-schema.org/draft/2019-09/vocab/format": false,
		            "https://json-schema.org/draft/2019-09/vocab/content": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Core and Validation specifications meta-schema",
		          "allOf": [
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/core",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/core": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Core vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "$id": {
		                  "type": "string",
		                  "format": "uri-reference",
		                  "$comment": "Non-empty fragments not allowed.",
		                  "pattern": "^[^#]*#?$"
		                },
		                "$schema": {
		                  "type": "string",
		                  "format": "uri"
		                },
		                "$anchor": {
		                  "type": "string",
		                  "pattern": "^[A-Za-z][-A-Za-z0-9.:_]*$"
		                },
		                "$ref": {
		                  "type": "string",
		                  "format": "uri-reference"
		                },
		                "$recursiveRef": {
		                  "type": "string",
		                  "format": "uri-reference"
		                },
		                "$recursiveAnchor": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "$vocabulary": {
		                  "type": "object",
		                  "propertyNames": {
		                    "type": "string",
		                    "format": "uri"
		                  },
		                  "additionalProperties": {
		                    "type": "boolean"
		                  }
		                },
		                "$comment": {
		                  "type": "string"
		                },
		                "$defs": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "default": {}
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/applicator",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/applicator": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Applicator vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "additionalItems": {
		                  "$recursiveRef": "#"
		                },
		                "unevaluatedItems": {
		                  "$recursiveRef": "#"
		                },
		                "items": {
		                  "anyOf": [
		                    {
		                      "$recursiveRef": "#"
		                    },
		                    {
		                      "type": "array",
		                      "minItems": 1,
		                      "items": {
		                        "$recursiveRef": "#"
		                      }
		                    }
		                  ]
		                },
		                "contains": {
		                  "$recursiveRef": "#"
		                },
		                "additionalProperties": {
		                  "$recursiveRef": "#"
		                },
		                "unevaluatedProperties": {
		                  "$recursiveRef": "#"
		                },
		                "properties": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "default": {}
		                },
		                "patternProperties": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "propertyNames": {
		                    "format": "regex"
		                  },
		                  "default": {}
		                },
		                "dependentSchemas": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "propertyNames": {
		                  "$recursiveRef": "#"
		                },
		                "if": {
		                  "$recursiveRef": "#"
		                },
		                "then": {
		                  "$recursiveRef": "#"
		                },
		                "else": {
		                  "$recursiveRef": "#"
		                },
		                "allOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "anyOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "oneOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "not": {
		                  "$recursiveRef": "#"
		                }
		              },
		              "$defs": {
		                "schemaArray": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/validation",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/validation": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Validation vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "multipleOf": {
		                  "type": "number",
		                  "exclusiveMinimum": 0
		                },
		                "maximum": {
		                  "type": "number"
		                },
		                "exclusiveMaximum": {
		                  "type": "number"
		                },
		                "minimum": {
		                  "type": "number"
		                },
		                "exclusiveMinimum": {
		                  "type": "number"
		                },
		                "maxLength": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minLength": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "pattern": {
		                  "type": "string",
		                  "format": "regex"
		                },
		                "maxItems": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minItems": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "uniqueItems": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "maxContains": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minContains": {
		                  "type": "integer",
		                  "minimum": 0,
		                  "default": 1
		                },
		                "maxProperties": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minProperties": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "required": {
		                  "type": "array",
		                  "items": {
		                    "type": "string"
		                  },
		                  "uniqueItems": true,
		                  "default": []
		                },
		                "dependentRequired": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "type": "array",
		                    "items": {
		                      "type": "string"
		                    },
		                    "uniqueItems": true,
		                    "default": []
		                  }
		                },
		                "const": true,
		                "enum": {
		                  "type": "array",
		                  "items": true
		                },
		                "type": {
		                  "anyOf": [
		                    {
		                      "enum": [
		                        "array",
		                        "boolean",
		                        "integer",
		                        "null",
		                        "number",
		                        "object",
		                        "string"
		                      ]
		                    },
		                    {
		                      "type": "array",
		                      "items": {
		                        "enum": [
		                          "array",
		                          "boolean",
		                          "integer",
		                          "null",
		                          "number",
		                          "object",
		                          "string"
		                        ]
		                      },
		                      "minItems": 1,
		                      "uniqueItems": true
		                    }
		                  ]
		                }
		              },
		              "$defs": {
		                "nonNegativeInteger": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "nonNegativeIntegerDefault0": {
		                  "type": "integer",
		                  "minimum": 0,
		                  "default": 0
		                },
		                "simpleTypes": {
		                  "enum": [
		                    "array",
		                    "boolean",
		                    "integer",
		                    "null",
		                    "number",
		                    "object",
		                    "string"
		                  ]
		                },
		                "stringArray": {
		                  "type": "array",
		                  "items": {
		                    "type": "string"
		                  },
		                  "uniqueItems": true,
		                  "default": []
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/meta-data",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/meta-data": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Meta-data vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "title": {
		                  "type": "string"
		                },
		                "description": {
		                  "type": "string"
		                },
		                "default": true,
		                "deprecated": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "readOnly": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "writeOnly": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "examples": {
		                  "type": "array",
		                  "items": true
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/format",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/format": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Format vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "format": {
		                  "type": "string"
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/content",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/content": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Content vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "contentMediaType": {
		                  "type": "string"
		                },
		                "contentEncoding": {
		                  "type": "string"
		                },
		                "contentSchema": {
		                  "$recursiveRef": "#"
		                }
		              }
		            }
		          ],
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "definitions": {
		              "$comment": "While no longer an official keyword as it is replaced by $defs, this keyword is retained in the meta-schema to prevent incompatible extensions as it remains in common use.",
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            },
		            "dependencies": {
		              "$comment": "\"dependencies\" is no longer a keyword, but schema authors should avoid redefining it to facilitate a smooth transition to \"dependentSchemas\" and \"dependentRequired\"",
		              "type": "object",
		              "additionalProperties": {
		                "anyOf": [
		                  {
		                    "$recursiveRef": "#"
		                  },
		                  {
		                    "type": "array",
		                    "items": {
		                      "type": "string"
		                    },
		                    "uniqueItems": true,
		                    "default": []
		                  }
		                ]
		              }
		            }
		          }
		        },
		        "script": {
		          "type": "string"
		        }
		      }
		    },
		    {
		      "title": "Deployer",
		      "type": "object",
		      "required": [
		        "apiVersion",
		        "kind",
		        "name",
		        "script"
		      ],
		      "additionalProperties": false,
		      "properties": {
		        "apiVersion": {
		          "description": "Version of the configuration format.",
		          "const": "g2a-cli/v2.0"
		        },
		        "kind": {
		          "const": "Deployer"
		        },
		        "name": {
		          "description": "Name of the object, unique within the kind.",
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$"
		        },
		        "schema": {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/schema",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/core": true,
		            "https://json-schema.org/draft/2019-09/vocab/applicator": true,
		            "https://json-schema.org/draft/2019-09/vocab/validation": true,
		            "https://json-schema.org/draft/2019-09/vocab/meta-data": true,
		            "https://json-schema.org/draft/2019-09/vocab/format": false,
		            "https://json-schema.org/draft/2019-09/vocab/content": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Core and Validation specifications meta-schema",
		          "allOf": [
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/core",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/core": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Core vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "$id": {
		                  "type": "string",
		                  "format": "uri-reference",
		                  "$comment": "Non-empty fragments not allowed.",
		                  "pattern": "^[^#]*#?$"
		                },
		                "$schema": {
		                  "type": "string",
		                  "format": "uri"
		                },
		                "$anchor": {
		                  "type": "string",
		                  "pattern": "^[A-Za-z][-A-Za-z0-9.:_]*$"
		                },
		                "$ref": {
		                  "type": "string",
		                  "format": "uri-reference"
		                },
		                "$recursiveRef": {
		                  "type": "string",
		                  "format": "uri-reference"
		                },
		                "$recursiveAnchor": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "$vocabulary": {
		                  "type": "object",
		                  "propertyNames": {
		                    "type": "string",
		                    "format": "uri"
		                  },
		                  "additionalProperties": {
		                    "type": "boolean"
		                  }
		                },
		                "$comment": {
		                  "type": "string"
		                },
		                "$defs": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "default": {}
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/applicator",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/applicator": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Applicator vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "additionalItems": {
		                  "$recursiveRef": "#"
		                },
		                "unevaluatedItems": {
		                  "$recursiveRef": "#"
		                },
		                "items": {
		                  "anyOf": [
		                    {
		                      "$recursiveRef": "#"
		                    },
		                    {
		                      "type": "array",
		                      "minItems": 1,
		                      "items": {
		                        "$recursiveRef": "#"
		                      }
		                    }
		                  ]
		                },
		                "contains": {
		                  "$recursiveRef": "#"
		                },
		                "additionalProperties": {
		                  "$recursiveRef": "#"
		                },
		                "unevaluatedProperties": {
		                  "$recursiveRef": "#"
		                },
		                "properties": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "default": {}
		                },
		                "patternProperties": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "propertyNames": {
		                    "format": "regex"
		                  },
		                  "default": {}
		                },
		                "dependentSchemas": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "propertyNames": {
		                  "$recursiveRef": "#"
		                },
		                "if": {
		                  "$recursiveRef": "#"
		                },
		                "then": {
		                  "$recursiveRef": "#"
		                },
		                "else": {
		                  "$recursiveRef": "#"
		                },
		                "allOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "anyOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "oneOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "not": {
		                  "$recursiveRef": "#"
		                }
		              },
		              "$defs": {
		                "schemaArray": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/validation",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/validation": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Validation vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "multipleOf": {
		                  "type": "number",
		                  "exclusiveMinimum": 0
		                },
		                "maximum": {
		                  "type": "number"
		                },
		                "exclusiveMaximum": {
		                  "type": "number"
		                },
		                "minimum": {
		                  "type": "number"
		                },
		                "exclusiveMinimum": {
		                  "type": "number"
		                },
		                "maxLength": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minLength": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "pattern": {
		                  "type": "string",
		                  "format": "regex"
		                },
		                "maxItems": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minItems": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "uniqueItems": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "maxContains": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minContains": {
		                  "type": "integer",
		                  "minimum": 0,
		                  "default": 1
		                },
		                "maxProperties": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minProperties": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "required": {
		                  "type": "array",
		                  "items": {
		                    "type": "string"
		                  },
		                  "uniqueItems": true,
		                  "default": []
		                },
		                "dependentRequired": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "type": "array",
		                    "items": {
		                      "type": "string"
		                    },
		                    "uniqueItems": true,
		                    "default": []
		                  }
		                },
		                "const": true,
		                "enum": {
		                  "type": "array",
		                  "items": true
		                },
		                "type": {
		                  "anyOf": [
		                    {
		                      "enum": [
		                        "array",
		                        "boolean",
		                        "integer",
		                        "null",
		                        "number",
		                        "object",
		                        "string"
		                      ]
		                    },
		                    {
		                      "type": "array",
		                      "items": {
		                        "enum": [
		                          "array",
		                          "boolean",
		                          "integer",
		                          "null",
		                          "number",
		                          "object",
		                          "string"
		                        ]
		                      },
		                      "minItems": 1,
		                      "uniqueItems": true
		                    }
		                  ]
		                }
		              },
		              "$defs": {
		                "nonNegativeInteger": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "nonNegativeIntegerDefault0": {
		                  "type": "integer",
		                  "minimum": 0,
		                  "default": 0
		                },
		                "simpleTypes": {
		                  "enum": [
		                    "array",
		                    "boolean",
		                    "integer",
		                    "null",
		                    "number",
		                    "object",
		                    "string"
		                  ]
		                },
		                "stringArray": {
		                  "type": "array",
		                  "items": {
		                    "type": "string"
		                  },
		                  "uniqueItems": true,
		                  "default": []
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/meta-data",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/meta-data": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Meta-data vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "title": {
		                  "type": "string"
		                },
		                "description": {
		                  "type": "string"
		                },
		                "default": true,
		                "deprecated": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "readOnly": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "writeOnly": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "examples": {
		                  "type": "array",
		                  "items": true
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/format",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/format": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Format vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "format": {
		                  "type": "string"
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/content",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/content": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Content vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "contentMediaType": {
		                  "type": "string"
		                },
		                "contentEncoding": {
		                  "type": "string"
		                },
		                "contentSchema": {
		                  "$recursiveRef": "#"
		                }
		              }
		            }
		          ],
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "definitions": {
		              "$comment": "While no longer an official keyword as it is replaced by $defs, this keyword is retained in the meta-schema to prevent incompatible extensions as it remains in common use.",
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            },
		            "dependencies": {
		              "$comment": "\"dependencies\" is no longer a keyword, but schema authors should avoid redefining it to facilitate a smooth transition to \"dependentSchemas\" and \"dependentRequired\"",
		              "type": "object",
		              "additionalProperties": {
		                "anyOf": [
		                  {
		                    "$recursiveRef": "#"
		                  },
		                  {
		                    "type": "array",
		                    "items": {
		                      "type": "string"
		                    },
		                    "uniqueItems": true,
		                    "default": []
		                  }
		                ]
		              }
		            }
		          }
		        },
		        "script": {
		          "type": "string"
		        }
		      }
		    },
		    {
		      "title": "Environment",
		      "type": "object",
		      "required": [
		        "apiVersion",
		        "kind",
		        "name"
		      ],
		      "properties": {
		        "apiVersion": {
		          "description": "Version of the configuration format.",
		          "const": "g2a-cli/v2.0"
		        },
		        "kind": {
		          "description": "Determines type of the document.",
		          "const": "Environment"
		        },
		        "name": {
		          "description": "Name of the object, unique within the kind.",
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$"
		        },
		        "deployServices": {
		          "description": "Default list of the services to deploy to this environment. It may be modified by using \"--services\" option.\n",
		          "type": "array",
		          "items": {
		            "description": "Name of the object, unique within the kind.",
		            "type": "string",
		            "minLength": 1,
		            "pattern": "^[a-z][A-Za-z0-9_-]*$"
		          }
		        },
		        "variables": {
		          "description": "Definitions of the variables to use in the configuration files. Names are case-insensitive.\n",
		          "examples": [
		            {
		              "name": "value"
		            }
		          ],
		          "type": "object",
		          "patternProperties": {
		            "^[a-zA-Z][a-zA-Z0-9]*$": {
		              "type": "string"
		            }
		          }
		        }
		      }
		    },
		    {
		      "title": "Project",
		      "description": null,
		      "type": "object",
		      "required": [
		        "apiVersion",
		        "kind",
		        "name"
		      ],
		      "properties": {
		        "apiVersion": {
		          "description": "Version of the configuration format.",
		          "const": "g2a-cli/v2.0"
		        },
		        "kind": {
		          "description": "Determines type of the document.",
		          "const": "Project"
		        },
		        "name": {
		          "description": "Name of the object, unique within the kind.",
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$",
		          "examples": [
		            "generic-api"
		          ]
		        },
		        "files": {
		          "description": "List of the configuration files to load.",
		          "oneOf": [
		            {
		              "type": "array",
		              "items": {
		                "oneOf": [
		                  {
		                    "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		                    "examples": [
		                      "services/*/service.yaml",
		                      "environments/*/environments.yaml"
		                    ],
		                    "type": "string",
		                    "minLength": 1
		                  },
		                  {
		                    "type": "object",
		                    "additionalProperties": false,
		                    "required": [
		                      "git"
		                    ],
		                    "properties": {
		                      "git": {
		                        "description": "Files may be also fetched from another git repository.\n",
		                        "type": "object",
		                        "additionalProperties": false,
		                        "required": [
		                          "url",
		                          "rev",
		                          "files"
		                        ],
		                        "properties": {
		                          "url": {
		                            "examples": [
		                              "https://github.com/g2a-com/cicd"
		                            ],
		                            "type": "string",
		                            "minLength": 1
		                          },
		                          "rev": {
		                            "type": "string",
		                            "examples": [
		                              "main"
		                            ],
		                            "minLength": 1
		                          },
		                          "files": {
		                            "examples": [
		                              "executors/*/*.yaml"
		                            ],
		                            "oneOf": [
		                              {
		                                "type": "array",
		                                "items": {
		                                  "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		                                  "examples": [
		                                    "services/*/service.yaml",
		                                    "environments/*/environments.yaml"
		                                  ],
		                                  "type": "string",
		                                  "minLength": 1
		                                }
		                              },
		                              {
		                                "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		                                "examples": [
		                                  "services/*/service.yaml",
		                                  "environments/*/environments.yaml"
		                                ],
		                                "type": "string",
		                                "minLength": 1
		                              }
		                            ]
		                          }
		                        }
		                      }
		                    }
		                  }
		                ]
		              }
		            },
		            {
		              "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		              "examples": [
		                "services/*/service.yaml",
		                "environments/*/environments.yaml"
		              ],
		              "type": "string",
		              "minLength": 1
		            }
		          ]
		        },
		        "variables": {
		          "description": "Definitions of the variables to use in the configuration files. Names are case-insensitive.\n",
		          "examples": [
		            {
		              "name": "value"
		            }
		          ],
		          "type": "object",
		          "patternProperties": {
		            "^[a-zA-Z][a-zA-Z0-9]*$": {
		              "type": "string"
		            }
		          }
		        },
		        "tasks": {
		          "description": "Definitions of the tasks used by commands \"prepare\", \"test\", \"lint\" and \"run\". These tasks may be also specified in services definitions.",
		          "type": "object",
		          "properties": {
		            "prepare": {
		              "description": "Defines steps required to preapre freshly clonned repository for development, tests or build. This definition is used by \"prepare\" command.\n",
		              "examples": [
		                [
		                  {
		                    "make": {
		                      "target": "prepare"
		                    }
		                  }
		                ]
		              ],
		              "type": "array",
		              "items": {
		                "oneOf": [
		                  {
		                    "type": "object",
		                    "minProperties": 1,
		                    "maxProperties": 1,
		                    "additionalProperties": true
		                  },
		                  {
		                    "type": "string"
		                  }
		                ]
		              }
		            },
		            "test": {
		              "description": "Defines how to run tests. This definition is used by \"test\" command.\n",
		              "examples": [
		                [
		                  {
		                    "script": {
		                      "sh": "go test ./..."
		                    }
		                  }
		                ]
		              ],
		              "type": "array",
		              "items": {
		                "oneOf": [
		                  {
		                    "type": "object",
		                    "minProperties": 1,
		                    "maxProperties": 1,
		                    "additionalProperties": true
		                  },
		                  {
		                    "type": "string"
		                  }
		                ]
		              }
		            },
		            "lint": {
		              "description": "Defines how to lint the code. Ideally should try to fix the issues. This definition is used by \"lint\" command.\n",
		              "examples": [
		                [
		                  "prettier"
		                ]
		              ],
		              "type": "array",
		              "items": {
		                "oneOf": [
		                  {
		                    "type": "object",
		                    "minProperties": 1,
		                    "maxProperties": 1,
		                    "additionalProperties": true
		                  },
		                  {
		                    "type": "string"
		                  }
		                ]
		              }
		            }
		          },
		          "additionalProperties": {
		            "description": "You are not bound to use pre-defined tasks. All tasks (including custom ones) may be run using \"run\" command.\n",
		            "examples": [
		              [
		                {
		                  "runnerName": {
		                    "some": "params"
		                  }
		                }
		              ]
		            ],
		            "tsType": "({ [k: string]: unknown; } | string )[] | undefined",
		            "type": "array",
		            "items": {
		              "oneOf": [
		                {
		                  "type": "object",
		                  "minProperties": 1,
		                  "maxProperties": 1,
		                  "additionalProperties": true
		                },
		                {
		                  "type": "string"
		                }
		              ]
		            }
		          },
		          "$defs": {
		            "task": {
		              "type": "array",
		              "items": {
		                "oneOf": [
		                  {
		                    "type": "object",
		                    "minProperties": 1,
		                    "maxProperties": 1,
		                    "additionalProperties": true
		                  },
		                  {
		                    "type": "string"
		                  }
		                ]
		              }
		            }
		          }
		        }
		      },
		      "$defs": {
		        "glob": {
		          "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		          "examples": [
		            "services/*/service.yaml",
		            "environments/*/environments.yaml"
		          ],
		          "type": "string",
		          "minLength": 1
		        },
		        "gitGlob": {
		          "type": "object",
		          "additionalProperties": false,
		          "required": [
		            "git"
		          ],
		          "properties": {
		            "git": {
		              "description": "Files may be also fetched from another git repository.\n",
		              "type": "object",
		              "additionalProperties": false,
		              "required": [
		                "url",
		                "rev",
		                "files"
		              ],
		              "properties": {
		                "url": {
		                  "examples": [
		                    "https://github.com/g2a-com/cicd"
		                  ],
		                  "type": "string",
		                  "minLength": 1
		                },
		                "rev": {
		                  "type": "string",
		                  "examples": [
		                    "main"
		                  ],
		                  "minLength": 1
		                },
		                "files": {
		                  "examples": [
		                    "executors/*/*.yaml"
		                  ],
		                  "oneOf": [
		                    {
		                      "type": "array",
		                      "items": {
		                        "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		                        "examples": [
		                          "services/*/service.yaml",
		                          "environments/*/environments.yaml"
		                        ],
		                        "type": "string",
		                        "minLength": 1
		                      }
		                    },
		                    {
		                      "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		                      "examples": [
		                        "services/*/service.yaml",
		                        "environments/*/environments.yaml"
		                      ],
		                      "type": "string",
		                      "minLength": 1
		                    }
		                  ]
		                }
		              }
		            }
		          }
		        }
		      }
		    },
		    {
		      "title": "Pusher",
		      "type": "object",
		      "required": [
		        "apiVersion",
		        "kind",
		        "name",
		        "script"
		      ],
		      "additionalProperties": false,
		      "properties": {
		        "apiVersion": {
		          "description": "Version of the configuration format.",
		          "const": "g2a-cli/v2.0"
		        },
		        "kind": {
		          "const": "Pusher"
		        },
		        "name": {
		          "description": "Name of the object, unique within the kind.",
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$"
		        },
		        "schema": {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/schema",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/core": true,
		            "https://json-schema.org/draft/2019-09/vocab/applicator": true,
		            "https://json-schema.org/draft/2019-09/vocab/validation": true,
		            "https://json-schema.org/draft/2019-09/vocab/meta-data": true,
		            "https://json-schema.org/draft/2019-09/vocab/format": false,
		            "https://json-schema.org/draft/2019-09/vocab/content": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Core and Validation specifications meta-schema",
		          "allOf": [
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/core",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/core": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Core vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "$id": {
		                  "type": "string",
		                  "format": "uri-reference",
		                  "$comment": "Non-empty fragments not allowed.",
		                  "pattern": "^[^#]*#?$"
		                },
		                "$schema": {
		                  "type": "string",
		                  "format": "uri"
		                },
		                "$anchor": {
		                  "type": "string",
		                  "pattern": "^[A-Za-z][-A-Za-z0-9.:_]*$"
		                },
		                "$ref": {
		                  "type": "string",
		                  "format": "uri-reference"
		                },
		                "$recursiveRef": {
		                  "type": "string",
		                  "format": "uri-reference"
		                },
		                "$recursiveAnchor": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "$vocabulary": {
		                  "type": "object",
		                  "propertyNames": {
		                    "type": "string",
		                    "format": "uri"
		                  },
		                  "additionalProperties": {
		                    "type": "boolean"
		                  }
		                },
		                "$comment": {
		                  "type": "string"
		                },
		                "$defs": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "default": {}
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/applicator",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/applicator": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Applicator vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "additionalItems": {
		                  "$recursiveRef": "#"
		                },
		                "unevaluatedItems": {
		                  "$recursiveRef": "#"
		                },
		                "items": {
		                  "anyOf": [
		                    {
		                      "$recursiveRef": "#"
		                    },
		                    {
		                      "type": "array",
		                      "minItems": 1,
		                      "items": {
		                        "$recursiveRef": "#"
		                      }
		                    }
		                  ]
		                },
		                "contains": {
		                  "$recursiveRef": "#"
		                },
		                "additionalProperties": {
		                  "$recursiveRef": "#"
		                },
		                "unevaluatedProperties": {
		                  "$recursiveRef": "#"
		                },
		                "properties": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "default": {}
		                },
		                "patternProperties": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "propertyNames": {
		                    "format": "regex"
		                  },
		                  "default": {}
		                },
		                "dependentSchemas": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "propertyNames": {
		                  "$recursiveRef": "#"
		                },
		                "if": {
		                  "$recursiveRef": "#"
		                },
		                "then": {
		                  "$recursiveRef": "#"
		                },
		                "else": {
		                  "$recursiveRef": "#"
		                },
		                "allOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "anyOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "oneOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "not": {
		                  "$recursiveRef": "#"
		                }
		              },
		              "$defs": {
		                "schemaArray": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/validation",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/validation": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Validation vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "multipleOf": {
		                  "type": "number",
		                  "exclusiveMinimum": 0
		                },
		                "maximum": {
		                  "type": "number"
		                },
		                "exclusiveMaximum": {
		                  "type": "number"
		                },
		                "minimum": {
		                  "type": "number"
		                },
		                "exclusiveMinimum": {
		                  "type": "number"
		                },
		                "maxLength": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minLength": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "pattern": {
		                  "type": "string",
		                  "format": "regex"
		                },
		                "maxItems": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minItems": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "uniqueItems": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "maxContains": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minContains": {
		                  "type": "integer",
		                  "minimum": 0,
		                  "default": 1
		                },
		                "maxProperties": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minProperties": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "required": {
		                  "type": "array",
		                  "items": {
		                    "type": "string"
		                  },
		                  "uniqueItems": true,
		                  "default": []
		                },
		                "dependentRequired": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "type": "array",
		                    "items": {
		                      "type": "string"
		                    },
		                    "uniqueItems": true,
		                    "default": []
		                  }
		                },
		                "const": true,
		                "enum": {
		                  "type": "array",
		                  "items": true
		                },
		                "type": {
		                  "anyOf": [
		                    {
		                      "enum": [
		                        "array",
		                        "boolean",
		                        "integer",
		                        "null",
		                        "number",
		                        "object",
		                        "string"
		                      ]
		                    },
		                    {
		                      "type": "array",
		                      "items": {
		                        "enum": [
		                          "array",
		                          "boolean",
		                          "integer",
		                          "null",
		                          "number",
		                          "object",
		                          "string"
		                        ]
		                      },
		                      "minItems": 1,
		                      "uniqueItems": true
		                    }
		                  ]
		                }
		              },
		              "$defs": {
		                "nonNegativeInteger": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "nonNegativeIntegerDefault0": {
		                  "type": "integer",
		                  "minimum": 0,
		                  "default": 0
		                },
		                "simpleTypes": {
		                  "enum": [
		                    "array",
		                    "boolean",
		                    "integer",
		                    "null",
		                    "number",
		                    "object",
		                    "string"
		                  ]
		                },
		                "stringArray": {
		                  "type": "array",
		                  "items": {
		                    "type": "string"
		                  },
		                  "uniqueItems": true,
		                  "default": []
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/meta-data",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/meta-data": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Meta-data vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "title": {
		                  "type": "string"
		                },
		                "description": {
		                  "type": "string"
		                },
		                "default": true,
		                "deprecated": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "readOnly": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "writeOnly": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "examples": {
		                  "type": "array",
		                  "items": true
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/format",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/format": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Format vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "format": {
		                  "type": "string"
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/content",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/content": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Content vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "contentMediaType": {
		                  "type": "string"
		                },
		                "contentEncoding": {
		                  "type": "string"
		                },
		                "contentSchema": {
		                  "$recursiveRef": "#"
		                }
		              }
		            }
		          ],
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "definitions": {
		              "$comment": "While no longer an official keyword as it is replaced by $defs, this keyword is retained in the meta-schema to prevent incompatible extensions as it remains in common use.",
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            },
		            "dependencies": {
		              "$comment": "\"dependencies\" is no longer a keyword, but schema authors should avoid redefining it to facilitate a smooth transition to \"dependentSchemas\" and \"dependentRequired\"",
		              "type": "object",
		              "additionalProperties": {
		                "anyOf": [
		                  {
		                    "$recursiveRef": "#"
		                  },
		                  {
		                    "type": "array",
		                    "items": {
		                      "type": "string"
		                    },
		                    "uniqueItems": true,
		                    "default": []
		                  }
		                ]
		              }
		            }
		          }
		        },
		        "script": {
		          "type": "string"
		        }
		      }
		    },
		    {
		      "title": "Service",
		      "type": "object",
		      "required": [
		        "apiVersion",
		        "kind",
		        "name"
		      ],
		      "properties": {
		        "apiVersion": {
		          "description": "Version of the configuration format.",
		          "const": "g2a-cli/v2.0"
		        },
		        "kind": {
		          "description": "Determines type of the document.",
		          "const": "Service"
		        },
		        "name": {
		          "description": "Name of the object, unique within the kind.",
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$",
		          "examples": [
		            "example-api"
		          ]
		        },
		        "artifacts": {
		          "description": "List of artifacts to produce by build command. Each entry describes single artifact like docker image or npm package.\n",
		          "type": "array",
		          "items": {
		            "examples": [
		              {
		                "docker": {
		                  "image": "example.com/test/image"
		                }
		              },
		              {
		                "hugo": {
		                  "dir": "{{ .Service.Dir }}/docs"
		                },
		                "push": {
		                  "artifactory": {
		                    "path": {
		                      "source": "{{ .Service.Dir }}/docs/public/*",
		                      "target": "docs-snapshot-local/generic-api/{{ .Tag }}/"
		                    }
		                  }
		                }
		              },
		              {
		                "docker": {
		                  "image": "example.com/test/image2"
		                },
		                "push": false
		              }
		            ],
		            "x-examplesDescriptions": [
		              "Each artifact definition contains a single property defining names of executors (builder and pusher) used to handle it. Format of the configuration within is determined by a schema attached to Builder definition. If there is a matching Pusher, configuration must conform to its schema as well.",
		              "If you want to use Pusher and Builder with different names or different configuration formats, add \"push\" property with a separate pusher definition.",
		              "If you don't want to push artifact, set \"push\" property to false."
		            ],
		            "oneOf": [
		              {
		                "oneOf": [
		                  {
		                    "type": "object",
		                    "minProperties": 1,
		                    "maxProperties": 1,
		                    "additionalProperties": true
		                  },
		                  {
		                    "type": "string"
		                  }
		                ]
		              },
		              {
		                "type": "object",
		                "minProperties": 2,
		                "maxProperties": 2,
		                "required": [
		                  "push"
		                ],
		                "properties": {
		                  "push": {
		                    "tsType": "false | Record<string, unknown>",
		                    "oneOf": [
		                      {
		                        "oneOf": [
		                          {
		                            "type": "object",
		                            "minProperties": 1,
		                            "maxProperties": 1,
		                            "additionalProperties": true
		                          },
		                          {
		                            "type": "string"
		                          }
		                        ]
		                      },
		                      {
		                        "const": false
		                      }
		                    ]
		                  }
		                },
		                "additionalProperties": true
		              }
		            ]
		          }
		        },
		        "tags": {
		          "description": "Describes how to generate tags used when pushing artifacts to registry.\n",
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ],
		            "examples": [
		              "gitSha",
		              "gitTag"
		            ]
		          }
		        },
		        "releases": {
		          "description": "List of releases to do by deploy command.\n",
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ],
		            "examples": [
		              {
		                "helm": {
		                  "name": "redis",
		                  "chartPath": "bitnami/redis",
		                  "valuesFiles": [
		                    "{{ .Environment.Dir }}/redis.yaml"
		                  ],
		                  "chartRepository": {
		                    "name": "bitnami",
		                    "url": "https://charts.bitnami.com/bitnami"
		                  }
		                }
		              }
		            ]
		          }
		        },
		        "tasks": {
		          "description": "Definitions of the tasks used by commands \"prepare\", \"test\", \"lint\" and \"run\". These tasks may be also specified in the Project definition.",
		          "type": "object",
		          "properties": {
		            "prepare": {
		              "description": "Defines steps required to preapre freshly clonned repository for development, tests or build. This definition is used by \"prepare\" command.\n",
		              "examples": [
		                [
		                  {
		                    "make": {
		                      "target": "prepare"
		                    }
		                  }
		                ]
		              ],
		              "type": "array",
		              "items": {
		                "oneOf": [
		                  {
		                    "type": "object",
		                    "minProperties": 1,
		                    "maxProperties": 1,
		                    "additionalProperties": true
		                  },
		                  {
		                    "type": "string"
		                  }
		                ]
		              }
		            },
		            "test": {
		              "description": "Defines how to run tests. This definition is used by \"test\" command.\n",
		              "examples": [
		                [
		                  {
		                    "script": {
		                      "sh": "go test ./..."
		                    }
		                  }
		                ]
		              ],
		              "type": "array",
		              "items": {
		                "oneOf": [
		                  {
		                    "type": "object",
		                    "minProperties": 1,
		                    "maxProperties": 1,
		                    "additionalProperties": true
		                  },
		                  {
		                    "type": "string"
		                  }
		                ]
		              }
		            },
		            "lint": {
		              "description": "Defines how to lint the code. Ideally should try to fix the issues. This definition is used by \"lint\" command.\n",
		              "examples": [
		                [
		                  "prettier"
		                ]
		              ],
		              "type": "array",
		              "items": {
		                "oneOf": [
		                  {
		                    "type": "object",
		                    "minProperties": 1,
		                    "maxProperties": 1,
		                    "additionalProperties": true
		                  },
		                  {
		                    "type": "string"
		                  }
		                ]
		              }
		            }
		          },
		          "additionalProperties": {
		            "description": "You are not bound to use pre-defined tasks. All tasks (including custom ones) may be run using \"run\" command.\n",
		            "examples": [
		              [
		                {
		                  "runnerName": {
		                    "some": "params"
		                  }
		                }
		              ]
		            ],
		            "tsType": "({ [k: string]: unknown; } | string )[] | undefined",
		            "type": "array",
		            "items": {
		              "oneOf": [
		                {
		                  "type": "object",
		                  "minProperties": 1,
		                  "maxProperties": 1,
		                  "additionalProperties": true
		                },
		                {
		                  "type": "string"
		                }
		              ]
		            }
		          },
		          "$defs": {
		            "task": {
		              "type": "array",
		              "items": {
		                "oneOf": [
		                  {
		                    "type": "object",
		                    "minProperties": 1,
		                    "maxProperties": 1,
		                    "additionalProperties": true
		                  },
		                  {
		                    "type": "string"
		                  }
		                ]
		              }
		            }
		          }
		        }
		      }
		    },
		    {
		      "title": "Tagger",
		      "type": "object",
		      "required": [
		        "apiVersion",
		        "kind",
		        "name",
		        "script"
		      ],
		      "additionalProperties": false,
		      "properties": {
		        "apiVersion": {
		          "description": "Version of the configuration format.",
		          "const": "g2a-cli/v2.0"
		        },
		        "kind": {
		          "const": "Tagger"
		        },
		        "name": {
		          "description": "Name of the object, unique within the kind.",
		          "type": "string",
		          "minLength": 1,
		          "pattern": "^[a-z][A-Za-z0-9_-]*$"
		        },
		        "schema": {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/schema",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/core": true,
		            "https://json-schema.org/draft/2019-09/vocab/applicator": true,
		            "https://json-schema.org/draft/2019-09/vocab/validation": true,
		            "https://json-schema.org/draft/2019-09/vocab/meta-data": true,
		            "https://json-schema.org/draft/2019-09/vocab/format": false,
		            "https://json-schema.org/draft/2019-09/vocab/content": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Core and Validation specifications meta-schema",
		          "allOf": [
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/core",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/core": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Core vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "$id": {
		                  "type": "string",
		                  "format": "uri-reference",
		                  "$comment": "Non-empty fragments not allowed.",
		                  "pattern": "^[^#]*#?$"
		                },
		                "$schema": {
		                  "type": "string",
		                  "format": "uri"
		                },
		                "$anchor": {
		                  "type": "string",
		                  "pattern": "^[A-Za-z][-A-Za-z0-9.:_]*$"
		                },
		                "$ref": {
		                  "type": "string",
		                  "format": "uri-reference"
		                },
		                "$recursiveRef": {
		                  "type": "string",
		                  "format": "uri-reference"
		                },
		                "$recursiveAnchor": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "$vocabulary": {
		                  "type": "object",
		                  "propertyNames": {
		                    "type": "string",
		                    "format": "uri"
		                  },
		                  "additionalProperties": {
		                    "type": "boolean"
		                  }
		                },
		                "$comment": {
		                  "type": "string"
		                },
		                "$defs": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "default": {}
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/applicator",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/applicator": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Applicator vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "additionalItems": {
		                  "$recursiveRef": "#"
		                },
		                "unevaluatedItems": {
		                  "$recursiveRef": "#"
		                },
		                "items": {
		                  "anyOf": [
		                    {
		                      "$recursiveRef": "#"
		                    },
		                    {
		                      "type": "array",
		                      "minItems": 1,
		                      "items": {
		                        "$recursiveRef": "#"
		                      }
		                    }
		                  ]
		                },
		                "contains": {
		                  "$recursiveRef": "#"
		                },
		                "additionalProperties": {
		                  "$recursiveRef": "#"
		                },
		                "unevaluatedProperties": {
		                  "$recursiveRef": "#"
		                },
		                "properties": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "default": {}
		                },
		                "patternProperties": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  },
		                  "propertyNames": {
		                    "format": "regex"
		                  },
		                  "default": {}
		                },
		                "dependentSchemas": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "propertyNames": {
		                  "$recursiveRef": "#"
		                },
		                "if": {
		                  "$recursiveRef": "#"
		                },
		                "then": {
		                  "$recursiveRef": "#"
		                },
		                "else": {
		                  "$recursiveRef": "#"
		                },
		                "allOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "anyOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "oneOf": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                },
		                "not": {
		                  "$recursiveRef": "#"
		                }
		              },
		              "$defs": {
		                "schemaArray": {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/validation",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/validation": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Validation vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "multipleOf": {
		                  "type": "number",
		                  "exclusiveMinimum": 0
		                },
		                "maximum": {
		                  "type": "number"
		                },
		                "exclusiveMaximum": {
		                  "type": "number"
		                },
		                "minimum": {
		                  "type": "number"
		                },
		                "exclusiveMinimum": {
		                  "type": "number"
		                },
		                "maxLength": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minLength": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "pattern": {
		                  "type": "string",
		                  "format": "regex"
		                },
		                "maxItems": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minItems": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "uniqueItems": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "maxContains": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minContains": {
		                  "type": "integer",
		                  "minimum": 0,
		                  "default": 1
		                },
		                "maxProperties": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "minProperties": {
		                  "default": 0,
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "required": {
		                  "type": "array",
		                  "items": {
		                    "type": "string"
		                  },
		                  "uniqueItems": true,
		                  "default": []
		                },
		                "dependentRequired": {
		                  "type": "object",
		                  "additionalProperties": {
		                    "type": "array",
		                    "items": {
		                      "type": "string"
		                    },
		                    "uniqueItems": true,
		                    "default": []
		                  }
		                },
		                "const": true,
		                "enum": {
		                  "type": "array",
		                  "items": true
		                },
		                "type": {
		                  "anyOf": [
		                    {
		                      "enum": [
		                        "array",
		                        "boolean",
		                        "integer",
		                        "null",
		                        "number",
		                        "object",
		                        "string"
		                      ]
		                    },
		                    {
		                      "type": "array",
		                      "items": {
		                        "enum": [
		                          "array",
		                          "boolean",
		                          "integer",
		                          "null",
		                          "number",
		                          "object",
		                          "string"
		                        ]
		                      },
		                      "minItems": 1,
		                      "uniqueItems": true
		                    }
		                  ]
		                }
		              },
		              "$defs": {
		                "nonNegativeInteger": {
		                  "type": "integer",
		                  "minimum": 0
		                },
		                "nonNegativeIntegerDefault0": {
		                  "type": "integer",
		                  "minimum": 0,
		                  "default": 0
		                },
		                "simpleTypes": {
		                  "enum": [
		                    "array",
		                    "boolean",
		                    "integer",
		                    "null",
		                    "number",
		                    "object",
		                    "string"
		                  ]
		                },
		                "stringArray": {
		                  "type": "array",
		                  "items": {
		                    "type": "string"
		                  },
		                  "uniqueItems": true,
		                  "default": []
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/meta-data",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/meta-data": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Meta-data vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "title": {
		                  "type": "string"
		                },
		                "description": {
		                  "type": "string"
		                },
		                "default": true,
		                "deprecated": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "readOnly": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "writeOnly": {
		                  "type": "boolean",
		                  "default": false
		                },
		                "examples": {
		                  "type": "array",
		                  "items": true
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/format",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/format": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Format vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "format": {
		                  "type": "string"
		                }
		              }
		            },
		            {
		              "$schema": "https://json-schema.org/draft/2019-09/schema",
		              "$id": "https://json-schema.org/draft/2019-09/meta/content",
		              "$vocabulary": {
		                "https://json-schema.org/draft/2019-09/vocab/content": true
		              },
		              "$recursiveAnchor": true,
		              "title": "Content vocabulary meta-schema",
		              "type": [
		                "object",
		                "boolean"
		              ],
		              "properties": {
		                "contentMediaType": {
		                  "type": "string"
		                },
		                "contentEncoding": {
		                  "type": "string"
		                },
		                "contentSchema": {
		                  "$recursiveRef": "#"
		                }
		              }
		            }
		          ],
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "definitions": {
		              "$comment": "While no longer an official keyword as it is replaced by $defs, this keyword is retained in the meta-schema to prevent incompatible extensions as it remains in common use.",
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            },
		            "dependencies": {
		              "$comment": "\"dependencies\" is no longer a keyword, but schema authors should avoid redefining it to facilitate a smooth transition to \"dependentSchemas\" and \"dependentRequired\"",
		              "type": "object",
		              "additionalProperties": {
		                "anyOf": [
		                  {
		                    "$recursiveRef": "#"
		                  },
		                  {
		                    "type": "array",
		                    "items": {
		                      "type": "string"
		                    },
		                    "uniqueItems": true,
		                    "default": []
		                  }
		                ]
		              }
		            }
		          }
		        },
		        "script": {
		          "type": "string"
		        }
		      }
		    }
		  ]
		}
	`),
	"g2a-cli/v2.0/Project": []byte(`
		{
		  "title": "Project",
		  "description": null,
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind",
		    "name"
		  ],
		  "properties": {
		    "apiVersion": {
		      "description": "Version of the configuration format.",
		      "const": "g2a-cli/v2.0"
		    },
		    "kind": {
		      "description": "Determines type of the document.",
		      "const": "Project"
		    },
		    "name": {
		      "examples": [
		        "generic-api"
		      ],
		      "description": "Name of the object, unique within the kind.",
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "files": {
		      "description": "List of the configuration files to load.",
		      "oneOf": [
		        {
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		                "examples": [
		                  "services/*/service.yaml",
		                  "environments/*/environments.yaml"
		                ],
		                "type": "string",
		                "minLength": 1
		              },
		              {
		                "type": "object",
		                "additionalProperties": false,
		                "required": [
		                  "git"
		                ],
		                "properties": {
		                  "git": {
		                    "description": "Files may be also fetched from another git repository.\n",
		                    "type": "object",
		                    "additionalProperties": false,
		                    "required": [
		                      "url",
		                      "rev",
		                      "files"
		                    ],
		                    "properties": {
		                      "url": {
		                        "examples": [
		                          "https://github.com/g2a-com/cicd"
		                        ],
		                        "type": "string",
		                        "minLength": 1
		                      },
		                      "rev": {
		                        "type": "string",
		                        "examples": [
		                          "main"
		                        ],
		                        "minLength": 1
		                      },
		                      "files": {
		                        "examples": [
		                          "executors/*/*.yaml"
		                        ],
		                        "oneOf": [
		                          {
		                            "type": "array",
		                            "items": {
		                              "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		                              "examples": [
		                                "services/*/service.yaml",
		                                "environments/*/environments.yaml"
		                              ],
		                              "type": "string",
		                              "minLength": 1
		                            }
		                          },
		                          {
		                            "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		                            "examples": [
		                              "services/*/service.yaml",
		                              "environments/*/environments.yaml"
		                            ],
		                            "type": "string",
		                            "minLength": 1
		                          }
		                        ]
		                      }
		                    }
		                  }
		                }
		              }
		            ]
		          }
		        },
		        {
		          "description": "Paths to files may include wildecards like \"*\" which matches single path segment.\n",
		          "examples": [
		            "services/*/service.yaml",
		            "environments/*/environments.yaml"
		          ],
		          "type": "string",
		          "minLength": 1
		        }
		      ]
		    },
		    "variables": {
		      "description": "Definitions of the variables to use in the configuration files. Names are case-insensitive.\n",
		      "examples": [
		        {
		          "name": "value"
		        }
		      ],
		      "type": "object",
		      "patternProperties": {
		        "^[a-zA-Z][a-zA-Z0-9]*$": {
		          "type": "string"
		        }
		      }
		    },
		    "tasks": {
		      "description": "Definitions of the tasks used by commands \"prepare\", \"test\", \"lint\" and \"run\". These tasks may be also specified in services definitions.",
		      "type": "object",
		      "properties": {
		        "prepare": {
		          "description": "Defines steps required to preapre freshly clonned repository for development, tests or build. This definition is used by \"prepare\" command.\n",
		          "examples": [
		            [
		              {
		                "make": {
		                  "target": "prepare"
		                }
		              }
		            ]
		          ],
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ]
		          }
		        },
		        "test": {
		          "description": "Defines how to run tests. This definition is used by \"test\" command.\n",
		          "examples": [
		            [
		              {
		                "script": {
		                  "sh": "go test ./..."
		                }
		              }
		            ]
		          ],
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ]
		          }
		        },
		        "lint": {
		          "description": "Defines how to lint the code. Ideally should try to fix the issues. This definition is used by \"lint\" command.\n",
		          "examples": [
		            [
		              "prettier"
		            ]
		          ],
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ]
		          }
		        }
		      },
		      "additionalProperties": {
		        "description": "You are not bound to use pre-defined tasks. All tasks (including custom ones) may be run using \"run\" command.\n",
		        "examples": [
		          [
		            {
		              "runnerName": {
		                "some": "params"
		              }
		            }
		          ]
		        ],
		        "tsType": "({ [k: string]: unknown; } | string )[] | undefined",
		        "type": "array",
		        "items": {
		          "oneOf": [
		            {
		              "type": "object",
		              "minProperties": 1,
		              "maxProperties": 1,
		              "additionalProperties": true
		            },
		            {
		              "type": "string"
		            }
		          ]
		        }
		      },
		      "$defs": {
		        "task": {
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ]
		          }
		        }
		      }
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Pusher": []byte(`
		{
		  "title": "Pusher",
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind",
		    "name",
		    "script"
		  ],
		  "additionalProperties": false,
		  "properties": {
		    "apiVersion": {
		      "description": "Version of the configuration format.",
		      "const": "g2a-cli/v2.0"
		    },
		    "kind": {
		      "const": "Pusher"
		    },
		    "name": {
		      "description": "Name of the object, unique within the kind.",
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "schema": {
		      "$schema": "https://json-schema.org/draft/2019-09/schema",
		      "$id": "https://json-schema.org/draft/2019-09/schema",
		      "$vocabulary": {
		        "https://json-schema.org/draft/2019-09/vocab/core": true,
		        "https://json-schema.org/draft/2019-09/vocab/applicator": true,
		        "https://json-schema.org/draft/2019-09/vocab/validation": true,
		        "https://json-schema.org/draft/2019-09/vocab/meta-data": true,
		        "https://json-schema.org/draft/2019-09/vocab/format": false,
		        "https://json-schema.org/draft/2019-09/vocab/content": true
		      },
		      "$recursiveAnchor": true,
		      "title": "Core and Validation specifications meta-schema",
		      "allOf": [
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/core",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/core": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Core vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "$id": {
		              "type": "string",
		              "format": "uri-reference",
		              "$comment": "Non-empty fragments not allowed.",
		              "pattern": "^[^#]*#?$"
		            },
		            "$schema": {
		              "type": "string",
		              "format": "uri"
		            },
		            "$anchor": {
		              "type": "string",
		              "pattern": "^[A-Za-z][-A-Za-z0-9.:_]*$"
		            },
		            "$ref": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveRef": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveAnchor": {
		              "type": "boolean",
		              "default": false
		            },
		            "$vocabulary": {
		              "type": "object",
		              "propertyNames": {
		                "type": "string",
		                "format": "uri"
		              },
		              "additionalProperties": {
		                "type": "boolean"
		              }
		            },
		            "$comment": {
		              "type": "string"
		            },
		            "$defs": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/applicator",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/applicator": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Applicator vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "additionalItems": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedItems": {
		              "$recursiveRef": "#"
		            },
		            "items": {
		              "anyOf": [
		                {
		                  "$recursiveRef": "#"
		                },
		                {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                }
		              ]
		            },
		            "contains": {
		              "$recursiveRef": "#"
		            },
		            "additionalProperties": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedProperties": {
		              "$recursiveRef": "#"
		            },
		            "properties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            },
		            "patternProperties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "propertyNames": {
		                "format": "regex"
		              },
		              "default": {}
		            },
		            "dependentSchemas": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              }
		            },
		            "propertyNames": {
		              "$recursiveRef": "#"
		            },
		            "if": {
		              "$recursiveRef": "#"
		            },
		            "then": {
		              "$recursiveRef": "#"
		            },
		            "else": {
		              "$recursiveRef": "#"
		            },
		            "allOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "anyOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "oneOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "not": {
		              "$recursiveRef": "#"
		            }
		          },
		          "$defs": {
		            "schemaArray": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/validation",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/validation": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Validation vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "multipleOf": {
		              "type": "number",
		              "exclusiveMinimum": 0
		            },
		            "maximum": {
		              "type": "number"
		            },
		            "exclusiveMaximum": {
		              "type": "number"
		            },
		            "minimum": {
		              "type": "number"
		            },
		            "exclusiveMinimum": {
		              "type": "number"
		            },
		            "maxLength": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minLength": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "pattern": {
		              "type": "string",
		              "format": "regex"
		            },
		            "maxItems": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minItems": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "uniqueItems": {
		              "type": "boolean",
		              "default": false
		            },
		            "maxContains": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minContains": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 1
		            },
		            "maxProperties": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minProperties": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "required": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            },
		            "dependentRequired": {
		              "type": "object",
		              "additionalProperties": {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            },
		            "const": true,
		            "enum": {
		              "type": "array",
		              "items": true
		            },
		            "type": {
		              "anyOf": [
		                {
		                  "enum": [
		                    "array",
		                    "boolean",
		                    "integer",
		                    "null",
		                    "number",
		                    "object",
		                    "string"
		                  ]
		                },
		                {
		                  "type": "array",
		                  "items": {
		                    "enum": [
		                      "array",
		                      "boolean",
		                      "integer",
		                      "null",
		                      "number",
		                      "object",
		                      "string"
		                    ]
		                  },
		                  "minItems": 1,
		                  "uniqueItems": true
		                }
		              ]
		            }
		          },
		          "$defs": {
		            "nonNegativeInteger": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "nonNegativeIntegerDefault0": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 0
		            },
		            "simpleTypes": {
		              "enum": [
		                "array",
		                "boolean",
		                "integer",
		                "null",
		                "number",
		                "object",
		                "string"
		              ]
		            },
		            "stringArray": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/meta-data",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/meta-data": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Meta-data vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "title": {
		              "type": "string"
		            },
		            "description": {
		              "type": "string"
		            },
		            "default": true,
		            "deprecated": {
		              "type": "boolean",
		              "default": false
		            },
		            "readOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "writeOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "examples": {
		              "type": "array",
		              "items": true
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/format",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/format": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Format vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "format": {
		              "type": "string"
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/content",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/content": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Content vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "contentMediaType": {
		              "type": "string"
		            },
		            "contentEncoding": {
		              "type": "string"
		            },
		            "contentSchema": {
		              "$recursiveRef": "#"
		            }
		          }
		        }
		      ],
		      "type": [
		        "object",
		        "boolean"
		      ],
		      "properties": {
		        "definitions": {
		          "$comment": "While no longer an official keyword as it is replaced by $defs, this keyword is retained in the meta-schema to prevent incompatible extensions as it remains in common use.",
		          "type": "object",
		          "additionalProperties": {
		            "$recursiveRef": "#"
		          },
		          "default": {}
		        },
		        "dependencies": {
		          "$comment": "\"dependencies\" is no longer a keyword, but schema authors should avoid redefining it to facilitate a smooth transition to \"dependentSchemas\" and \"dependentRequired\"",
		          "type": "object",
		          "additionalProperties": {
		            "anyOf": [
		              {
		                "$recursiveRef": "#"
		              },
		              {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            ]
		          }
		        }
		      }
		    },
		    "script": {
		      "type": "string"
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Service": []byte(`
		{
		  "title": "Service",
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind",
		    "name"
		  ],
		  "properties": {
		    "apiVersion": {
		      "description": "Version of the configuration format.",
		      "const": "g2a-cli/v2.0"
		    },
		    "kind": {
		      "description": "Determines type of the document.",
		      "const": "Service"
		    },
		    "name": {
		      "description": "Unique name used to identify service.",
		      "examples": [
		        "example-api"
		      ],
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "artifacts": {
		      "description": "List of artifacts to produce by build command. Each entry describes single artifact like docker image or npm package.\n",
		      "type": "array",
		      "items": {
		        "examples": [
		          {
		            "docker": {
		              "image": "example.com/test/image"
		            }
		          },
		          {
		            "hugo": {
		              "dir": "{{ .Service.Dir }}/docs"
		            },
		            "push": {
		              "artifactory": {
		                "path": {
		                  "source": "{{ .Service.Dir }}/docs/public/*",
		                  "target": "docs-snapshot-local/generic-api/{{ .Tag }}/"
		                }
		              }
		            }
		          },
		          {
		            "docker": {
		              "image": "example.com/test/image2"
		            },
		            "push": false
		          }
		        ],
		        "x-examplesDescriptions": [
		          "Each artifact definition contains a single property defining names of executors (builder and pusher) used to handle it. Format of the configuration within is determined by a schema attached to Builder definition. If there is a matching Pusher, configuration must conform to its schema as well.",
		          "If you want to use Pusher and Builder with different names or different configuration formats, add \"push\" property with a separate pusher definition.",
		          "If you don't want to push artifact, set \"push\" property to false."
		        ],
		        "oneOf": [
		          {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ]
		          },
		          {
		            "type": "object",
		            "minProperties": 2,
		            "maxProperties": 2,
		            "required": [
		              "push"
		            ],
		            "properties": {
		              "push": {
		                "tsType": "false | Record<string, unknown>",
		                "oneOf": [
		                  {
		                    "oneOf": [
		                      {
		                        "type": "object",
		                        "minProperties": 1,
		                        "maxProperties": 1,
		                        "additionalProperties": true
		                      },
		                      {
		                        "type": "string"
		                      }
		                    ]
		                  },
		                  {
		                    "const": false
		                  }
		                ]
		              }
		            },
		            "additionalProperties": true
		          }
		        ]
		      }
		    },
		    "tags": {
		      "description": "Describes how to generate tags used when pushing artifacts to registry.\n",
		      "type": "array",
		      "items": {
		        "oneOf": [
		          {
		            "type": "object",
		            "minProperties": 1,
		            "maxProperties": 1,
		            "additionalProperties": true
		          },
		          {
		            "type": "string"
		          }
		        ],
		        "examples": [
		          "gitSha",
		          "gitTag"
		        ]
		      }
		    },
		    "releases": {
		      "description": "List of releases to do by deploy command.\n",
		      "type": "array",
		      "items": {
		        "oneOf": [
		          {
		            "type": "object",
		            "minProperties": 1,
		            "maxProperties": 1,
		            "additionalProperties": true
		          },
		          {
		            "type": "string"
		          }
		        ],
		        "examples": [
		          {
		            "helm": {
		              "name": "redis",
		              "chartPath": "bitnami/redis",
		              "valuesFiles": [
		                "{{ .Environment.Dir }}/redis.yaml"
		              ],
		              "chartRepository": {
		                "name": "bitnami",
		                "url": "https://charts.bitnami.com/bitnami"
		              }
		            }
		          }
		        ]
		      }
		    },
		    "tasks": {
		      "description": "Definitions of the tasks used by commands \"prepare\", \"test\", \"lint\" and \"run\". These tasks may be also specified in the Project definition.",
		      "type": "object",
		      "properties": {
		        "prepare": {
		          "description": "Defines steps required to preapre freshly clonned repository for development, tests or build. This definition is used by \"prepare\" command.\n",
		          "examples": [
		            [
		              {
		                "make": {
		                  "target": "prepare"
		                }
		              }
		            ]
		          ],
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ]
		          }
		        },
		        "test": {
		          "description": "Defines how to run tests. This definition is used by \"test\" command.\n",
		          "examples": [
		            [
		              {
		                "script": {
		                  "sh": "go test ./..."
		                }
		              }
		            ]
		          ],
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ]
		          }
		        },
		        "lint": {
		          "description": "Defines how to lint the code. Ideally should try to fix the issues. This definition is used by \"lint\" command.\n",
		          "examples": [
		            [
		              "prettier"
		            ]
		          ],
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ]
		          }
		        }
		      },
		      "additionalProperties": {
		        "description": "You are not bound to use pre-defined tasks. All tasks (including custom ones) may be run using \"run\" command.\n",
		        "examples": [
		          [
		            {
		              "runnerName": {
		                "some": "params"
		              }
		            }
		          ]
		        ],
		        "tsType": "({ [k: string]: unknown; } | string )[] | undefined",
		        "type": "array",
		        "items": {
		          "oneOf": [
		            {
		              "type": "object",
		              "minProperties": 1,
		              "maxProperties": 1,
		              "additionalProperties": true
		            },
		            {
		              "type": "string"
		            }
		          ]
		        }
		      },
		      "$defs": {
		        "task": {
		          "type": "array",
		          "items": {
		            "oneOf": [
		              {
		                "type": "object",
		                "minProperties": 1,
		                "maxProperties": 1,
		                "additionalProperties": true
		              },
		              {
		                "type": "string"
		              }
		            ]
		          }
		        }
		      }
		    }
		  }
		}
	`),
	"g2a-cli/v2.0/Tagger": []byte(`
		{
		  "title": "Tagger",
		  "type": "object",
		  "required": [
		    "apiVersion",
		    "kind",
		    "name",
		    "script"
		  ],
		  "additionalProperties": false,
		  "properties": {
		    "apiVersion": {
		      "description": "Version of the configuration format.",
		      "const": "g2a-cli/v2.0"
		    },
		    "kind": {
		      "const": "Tagger"
		    },
		    "name": {
		      "description": "Name of the object, unique within the kind.",
		      "type": "string",
		      "minLength": 1,
		      "pattern": "^[a-z][A-Za-z0-9_-]*$"
		    },
		    "schema": {
		      "$schema": "https://json-schema.org/draft/2019-09/schema",
		      "$id": "https://json-schema.org/draft/2019-09/schema",
		      "$vocabulary": {
		        "https://json-schema.org/draft/2019-09/vocab/core": true,
		        "https://json-schema.org/draft/2019-09/vocab/applicator": true,
		        "https://json-schema.org/draft/2019-09/vocab/validation": true,
		        "https://json-schema.org/draft/2019-09/vocab/meta-data": true,
		        "https://json-schema.org/draft/2019-09/vocab/format": false,
		        "https://json-schema.org/draft/2019-09/vocab/content": true
		      },
		      "$recursiveAnchor": true,
		      "title": "Core and Validation specifications meta-schema",
		      "allOf": [
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/core",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/core": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Core vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "$id": {
		              "type": "string",
		              "format": "uri-reference",
		              "$comment": "Non-empty fragments not allowed.",
		              "pattern": "^[^#]*#?$"
		            },
		            "$schema": {
		              "type": "string",
		              "format": "uri"
		            },
		            "$anchor": {
		              "type": "string",
		              "pattern": "^[A-Za-z][-A-Za-z0-9.:_]*$"
		            },
		            "$ref": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveRef": {
		              "type": "string",
		              "format": "uri-reference"
		            },
		            "$recursiveAnchor": {
		              "type": "boolean",
		              "default": false
		            },
		            "$vocabulary": {
		              "type": "object",
		              "propertyNames": {
		                "type": "string",
		                "format": "uri"
		              },
		              "additionalProperties": {
		                "type": "boolean"
		              }
		            },
		            "$comment": {
		              "type": "string"
		            },
		            "$defs": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/applicator",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/applicator": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Applicator vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "additionalItems": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedItems": {
		              "$recursiveRef": "#"
		            },
		            "items": {
		              "anyOf": [
		                {
		                  "$recursiveRef": "#"
		                },
		                {
		                  "type": "array",
		                  "minItems": 1,
		                  "items": {
		                    "$recursiveRef": "#"
		                  }
		                }
		              ]
		            },
		            "contains": {
		              "$recursiveRef": "#"
		            },
		            "additionalProperties": {
		              "$recursiveRef": "#"
		            },
		            "unevaluatedProperties": {
		              "$recursiveRef": "#"
		            },
		            "properties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "default": {}
		            },
		            "patternProperties": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              },
		              "propertyNames": {
		                "format": "regex"
		              },
		              "default": {}
		            },
		            "dependentSchemas": {
		              "type": "object",
		              "additionalProperties": {
		                "$recursiveRef": "#"
		              }
		            },
		            "propertyNames": {
		              "$recursiveRef": "#"
		            },
		            "if": {
		              "$recursiveRef": "#"
		            },
		            "then": {
		              "$recursiveRef": "#"
		            },
		            "else": {
		              "$recursiveRef": "#"
		            },
		            "allOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "anyOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "oneOf": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            },
		            "not": {
		              "$recursiveRef": "#"
		            }
		          },
		          "$defs": {
		            "schemaArray": {
		              "type": "array",
		              "minItems": 1,
		              "items": {
		                "$recursiveRef": "#"
		              }
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/validation",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/validation": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Validation vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "multipleOf": {
		              "type": "number",
		              "exclusiveMinimum": 0
		            },
		            "maximum": {
		              "type": "number"
		            },
		            "exclusiveMaximum": {
		              "type": "number"
		            },
		            "minimum": {
		              "type": "number"
		            },
		            "exclusiveMinimum": {
		              "type": "number"
		            },
		            "maxLength": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minLength": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "pattern": {
		              "type": "string",
		              "format": "regex"
		            },
		            "maxItems": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minItems": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "uniqueItems": {
		              "type": "boolean",
		              "default": false
		            },
		            "maxContains": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minContains": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 1
		            },
		            "maxProperties": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "minProperties": {
		              "default": 0,
		              "type": "integer",
		              "minimum": 0
		            },
		            "required": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            },
		            "dependentRequired": {
		              "type": "object",
		              "additionalProperties": {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            },
		            "const": true,
		            "enum": {
		              "type": "array",
		              "items": true
		            },
		            "type": {
		              "anyOf": [
		                {
		                  "enum": [
		                    "array",
		                    "boolean",
		                    "integer",
		                    "null",
		                    "number",
		                    "object",
		                    "string"
		                  ]
		                },
		                {
		                  "type": "array",
		                  "items": {
		                    "enum": [
		                      "array",
		                      "boolean",
		                      "integer",
		                      "null",
		                      "number",
		                      "object",
		                      "string"
		                    ]
		                  },
		                  "minItems": 1,
		                  "uniqueItems": true
		                }
		              ]
		            }
		          },
		          "$defs": {
		            "nonNegativeInteger": {
		              "type": "integer",
		              "minimum": 0
		            },
		            "nonNegativeIntegerDefault0": {
		              "type": "integer",
		              "minimum": 0,
		              "default": 0
		            },
		            "simpleTypes": {
		              "enum": [
		                "array",
		                "boolean",
		                "integer",
		                "null",
		                "number",
		                "object",
		                "string"
		              ]
		            },
		            "stringArray": {
		              "type": "array",
		              "items": {
		                "type": "string"
		              },
		              "uniqueItems": true,
		              "default": []
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/meta-data",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/meta-data": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Meta-data vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "title": {
		              "type": "string"
		            },
		            "description": {
		              "type": "string"
		            },
		            "default": true,
		            "deprecated": {
		              "type": "boolean",
		              "default": false
		            },
		            "readOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "writeOnly": {
		              "type": "boolean",
		              "default": false
		            },
		            "examples": {
		              "type": "array",
		              "items": true
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/format",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/format": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Format vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "format": {
		              "type": "string"
		            }
		          }
		        },
		        {
		          "$schema": "https://json-schema.org/draft/2019-09/schema",
		          "$id": "https://json-schema.org/draft/2019-09/meta/content",
		          "$vocabulary": {
		            "https://json-schema.org/draft/2019-09/vocab/content": true
		          },
		          "$recursiveAnchor": true,
		          "title": "Content vocabulary meta-schema",
		          "type": [
		            "object",
		            "boolean"
		          ],
		          "properties": {
		            "contentMediaType": {
		              "type": "string"
		            },
		            "contentEncoding": {
		              "type": "string"
		            },
		            "contentSchema": {
		              "$recursiveRef": "#"
		            }
		          }
		        }
		      ],
		      "type": [
		        "object",
		        "boolean"
		      ],
		      "properties": {
		        "definitions": {
		          "$comment": "While no longer an official keyword as it is replaced by $defs, this keyword is retained in the meta-schema to prevent incompatible extensions as it remains in common use.",
		          "type": "object",
		          "additionalProperties": {
		            "$recursiveRef": "#"
		          },
		          "default": {}
		        },
		        "dependencies": {
		          "$comment": "\"dependencies\" is no longer a keyword, but schema authors should avoid redefining it to facilitate a smooth transition to \"dependentSchemas\" and \"dependentRequired\"",
		          "type": "object",
		          "additionalProperties": {
		            "anyOf": [
		              {
		                "$recursiveRef": "#"
		              },
		              {
		                "type": "array",
		                "items": {
		                  "type": "string"
		                },
		                "uniqueItems": true,
		                "default": []
		              }
		            ]
		          }
		        }
		      }
		    },
		    "script": {
		      "type": "string"
		    }
		  }
		}
	`),
}
