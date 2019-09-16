package exec

import (
	"os"
	"strings"
)

func ExampleExecute() {
	Execute(strings.Join(os.Args[1:], " "))
	// output:
}
