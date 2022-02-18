---
title: Environment Configuration
menuTitle: Environment
weight: 40
---

An environment defines a configuration for the deploy command. The document itself contains only to
types of information: list of services to deploy and variables to use in services configuration, but
you may easily add additional configuration files (like values files for Helm) to the same directory
and use `{{ .Environment.Dir }}` placeholder to pick a relevant file.

## Example

{{< yaml-table "/schemas/g2a-cli/v2.0/environment.json" >}}
