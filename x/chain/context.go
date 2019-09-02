package chain

import "github.com/openjw/genter/x/context"

type chainKey struct {
}

var activeKey = chainKey{}

func ContextWithValue(ctx *context.Context, val interface{}) *context.Context {
	return context.WithValue(ctx, activeKey, val)
}
