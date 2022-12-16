//go:build !wasi && !wasm
// +build !wasi,!wasm

package databaseSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var (
	NewDatabase = func(name string, databaseId *uint32) (error errno.Error) {
		return 0
	}

	DatabaseGet = func(databaseId uint32, key string, data *byte) (error errno.Error) {
		return 0
	}

	DatabaseGetSize = func(databaseId uint32, key string, size *uint32) (error errno.Error) {
		return 0
	}

	DatabasePut = func(databaseId uint32, key string, data *byte, dataSize uint32) (error errno.Error) {
		return 0
	}

	DatabaseClose = func(databaseId uint32) (error errno.Error) {
		return 0
	}

	DatabaseDelete = func(databaseId uint32, key string) (error errno.Error) {
		return 0
	}

	DatabaseList = func(databaseId uint32, key string, data *byte) (error errno.Error) {
		return 0
	}

	DatabaseListSize = func(databaseId uint32, key string, size *uint32) (error errno.Error) {
		return 0
	}
)
