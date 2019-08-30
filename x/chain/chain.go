package chain

import "net/http"

type Chain struct {
	hs []func(http.Handler) http.Handler
}

func New(handlers ...func(http.Handler) http.Handler) *Chain {
	return &Chain{hs: handlers}
}

func (c *Chain) Append(handlers ...func(http.Handler) http.Handler) *Chain {
	c = New(appendHandlers(c.hs, handlers...)...)

	return c
}

// Merge receives one or more Chain instances, and returns a merged Chain.
func (c *Chain) Merge(chains ...*Chain) *Chain {
	for k := range chains {
		c = New(appendHandlers(c.hs, chains[k].hs...)...)
	}

	return c
}

func appendHandlers(hs []func(http.Handler) http.Handler, ahs ...func(http.Handler) http.Handler) []func(http.Handler) http.Handler {
	lcur := len(hs)
	ltot := lcur + len(ahs)
	if ltot > cap(hs) {
		nhs := make([]func(http.Handler) http.Handler, ltot)
		copy(nhs, hs)
		hs = nhs
	}

	copy(hs[lcur:], ahs)

	return hs
}
