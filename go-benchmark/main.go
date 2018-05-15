package main

import (
	"runtime"
	"sync"
)

const (
	threads_count = 1000
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}
	wg.Add(threads_count)

	for i := 0; i < threads_count; i++ {
		go func() {
			defer wg.Done()

			x := 0
			for j := 0; j < 5000000; j++ {
				x += 1
			}
		}()
	}

	wg.Wait()
}
