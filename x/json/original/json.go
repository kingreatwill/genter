package original

import "encoding/json"

type original struct {
}

func New() *original {
	return &original{}
}

func (j *original) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *original) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
