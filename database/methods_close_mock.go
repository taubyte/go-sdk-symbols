//go:build !wasi && !wasm
// +build !wasi,!wasm

package databaseSym

import (
	"github.com/taubyte/go-sdk/errno"
)

func MockClose(expectedId uint32) {
	DatabaseClose = func(databaseId uint32) (error errno.Error) {
		if expectedId != databaseId {
			return 1
		}

		return 0
	}
}
