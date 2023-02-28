//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpClientSym

import (
	"io"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

func MockResponse(testClientId, testRequestId uint32, responseBody io.Reader) {
	ReadHttpResponseBody = func(clientId, requestId uint32, buf *byte, bufSize uint32, countPtr *uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		_buf := unsafe.Slice(buf, bufSize)
		n, err := responseBody.Read(_buf)
		if err != nil && err != io.EOF {
			return errno.ErrorHttpReadBody
		}

		*countPtr = uint32(n)

		if err == io.EOF {
			return errno.ErrorEOF
		}

		return 0
	}

	CloseHttpResponseBody = func(clientId, requestId uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		return 0
	}

	DoHttpRequest = func(clientId, requestId uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		return 0
	}
}

func MockResponseHeaders(testClientId, testRequestId uint32, testHeaders map[string][]string) {
	MockResponseHeadersGet(testClientId, testRequestId, testHeaders)
	MockResponseHeadersGetKeys(testClientId, testRequestId, testHeaders)
}

func MockResponseHeadersGet(testClientId uint32, testRequestId uint32, testHeaders map[string][]string) {
	GetHttpResponseHeaderSize = func(clientId uint32, requestId uint32, key string, sizePtr *uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		var value []byte
		err := codec.Convert(testHeaders[key]).To(&value)
		if err != nil {
			return errno.ErrorByteConversionFailed
		}

		*sizePtr = uint32(len(value))
		return 0
	}

	GetHttpResponseHeader = func(clientId uint32, requestId uint32, key string, headerPtr *byte) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		var value []byte
		err := codec.Convert(testHeaders[key]).To(&value)
		if err != nil {
			return errno.ErrorByteConversionFailed
		}

		_data := unsafe.Slice(headerPtr, len(value))
		copy(_data, value)
		return 0
	}
}

func MockResponseHeadersGetKeys(testClientId uint32, testRequestId uint32, testHeaders map[string][]string) {
	GetHttpResponseHeaderKeysSize = func(clientId uint32, requestId uint32, sizePtr *uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		keys := make([]string, 0)
		for key := range testHeaders {
			keys = append(keys, key)
		}

		var value []byte
		err := codec.Convert(keys).To(&value)
		if err != nil {
			return errno.ErrorByteConversionFailed
		}

		*sizePtr = uint32(len(value))
		return 0
	}

	GetHttpResponseHeaderKeys = func(clientId uint32, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		keys := make([]string, 0)
		for key := range testHeaders {
			keys = append(keys, key)
		}

		var value []byte
		err := codec.Convert(keys).To(&value)
		if err != nil {
			return errno.ErrorByteConversionFailed
		}

		_data := unsafe.Slice(headerPtr, len(value))
		copy(_data, value)
		return 0
	}
}
