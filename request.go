package listrak

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

type BaseRequest struct {
	Payload  io.Reader
	request  *http.Request
	client   http.Client
	basePath string
}

func (b *BaseRequest) NewRequest() {
	b.Payload = bytes.NewBuffer(nil)
	b.client = http.Client{}
}

func (b BaseRequest) setHeaders() (err error) {
	token, err := getAccessToken()
	if err != nil {
		return
	}
	b.request.Header.Set("Content-Type", "application/json")
	b.request.Header.Add("Authorization", "Bearer "+token)

	return
}

func (b *BaseRequest) SetBasePath(basePath string) {
	b.basePath = basePath
}

func (b BaseRequest) SendRequest(method string, path string) (response *http.Response, err error) {
	method = strings.ToUpper(method)
	firstCharacter := path[0:1]
	if firstCharacter != "/" {
		path = "/" + path
	}
	fullPath := b.basePath + path
	b.request, err = http.NewRequest(method, fullPath, b.Payload)
	if err != nil {
		return
	}
	err = b.setHeaders()
	if err != nil {
		return
	}
	response, err = b.client.Do(b.request)

	return
}
