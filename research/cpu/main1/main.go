package main

import "runtime"

func main(){
	println(runtime.NumCPU(),runtime.GOMAXPROCS(-1))
	// 可以看到输出NumCpu为4，GOMAXPROCS也为4
	// linux 可以查看 /proc/cpuinfo
}
