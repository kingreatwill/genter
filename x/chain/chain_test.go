package chain

import (
	"fmt"
	"github.com/openjw/genter/x/context"
	"github.com/openjw/genter/x/errors"
)

func Example_1() {
	err := New(func(ctx *context.Context) (next bool, err error) {
		fmt.Println(ctx.Get())
		ctx.Set(2)
		return true, nil
	}, func(ctx *context.Context) (next bool, err error) {
		fmt.Println(ctx.Get())
		return true, errors.New("can't next.")
	}, func(ctx *context.Context) (next bool, err error) {
		fmt.Println(ctx.Get())
		return true, nil
	}).Run(1)
	fmt.Println(err)
	// Output:
}
