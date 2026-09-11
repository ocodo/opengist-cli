
# Opengist-CLI


A small, straightforward command-line client for [Opengist](https://opengist.io)

opengist-cli lets you manage your Opengist snippets directly from the terminal, without needing to open a browser or use the web interface. Create, edit, delete, and list gists using simple commands, making it useful for quick note-taking, sharing code, scripting, and integrating gists into your everyday CLI workflow.

It is designed to be lightweight and Unix-friendly, with configuration provided through environment variables so it can easily fit into shell scripts, dotfiles, and automated workflows.

## Install

Requires [uv](https://github.com/astral-sh/uv)

```bash
git clone git@github.com:ocodo/opengist-cli
```

## Usage

copy opengist-cli to your path e.g. `~/.local/bin/`

```bash
opengist-cli <command> [options]
```
Commands:

- `add`:       Create a new gist
- `edit`:      Edit an existing gist
- `delete`:    Delete a gist
- `list`:      List gists

Config:

Requires an opengist host url and access token.  Use env vars

```bash
OPENGIST_CLI_TOKEN # opengist user access token
OPENGIST_CLI_URL # opengist host url
```

for command-specific help run:

```bash
opengist-cli <command> --help
```


### Licence

AGPL v3

![License: GNU AGPL v3](https://img.shields.io/badge/License-agpl3.0-blue.svg)