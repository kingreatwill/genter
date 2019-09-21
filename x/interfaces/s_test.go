package interfaces

type X2 struct {
	A string
}

func f() *X2 {
	return nil
}

func ExampleIsZero() {
	println(IsZero(f()))
	// output:
}
