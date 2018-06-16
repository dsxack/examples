all: zephir libsum build

libsum:
	go build -buildmode=c-archive -o zephir/libsum.a sum.go

zephir:
	zephir init sum
	mv sum zephir
	cp sum.zep zephir/sum/

build:
	cd zephir; zephir build

test:
	php -d extension=sum.so ./test.php
