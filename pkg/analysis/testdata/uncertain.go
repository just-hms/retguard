package testdata

import "fmt"

// even if val is never been assigned it has been modified, which is good enough for me
// this count also for interfaces and types which are pointer by definition
func FunctionComplexNegativeCase(a int) (val *int, err int) {
	hCreate := func(input *int) {
		*input = 0
	}
	hCreate(val)
	err = 2
	return // want "named return value val is never assigned"
}

type Scope int

func (s *Scope) init() {}

func InitFunc() (s Scope) {
	s.init()
	return s // want "named return value s is never assigned"
}

func FunctionPositiveCase(a int) (val, err int) {
	fmt.Println(val)
	err = 2
	return // want "named return value val is never assigned"
}
