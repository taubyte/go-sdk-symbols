//go:build !wasi && !wasm
// +build !wasi,!wasm

package eventSym

var GetEventType = func(eventId uint32, typeid *uint32) {}
