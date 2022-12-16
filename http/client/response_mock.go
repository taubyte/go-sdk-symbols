//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpClientSym

import (
	"io"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func MockResponse(testClientId, testRequestId uint32, responseBody io.Reader) {
	ReadHttpResponseBody = func(clientId, requestId uint32, buf *byte, bufSize uint32, countPtr *uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		_buf := unsafe.Slice(buf, bufSize)
		n, err := responseBody.Read(_buf)
		if err != nil && err != io.EOF {
			return errno.ErrorHttpReadBody
		}

		*countPtr = uint32(n)

		if err == io.EOF {
			return errno.ErrorEOF
		}

		return 0
	}

	CloseHttpResponseBody = func(clientId, requestId uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		return 0
	}

	DoHttpRequest = func(clientId, requestId uint32) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		return 0
	}
}
