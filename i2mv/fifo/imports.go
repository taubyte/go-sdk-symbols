//go:build !wasi && !wasm
// +build !wasi,!wasm

package fifoSym

import "github.com/taubyte/go-sdk/errno"

var FifoNew = func(readCloser uint32) (id uint32) {
	return 0
}

var FifoPush = func(id uint32, buf byte) (error errno.Error) {
	return 0
}

var FifoPop = func(id uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var FifoIsCloser = func(id uint32, isCloser *uint32) (error errno.Error) {
	return 0
}

var FifoClose = func(id uint32) {}
