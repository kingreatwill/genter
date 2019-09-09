package env

import (
	"github.com/openjw/genter/x/env"
	"github.com/openjw/genter/x/strings"
	"os"
	"sync"
)

type EnvConfig struct {
	envName string
	reg     bool
	call    func()

	lock  sync.Mutex
	confs map[string]string
}

// 支持单个环境变量, 多个分号隔开.
func New(envName string, reg bool, call func()) *EnvConfig {
	return &EnvConfig{envName: envName, reg: reg, call: call, confs: map[string]string{}}
}

func (c *EnvConfig) Gets() map[string]string {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.confs
}

func (c *EnvConfig) Parse(v interface{}) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return env.Parse(v)
}

func (c *EnvConfig) Get(key string) string {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.confs[key]
}

func (c *EnvConfig) Reload() error {
	c.lock.Lock()
	nconfs := map[string]string{}
	es := os.Environ()
	for _, cf := range es {
		cfs := strings.Split(cf, "=")
		if cfs[0] != "" {
			nconfs[cfs[0]] = os.Getenv(cfs[0])
		}
	}
	c.confs = nconfs
	c.lock.Unlock()
	if c.call != nil {
		c.call()
	}
	return nil
}
