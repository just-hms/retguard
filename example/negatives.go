package example

func SimpleNegativeCase(a int) (val int) {
	val = 1
	return
}

func ReturnNegativeCase(a int) (val int) {
	return 1
}

func DuplNegativeCase(a int) (val, err int) {
	val = 1
	err = 2
	return
}

func hCreate(input *int) {
	*input = 0
}

// even if val is never been assigned it has been modified, which is good enough for me
func FunctionNegativeCase(a int) (val, err int) {
	hCreate(&val)
	err = 2
	return
}

// even if val is never been assigned it has been modified, which is good enough for me
// this count also for interfaces and types which are pointer by definition
func FunctionComplexNegativeCase(a int) (val *int, err int) {
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

func hGetSequence(_ int) ([]string, error) { return []string{"A", "B"}, nil }

func RealExample(contextID int) (item string, err error) {

	// get sequence by context ID
	seq, err := hGetSequence(contextID)
	if err != nil {
		return "", err
	}

	return seq[0], nil
}
