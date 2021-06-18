package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/panjf2000/gnet"
	"time"
)

const (
	headerLen = 8
)

type handler struct {
	*gnet.EventServer
}

func (h *handler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Printf("received %d bytes\n", len(frame))
	fmt.Printf("%s\n", string(frame))
	out = frame
	return
}

type codec struct{}

func (c2 *codec) Encode(_ gnet.Conn, buf []byte) ([]byte, error) {
	var newBuf bytes.Buffer
	err := binary.Write(&newBuf, binary.LittleEndian, uint64(len(buf)))
	if err != nil {
		return nil, err
	}
	newBuf.Write(buf)

	return newBuf.Bytes(), nil
}

func (c2 *codec) Decode(c gnet.Conn) ([]byte, error) {
	headerSize, header := c.ReadN(headerLen)
	if headerSize != headerLen {
		return nil, errors.New("not enough header data")
	}

	byteBuffer := bytes.NewBuffer(header)
	payloadLength, err := binary.ReadUvarint(byteBuffer)
	if err != nil {
		return nil, fmt.Errorf("error read payload size: %v", err)
	}

	payloadLen := int(payloadLength)
	protocolLen := headerLen + payloadLen

	dataSize, data := c.ReadN(protocolLen)
	if dataSize != protocolLen {
		return nil, errors.New("not enough payload data")
	}

	c.ShiftN(protocolLen)

	return data[headerLen:], nil
}

func main() {
	handler := &handler{}
	codec := &codec{}

	err := gnet.Serve(
		handler,
		"tcp://0.0.0.0:8001",
		gnet.WithMulticore(true),
		gnet.WithTCPKeepAlive(time.Minute*5),
		gnet.WithCodec(codec),
	)
	if err != nil {
		panic(err)
	}
}
