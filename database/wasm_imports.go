//go:build wasi || wasm
// +build wasi wasm

package databaseSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export newDatabase
func NewDatabase(name string, databaseId *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseGet
func DatabaseGet(databaseId uint32, key string, data *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseGetSize
func DatabaseGetSize(databaseId uint32, key string, size *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databasePut
func DatabasePut(databaseId uint32, key string, data *byte, dataSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseClose
func DatabaseClose(databaseId uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseDelete
func DatabaseDelete(databaseId uint32, key string) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseList
func DatabaseList(databaseId uint32, key string, data *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseListSize
func DatabaseListSize(databaseId uint32, key string, size *uint32) (error errno.Error)
