package example

import "fmt"

func SimplePositiveCase(a int) (val int) {
	return // want "named return value val is never assigned"
}

func DuplPositiveCase(a int) (val, err int) {
	err = 2
	return // want "named return value val is never assigned"
}

func ShadowingPositiveCase(a int) (val, err int) {
	{
		val := 2
		// fmt.Println(val)
		for range val {
			fmt.Println()
		}
	}
	err = 2
	return // want "named return value val is never assigned"
}

// func VarShadowingPositiveCase(a int) (val, err int) {
// 	{
// 		var val = 2
// 		fmt.Println(val)
// 	}
// 	err = 2
// 	return
// }

// func MultiVarShadowingPositiveCase(a int) (val, err int) {
// 	{
// 		var (
// 			val  = 2
// 			val2 = "err"
// 		)
// 		fmt.Println(val, val2)
// 	}
// 	err = 2
// 	return
// }

// func VarShadowingWithAssignmentPositiveCase(a int) (val, err int) {
// 	{
// 		var val = 2
// 		val = 3
// 		fmt.Println(val)
// 	}
// 	err = 2
// 	return
// }

// func FunctionPositiveCase(a int) (val, err int) {
// 	fmt.Println(val)
// 	err = 2
// 	return
// }
