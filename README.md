
# Opengist-CLI

A small, straightforward command-line client for [Opengist](https://opengist.io).

`opengist-cli` lets you manage your Opengist snippets directly from the terminal, without needing to open a browser or use the web interface. Create, edit, delete, and list gists using simple commands, making it useful for quick note-taking, sharing code, scripting, and integrating gists into your everyday CLI workflow.

It is designed to be a lightweight, single-binary tool with configuration provided through environment variables so it can easily fit into shell scripts, dotfiles, and automated workflows.

## Features

- **Zero dependencies:** Single, cross-compiled binary.
- **Cross-platform:** Native support for Linux, macOS (Apple Silicon & Intel), and Windows.
- **Shell Completion:** Auto-generated completion for Bash, Zsh, Fish, and PowerShell.

## Configuration

`opengist-cli` requires an Opengist host URL and user access token. Set the following environment variables:

```bash
export OPENGIST_CLI_TOKEN="your_access_token_here"
export OPENGIST_CLI_URL="https://your-opengist-instance"
```

## Usage

```bash
opengist-cli <command> [options]
```

### Commands

* `add`: Create a new gist from one or more files
* `edit`: Edit an existing gist
* `delete`: Delete an existing gist
* `list`: List gists (supports plain text and `--markdown` table format)
* `completion`: Generate autocompletion script for your shell

For command-specific help run:

```bash
opengist-cli <command> --help
```

## Installation

Get a single binary download from the [latest release](https://github.com/ocodo/opengist-cli/releases/latest)

### Build From Source

Requires [Go 1.22+](https://go.dev/).

```bash
git clone git@github.com:ocodo/opengist-cli.git
cd opengist-cli/pkg
make build
cp bin/opengist-cli ~/.local/bin/
```

To cross-compile binaries for all supported platforms:

```bash
make build-all
```

The resulting binaries will be in the `./bin/` directory.

### Shell Completion

Generate shell autocompletion for your preferred shell:

```bash
# Zsh
opengist-cli completion zsh > "${fpath[1]}/_opengist-cli"

# Bash
opengist-cli completion bash > /etc/bash_completion.d/opengist-cli

# Fish
opengist-cli completion fish > ~/.config/fish/completions/opengist-cli.fish
```

## License

AGPL v3


