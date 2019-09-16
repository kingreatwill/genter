package examples

import (
	"github.com/openjw/genter/x/conf"
	"github.com/openjw/genter/x/conf/yaml"
)

func Example_test1() {
	conf.Add(yaml.YamlConfig{Data: ""})
	// output:
}
