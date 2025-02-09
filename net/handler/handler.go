package handler

import (
	"encoding/json"

	"github.com/toumakido/gohttp/net/request"
	"github.com/toumakido/gohttp/net/response"
)

func Index(req *request.Request) *response.Response {
	// リクエスト内容をjsonで返す。（200）
	body, err := json.Marshal(req.Header)
	if err != nil {
		return response.NewErrorResponse(err)
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	return response.NewResponse(header, string(body))
}

type getHelloResponse struct {
	Messsage string `json:"message"`
}

func Hello(*request.Request) *response.Response {
	var b getHelloResponse
	b.Messsage = "hello!"
	body, err := json.Marshal(b)
	if err != nil {
		return response.NewErrorResponse(err)
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	return response.NewResponse(header, string(body))
}
