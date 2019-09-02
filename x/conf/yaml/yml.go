package yaml

import "gopkg.in/yaml.v2"

type YamlConfig struct {
	data string
}

func (y *YamlConfig) Parse(v interface{}) error {
	return yaml.Unmarshal([]byte(y.data), v)
}

func (y *YamlConfig) Get(key string) interface{} {

	return nil
}
