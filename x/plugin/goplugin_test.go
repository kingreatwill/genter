package goplugin

import (
	"testing"
)

// go build -buildmode=plugin -o aplugin.so aplugin.go
func Test_1(t *testing.T) {
	f, e := GetHandler("xxx.so", "GetHandler")
	t.Log(f, e)
}
