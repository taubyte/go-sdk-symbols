//go:build !wasi && !wasm
// +build !wasi,!wasm

package databaseSym

import (
	"github.com/taubyte/go-sdk/errno"
)

func MockNew(testId uint32, testName string) {
	// Override vm method used in New
	NewDatabase = func(name string, databaseId *uint32) (error errno.Error) {
		*databaseId = uint32(testId)
		if testName != name {
			return 1
		}

		return 0
	}
}

func Mock(testId uint32, testName string, testData map[string][]byte) {
	MockClose(testId)
	MockDelete(testId, "", testData)
	MockGet(testId, testData)
	MockList(testId, "", testData)
	MockNew(testId, testName)
	MockPut(testId, testData)
}
