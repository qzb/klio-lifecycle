---
title: G2A Command Line Interface
date: 2021-03-02T17:55:28Z
---
# G2A Command Line Interface

The _G2A CLI_ is a tool designed to get you working quickly and efficiently with G2A services, with an emphasis on automation.
It unifies and generalizes all work you need to do to run, build and deploy G2A services.
G2A CLI is intended to be used the same way on a developer's local machine or on CI/CD servers like Jenkins, Bamboo or TeamCity.

## Requirements

_G2A CLI_ core doesn't require any dependencies to work.
It is provided to you as a standalone binary file, which you include in your `PATH`.
However, additional command can depend on various of tools including [docker](https://www.docker.com/), 
[kubectl](https://kubernetes.io/docs/tasks/tools/), [helm](https://helm.sh/) or even some language runtimes like 
[node](https://nodejs.org/en/) or [python](https://www.python.org/).

## Installation

There are 3 ways you can install _G2A CLI_

#### Automatic installation

Run this one-liner to download the latest version of _G2A CLI_. 

{{% tabs %}}
{{% tab "Linux/MacOS"  %}}
```bash
curl -sL https://cli.code.g2a.com/install | sh
```
{{% /tab  %}}
{{% tab "Windows" %}}
```powershell
. { iwr -useb https://cli.code.g2a.com/install.ps1 } | iex
```
{{% /tab  %}}
{{% /tabs %}}

This script prompts you to extend `PATH` variable with a g2a cli home directory.
Do it before the first use of the cli.

#### Manual installation

Download the latest version of _G2A CLI_ from [releases site](http://cli.code.g2a.com/), rename it to `g2a` (or `g2a.exe` on Windows) and put it in directory of your choice e.g.,

{{% tabs %}}
{{% tab "Linux/MacOS"  %}}
```bash
curl -Lo g2a https://cli.code.g2a.com/v2.3.8-darwin-amd64 && chmod +x g2a && sudo mv g2a /usr/local/bin
```
{{% /tabs %}}

Remember to extend your `PATH` variable with the directory.

#### From Source 

You can get and build _G2A CLI_ for yourself. Such an installation requires Go version 1.12 or higher (with support of Go modules).
To download _G2A CLI_ and make `g2a` executable, use `go get` command:

```bash
go get -u stash.code.g2a.com/cli/g2a
```

## Verify installation

After installing `g2a` with one of methods described above, you should verify the version of your installation.
Check your version with following:

```bash
g2a --version
```

{{% notice warning %}}
If version is < 2.0 you have old Python's CLI installed which is considered **DEPRECATED** and should not be used anymore.
Please ensure you are using 2nd iteration of _G2A CLI_.
{{% /notice %}}

## Contribution

G2A CLI is brought to you by Service Delivery Platform. tool used across all technology we invite you to contribution.
All_G2A CLI_repositories are stored in [Stash](https://stash.code.g2a.com/projects/CLI) and publicly are available.
PRs to any cli module are more than welcome.
Change Requests can be reported in Jira [Klio project](https://jira.code.g2a.com/secure/RapidBoard.jspa?projectKey=KLIO&rapidView=1098).
In case of questions regarding CLI, join [CLI User Group](https://teams.microsoft.com/l/channel/19%3ad71d35810e844962bfc7ad89b27578a0%40thread.tacv2/CLI%2520(klio)?groupId=ba9389f8-dd40-41b6-a587-26dec00528a8&tenantId=a78c2a2a-b9d8-453e-a193-967884425b10) available on Microsoft Teams.
