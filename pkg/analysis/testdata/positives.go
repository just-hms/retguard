package testdata

import "fmt"

func SimplePositiveCase(a int) (val int) {
	return // want "named return value val is never assigned"
}

func DuplPositiveCase(a int) (val, err int) {
	err = 2
	return // want "named return value val is never assigned"
}

func DuplNotAssignedPositiveCase(a int) (val, err int) {
	return // want "named return values val, err are never assigned"
}

func ShadowingPositiveCase(a int) (val, err int) {
	{
		val := 2
		for range val {
			fmt.Println()
		}
	}
	err = 2
	return // want "named return value val is never assigned"
}

func VarShadowingPositiveCase(a int) (val, err int) {
	{
		val := 2
		for range val {
			fmt.Println()
		}
	}
	err = 2
	return // want "named return value val is never assigned"
}

func MultiVarShadowingPositiveCase(a int) (val, err int) {
	{
		var (
			val  = 2
			val2 = "err"
		)
		for range val {
			fmt.Println()
		}
		for range val2 {
			fmt.Println()
		}
	}
	err = 2
	return // want "named return value val is never assigned"
}

func VarShadowingWithAssignmentPositiveCase(a int) (val, err int) {
	{
		var val = 2
		val = 3
		for range val {
			fmt.Println()
		}
	}
	err = 2
	return // want "named return value val is never assigned"
}
