package main

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -lhello2

void hello();
*/
import "C"
import "fmt"

func Hello() {
	fmt.Println("Hello from Go plugin hello2.go")
	fmt.Println("Call С hello2.c")
	C.hello()
}
