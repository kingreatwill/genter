package conf

type Config interface {
	Parse(v interface{}) error
	Get(key string) interface{}
}
