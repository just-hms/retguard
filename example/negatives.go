package example

func SimpleNegativeCase(a int) (val int) {
	val = 1
	return
}

func DuplNegativeCase(a int) (val, err int) {
	val = 1
	err = 2
	return
}

func create(input *int) {
	*input = 0
}

// even if val is never been assigned it has been modified, which is good enough for me
func FunctionNegativeCase(a int) (val, err int) {
	create(&val)
	err = 2
	return
}

// even if val is never been assigned it has been modified, which is good enough for me
// this count also for interfaces and types which are pointer by definition
func FunctionComplexNegativeCase(a int) (val *int, err int) {
	create(val)
	err = 2
	return
}

type obj struct {
	Content any
}

func ComplexNegativeCase(a int) (val obj, err int) {
	val.Content = 3
	err = 2
	return
}
