package main

import (
	"fmt"
	"sync"
)

func mockSendToServer(url string) {
	fmt.Printf("server url: %s\n", url)
}

func main() {
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
/*
output:
server url: 0.0.0.0:7000
server url: 0.0.0.0:7000
server url: 0.0.0.0:7000

*/
