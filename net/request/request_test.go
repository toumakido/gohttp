package request

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewRequest(t *testing.T) {
	t.Parallel()

	req := `GET /oreilly-japan/real-world-http HTTP/1.1
		Host: github.com
		Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8
		Accept-Encoding: gzip, deflate, br
		Accept-Language: ja-JP,ja;q=0.9,en-US;q=0.8,en;q=0.7
		Content-Type: application/json
		User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko)Chrome/98.0.4758.82 Safari/537.36
		
		{
			"message": "hello"
		}
		`

	got, err := NewRequest([]byte(req))
	require.NoError(t, err)
	require.Equal(t, HTTPGet, got.Method)
	require.Equal(t, "/oreilly-japan/real-world-http", got.Endpoint)
	require.Equal(t, "HTTP/1.1", got.Version)
	require.Equal(t, 6, len(got.Header))

	type body struct {
		Message string  `json:"message"`
		Item    *string `json:"item"`
	}
	var b body
	require.NoError(t, json.Unmarshal([]byte(got.Body), &b))
	require.Equal(t, "hello", b.Message)
	require.Nil(t, b.Item)
}
