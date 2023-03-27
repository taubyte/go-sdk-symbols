//go:build !wasi && !wasm
// +build !wasi,!wasm

package ecdsaSym

import (
	"bytes"
	"crypto/ecdsa"
	"unsafe"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taubyte/go-sdk/errno"
)

var EthSignMessage = func(messagePtr *byte, messageSize uint32, privKeyPtr *byte, privKeySize uint32, signaturePtr *byte) (error errno.Error) {
	message := unsafe.Slice(messagePtr, messageSize)
	privKey := unsafe.Slice(privKeyPtr, privKeySize)

	pk, err := crypto.ToECDSA(privKey)
	if err != nil {
		return errno.ErrorEthereumInvalidPrivateKey
	}

	hash := crypto.Keccak256Hash([]byte(message))

	sig, err := crypto.Sign(hash.Bytes(), pk)
	if err != nil {
		return errno.ErrorEthereumSignFailed
	}

	data := unsafe.Slice(signaturePtr, len(sig))
	copy(data, sig)

	return 0
}

var EthVerifySignature = func(messagePtr *byte, messageSize uint32, pubKeyPtr *byte, pubKeySize uint32, signaturePtr *byte, verifiedPtr *uint32) (error errno.Error) {
	message := unsafe.Slice(messagePtr, messageSize)
	pubKey := unsafe.Slice(pubKeyPtr, pubKeySize)
	signature := unsafe.Slice(signaturePtr, 65)

	sigPubKey, err := crypto.Ecrecover(crypto.Keccak256Hash([]byte(message)).Bytes(), signature)
	if err != nil {
		return errno.ErrorEthereumRecoverPubKeyFailed
	}

	if bytes.Equal(sigPubKey, pubKey) {
		*verifiedPtr = 1
	}

	return 0
}

var EthHexToECDSA = func(hexString string, bufPtr *byte) (error errno.Error) {
	privKey, err := crypto.HexToECDSA(hexString)
	if err != nil {
		return errno.ErrorEthereumInvalidHexKey
	}

	privKeyBytes := privKey.D.Bytes()
	data := unsafe.Slice(bufPtr, len(privKeyBytes))
	copy(data, privKeyBytes)

	return 0
}

var EthPubFromPriv = func(privKeyPtr *byte, privKeySize uint32, bufPtr *byte) (error errno.Error) {
	pkBytes := unsafe.Slice(privKeyPtr, privKeySize)
	pk, err := crypto.ToECDSA(pkBytes)
	if err != nil {
		return errno.ErrorEthereumInvalidPrivateKey
	}

	publicKey, ok := pk.Public().(*ecdsa.PublicKey)
	if !ok {
		return errno.ErrorEthereumInvalidPublicKey
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKey)

	data := unsafe.Slice(bufPtr, len(publicKeyBytes))
	copy(data, publicKeyBytes)

	return 0
}

var EthPubKeyFromSignedMessage = func(messsagePtr *byte, messageSize uint32, signaturePtr *byte, signatureSize uint32, pubKeyPtr *byte) (error errno.Error) {
	message := unsafe.Slice(messsagePtr, messageSize)
	signature := unsafe.Slice(signaturePtr, signatureSize)

	messageHash := crypto.Keccak256Hash(message)

	publicKey, err := crypto.Ecrecover(messageHash.Bytes(), signature)
	if err != nil {
		return errno.ErrorEthereumRecoverPubKeyFailed
	}

	data := unsafe.Slice(pubKeyPtr, len(publicKey))
	copy(data, publicKey)

	return 0
}
