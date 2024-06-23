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

the linter is pretty new, checkout https://github.com/just-hms/retguard/labels/falsepositive and https://github.com/just-hms/retguard/labels/falsenegative issues before thinking that your code is broken
