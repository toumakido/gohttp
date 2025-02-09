package response

import (
	"encoding/json"
	"fmt"

	"github.com/toumakido/gohttp/net/request"
)

type Response struct {
	Version string
	Status  Status
	Header  map[string]string
	Body    string
}

type Status string

const (
	StatusOK       Status = "200 OK"
	Status500Error Status = "500 Internal Server Error"
)

func NewResponse(req *request.Request) *Response {
	// リクエスト内容をjsonで返す。（200）
	body, err := json.Marshal(req.Header)
	if err != nil {
		return NewErrorResponse(err)
	}
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	return &Response{
		Version: "HTTP/1.1",
		Status:  StatusOK,
		Header:  header,
		Body:    string(body),
	}
}

func NewErrorResponse(err error) *Response {
	return &Response{
		Version: "HTTP/1.1",
		Status:  Status500Error,
		Body:    err.Error(),
	}
}

/*
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Type: text/html; charset=utf-8
Date: Sat, 27 May 2023 02:54:02 GMT
Server: GitHub.com
<!DOCTYPE html>
<html lang="en" data-color-mode="auto" data-light-theme="light" data-dark-theme="dark" dataa11y-animated-images="system">
<head>
<meta charset="utf-8">
<script type="application/javascript" src="https://github.githubassets.com/assets/code-menu.
js"></script>
: （略）
*/

func (res *Response) String() string {
	ret := fmt.Sprintf("%s %s\n", res.Version, string(res.Status)) // 1行目
	for key, val := range res.Header {
		ret += fmt.Sprintf("%s: %s\n", key, val) // header
	}
	ret += fmt.Sprintf("\n%s", res.Body) // 1行あけて、body
	return ret
}
