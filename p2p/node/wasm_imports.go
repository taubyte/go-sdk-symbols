//go:build wasi || wasm
// +build wasi wasm

package nodeSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export newCommand
func NewCommand(protocol, command string, id *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export sendCommand
func SendCommand(id uint32, data *byte, dataSize uint32, responseSize *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export readCommandResponse
func ReadCommandResponse(id uint32, data *byte, dataSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export listenToProtocolSize
func ListenToProtocolSize(protocol string, responseSize *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export listenToProtocol
func ListenToProtocol(protocol string, response *byte, responseSize uint32) (error errno.Error)
