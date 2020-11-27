# go-plugins-cgo-with-same-symbol-names

Example demonstrate using go plugins that statically linked C libraries with CGO with same symbol names
with using `-fvisibility=hidden` GCC flag

https://github.com/golang/go/issues/42854

#### Build

```sh
make
```

#### Clean

```sh
make clean
```

#### Output

```sh
./main
Load Go plugin hello1.so
Load Go plugin hello2.so

Call Go plugin hello1.go
Hello from Go plugin hello1.go
Call C hello1.c
Hello from C hello1.c

Call Go plugin hello2.go
Hello from Go plugin hello2.go
Call С hello2.c
Hello from C hello2.c
```
