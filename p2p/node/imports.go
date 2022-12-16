//go:build !wasi && !wasm
// +build !wasi,!wasm

package nodeSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var NewCommand = func(protocol, command string, id *uint32) (error errno.Error) {
	return 0
}

var SendCommand = func(id uint32, data *byte, dataSize uint32, responseSize *uint32) (error errno.Error) {
	return 0
}

var ReadCommandResponse = func(id uint32, data *byte, dataSize uint32) (error errno.Error) {
	return 0
}

var ListenToProtocol = func(protocol string, response *byte, responseSize uint32) (error errno.Error) {
	return 0
}

var ListenToProtocolSize = func(protocol string, responseSize *uint32) (error errno.Error) {
	return 0
}
