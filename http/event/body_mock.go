//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpEventSym

import (
	"io"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func MockBody(testEventId uint32, responseBody io.Reader) {
	ReadHttpEventBody = func(eventId uint32, buf *byte, bufSize uint32, countPtr *uint32) (error errno.Error) {
		if testEventId != eventId {
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

	CloseHttpEventBody = func(eventId uint32) (error errno.Error) {
		if testEventId != eventId {
			return 1
		}

		return 0
	}
}
