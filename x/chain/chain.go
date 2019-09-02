package chain

import (
	"github.com/openjw/genter/x/context"
)

type Chain struct {
	fns []func(*context.Context) (next bool, err error)
}

func New(fns ...func(*context.Context) (bool, error)) *Chain {
	return &Chain{fns: fns}
}

func (c *Chain) Append(fns ...func(*context.Context) (bool, error)) *Chain {
	c.fns = append(c.fns, fns...)
	return c
}

func (c *Chain) Merge(chains ...*Chain) *Chain {
	for k := range chains {
		c.Append(chains[k].fns...)
	}
	return c
}

func (c *Chain) Execute(context *context.Context) error {
	for _, fn := range c.fns {
		b, err := fn(context)
		if !b {
			return err
		}
	}
	return nil
}

func (c *Chain) Run(val interface{}) error {
	ctx := ContextWithValue(context.Background(), val)
	return c.Execute(ctx)
}
