//go:build wasi || wasm
// +build wasi wasm

package storageSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export storageNew
func StorageNew(storageName string, idPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageGet
func StorageGet(storageName string, idPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageAddFile
func StorageAddFile(storageId uint32, fileName string, versionPtr *uint32, bufPtr *byte, bufLen uint32, overWrite uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageGetFile
func StorageGetFile(storageId uint32, fileName string, version uint32, fdPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageReadFile
func StorageReadFile(storageId uint32, fd uint32, buf *byte, bufSize uint32, count *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageCloseFile
func StorageCloseFile(storageId uint32, fd uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageDeleteFile
func StorageDeleteFile(storageId uint32, fileName string, version uint32, all uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageListVersionsSize
func StorageListVersionsSize(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageListVersions
func StorageListVersions(storageId uint32, fileName string, versionPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageUsedSize
func StorageUsedSize(storageId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageUsed
func StorageUsed(storageId uint32, usedPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageCapacitySize
func StorageCapacitySize(storageId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageCapacity
func StorageCapacity(storageId uint32, capacityPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageCid
func StorageCid(storageId uint32, fileName string, cidPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageCurrentVersionSize
func StorageCurrentVersionSize(storageId uint32, fileName string, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageCurrentVersion
func StorageCurrentVersion(filename string, versionPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageListFilesSize
func StorageListFilesSize(storageId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageListFiles
func StorageListFiles(storageId uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageNewContent
func StorageNewContent(contentIdPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export storageOpenCid
func StorageOpenCid(contentIdPtr *uint32, cid *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export contentCloseFile
func ContentCloseFile(contentId uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export contentFileCid
func ContentFileCid(contentId uint32, cidPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export contentReadFile
func ContentReadFile(contentId uint32, buf *byte, bufLen uint32, countPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export contentWriteFile
func ContentWriteFile(contentId uint32, buf *byte, bufLen uint32, writePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export contentPushFile
func ContentPushFile(contentId uint32, cidPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export contentSeekFile
func ContentSeekFile(contentId uint32, offset int64, whence int, offsetPtr *int) (error errno.Error)
