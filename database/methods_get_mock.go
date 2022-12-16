//go:build !wasi && !wasm
// +build !wasi,!wasm

package databaseSym

import (
	"github.com/taubyte/go-sdk-symbols/mocks"
)

func MockGet(id uint32, data map[string][]byte) {
	DatabaseGetSize = mocks.SizeIdKeyFunction(id, data)
	DatabaseGet = mocks.DataIdKeyFunction(id, data)
}
