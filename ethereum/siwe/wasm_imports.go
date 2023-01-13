//go:build wasi || wasm
// +build wasi wasm

package siweSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export siweInitMessageLen
func SiweInitMessageLen(domain string, uri string, address string, nonce string, optionsPtr *byte, optionsLen uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export siweInitMessage
func SiweInitMessage(domain string, uri string, address string, nonce string, optionsPtr *byte, optionsLen uint32, dataPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export siweVerifyEIP191
func SiweVerifyEIP191(message string, signature string) (error errno.Error)
