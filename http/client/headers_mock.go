//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpClientSym

import (
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

func MockHeaders(testClientId uint32, testRequestId uint32, testHeaders map[string][]string) {
	SetHttpRequestHeader = func(clientId uint32, requestId uint32, key string, valuesPtr *byte, valuesSize uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		data := unsafe.Slice(valuesPtr, valuesSize)

		var values []string
		err := codec.Convert(data).To(&values)
		if err != nil {
			return errno.ErrorByteConversionFailed
		}

		testHeaders[key] = values
		return 0
	}

	GetHttpRequestHeaderSize = func(clientId uint32, requestId uint32, key string, sizePtr *uint32) (error errno.Error) {
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

	GetHttpRequestHeader = func(clientId uint32, requestId uint32, key string, headerPtr *byte) (error errno.Error) {
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

	AddHttpRequestHeader = func(clientId uint32, requestId uint32, key, value string) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		headers, ok := testHeaders[key]
		if ok {
			testHeaders[key] = append(headers, value)
		} else {
			testHeaders[key] = []string{value}
		}

		return 0
	}

	GetHttpRequestHeaderKeysSize = func(clientId uint32, requestId uint32, sizePtr *uint32) (error errno.Error) {
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

	GetHttpRequestHeaderKeys = func(clientId uint32, requestId uint32, headerPtr *byte, headerSize uint32) (error errno.Error) {
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
