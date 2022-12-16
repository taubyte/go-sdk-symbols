//go:build !wasi && !wasm
// +build !wasi,!wasm

package storageSym

import (
	"bytes"
	"strconv"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"

	"github.com/taubyte/go-sdk/utils/codec"
)

type FileMockData struct {
	StorageId   uint32
	StorageName string
	Id          uint32
	FileName    string
	Version     int
	FileId      uint32
	Versions    map[string][]string
	Overwrite   bool
	Read        int
	Data        []byte
}

func (_m FileMockData) Mock() *FileMockData {
	m := &_m
	m.MockNew()
	m.MockGetFile()
	m.MockCurrentVersion()
	m.MockAdd()
	m.MockVersions()
	m.MockDelete()
	m.MockDeleteAllVersions()
	return m
}

func (m FileMockData) MockNew() {
	StorageNew = func(storageName string, idPtr *uint32) (error errno.Error) {
		if m.StorageName != storageName {
			return 1
		}

		*idPtr = uint32(m.StorageId)
		return 0
	}
}

func (m FileMockData) MockGetFile() {
	StorageGetFile = func(storageId uint32, fileName string, version uint32, fdPtr *uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		if fileName != m.FileName {
			return 1
		}

		*fdPtr = m.FileId
		return 0
	}
}

func (m FileMockData) MockCurrentVersion() {
	StorageCurrentVersionSize = func(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		if fileName != m.FileName {
			return 1
		}

		*sizePtr = uint32(len(strconv.Itoa(m.Version)))
		return 0
	}

	StorageCurrentVersion = func(filename string, versionPtr *byte) (error errno.Error) {
		if filename != m.FileName {
			return 1
		}

		data := unsafe.Slice(versionPtr, byte(m.Version))
		copy(data, []byte(strconv.Itoa(m.Version)))
		return 0
	}
}

func (m FileMockData) MockAdd() {
	StorageAddFile = func(storageId uint32, fileName string, versionPtr *uint32, bufPtr *byte, bufLen, overWrite uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		if fileName != m.FileName {
			return 1
		}

		data := unsafe.Slice(bufPtr, bufLen)
		if bytes.Compare(data, m.Data) != 0 {
			return 1
		}

		if overWrite == 0 {
			*versionPtr = uint32(m.Version)
		} else {
			*versionPtr = uint32(m.Version + 1)
		}

		return 0
	}
}

func (m FileMockData) MockDelete() {
	StorageDeleteFile = func(storageId uint32, fileName string, version, all uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		return 0
	}
}

func (m FileMockData) MockDeleteAllVersions() {
	StorageDeleteFile = func(storageId uint32, fileName string, version, all uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		return 0
	}
}

func (m FileMockData) MockVersions() {
	StorageListVersionsSize = func(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		if m.FileName != fileName {
			return 1
		}

		versions, ok := m.Versions[fileName]
		if ok == false {
			return 1
		}

		var encoded []byte
		err := codec.Convert(versions).To(&encoded)
		if err != nil {
			return 1
		}

		*sizePtr = uint32(len(encoded))
		return 0
	}

	StorageListVersions = func(storageId uint32, fileName string, versionPtr *byte) (error errno.Error) {
		if storageId != m.StorageId {
			return 1
		}

		if m.FileName != fileName {
			return 1
		}

		versions, ok := m.Versions[fileName]
		if ok == false {
			return 1
		}

		var encoded []byte
		err := codec.Convert(versions).To(&encoded)
		if err != nil {
			return 1
		}

		data := unsafe.Slice(versionPtr, len(encoded))
		copy(data, encoded)

		return 0
	}

}
