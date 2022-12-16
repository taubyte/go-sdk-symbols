//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
	"math/rand"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func MockSign(testMessage string) {
	EthSignMessage = func(message, privKey string, signaturePtr *byte) (error errno.Error) {
		if message != testMessage {
			return 1
		}

		fakeSig := make([]byte, 65)
		rand.Read(fakeSig)

		data := unsafe.Slice(signaturePtr, len(fakeSig))
		copy(data, fakeSig)

		return 0
	}
}

func MockVerify(testMessage string, verified bool) {
	EthVerifySignature = func(message string, signaturePtr *byte, privKey string, verifiedPtr *uint32) (error errno.Error) {
		if message != testMessage {
			return 1
		}

		if verified {
			*verifiedPtr = 1
		}

		return
	}
}
