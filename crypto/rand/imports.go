//go:build !wasi && !wasm
// +build !wasi,!wasm

package randSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var CryptoRead = func(bufPtr *byte, bufLen uint32, readPtr *uint64) (error errno.Error) {
	return 0
}
