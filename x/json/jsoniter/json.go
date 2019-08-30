package jsoniter

import (
	"github.com/json-iterator/go"
)

type jsonConf struct {
	api jsoniter.API
}

func New() *jsonConf {
	return &jsonConf{
		jsoniter.ConfigFastest,
	}
}

func (j *jsonConf) Unmarshal(data []byte, v interface{}) error {
	return j.api.Unmarshal(data, v)
}

func (j *jsonConf) Marshal(v interface{}) ([]byte, error) {
	return j.api.Marshal(v)
}
