package request

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  HttpMethod
	Version string
	Header  map[string]string
}

func NewRequest(req []byte) (*Request, error) {
	reqLines := strings.Split(string(req), "\n")

	method, ver, err := parseFirstLine(reqLines[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parseFirstLine: %w", err)
	}

	header := make(map[string]string, 0)
	for _, line := range reqLines[1:] {
		splitted := strings.Split(strings.Trim(line, " "), ":")
		if len(splitted) == 2 {
			header[splitted[0]] = splitted[1]
		}
	}

	return &Request{
		Method:  method,
		Version: ver,
		Header:  header,
	}, nil
}

type HttpMethod string

const (
	httpHead HttpMethod = "HEAD"
	httpGet  HttpMethod = "GET"
	httpPost HttpMethod = "POST"
)

func newHTTPMethod(s string) (HttpMethod, error) {
	switch s {
	case "HEAD":
		return httpHead, nil
	case "GET":
		return httpGet, nil
	case "POST":
		return httpGet, nil
	}
	return "", fmt.Errorf("unknown method %s", s)
}

func parseFirstLine(line string) (HttpMethod, string, error) {
	splitted := strings.Split(line, " ")
	if len(splitted) == 3 {
		method, err := newHTTPMethod(splitted[0])
		if err != nil {
			return "", "", err
		}
		return method, splitted[2], nil
	}
	return "", "", fmt.Errorf("failed to parse request's first line :%s", line)
}
