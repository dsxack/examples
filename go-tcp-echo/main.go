package main

import (
	"io"
	"log"
	"net"
)

const (
	listenAddress = "0.0.0.0:8000"
)

func handle(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	log.Println("Connection opened.")

	_, err := io.CopyBuffer(conn, conn, make([]byte, 1))
	if err != nil {
		panic(err)
	}

	err = conn.Close()
	if err != nil {
		panic(err)
	}

	log.Println("Connection closed.")
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		panic(err)
	}

	log.Printf("Start listen address %s", listenAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}
