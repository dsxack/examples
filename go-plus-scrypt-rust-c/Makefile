ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

export CGO_CPPFLAGS=-I $(ROOT_DIR)include
export CGO_LDFLAGS=-L$(ROOT_DIR)lib

libscryptbindings = lib/libscryptbindings.a
cmain = src/cmain/main

$(libscryptbindings):
	cd src/rustscryptbindings; $(MAKE);
	cp src/rustscryptbindings/target/release/librustscryptbindings.a ./lib

$(cmain): $(libscryptbindings)
	cd src/cmain; cmake CMakeLists.txt; $(MAKE)

.PHONY: gotest
gotest: $(libscryptbindings)
	cd src/goscrypt; go test .;

.PHONY: gotest
gobench: $(libscryptbindings)
	cd src/goscrypt; go test -bench .;
