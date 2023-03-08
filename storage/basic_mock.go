//go:build !wasi && !wasm
// +build !wasi,!wasm

package storageSym

import (
	"unsafe"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-sdk/errno"
)

func MockNew(testId uint32, testName string) {
	StorageNew = func(storageName string, idPtr *uint32) (error errno.Error) {
		if testName != storageName {
			return 1
		}

		*idPtr = uint32(testId)
		return 0
	}
}

func MockCreate(testId uint32) {
	StorageNewContent = func(contentIdPtr *uint32) (error errno.Error) {
		*contentIdPtr = testId
		return 0
	}
}

func MockGet(testData map[string]uint32, expectedCid uint32) {
	StorageGet = func(storageName string, idPtr *uint32) (error errno.Error) {
		var ok bool
		*idPtr, ok = testData[storageName]
		if ok == false {
			return 1

		}

		if *idPtr != expectedCid {
			return 1
		}

		return 0
	}
}

func MockOpen(testId uint32, expectedCid string) {
	StorageOpenCid = func(contentIdPtr *uint32, _cid *byte) (error errno.Error) {
		cidBytes := unsafe.Slice(_cid, 64)
		_, testCid, err := cid.CidFromBytes(cidBytes)
		if err != nil {
			return 1
		}

		if testCid.String() != expectedCid {
			return 1
		}

		*contentIdPtr = testId
		return 0
	}
}
