# Dev notes

If the current IDE used doesn't recognize PEP723.

I setup a `pyproject.toml` and copy the PEP723 block from `opengist-cli` and add a `.python-version` if tooling needs it.

Running:

```
uv venv
uv sync
```

Will add the dependencies to local `.venv` and install a Python 3.12 interpreter there.

```
. .venv/bin/activate
```

