//go:build wasi || wasm
// +build wasi wasm

package i2mvSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export memoryViewNew
func MemoryViewNew(bufPtr *byte, size uint32, readCloser uint32, idPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export memoryViewRead
func MemoryViewRead(id uint32, offset uint32, count uint32, bufPtr *byte, nPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export memoryViewClose
func MemoryViewClose(id uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export memoryViewSize
func MemoryViewSize(id uint32, sizePtr *uint32) (error errno.Error)
