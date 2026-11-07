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

//go:wasm-module taubyte/sdk
//export databaseBatchNew
func DatabaseBatchNew(databaseId uint32, batchId *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseBatchPut
func DatabaseBatchPut(batchId uint32, key string, data *byte, dataSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseBatchDelete
func DatabaseBatchDelete(batchId uint32, key string) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseBatchCommit
func DatabaseBatchCommit(batchId uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseListEntries
func DatabaseListEntries(databaseId uint32, prefix string, data *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseListEntriesSize
func DatabaseListEntriesSize(databaseId uint32, prefix string, size *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseListMatch
func DatabaseListMatch(databaseId uint32, prefix string, pattern string, data *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseListMatchSize
func DatabaseListMatchSize(databaseId uint32, prefix string, pattern string, size *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseListPage
func DatabaseListPage(databaseId uint32, prefix string, cursor string, limit uint32, data *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseListPageSize
func DatabaseListPageSize(databaseId uint32, prefix string, cursor string, limit uint32, size *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export databaseCount
func DatabaseCount(databaseId uint32, prefix string, count *uint32) (error errno.Error)
