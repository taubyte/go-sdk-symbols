//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpClientSym

import (
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func MockBody(m *MockData) {
	SetHttpRequestBody = func(clientId, requestId uint32, data *byte, dataSize uint32) (error errno.Error) {
		if clientId != m.ClientId || requestId != m.RequestId {
			return 1
		}

		_data := unsafe.Slice(data, dataSize)
		m.RequestBody = make([]byte, 0)
		m.RequestBody = append(m.RequestBody, _data...)
		return 0
	}
}
