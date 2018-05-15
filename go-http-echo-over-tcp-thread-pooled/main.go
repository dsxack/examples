package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
)

const address = "0.0.0.0:8000"
const bufSize = 512

func worker(input <-chan net.Conn) {
	for conn := range input {
		handler(conn)
	}
}

func handler(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r.(error))
		}
	}()
	defer conn.Close()

	conn.Write([]byte("HTTP/1.1 200 OK\n"))
	conn.Write([]byte("Transfer-Encoding: chunked\n"))
	conn.Write([]byte("\r\n"))

	buf := make([]byte, bufSize)

	for {
		m, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		fmt.Fprintf(conn, "%X\r\n", m)
		conn.Write(buf[0:m])
		conn.Write([]byte("\r\n"))

		if m < bufSize {
			break
		}
	}

	conn.Write([]byte("0\r\n\r\n"))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Start listen %s for new connections\n", address)

	input := make(chan net.Conn, 100)
	for i := 1; i <= runtime.NumCPU(); i += 1 {
		go worker(input)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		input <- conn
	}
}
