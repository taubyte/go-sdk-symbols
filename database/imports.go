//go:build !wasi && !wasm
// +build !wasi,!wasm

package databaseSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var (
	NewDatabase = func(name string, databaseId *uint32) (error errno.Error) {
		return 0
	}

	DatabaseGet = func(databaseId uint32, key string, data *byte) (error errno.Error) {
		return 0
	}

	DatabaseGetSize = func(databaseId uint32, key string, size *uint32) (error errno.Error) {
		return 0
	}

	DatabasePut = func(databaseId uint32, key string, data *byte, dataSize uint32) (error errno.Error) {
		return 0
	}

	DatabaseClose = func(databaseId uint32) (error errno.Error) {
		return 0
	}

	DatabaseDelete = func(databaseId uint32, key string) (error errno.Error) {
		return 0
	}

	DatabaseList = func(databaseId uint32, key string, data *byte) (error errno.Error) {
		return 0
	}

	DatabaseListSize = func(databaseId uint32, key string, size *uint32) (error errno.Error) {
		return 0
	}

	DatabaseBatchNew = func(databaseId uint32, batchId *uint32) (error errno.Error) {
		return 0
	}

	DatabaseBatchPut = func(batchId uint32, key string, data *byte, dataSize uint32) (error errno.Error) {
		return 0
	}

	DatabaseBatchDelete = func(batchId uint32, key string) (error errno.Error) {
		return 0
	}

	DatabaseBatchCommit = func(batchId uint32) (error errno.Error) {
		return 0
	}

	DatabaseListEntries = func(databaseId uint32, prefix string, data *byte) (error errno.Error) {
		return 0
	}

	DatabaseListEntriesSize = func(databaseId uint32, prefix string, size *uint32) (error errno.Error) {
		return 0
	}

	DatabaseListMatch = func(databaseId uint32, prefix string, pattern string, data *byte) (error errno.Error) {
		return 0
	}

	DatabaseListMatchSize = func(databaseId uint32, prefix string, pattern string, size *uint32) (error errno.Error) {
		return 0
	}

	DatabaseListPage = func(databaseId uint32, prefix string, cursor string, limit uint32, data *byte) (error errno.Error) {
		return 0
	}

	DatabaseListPageSize = func(databaseId uint32, prefix string, cursor string, limit uint32, size *uint32) (error errno.Error) {
		return 0
	}

	DatabaseCount = func(databaseId uint32, prefix string, count *uint32) (error errno.Error) {
		return 0
	}
)
