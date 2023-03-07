//go:build !wasi && !wasm
// +build !wasi,!wasm

package fifoSym

import (
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/booleans"
)

func MockAll(testId uint32, closable bool, fakeData []byte) {
	MockNew(testId)
	MockPop(testId, fakeData)
	MockPush(testId)
	MockOpen(testId, closable)
}

func MockNew(testId uint32) {
	FifoNew = func(readCloser uint32) (id uint32) {
		return testId
	}
}

func MockPop(testId uint32, fakeData []byte) {
	var idx int
	FifoPop = func(id uint32, bufPtr *byte) (error errno.Error) {
		if id != testId {
			return 1
		}

		if idx == len(fakeData) {
			return errno.ErrorEOF
		}

		*bufPtr = fakeData[idx]
		idx++
		return 0
	}
}

func MockPush(testId uint32) {
	FifoPush = func(id uint32, buf byte) (error errno.Error) {
		if id != testId {
			return 1
		}

		return 0
	}
}

func MockOpen(testId uint32, isClosable bool) {
	FifoIsCloser = func(id uint32, isCloser *uint32) (error errno.Error) {
		if testId != id {
			return 1
		}

		*isCloser = booleans.FromBool(isClosable)

		return 0
	}
}
