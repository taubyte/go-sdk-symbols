//go:build wasi || wasm
// +build wasi wasm

package pubsubSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export getMessageData
func GetMessageData(eventId uint32, buf *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getMessageDataSize
func GetMessageDataSize(eventId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getMessageChannel
func GetMessageChannel(eventId uint32, channelPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getMessageChannelSize
func GetMessageChannelSize(eventId uint32, channelSizePtr *uint32) (error errno.Error)
