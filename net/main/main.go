package main

import (
	"log"
	"net"

	"github.com/toumakido/gohttp/net/req"
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
	defer conn.Close()

	b := make([]byte, 1024)
	_, err := conn.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	_, err = req.NewRequest(b)
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) error {
	return err
}
