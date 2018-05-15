package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"io"
)

var (
	pipe = newPipe()
)

func handleReader(reader io.Reader, length int64) {
	pipe.Read(reader, length)
}

func handleWriter(writer io.Writer) {
	pipe.Write(writer)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		length, err := strconv.ParseInt(r.Header.Get("content-length"), 10, 64)
		if err != nil {
			panic(err)
		}
		handleReader(r.Body, length)
		return
	}

	if r.Method == http.MethodGet {
		handleWriter(w)
		return
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	var port int

	flag.IntVar(&port, "port", 0, "port")
	flag.Parse()

	if port <= 0 {
		panic(errors.New("port param is required"))
	}
	addr := fmt.Sprintf("0.0.0.0:%d", port)

	log.Printf("start listen address %s", addr)

	http.ListenAndServe(addr, http.HandlerFunc(handle))
}
