//go:build !wasi && !wasm
// +build !wasi,!wasm

package eventSym

import (
	"github.com/taubyte/go-sdk/common"
)

func MockEventType(testId uint32, _type common.EventType) {
	GetEventType = func(eventId uint32, typeid *uint32) {
		if testId != eventId {
			*typeid = 0
			return
		}

		*typeid = uint32(_type)
	}
}
