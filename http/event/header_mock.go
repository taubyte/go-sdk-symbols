//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpEventSym

import (
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

func getHeaderBytes(testHeaders map[string]string) ([]byte, errno.Error) {
	headerList := make([]string, 0)
	for key := range testHeaders {
		headerList = append(headerList, key)
	}

	var headerBytes []byte
	err := codec.Convert(headerList).To(&headerBytes)
	if err != nil {
		return nil, errno.ErrorByteConversionFailed
	}

	return headerBytes, 0
}

func MockHeaders(testEventId uint32, testHeaders map[string]string) {
	EventHttpHeaderAdd = func(eventId uint32, key, val string) (error errno.Error) {
		if eventId != testEventId {
			return 1
		}

		testHeaders[key] = val
		return 0
	}

	GetHttpEventHeaderByNameSize = func(eventId uint32, size *uint32, key string) (error errno.Error) {
		if eventId != testEventId {
			return 1
		}

		*size = uint32(len(testHeaders[key]))
		return 0
	}

	GetHttpEventHeaderByName = func(eventId uint32, key string, bufPtr *byte, bufSize uint32) (error errno.Error) {
		if eventId != testEventId {
			return 1
		}

		_data := unsafe.Slice(bufPtr, bufSize)
		copy(_data, testHeaders[key])
		return 0
	}

	GetHttpEventRequestHeaderKeysSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
		if eventId != testEventId {
			return 1
		}

		headerBytes, err := getHeaderBytes(testHeaders)
		if err != 0 {
			return err
		}

		*sizePtr = uint32(len(headerBytes))
		return 0
	}

	GetHttpEventRequestHeaderKeys = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		if eventId != testEventId {
			return 1
		}

		headerBytes, err := getHeaderBytes(testHeaders)
		if err != 0 {
			return err
		}

		dataPtr := unsafe.Slice(bufPtr, len(headerBytes))
		copy(dataPtr, headerBytes)

		return 0
	}
}
