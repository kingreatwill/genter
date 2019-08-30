package main

import (
	"fmt"
	"github.com/openjw/genter/x/json"
)

func main() {
	fmt.Println(json.Marshal(struct {
	}{}))
}
