# retguard

`retguard` is a go linter that checks that every named return value has been assigned at least one


## Install

```bash
go install github.com/just-hms/retguard@latest
```

## Usage

```bash
retguard ./...
```

## Issues

the linter is pretty new, checkout https://img.shields.io/github/labels/just-hms/retguard/falsepositive and https://img.shields.io/github/labels/just-hms/retguard/falsenegative issues before thinking that your code is broken
