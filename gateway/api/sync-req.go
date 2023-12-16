package api

import (
	"bytes"
	"fmt"
	"net/http"
)

func CommunicateSync(host, port, uri, method string, req []byte) (*http.Response, error) {
	httpURL := fmt.Sprintf("http://%s:%s/%s", host, port, uri)
	request, err := http.NewRequest(method, httpURL, bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(request)
}
