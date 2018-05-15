package main

import (
	"io"
	"log"
)

type Pipe struct {
	pipeReader    io.Reader
	pipeWriter    io.Writer
	contentLen    chan int64
	writerIsReady chan struct{}
	readerIsReady chan struct{}
}

func (pipe *Pipe) Write(c io.Writer) {
	done := make(chan struct{})

	go func() {
		log.Println("writer waiting for reader")
		<-pipe.readerIsReady

		length := <-pipe.contentLen
		log.Printf("start write %d bytes", length)
		io.CopyN(c, pipe.pipeReader, length)
		log.Printf("%d bytes writed", length)

		done <- struct{}{}
	}()

	pipe.writerIsReady <- struct{}{}
	log.Println("writer is ready")

	<-done

}

func (pipe *Pipe) Read(r io.Reader, length int64) {
	done := make(chan struct{})

	go func() {
		log.Println("reader waiting for writer")
		<-pipe.writerIsReady

		log.Printf("start read %d bytes", length)
		io.CopyN(pipe.pipeWriter, r, length)
		log.Printf("%d bytes readed", length)

		done <- struct{}{}
	}()

	pipe.readerIsReady <- struct{}{}
	log.Println("reader is ready")
	pipe.contentLen <- length

	<-done
}

func newPipe() *Pipe {
	pipeReader, pipeWriter := io.Pipe()

	return &Pipe{
		pipeReader:    pipeReader,
		pipeWriter:    pipeWriter,
		contentLen:    make(chan int64),
		writerIsReady: make(chan struct{}),
		readerIsReady: make(chan struct{}),
	}
}
