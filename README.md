![License: GNU AGPL v3](https://img.shields.io/badge/License-agpl3.0-blue.svg)

# Opengist-CLI

CLI for [Opengist](https://opengist.io)

## Install

Requires [uv](https://github.com/astral-sh/uv)

```bash
git clone git@github.com:ocodo/opengist-cli
```
## Usage

```
# copy opengist-cli to your path e.g. `~/.local/bin/`

opengist-cli <command> [options]
```
Commands:

- `add`:       Create a new gist
- `edit`:      Edit an existing gist
- `delete`:    Delete a gist
- `list`:      List gists
	
Config:

- Requires an opengist host url and access token.
- set env vars OPENGIST_CLI_TOKEN OPENGIST_CLI_URL	

Run `opengist-cli <command> --help` for command-specific help.

### Licence

AGPL v3
