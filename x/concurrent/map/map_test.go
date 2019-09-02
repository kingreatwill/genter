package cmap

import "fmt"

func Example_1() {
	m := New()
	// Sets item within map, sets "bar" under key "foo"
	m.Set("foo", "bar")
	// Retrieve item from map.
	if tmp, ok := m.Get("foo"); ok {
		bar := tmp.(string)
		fmt.Println(bar)
	}
	// Removes item under key "foo"
	m.Remove("foo")
	// Output:bar
}
