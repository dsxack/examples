package main

import "C"

func main() {}

//export Sum
func Sum(a, b C.int) C.int { return C.int(int(a) + int(b)) }
