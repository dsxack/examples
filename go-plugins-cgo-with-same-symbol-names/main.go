package main

import (
	"fmt"
	"plugin"
)

func main() {
	p1, err := plugin.Open("hello1/hello1.so")
	if err != nil {
		panic(err)
	}
	fmt.Println("Load Go plugin hello1.so")

	p2, err := plugin.Open("hello2/hello2.so")
	if err != nil {
		panic(err)
	}
	fmt.Println("Load Go plugin hello2.so\n")

	h1, err := p1.Lookup("Hello")
	if err != nil {
		panic(err)
	}

	h2, err := p2.Lookup("Hello")
	if err != nil {
		panic(err)
	}

	fmt.Println("Call Go plugin hello1.go")
	h1.(func())()
	fmt.Println("\nCall Go plugin hello2.go")
	h2.(func())()
}
