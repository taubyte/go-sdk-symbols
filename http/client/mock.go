//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpClientSym

import "bytes"

type MockData struct {
	ClientId      uint32
	RequestId     uint32
	RequestUrl    string
	RequestMethod *string
	RequestBody   []byte
	ResponseBody  []byte
	Headers       map[string][]string
}

func (_m MockData) Mock() *MockData {
	m := &_m

	if m.Headers == nil {
		m.Headers = make(map[string][]string)
	}
	if m.RequestMethod == nil {
		m.RequestMethod = new(string)
	}
	if m.RequestBody == nil {
		m.RequestBody = []byte{}
	}

	MockNewClient(m.ClientId)
	MockNewRequest(m.ClientId, m.RequestId)
	MockNewRequestURL(m.ClientId, m.RequestId, m.RequestUrl)
	MockHeaders(m.ClientId, m.RequestId, m.Headers)
	MockMethod(m.ClientId, m.RequestId, m.RequestMethod)
	MockBody(m)

	responseBodyReader := bytes.NewReader(m.ResponseBody)
	MockResponse(m.ClientId, m.RequestId, responseBodyReader)

	return m
}
