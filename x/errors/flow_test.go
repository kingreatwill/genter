package errors

import (
	"fmt"
)

func Example_flow() {
	err := Flow(func() error {
		fmt.Println(1)
		return New("whoops")
	}, func() error {
		fmt.Println(2)
		return nil
	})
	if err != nil {
		fmt.Printf("%+v", err)
	}
	// Output:
}
