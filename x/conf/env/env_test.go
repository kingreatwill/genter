package env

import "fmt"

func ExampleParse() {
	c := New("", false, func() {
		fmt.Println("call")
	})
	c.Reload()
	fmt.Println(c.confs)
	// Output: {Home:/tmp/fakehome Port:3000 IsProduction:false Inner:{Foo:foobar}}
}
