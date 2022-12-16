//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpClientSym

import (
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/convert"
)

func MockMethod(testClientId, testRequestId uint32, setPtr *string) {
	SetHttpRequestMethod = func(clientId uint32, requestId, method uint32) errno.Error {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		*setPtr, _ = convert.MethodUintToString(method)
		return 0
	}

	GetHttpRequestMethod = func(clientId, requestId uint32, methodPtr *uint32) errno.Error {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		*methodPtr, _ = convert.MethodStringToUint(*setPtr)
		return 0
	}
}
