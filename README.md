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

## False positives

For now this is not considered an assignment, even if val can be modified in the `create` function, see: https://github.com/just-hms/retguard/issues/1

```go
func create(input *int) {
	*input = 0
}

// false positive
func FunctionComplexNegativeCase(a int) (val *int, err int) {
	create(val)
	err = 2
	return
}
```
