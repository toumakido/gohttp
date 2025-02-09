package request

import (
	"fmt"
	"strings"
)

type Request struct {
	Method   httpMethod
	Endpoint string
	Version  string
	Header   map[string]string
	Body     string
}

func NewRequest(req []byte) (*Request, error) {
	reqLines := strings.Split(string(req), "\n")

	method, endopoint, ver, err := parseFirstLine(reqLines[0])
	if err != nil {
		return nil, err
	}

	header := make(map[string]string, 0)
	var body string
	var bodyStartFromNext bool
	for _, line := range reqLines[1:] {
		if !bodyStartFromNext {
			splitted := strings.Split(strings.Trim(line, " "), ":")
			if len(splitted) >= 2 {
				header[splitted[0]] = strings.Join(splitted[1:], "")
			} else {
				bodyStartFromNext = true
			}
		} else {
			body += line
		}
	}

	return &Request{
		Method:   method,
		Endpoint: endopoint,
		Version:  ver,
		Header:   header,
		Body:     body,
	}, nil
}

type httpMethod string

const (
	HTTPHead httpMethod = "HEAD"
	HTTPGet  httpMethod = "GET"
	HTTPPost httpMethod = "POST"
)

func hewHTTPMethod(s string) (httpMethod, error) {
	switch s {
	case "HEAD":
		return HTTPHead, nil
	case "GET":
		return HTTPGet, nil
	case "POST":
		return HTTPPost, nil
	}
	return "", fmt.Errorf("unknown method %s", s)
}

func parseFirstLine(line string) (httpMethod, string, string, error) {
	splitted := strings.Split(line, " ")
	if len(splitted) == 3 {
		method, err := hewHTTPMethod(splitted[0])
		if err != nil {
			return "", "", "", err
		}
		return method, splitted[1], splitted[2], nil
	}
	return "", "", "", fmt.Errorf("failed to parse request's first line :%s", line)
}
