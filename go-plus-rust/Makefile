ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
LIBS := hello\
	threaded

build:
	for lib in ${LIBS}; do \
		cd $(ROOT_DIR)lib/$${lib} && cargo build --release ; \
		cp $(ROOT_DIR)lib/$${lib}/target/release/lib$${lib}.dylib $(ROOT_DIR)lib/ ; \
	done
	go build -ldflags="-r $(ROOT_DIR)lib" main.go

run: build
	./main
