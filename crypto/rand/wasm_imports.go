//go:build wasi || wasm
// +build wasi wasm

package randSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export cryptoRead
func CryptoRead(bufPtr *byte, bufLen uint32, readPtr *uint64) (error errno.Error)
