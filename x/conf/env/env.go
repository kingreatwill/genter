package env

import (
	"github.com/openjw/genter/x/env"
	"os"
)

type EnvConfig struct {
}

func (c *EnvConfig) Parse(v interface{}) error {
	return env.Parse(v)
}

func (c *EnvConfig) Get(key string) interface{} {
	return DefaultCof(key, "")
}

func DefaultCof(env, value string) string {
	v := os.Getenv(env)
	if v == "" {
		return value
	}
	return v
}
