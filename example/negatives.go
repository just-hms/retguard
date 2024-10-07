package example

import "fmt"

func SimpleNegativeCase(a int) (val int) {
	val = 1
	return
}

func ReturnNegativeCase(a int) (val int) {
	return 1
}

func IncNegativeCase() (val int) {
	val++
	return val
}

func DuplNegativeCase(a int) (val, err int) {
	val = 1
	err = 2
	return
}

func ShadowedAndNot(a int) (val int) {
	{
		val := 1
		fmt.Println(val)
	}
	val = 1
	return
}

// even if val is never been assigned it has been modified, which is good enough for me
func FunctionNegativeCase(a int) (val, err int) {
	hCreate := func(input *int) {
		*input = 0
	}
	hCreate(&val)
	err = 2
	return
}

// even if val is never been assigned it has been modified, which is good enough for me
// this count also for interfaces and types which are pointer by definition
func FunctionComplexNegativeCase(a int) (val *int, err int) {
	hCreate := func(input *int) {
		*input = 0
	}
	hCreate(val)
	err = 2
	return
}

type hObj struct {
	Content any
}

func ComplexNegativeCase(a int) (val hObj, err int) {
	val.Content = 3
	err = 2
	return
}

func Override(a int) (err int) {
	{
		err := 2
		return err
	}
}

func RealExample(contextID int) (item string, err error) {
	hGetSequence := func(_ int) ([]string, error) { return []string{"A", "B"}, nil }

	// get sequence by context ID
	seq, err := hGetSequence(contextID)
	if err != nil {
		return "", err
	}

	return seq[0], nil
}

type Scope int

func (s Scope) init() {}

func InitFunc() (s Scope) {
	s.init()
	return s
}
