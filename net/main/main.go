package main

import (
	"log"
	"net"

	"github.com/toumakido/gohttp/net/handler"
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
			res = route(req)
		}
		conn.Write([]byte(res.String()))
	}
}

func route(req *request.Request) *response.Response {
	switch req.Method {
	case request.HTTPGet, request.HTTPHead:
		switch req.Endpoint {
		case "/hello", "/hello/":
			return handler.GetHello()
		default:
			return handler.GetIndex(req)
		}
	case request.HTTPPost:
		switch req.Endpoint {
		default:
			return handler.PostIndex(req)
		}
	default:
		return handler.GetIndex(req)
	}
}
