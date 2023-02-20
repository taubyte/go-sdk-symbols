//go:build wasi || wasm
// +build wasi wasm

package eventSym

//go:wasm-module taubyte/sdk
//export getEventType
func GetEventType(eventId uint32, typeid *uint32)
