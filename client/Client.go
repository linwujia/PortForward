package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type Client struct {
	address     string
	port        string
	forwardPort string
}

func NewClient(address, port, forwardPort string) *Client {
	client := new(Client)
	client.address = address
	client.port = port
	client.forwardPort = forwardPort
	return client
}

func (c *Client) Accept() {

	dialer := c.clientBackendDialer()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", c.address, c.port), 10*time.Second)
	if err != nil {
		log.Fatal(err)
		return
	}

	accept, err := dialer()
	if err != nil {
		log.Printf("client backen dialer err %v", err)
		return
	}

	group := &sync.WaitGroup{}
	group.Add(2)
	go transfer(conn, accept, group)
	go transfer(accept, conn, group)
	group.Wait()
}

func (c *Client) clientBackendDialer() func() (net.Conn, error) {
	return func() (conn net.Conn, err error) {
		return net.DialTimeout("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", c.forwardPort), 10*time.Second)
	}
}

func transfer(writer io.Writer, reader io.Reader, group *sync.WaitGroup) {
	defer group.Done()
	io.Copy(writer, reader)
}
