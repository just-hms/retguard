package example

import "fmt"

// should fail val is not assigned
func SimplePositiveCase(a int) (val int) {
	return
}

// should fail val is not assigned
func DuplPositiveCase(a int) (val, err int) {
	err = 2
	return
}

// should fail val is not assigned
func ShadowingPositiveCase(a int) (val, err int) {
	{
		val := 2
		fmt.Println(val)
	}
	err = 2
	return
}

// should fail val is not assigned
func VarShadowingPositiveCase(a int) (val, err int) {
	{
		var val = 2
		fmt.Println(val)
	}
	err = 2
	return
}

// should fail val is not assigned
func MultiVarShadowingPositiveCase(a int) (val, err int) {
	{
		var (
			val  = 2
			val2 = "err"
		)
		fmt.Println(val, val2)
	}
	err = 2
	return
}

// should fail val is not assigned
func VarShadowingWithAssignmentPositiveCase(a int) (val, err int) {
	{
		var val = 2
		val = 3
		fmt.Println(val)
	}
	err = 2
	return
}

// helpNotModifier couldn't modify the val in any case
func FunctionPositiveCase(a int) (val, err int) {
	fmt.Println(val)
	err = 2
	return
}
