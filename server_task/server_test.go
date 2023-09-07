package main

import (
	"bufio"
	"log"
	"net"
	"sync"
	"testing"
)

func Test_handleConn(t *testing.T) {
	srv, cl := net.Pipe()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		handleConn(srv)
		srv.Close()
	}()

	_, err := cl.Write([]byte("concurerncy not parallelism\n"))
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(cl)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	_ = b
	// TODO check response

	wg.Wait()
	cl.Close()
}
