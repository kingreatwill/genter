package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func mockSendToServer(url string) {
	fmt.Printf("server url: %s\n", url)
}

/*
go run main.go  ，将在同级目录生成trace.out文件。
此时，执行go tool命令
go tool trace trace.out

在这里，我们只关心两项指标:
第一行View trace（可视化整个程序的调度流程）
第二行Groutine analysis。

进入Goroutine analysis项。
```
Goroutines:
main.main.func1 N=3
runtime/trace.Start.func1 N=1
runtime.main N=1
N=4
```
可以看到，程序一共有5个goroutine，分别是三个for循环里启动的匿名go func()、一个trace.Start.func1和runtime.main。

进入main.main.func1 ,记住三个Goroutine的编号（我这里是7,8,9），此时，退回到http://127.0.0.1:xx页面，进入View trace项。
可以看到Goroutine信息
*/
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	urls := []string{"0.0.0.0:5000", "0.0.0.0:6000", "0.0.0.0:7000"}
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mockSendToServer(url)
		}()
	}
	wg.Wait()
}