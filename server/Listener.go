package main

import (
	"fmt"
	"github.com/ghostunnel/ghostunnel/socket"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type Listener struct {
	address string
	port    string

	targetAddress string
	targetPort    string

	handers sync.WaitGroup
}

func NewListener(address, port string, targetAddress, targetPort string) *Listener {
	listener := new(Listener)
	listener.address = address
	listener.port = port
	listener.targetAddress = targetAddress
	listener.targetPort = targetPort
	return listener
}

func (l *Listener) Accept() {

	dialer := l.clientBackendDialer()

	listener, err := socket.ParseAndOpen(fmt.Sprintf("%s:%s", l.address, l.port))
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			continue
		}

		accept, err2 := dialer()
		if err2 != nil {
			log.Printf("client backen dialer err %v", err)
			break
		}

		go l.transfer(conn, accept)
	}

	l.Wait()
}

func (l *Listener) clientBackendDialer() func() (net.Conn, error) {
	return func() (conn net.Conn, err error) {
		return net.DialTimeout("tcp", fmt.Sprintf("%s:%s", l.targetAddress, l.targetPort), 10*time.Second)
	}
}

func (l *Listener) transfer(listener net.Conn, backConn net.Conn) {
	l.handers.Add(2)
	go l.transferData(listener, backConn)
	go l.transferData(backConn, listener)
	l.handers.Wait()
}

func (l *Listener) Wait() {
	l.handers.Wait()
}

func (l *Listener) transferData(writer io.Writer, reader io.Reader) {
	defer l.handers.Done()
	io.Copy(writer, reader)
}
