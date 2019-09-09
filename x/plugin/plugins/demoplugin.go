package plugins

import (
	"net/http"
)

// http://www.361way.com/go-plugin/5925.html
// go build -buildmode=plugin -o aplugin.so aplugin.go
//除了上面提到的buildmode=plugin外，还有一种用法就是 buildmode=c-shared ，使用该参数时会生成出来两个文件，一个.so文件，一个.h头文件 ，使用起来就和使用c 生成的库文件和模块文件一样使用
// go build -o awesome.so -buildmode=c-shared awesome.go
func GetHandler() http.HandlerFunc {
	return nil
}
