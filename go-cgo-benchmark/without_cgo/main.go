package main

func main() {
	for i := 1; i <= 10000; i++ {
		fibonacci(i)
	}
}

func fibonacci(n int) int {
	var i int
	var t1 = 0
	var t2 = 1
	var nextTerm int

	for i = 1; i <= n; i++ {
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
	}

	return t2
}
