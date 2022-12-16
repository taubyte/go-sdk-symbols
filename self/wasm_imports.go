//go:build wasi || wasm
// +build wasi wasm

package selfSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export selfApplicationSize
func ApplicationSize(sizePtr *uint32) errno.Error

//go:wasm-module taubyte/sdk
//export selfApplication
func Application(ptr *byte) errno.Error

//go:wasm-module taubyte/sdk
//export selfProjectSize
func ProjectSize(sizePtr *uint32) errno.Error

//go:wasm-module taubyte/sdk
//export selfProject
func Project(ptr *byte) errno.Error

//go:wasm-module taubyte/sdk
//export selfIdSize
func IdSize(sizePtr *uint32) errno.Error

//go:wasm-module taubyte/sdk
//export selfId
func Id(ptr *byte) errno.Error

//go:wasm-module taubyte/sdk
//export selfCommitSize
func CommitSize(sizePtr *uint32) errno.Error

//go:wasm-module taubyte/sdk
//export selfCommit
func Commit(ptr *byte) errno.Error

//go:wasm-module taubyte/sdk
//export selfBranchSize
func BranchSize(sizePtr *uint32) errno.Error

//go:wasm-module taubyte/sdk
//export selfBranch
func Branch(ptr *byte) errno.Error
