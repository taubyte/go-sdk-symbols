//go:build !wasi && !wasm
// +build !wasi,!wasm

package selfSym

import "github.com/taubyte/go-sdk/errno"

var ApplicationSize = func(sizePtr *uint32) errno.Error {
	return 0
}

var Application = func(ptr *byte) errno.Error {
	return 0
}

var ProjectSize = func(sizePtr *uint32) errno.Error {
	return 0
}

var Project = func(ptr *byte) errno.Error {
	return 0
}

var IdSize = func(sizePtr *uint32) errno.Error {
	return 0
}

var Id = func(ptr *byte) errno.Error {
	return 0
}

var CommitSize = func(sizePtr *uint32) errno.Error {
	return 0
}

var Commit = func(ptr *byte) errno.Error {
	return 0
}

var BranchSize = func(sizePtr *uint32) errno.Error {
	return 0
}

var Branch = func(ptr *byte) errno.Error {
	return 0
}
