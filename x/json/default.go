package json

import (
	"github.com/openjw/genter/x/json/original"
)

var j Ijson = original.New()

func SetProvider(i Ijson) {
	j = i
}

func Marshal(v interface{}) ([]byte, error) {
	return j.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return j.Unmarshal(data, v)
}
