//go:build !wasi && !wasm
// +build !wasi,!wasm

package globalSym

import "github.com/taubyte/go-sdk/errno"

var GetGlobalValueSize = func(
	name string,
	application, function uint32,
	valueSizePtr *uint32,
) errno.Error {
	return errno.ErrorNone
}

var GetGlobalValue = func(
	name string,
	application, function uint32,
	valuePtr *byte,
) errno.Error {
	return errno.ErrorNone
}

var GetOrCreateGlobalValueSize = func(
	name string,
	application, function uint32,
	valueSizePtr *uint32,
) errno.Error {
	return errno.ErrorNone
}

var PutGlobalValue = func(
	name string,
	application, function uint32,
	value []byte,
) errno.Error {
	return errno.ErrorNone
}
