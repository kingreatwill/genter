package conf

import "reflect"

type Config interface {
	Gets() map[string]string
	Get(key string) string
	GetT(key string, t reflect.Type) interface{}
	Parse(v interface{}) error
}
