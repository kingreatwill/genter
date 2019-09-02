package context

import (
	"context"
)

type Context struct {
	context.Context
	key, val interface{}
}

func Background() *Context {
	return &Context{Context: context.Background()}
}

func WithValue(parent *Context, key, val interface{}) *Context {
	if key == nil {
		panic("nil key")
	}
	return &Context{parent, key, val}
}

func (c *Context) Get() interface{} {
	return c.val
}

func (c *Context) Set(val interface{}) {
	c.val = val
}

func (c *Context) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}
