//go:build !wasi && !wasm
// +build !wasi,!wasm

package databaseSym

import (
	"github.com/taubyte/go-sdk/errno"
)

func MockDelete(expectedId uint32, expectedKey string, expectedData map[string][]byte) {
	DatabaseDelete = func(databaseId uint32, key string) (error errno.Error) {
		// Only check expectedKey if provided
		if expectedKey != "" && key != expectedKey {
			return 1
		}
		if databaseId != expectedId {
			return 1
		}

		delete(expectedData, key)

		return 0
	}
}
