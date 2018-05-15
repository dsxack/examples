package main

/*
#cgo LDFLAGS: -L./fibonacci -lfibonacci
#include "fibonacci/fibonacci.h"
 */
import "C"

func main() {
	for i := 1; i <= 10000; i++ {
		C.fibonacci(C.int(i))
	}
}
