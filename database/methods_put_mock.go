//go:build !wasi && !wasm
// +build !wasi,!wasm

package databaseSym

import (
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func MockPut(testId uint32, putData map[string][]byte) {
	DatabasePut = func(databaseId uint32, key string, data *byte, dataSize uint32) (error errno.Error) {
		if databaseId != testId {
			return 1
		}

		putData[key] = unsafe.Slice(data, dataSize)
		return 0
	}
}
