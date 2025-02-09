package main

import (
	"log"
	"net"

	"github.com/toumakido/gohttp/net/obj"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen tcp: %s", err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to accept: %s", err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	b := make([]byte, 1024)
	_, err := conn.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	req, err := obj.NewRequest(b)
	// conn.Write([]byte("response"))
	// conn.Close()
}
