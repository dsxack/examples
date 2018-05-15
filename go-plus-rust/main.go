package main

/*
#cgo LDFLAGS: -L./lib -lhello -lthreaded
#include "./lib/hello.h"
#include "./lib/threaded.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	C.hello(C.CString("John Smith"))
	C.hello_int(C.int(5))

	number := 50000000
	threads := runtime.NumCPU()

	now := time.Now()
	C.process(C.int(number), C.int(threads))
	fmt.Printf("time elapsed: %s\n", time.Since(now))

	group := sync.WaitGroup{}
	group.Add(threads)
	now = time.Now()
	for t := 0; t < threads; t++ {
		go func() {
			x := 0
			for n := 0; n < number; n++ {
				x += 1
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Printf("time elapsed: %s\n", time.Since(now))
}
