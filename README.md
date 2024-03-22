# Description

Dump and restore your `gh` CLI extensions.

## Prerequisites

`gh` CLI installed and authenticated (for `install` command)

## Installation

```
gh ext install williammartin/gh-exact
```

## Usage

```
gh exact dump -f manifest.yaml
```

```
gh exact restore -f manifest.yaml [--pin]
```

Without `--pin` the latest version of each extension will be installed. With `--pin` the exact version you dumped will be restored, but this will also be **pinned** as if you did `gh ext install --pin`.

## Limitations

* Doesn't handle local extension installation.
* Doesn't handle installing specific extension versions without pinning.
* Doesn't really do very good error reporting or general communication about what it is doing really
