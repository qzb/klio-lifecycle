---
title: Placeholders & Variables
menuTitle: Placeholders
weight: 50
---

In some places in the configuration files, you can use `{{ }}` placeholders. They can be used only
in strings, furthermore some properties (like "name") doesn't allow using placeholders at all.

All placeholder names are case-insensitive.

| Placeholder                 | Description                                              | Restrictions                         |
| --------------------------- | -------------------------------------------------------- | ------------------------------------ |
| `{{ .Environment.Dir }}`    | Directory of the file containing environment definition. | Environment, Service (only releases) |
| `{{ .Environment.Name }}`   | Name of the environment.                                 | Environment, Service (only releases) |
| `{{ .Environment.Vars.* }}` | Variables defined in the environment.                    | Environment, Service (only releases) |
| `{{ .Service.Dir }}`        | Directory of the file containing a service definition.   | Service                              |
| `{{ .Service.Name }}`       | Name of the service.                                     | Service                              |
| `{{ .Project.Dir }}`        | Directory of the file containing project definition.     |                                      |
| `{{ .Project.Name }}`       | Name of the project.                                     |                                      |
| `{{ .Project.Vars.* }}`     | Variables defined in the project                         |                                      |
| `{{ .Params.* }}`           | Params specified using `--param` command-line option.    |                                      |
| `{{ .Tag }}`                | Tag specified using `--tag` command-line option.         | Environment, Service (only releases) |

## Migration from `g2a-cli/v1beta4`

| g2a-cli/v2.0                | g2a-cli/v1beta4           |
| --------------------------- | ------------------------- |
| `{{ .Environment.Dir }}`    | `{{ .Dirs.Environment }}` |
| `{{ .Environment.Name }}`   | _n/a_                     |
| `{{ .Environment.Vars.* }}` | `{{ .Env.* }}`            |
| `{{ .Service.Dir }}`        | `{{ .Dirs.Service }}`     |
| `{{ .Service.Name }}`       | _n/a_                     |
| `{{ .Project.Dir }}`        | `{{ .Dirs.Project }}`     |
| `{{ .Project.Name }}`       | _n/a_                     |
| `{{ .Project.Vars.* }}`     | _n/a_                     |
| `{{ .Params.* }}`           | `{{ .Params.* }}`         |
| `{{ .Tag }}`                | `{{ .Opts.Tag }}`         |
