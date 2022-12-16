//go:build !wasi && !wasm
// +build !wasi,!wasm

package storageSym

import (
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

type StorageFileMockData struct {
	StorageId   uint32
	StorageName string
	FileId      uint32
	Data        []byte
}

func (_m StorageFileMockData) Mock() *StorageFileMockData {
	m := &_m

	m.MockNew()
	m.MockRead()
	m.MockClose()
	return m
}

func (m StorageFileMockData) MockNew() {
	StorageNew = func(storageName string, idPtr *uint32) (error errno.Error) {
		*idPtr = uint32(m.StorageId)
		if m.StorageName != storageName {
			return 1
		}

		return 0
	}
}

func (m StorageFileMockData) MockClose() {
	StorageCloseFile = func(storageId, fd uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		return 0
	}
}

func (m StorageFileMockData) MockRead() {
	StorageReadFile = func(storageId, fd uint32, buf *byte, bufSize uint32, count *uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		data := unsafe.Slice(buf, bufSize)
		copy(data, []byte(m.Data))
		*count = uint32(len(data))

		return 0
	}
}
