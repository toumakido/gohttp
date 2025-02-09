package response

import (
	"encoding/json"
	"fmt"

	"github.com/toumakido/gohttp/net/request"
)

type Response struct {
	Version string
	Status  status
	Header  map[string]string
	Body    string
}

const httpVersion string = "HTTP/1.1"

type status string

const (
	statusOK       status = "200 OK"
	status500Error status = "500 Internal Server Error"
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
		Version: httpVersion,
		Status:  statusOK,
		Header:  header,
		Body:    string(body),
	}
}

func NewErrorResponse(err error) *Response {
	fmt.Println("error response: ", err.Error())

	return &Response{
		Version: httpVersion,
		Status:  status500Error,
		Body:    err.Error(),
	}
}

func (res *Response) String() string {
	ret := fmt.Sprintf("%s %s\n", res.Version, string(res.Status)) // 1行目

	for key, val := range res.Header {
		ret += fmt.Sprintf("%s: %s\n", key, val) // headerを追加
	}

	ret += fmt.Sprintf("\n%s", res.Body) // 1行あけて、body

	return ret
}
