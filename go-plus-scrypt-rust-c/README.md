# Rust scrypt bindings for go

Attempt using rust scrypt library in go instead of own go library
 to reduce garbage collection.

## test

```
make gotest
```

## bench

```
make gobench
```

## Result

It is fault. CGO is too slow. 
Maybe I will try to send multiple input to rust from go and parallelize it.

```
cd src/goscrypt; go test -bench .;
goos: darwin
goarch: amd64
pkg: scrypt
BenchmarkRustScryptKey-8               6         177050621 ns/op
BenchmarkGoScryptKey-8             14812             80165 ns/op
PASS
ok      scrypt  3.900s
```