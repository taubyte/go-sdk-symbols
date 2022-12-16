//go:build !wasi && !wasm
// +build !wasi,!wasm

package ipfsClientSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var NewIPFSClient = func(clientId *uint32) (error errno.Error) {
	return 0
}

var IpfsNewContent = func(clientId uint32, contentIdPtr *uint32) (error errno.Error) {
	return 0
}

var IpfsOpenFile = func(clientId uint32, contentIdPtr *uint32, cid *byte, cidSize uint32) (error errno.Error) {
	return 0
}

var IpfsCloseFile = func(clientId uint32, contentId uint32) (error errno.Error) {
	return 0
}

var IpfsFileCid = func(clientId uint32, contentId uint32, cidPtr *byte) (error errno.Error) {
	return 0
}

var IpfsReadFile = func(clientId uint32, contentId uint32, buf *byte, bufLen uint32, countPtr *uint32) (error errno.Error) {
	return 0
}

var IpfsWriteFile = func(clientId uint32, contentId uint32, buf *byte, bufLen uint32, writePtr *uint32) (error errno.Error) {
	return 0
}

var IpfsPushFile = func(clientId uint32, contentId uint32, cidPtr *byte) (error errno.Error) {
	return 0
}
var IpfsSeekFile = func(clientId uint32, contentId uint32, offset int64, whence int, offsetPtr *int) (error errno.Error) {
	return 0
}
