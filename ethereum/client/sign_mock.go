//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
	"math/rand"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func MockSign(testMessage []byte) {
	EthSignMessage = func(messagePtr *byte, messageSize uint32, privKeyPtr *byte, privKeySize uint32, signaturePtr *byte) (error errno.Error) {
		if uint32(len(testMessage)) != messageSize {
			return 1
		}

		fakeSig := make([]byte, 65)
		rand.Read(fakeSig)

		data := unsafe.Slice(signaturePtr, len(fakeSig))
		copy(data, fakeSig)

		return 0
	}
}

func MockVerify(testMessage []byte, verified bool) {
	EthVerifySignature = func(messagePtr *byte, messageSize uint32, pubKeyPtr *byte, pubKeySize uint32, signaturePtr *byte, verifiedPtr *uint32) (error errno.Error) {
		if uint32(len(testMessage)) != messageSize {
			return 1
		}

		if verified {
			*verifiedPtr = 1
		}

		return
	}
}
