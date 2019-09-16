package yaml

import "gopkg.in/yaml.v2"

type YamlConfig struct {
	Data string
}

func (y *YamlConfig) Parse(v interface{}) error {
	return yaml.Unmarshal([]byte(y.Data), v)
}

func (y *YamlConfig) Get(key string) interface{} {

	return nil
}
