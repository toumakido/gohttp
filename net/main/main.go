package main

import (
	"log"
	"net"

	"github.com/toumakido/gohttp/net/request"
	"github.com/toumakido/gohttp/net/response"
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
	n, err := conn.Read(b)
	if err != nil {
		log.Fatalf("failed to read connection :%s", err.Error())
	}

	if n > 0 {
		var res *response.Response
		req, err := request.NewRequest(b)
		if err != nil {
			res = response.NewErrorResponse(err)
		} else {
			res = response.NewResponse(req)
		}
		conn.Write([]byte(res.String()))
	}
}
