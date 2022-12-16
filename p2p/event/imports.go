//go:build !wasi && !wasm
// +build !wasi,!wasm

package p2pEventSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var GetP2PEventData = func(eventId uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var GetP2PEventDataSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var GetP2PEventCommand = func(eventId uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var GetP2PEventCommandSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var GetP2PEventProtocol = func(eventId uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var GetP2PEventProtocolSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var GetP2PEventFrom = func(eventId uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var GetP2PEventFromSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var GetP2PEventTo = func(eventId uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var GetP2PEventToSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var WriteP2PResponse = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
	return 0
}
