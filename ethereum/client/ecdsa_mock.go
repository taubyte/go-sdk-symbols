//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func MockHexToECDSA(testString string, testPrivateKey []byte, testPrivKeySize uint32) {
	EthHexToECDSASize = func(hexString string, sizePtr *uint32) (error errno.Error) {
		if hexString != testString {
			return 1
		}

		*sizePtr = testPrivKeySize

		return
	}

	EthHexToECDSA = func(hexString string, bufPtr *byte) (error errno.Error) {
		if uint32(len(testPrivateKey)) != testPrivKeySize {
			return 1
		}

		data := unsafe.Slice(bufPtr, testPrivKeySize)
		copy(data, testPrivateKey)

		return 0
	}
}

func MockPublicKeyFromPrivate(testPrivateKey []byte) {
	EthPubFromPriv = func(privKeyPtr *byte, privKeySize uint32, bufPtr *byte) (error errno.Error) {
		if uint32(len(testPrivateKey)) != privKeySize {
			return 1
		}

		return
	}
}

func MockPublicKeyFromSignedMessage(testMessage []byte) {
	EthPubKeyFromSignedMessage = func(messsagePtr *byte, messageSize uint32, signaturePtr *byte, signatureSize uint32, pubKeyPtr *byte) (error errno.Error) {
		if uint32(len(testMessage)) != messageSize {
			return 1
		}

		return
	}
}
