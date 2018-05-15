package main

import (
	"io"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.ListenAndServe("0.0.0.0:8001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		w.WriteHeader(http.StatusOK)
		io.Copy(w, r.Body)
	}))
}
