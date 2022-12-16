//go:build !wasi && !wasm
// +build !wasi,!wasm

package mocks

import "github.com/taubyte/go-sdk/errno"

type idKeyPtrFunction func(id uint32, key string, sizePtr *uint32) (error errno.Error)

type idSizeFunction func(id uint32, sizePtr *uint32) (error errno.Error)
type idDataFunction func(id uint32, bufPtr *byte) (error errno.Error)
type idDataFunctionSafe func(id uint32, bufPtr *byte, bufSize uint32) (error errno.Error)
