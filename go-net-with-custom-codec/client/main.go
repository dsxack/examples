package main

import (
	"encoding/binary"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		panic(err)
	}

	go func() {
		_, _ = io.Copy(os.Stdout, conn)
	}()

	err = binary.Write(conn, binary.LittleEndian, uint64(4))
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("test"))
	if err != nil {
		panic(err)
	}

	err = binary.Write(conn, binary.LittleEndian, uint64(5))
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("test2"))
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 1)

	err = conn.Close()
	if err != nil {
		panic(err)
	}
}
