package main

import (
	"compress/gzip"
	"fileless/resources"
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"os"
	"os/exec"
)

type MemFile struct {
	fd int
}

func (f MemFile) Write(b []byte) (int, error) {
	return unix.Write(f.fd, b)
}

func main() {
	fd, err := unix.MemfdCreate("embeded", 0)
	if err != nil {
		panic(err)
	}

	gEmbeded, err := resources.Box().Open("embeded.gz")
	if err != nil {
		panic(err)
	}

	embeded, err := gzip.NewReader(gEmbeded)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(MemFile{fd}, embeded)
	if err != nil {
		panic(err)
	}

	pid := os.Getpid()
	filepath := fmt.Sprintf("/proc/%d/fd/%d", pid, fd)

	cmd := exec.Command(filepath)
	cmd.Run()
	cmd.Wait()
}
