//go:build wasi || wasm
// +build wasi wasm

package ecdsaSym

import "github.com/taubyte/go-sdk/errno"

//go:wasm-module taubyte/sdk
//export ethHexToECDSA
func EthHexToECDSA(hexString string, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethPubKeyFromSignedMessage
func EthPubKeyFromSignedMessage(messsagePtr *byte, messageSize uint32, signaturePtr *byte, signatureSize uint32, pubKeyPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethPubFromPriv
func EthPubFromPriv(privKeyPtr *byte, privKeySize uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethSignMessage
func EthSignMessage(messagePtr *byte, messageSize uint32, privKeyPtr *byte, privKeySize uint32, signaturePtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethVerifySignature
func EthVerifySignature(messagePtr *byte, messageSize uint32, pubKeyPtr *byte, pubKeySize uint32, signaturePtr *byte, verifiedPtr *uint32) (error errno.Error)
