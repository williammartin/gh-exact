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
gh exact dump -f extensions.yaml
```

```
gh exact install -f extensions.yaml [--pin]
```

Without `--pin` the latest version of each extension will be installed. With `--pin` the exact version you dumped will be restored, but this will also be **pinned** as if you did `gh ext install --pin`.

## Limitations

Doesn't handle local extension installation.
Doesn't handle installing specific extensions without pinning.
Doesn't really do very good error reporting.
