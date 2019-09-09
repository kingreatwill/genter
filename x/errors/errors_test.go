package errors

import (
	"errors"
	"fmt"
)

// go get github.com/pkg/errors
func ExampleWithStack_printf() {
	cause := New("whoops")
	fmt.Printf("%+v", cause)
	// Output:
}

func ExampleWithIs() {
	cause := New("whoops")
	errors.Is(cause, nil)
	fmt.Printf("%+v", cause)
	// Output:
}

func Example_1() {
	w := fmt.Errorf("FFF:%w", errors.New("123"))
	fmt.Println(w)
	// Output:
}
