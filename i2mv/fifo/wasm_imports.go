//go:build wasi || wasm
// +build wasi wasm

package fifoSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export fifoNew
func FifoNew(readCloser uint32) (id uint32)

//go:wasm-module taubyte/sdk
//export fifoPush
func FifoPush(id uint32, buf byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export fifoPop
func FifoPop(id uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export fifoIsCloser
func FifoIsCloser(id uint32, isCloser *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export fifoClose
func FifoClose(id uint32) {}
