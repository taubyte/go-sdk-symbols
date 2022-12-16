//go:build !wasi && !wasm
// +build !wasi,!wasm

package eventSym

import (
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

func getQueryBytes(testQueries map[string]string) ([]byte, errno.Error) {
	queryList := make([]string, 0)
	for key := range testQueries {
		queryList = append(queryList, key)
	}

	var queryBytes []byte
	err := codec.Convert(queryList).To(&queryBytes)
	if err != nil {
		return nil, errno.ErrorByteConversionFailed
	}

	return queryBytes, 0
}

func MockQueries(testEventId uint32, testQueries map[string]string) {
	GetHttpEventQueryValueByNameSize = func(eventId uint32, size *uint32, key string) (error errno.Error) {
		if eventId != testEventId {
			return 1
		}

		q, ok := testQueries[key]
		if ok == false {
			return 1
		}

		*size = uint32(len(q))
		return 0
	}

	GetHttpEventQueryValueByName = func(eventId uint32, key string, bufPtr *byte, bufSize uint32) (error errno.Error) {
		if eventId != testEventId {
			return 1
		}

		q, ok := testQueries[key]
		if ok == false {
			return 1
		}

		data := unsafe.Slice(bufPtr, bufSize)
		copy(data, q)
		return 0
	}

	GetHttpEventRequestQueryKeysSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
		if eventId != testEventId {
			return 1
		}

		queryBytes, err := getQueryBytes(testQueries)
		if err != 0 {
			return err
		}

		*sizePtr = uint32(len(queryBytes))
		return 0
	}

	GetHttpEventRequestQueryKeys = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		if eventId != testEventId {
			return 1
		}

		queryBytes, err := getQueryBytes(testQueries)
		if err != 0 {
			return err
		}

		dataPtr := unsafe.Slice(bufPtr, len(queryBytes))
		copy(dataPtr, queryBytes)
		return 0
	}
}
