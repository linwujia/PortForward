package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			log.Fatal(err2)
		}

		connForward, err2 := net.Dial("tcp", "www.baidu.com:80")
		if err2 != nil {
			log.Fatal(err2)
		}

		go handleConnection(conn, connForward)
		go handleConnection(connForward, conn)
	}
}

func handleConnection(r, w net.Conn) {
	defer r.Close()
	defer w.Close()

	io.Copy(r, w)
}
