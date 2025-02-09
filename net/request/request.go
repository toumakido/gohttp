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
}

func NewRequest(req []byte) (*Request, error) {
	reqLines := strings.Split(string(req), "\n")

	method, endopoint, ver, err := parseFirstLine(reqLines[0])
	if err != nil {
		return nil, err
	}

	header := make(map[string]string, 0)
	for _, line := range reqLines[1:] {
		splitted := strings.Split(strings.Trim(line, " "), ":")
		if len(splitted) >= 2 {
			header[splitted[0]] = strings.Join(splitted[1:], "")
		}
	}

	return &Request{
		Method:   method,
		Endpoint: endopoint,
		Version:  ver,
		Header:   header,
	}, nil
}

type httpMethod string

const (
	httpHead httpMethod = "HEAD"
	httpGet  httpMethod = "GET"
	httpPost httpMethod = "POST"
)

func hewHTTPMethod(s string) (httpMethod, error) {
	switch s {
	case "HEAD":
		return httpHead, nil
	case "GET":
		return httpGet, nil
	case "POST":
		return httpPost, nil
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
