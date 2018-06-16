all: lib

lib:
	go build -buildmode=c-shared -o sum.a sum.go

test:
	python test.py
