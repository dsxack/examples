all: lib extension test

lib:
	go build -buildmode=c-archive -o libsum.a sum.go

extension:
	python3 setup.py build_ext --inplace

test:
	pytest
