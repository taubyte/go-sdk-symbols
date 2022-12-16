//go:build !wasi && !wasm
// +build !wasi,!wasm

package databaseSym

import (
	"strings"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

func MockList(testId uint32, expectedKey string, testData map[string][]byte) {
	DatabaseListSize = func(databaseId uint32, key string, size *uint32) (error errno.Error) {
		// Only check expectedKey if provided
		if expectedKey != "" && key != expectedKey {
			return 1
		}

		if databaseId != testId {
			return 1
		}

		keys := []string{}
		for _key := range testData {
			if strings.HasPrefix(_key, key) {
				keys = append(keys, _key)
			}
		}

		var bytes []byte
		err := codec.Convert(keys).To(&bytes)
		if err != nil {
			return errno.ErrorByteConversionFailed
		}

		*size = uint32(len(bytes))
		return 0
	}

	DatabaseList = func(databaseId uint32, key string, data *byte) (error errno.Error) {
		keys := []string{}
		for _key := range testData {
			if strings.HasPrefix(_key, key) {
				keys = append(keys, _key)
			}
		}

		var bytes []byte
		err := codec.Convert(keys).To(&bytes)
		if err != nil {
			return errno.ErrorByteConversionFailed
		}

		dataPtr := unsafe.Slice(data, len(bytes))
		copy(dataPtr, bytes)

		return 0
	}
}
