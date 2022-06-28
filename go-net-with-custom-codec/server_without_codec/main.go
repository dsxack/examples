package main

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"time"
)

type handler struct {
	*gnet.EventServer
}

func (h *handler) React(frame []byte, _ gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Printf("received %d bytes\n", len(frame))
	fmt.Printf("%s\n", string(frame))
	out = frame
	return
}

func main() {
	handler := &handler{}

	err := gnet.Serve(
		handler,
		"tcp://0.0.0.0:8001",
		gnet.WithMulticore(true),
		gnet.WithTCPKeepAlive(time.Minute*5),
	)
	if err != nil {
		panic(err)
	}
}
