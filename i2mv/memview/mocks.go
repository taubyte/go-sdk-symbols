//go:build !wasi && !wasm
// +build !wasi,!wasm

package memviewSym

import (
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/booleans"
)

func MockNew(testId int32) {
	MemoryViewNew = func(bufPtr *byte, size, readCloser uint32, idPtr *uint32) (error errno.Error) {
		if testId < 0 {
			return 1
		}

		*idPtr = uint32(testId)

		return 0
	}
}

func MockRead(testId uint32) {
	MemoryViewRead = func(id, offset, count uint32, bufPtr *byte, nPtr *uint32) (error errno.Error) {
		if testId != id {
			return 1
		}

		*nPtr = count

		return 0
	}
}

func MockOpen(testSize uint32, isClosable bool) {
	MemoryViewOpen = func(id uint32, isClosablePtr, sizePtr *uint32) (error errno.Error) {
		if testSize == 0 {
			return 1
		}

		*isClosablePtr = booleans.FromBool(isClosable)
		*sizePtr = testSize

		return 0
	}
}
