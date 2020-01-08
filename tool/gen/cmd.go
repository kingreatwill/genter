package gen

import (
	// main package
	"github.com/clipperhouse/typewriter"
)

func main() {
	app, err := typewriter.NewApp("+gen")
	if err != nil {
		panic(err)
	}
	app.WriteAll()
}
