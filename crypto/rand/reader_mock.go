//go:build !wasi && !wasm
// +build !wasi,!wasm

package randSym

import "github.com/taubyte/go-sdk/errno"

func MockReader(testBufLen uint32, testRead uint64) {
	CryptoRead = func(bufPtr *byte, bufLen uint32, readPtr *uint64) (error errno.Error) {
		if bufLen != testBufLen {
			return 1
		}

		*readPtr = testRead
		return 0
	}
}
