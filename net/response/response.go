package response

import (
	"fmt"
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

func NewResponse(header map[string]string, body string) *Response {
	return &Response{
		Version: httpVersion,
		Status:  statusOK,
		Header:  header,
		Body:    body,
	}
}

func NewErrorResponse(err error) *Response {
	fmt.Printf("error response: %s\n", err.Error())

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
