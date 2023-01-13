//go:build !wasi && !wasm
// +build !wasi,!wasm

package siweSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var SiweInitMessageLen = func(domain string, uri string, address string, nonce string, optionsPtr *byte, optionsLen uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}
var SiweInitMessage = func(domain string, uri string, address string, nonce string, optionsPtr *byte, optionsLen uint32, dataPtr *byte) (error errno.Error) {
	return 0
}
var SiweVerifyEIP191 = func(message string, signature string) (error errno.Error) {
	return 0
}
