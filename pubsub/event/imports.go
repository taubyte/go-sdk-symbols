//go:build !wasi && !wasm
// +build !wasi,!wasm

package pubsubSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var GetMessageData = func(eventId uint32, buf *byte) (error errno.Error) {
	return 0
}

var GetMessageDataSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var GetMessageChannel = func(eventId uint32, channelPtr *byte) (error errno.Error) {
	return 0
}

var GetMessageChannelSize = func(eventId uint32, channelSizePtr *uint32) (error errno.Error) {
	return 0
}
