//go:build !wasi && !wasm
// +build !wasi,!wasm

package storageSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var StorageNew = func(storageName string, idPtr *uint32) (error errno.Error) {
	return 0
}

var StorageGet = func(storageName string, idPtr *uint32) (error errno.Error) {
	return 0
}

var StorageAddFile = func(storageId uint32, fileName string, versionPtr *uint32, bufPtr *byte, bufLen uint32, overWrite uint32) (error errno.Error) {
	return 0
}

var StorageGetFile = func(storageId uint32, fileName string, version uint32, fdPtr *uint32) (error errno.Error) {
	return 0
}

var StorageReadFile = func(storageId uint32, fd uint32, buf *byte, bufSize uint32, count *uint32) (error errno.Error) {
	return 0
}

var StorageCloseFile = func(storageId uint32, fd uint32) (error errno.Error) {
	return 0
}

var StorageDeleteFile = func(storageId uint32, fileName string, version uint32, all uint32) (error errno.Error) {
	return 0
}

var StorageListVersionsSize = func(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error) {
	return 0
}

var StorageListVersions = func(storageId uint32, fileName string, versionPtr *byte) (error errno.Error) {
	return 0
}

var StorageUsedSize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var StorageUsed = func(storageId uint32, usedPtr *byte) (error errno.Error) {
	return 0
}

var StorageCapacitySize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var StorageCapacity = func(storageId uint32, capacityPtr *byte) (error errno.Error) {
	return 0
}

var StorageCid = func(storageId uint32, fileName string, cidPtr *byte) (error errno.Error) {
	return 0
}

var StorageCurrentVersionSize = func(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error) {
	return 0
}

var StorageCurrentVersion = func(filename string, versionPtr *byte) (error errno.Error) {
	return 0
}

var StorageListFilesSize = func(storageId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var StorageListFiles = func(storageId uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var StorageNewContent = func(contentIdPtr *uint32) (error errno.Error) {
	return 0
}

var StorageOpenCid = func(contentIdPtr *uint32, cid *byte, cidSize uint32) (error errno.Error) {
	return 0
}

var ContentCloseFile = func(contentId uint32) (error errno.Error) {
	return 0
}

var ContentFileCid = func(contentId uint32, cidPtr *byte) (error errno.Error) {
	return 0
}

var ContentReadFile = func(contentId uint32, buf *byte, bufLen uint32, countPtr *uint32) (error errno.Error) {
	return 0
}

var ContentWriteFile = func(contentId uint32, buf *byte, bufLen uint32, writePtr *uint32) (error errno.Error) {
	return 0
}

var ContentPushFile = func(contentId uint32, cidPtr *byte) (error errno.Error) {
	return 0
}

var ContentSeekFile = func(contentId uint32, offset int64, whence int, offsetPtr *int) (error errno.Error) {
	return 0
}
