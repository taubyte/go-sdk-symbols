//go:build wasi || wasm
// +build wasi wasm

package httpClientSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export newHttpClient
func NewHttpClient(clientId *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export newHttpRequest
func NewHttpRequest(clientId uint32, requestIdPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpRequestHeaderSize
func GetHttpRequestHeaderSize(clientId uint32, requestId uint32,
	key string,
	sizePtr *uint32,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpRequestHeader
func GetHttpRequestHeader(clientId uint32, requestId uint32,
	key string,
	headerPtr *byte,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export setHttpRequestHeader
func SetHttpRequestHeader(clientId uint32, requestId uint32,
	key string,
	valuesPtr *byte, valuesSize uint32,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export deleteHttpRequestHeader
func DeleteHttpRequestHeader(clientId uint32, requestId uint32,
	key string,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export addHttpRequestHeader
func AddHttpRequestHeader(clientId uint32, requestId uint32,
	key string,
	value string,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpRequestHeaderKeysSize
func GetHttpRequestHeaderKeysSize(clientId uint32, requestId uint32,
	sizePtr *uint32,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpRequestHeaderKeys
func GetHttpRequestHeaderKeys(clientId uint32, requestId uint32,
	headerPtr *byte, headerSize uint32,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export setHttpRequestMethod
func SetHttpRequestMethod(clientId uint32, requestId uint32, method uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpRequestMethod
func GetHttpRequestMethod(clientId uint32, requestId uint32, methodPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export setHttpRequestBody
func SetHttpRequestBody(clientId uint32, requestId uint32, data *byte, dataSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export setHttpRequestURL
func SetHttpRequestURL(clientId uint32, requestId uint32, url string) (error errno.Error)

//go:wasm-module taubyte/sdk
//export doHttpRequest
func DoHttpRequest(clientId uint32, requestId uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export readHttpResponseBody
func ReadHttpResponseBody(clientId uint32, requestId uint32, buf *byte, bufSize uint32, countPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export closeHttpResponseBody
func CloseHttpResponseBody(clientId uint32, requestId uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpResponseHeaderSize
func GetHttpResponseHeaderSize(clientId uint32, requestId uint32,
	key string,
	sizePtr *uint32,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpResponseHeader
func GetHttpResponseHeader(clientId uint32, requestId uint32,
	key string,
	headerPtr *byte,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpResponseHeaderKeysSize
func GetHttpResponseHeaderKeysSize(clientId uint32, requestId uint32,
	sizePtr *uint32,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpResponseHeaderKeys
func GetHttpResponseHeaderKeys(clientId uint32, requestId uint32,
	headerPtr *byte, headerSize uint32,
) (error errno.Error)