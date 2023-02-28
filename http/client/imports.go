//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpClientSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var NewHttpClient = func(clientId *uint32) (error errno.Error) {
	return 0
}

var NewHttpRequest = func(clientId uint32, requestIdPtr *uint32) (error errno.Error) {
	return 0
}

var GetHttpRequestHeaderSize = func(clientId uint32, requestId uint32,
	key string,
	sizePtr *uint32,
) (error errno.Error) {
	return 0
}

var GetHttpRequestHeader = func(clientId uint32, requestId uint32,
	key string,
	headerPtr *byte,
) (error errno.Error) {
	return 0
}

var SetHttpRequestHeader = func(clientId uint32, requestId uint32,
	key string,
	valuesPtr *byte, valuesSize uint32,
) (error errno.Error) {
	return 0
}

var DeleteHttpRequestHeader = func(clientId uint32, requestId uint32,
	key string,
) (error errno.Error) {
	return 0
}

var AddHttpRequestHeader = func(clientId uint32, requestId uint32,
	key string,
	value string,
) (error errno.Error) {
	return 0
}

var GetHttpRequestHeaderKeysSize = func(clientId uint32, requestId uint32,
	sizePtr *uint32,
) (error errno.Error) {
	return 0
}

var GetHttpRequestHeaderKeys = func(clientId uint32, requestId uint32,
	headerPtr *byte, headerSize uint32,
) (error errno.Error) {
	return 0
}

var SetHttpRequestMethod = func(clientId uint32, requestId uint32, method uint32) (error errno.Error) {
	return 0
}

var GetHttpRequestMethod = func(clientId uint32, requestId uint32, methodPtr *uint32) (error errno.Error) {
	return 0
}

var SetHttpRequestBody = func(clientId uint32, requestId uint32, data *byte, dataSize uint32) (error errno.Error) {
	return 0
}

var SetHttpRequestURL = func(clientId uint32, requestId uint32, url string) (error errno.Error) {
	return 0
}

var DoHttpRequest = func(clientId uint32, requestId uint32) (error errno.Error) {
	return 0
}

var ReadHttpResponseBody = func(clientId uint32, requestId uint32, buf *byte, bufSize uint32, countPtr *uint32) (error errno.Error) {
	return 0
}

var CloseHttpResponseBody = func(clientId uint32, requestId uint32) (error errno.Error) {
	return 0
}

var GetHttpResponseHeaderSize = func(clientId uint32, requestId uint32,
	key string,
	sizePtr *uint32,
) (error errno.Error) {
	return 0
}

var GetHttpResponseHeader = func(clientId uint32, requestId uint32,
	key string,
	headerPtr *byte,
) (error errno.Error) {
	return 0
}

var GetHttpResponseHeaderKeysSize = func(clientId uint32, requestId uint32,
	sizePtr *uint32,
) (error errno.Error) {
	return 0
}

var GetHttpResponseHeaderKeys = func(clientId uint32, requestId uint32,
	headerPtr *byte, headerSize uint32,
) (error errno.Error) {
	return 0
}