package handler

import (
	"bytes"
	"encoding/json"

	"github.com/toumakido/gohttp/net/request"
	"github.com/toumakido/gohttp/net/response"
)

func GetIndex(req *request.Request) *response.Response {
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
	Message string `json:"message"`
}

func GetHello() *response.Response {
	b := getHelloResponse{Message: "hello!"}
	body, err := json.Marshal(b)
	if err != nil {
		return response.NewErrorResponse(err)
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	return response.NewResponse(header, string(body))
}

type postIndexRequestBody struct {
	Message string `json:"message"`
}

type postIndexResponse struct {
	Message    string `json:"message"`
	ReqMessage string `json:"reqMessage"`
}

func PostIndex(req *request.Request) *response.Response {
	var msg postIndexRequestBody
	err := json.Unmarshal(bytes.Trim([]byte(req.Body), "\x00"), &msg)
	if err != nil {
		return response.NewErrorResponse(err)
	}
	b := postIndexResponse{Message: "hello!", ReqMessage: msg.Message}
	body, err := json.Marshal(b)
	if err != nil {
		return response.NewErrorResponse(err)
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	return response.NewResponse(header, string(body))
}
