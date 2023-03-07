//go:build !wasi && !wasm
// +build !wasi,!wasm

package memoryViewSym

import "github.com/taubyte/go-sdk/errno"

var MemoryViewNew = func(bufPtr *byte, size uint32, readCloser uint32, idPtr *uint32) (error errno.Error) {
	return 0
}

var MemoryViewSize = func(id uint32, isClosablePtr *uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var MemoryViewRead = func(id uint32, offset uint32, count uint32, bufPtr *byte, nPtr *uint32) (error errno.Error) {
	return 0
}

var MemoryViewClose = func(id uint32) {}
