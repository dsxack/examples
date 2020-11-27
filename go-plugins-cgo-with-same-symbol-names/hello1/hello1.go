package main

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -lhello1

void hello();
*/
import "C"
import "fmt"

func Hello() {
	fmt.Println("Hello from Go plugin hello1.go")
	fmt.Println("Call C hello1.c")
	C.hello()
}
