all: sum.a main

sum.a:
	go build -buildmode=c-archive -o sum.a sum.go

main:
	gcc -o main main.c sum.a
