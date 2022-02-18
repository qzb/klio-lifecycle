---
title: Executors Configurations
menuTitle: Executors
weight: 50
---

Since Service defines _what_ to build or deploy, you need an executor to specify _how_ to do it.
There are few kinds of executors:

- Tagger
- Builder
- Pusher
- Deployer

Each of those documents contains a schema for input parameters and short JavaScript code to run when
the corresponding action is executed.

## Example

{{< yaml-table "/schemas/g2a-cli/v2.0/executor.json" >}}
