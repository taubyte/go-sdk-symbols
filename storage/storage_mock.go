//go:build !wasi && !wasm
// +build !wasi,!wasm

package storageSym

import (
	"strconv"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

type StorageMockData struct {
	StorageId   uint32
	StorageName string
	FileId      uint32
	FileName    string
	Cid         string
	Files       string
	Used        int
	Capacity    int
}

func (_m StorageMockData) Mock() *StorageMockData {
	m := &_m

	m.MockNew()
	m.MockCid()
	m.MockRemainingCapacity()
	m.MockListFiles()
	return m
}

func (m StorageMockData) MockNew() {
	StorageNew = func(storageName string, idPtr *uint32) (error errno.Error) {
		if m.StorageName != storageName {
			return 1
		}

		*idPtr = uint32(m.StorageId)
		return 0
	}
}

func (m StorageMockData) MockCid() {
	StorageCidSize = func(storageId uint32, fileName string, idPtr *uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		if m.FileName != fileName {
			return 1
		}

		*idPtr = m.FileId
		return 0
	}

	StorageCid = func(cidPtr *byte, idPtr *uint32) (error errno.Error) {
		if *idPtr != m.FileId {
			return 1
		}

		data := unsafe.Slice(cidPtr, len(m.Cid))
		copy(data, []byte(m.Cid))

		return 0
	}
}

func (m StorageMockData) MockListFiles() {
	StorageListFilesSize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		*sizePtr = uint32(len(m.Files))
		return 0
	}

	StorageListFiles = func(storageId uint32, bufPtr *byte) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		data := unsafe.Slice(bufPtr, len(m.Files))
		copy(data, []byte(m.Files))

		return 0
	}
}

func (m StorageMockData) MockRemainingCapacity() {
	StorageUsedSize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
		if m.StorageId != storageId {
			return 1
		}

		*sizePtr = uint32(len(strconv.Itoa(m.Used)))
		return 0
	}

	StorageUsed = func(storageId uint32, usedPtr *byte) (error errno.Error) {
		if m.StorageId != storageId {
			return 1
		}

		data := unsafe.Slice(usedPtr, byte(m.Used))
		copy(data, []byte(strconv.Itoa(m.Used)))
		return 0
	}

	StorageCapacitySize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
		if m.StorageId != storageId {
			return 1
		}

		*sizePtr = uint32(len(strconv.Itoa(m.Capacity)))
		return 0
	}

	StorageCapacity = func(storageId uint32, capacityPtr *byte) (error errno.Error) {
		if m.StorageId != storageId {
			return 1
		}

		data := unsafe.Slice(capacityPtr, byte(m.Capacity))
		copy(data, []byte(strconv.Itoa(m.Capacity)))
		return 0
	}
}
