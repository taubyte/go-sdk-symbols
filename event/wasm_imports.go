//go:build wasi || wasm
// +build wasi wasm

package eventSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export getEventType
func GetEventType(eventId uint32, typeid *uint32)
