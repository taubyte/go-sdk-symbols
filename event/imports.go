//go:build !wasi && !wasm
// +build !wasi,!wasm

package eventSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var GetEventType = func(eventId uint32, typeid *uint32) {}

var GetHttpEventMethodSize = func(eventId uint32, size *uint32) (error errno.Error) {
	return 0
}

var GetHttpEventMethod = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
	return 0
}

var EventHttpWrite = func(eventId uint32, bufPtr *byte, bufSize uint32, n *uint32) (error errno.Error) {
	return 0
}

var EventHttpRetCode = func(eventId uint32, code uint32) (error errno.Error) {
	return 0
}

var EventHttpHeaderAdd = func(eventId uint32, key string, val string) (error errno.Error) {
	return 0
}

var ReadHttpEventBody = func(eventId uint32, buf *byte, bufSize uint32, countPtr *uint32) (error errno.Error) {
	return 0
}

var GetHttpEventHostSize = func(eventId uint32, size *uint32) (error errno.Error) {
	return 0
}

var GetHttpEventHost = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
	return 0
}

var GetHttpEventPathSize = func(eventId uint32, size *uint32) (error errno.Error) {
	return 0
}

var GetHttpEventPath = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
	return 0
}

var GetHttpEventUserAgentSize = func(eventId uint32, size *uint32) (error errno.Error) {
	return 0
}

var GetHttpEventUserAgent = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
	return 0
}

var GetHttpEventQueryValueByNameSize = func(eventId uint32, size *uint32, key string) (error errno.Error) {
	return 0
}

var GetHttpEventQueryValueByName = func(eventId uint32, key string, bufPtr *byte, bufSize uint32) (error errno.Error) {
	return 0
}

var GetHttpEventRequestQueryKeys = func(eventId uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var GetHttpEventRequestQueryKeysSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var GetHttpEventHeaderByNameSize = func(eventId uint32, size *uint32, key string) (error errno.Error) {
	return 0
}

var GetHttpEventHeaderByName = func(eventId uint32, key string, bufPtr *byte, bufSize uint32) (error errno.Error) {
	return 0
}

var GetHttpEventRequestHeaderKeys = func(eventId uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var GetHttpEventRequestHeaderKeysSize = func(eventId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var CloseHttpEventBody = func(eventId uint32) (error errno.Error) {
	return 0
}
