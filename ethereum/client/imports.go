//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
	"bytes"
	"crypto/ecdsa"
	"unsafe"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taubyte/go-sdk/errno"
)

var EthBlockByNumber = func(clientId uint32, size uint32, bufPtr *byte, blockIdPtr *uint64) (error errno.Error) {
	return 0
}

var EthCurrentBlockNumber = func(clientId uint32, blockNumberPtr *uint64) (error errno.Error) {
	return 0
}

var EthBlockNumberFromIdSize = func(clientId uint32, blockIdPtr *uint64, sizePtr *uint32) (error errno.Error) {
	return 0
}

var EthBlockNumberFromId = func(clientId uint32, blockIdPtr *uint64, bufPtr *byte) (error errno.Error) {
	return 0
}

var EthCurrentChainIdSize = func(clientId uint32, sizePtr *uint32) (error errno.Error) {
	return 0
}

var EthCurrentChainId = func(clientId uint32, bufPtr *byte) (error errno.Error) {
	return 0
}

var EthNew = func(clientIdPtr *uint32, url string) (error errno.Error) {
	return 0
}

var EthGetTransactionFromBlockByHash = func(clientId uint32, blockIdPtr *uint64, idPtr *uint32, hashPtr *byte) (error errno.Error) {
	return 0
}

var EthGetTransactionsFromBlockSize = func(clientId uint32, blockIdPtr *uint64, sizePtr *uint32, arraySize *uint32) (error errno.Error) {
	return 0
}

var EthGetTransactionsFromBlock = func(clientId uint32, blockIdPtr *uint64, bufPtr *byte) (error errno.Error) {
	return 0
}

var EthGetTransactionMethodSize = func(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, method string, sizePtr *uint32) (error errno.Error) {
	return 0
}

var EthGetTransactionMethodBytes = func(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, method string, bufPtr *byte) (error errno.Error) {
	return 0
}

var EthTransactionRawSignaturesSize = func(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, vSigSizePtr *uint32, rSigSizePtr *uint32, sSigSizePtr *uint32) (error errno.Error) {
	return 0
}

var EthTransactionRawSignatures = func(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, vSigBufPtr *byte, rSigBufPtr *byte, sSigBufPtr *byte) (error errno.Error) {
	return 0
}

var EthSendTransaction = func(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32) (error errno.Error) {
	return 0
}

var EthGetTransactionMethodUint64 = func(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, method string, numPtr *uint64) (error errno.Error) {
	return 0
}

var EthNewContractSize = func(clientId uint32, abiPtr *byte, abiSize uint32, address string, methodsSizePtr *uint32, contractPtr *uint32) (error errno.Error) {
	return 0
}

var EthNewContract = func(clientId uint32, contractId uint32, methodsPtr *byte) (error errno.Error) {
	return 0
}

var EthCallContractSize = func(clientId uint32, contractId uint32, method string, inputsPtr *byte, inputsSize uint32, outputSizePtr *uint32) (error errno.Error) {
	return 0
}

var EthCallContract = func(clientId uint32, contractId uint32, method string, outputPtr *byte) (error errno.Error) {
	return 0
}

var EthGetContractMethodSize = func(clientId uint32, contractId uint32, method string, inputSizePtr *uint32, outputSizePtr *uint32) (error errno.Error) {
	return 0
}

var EthGetContractMethod = func(clientId uint32, contractId uint32, method string, inputPtr *byte, outputPtr *byte) (error errno.Error) {
	return 0
}

var EthTransactContract = func(clientId uint32, contractId uint32, chainIdPtr *byte, chainIdSize uint32, method string, privKeyPtr *byte, privKeySize uint32, inputPtr *byte, inputSize uint32, transactionIdPtr *uint32) (error errno.Error) {
	return 0
}

var EthCloseClient = func(clientId uint32) (error errno.Error) {
	return 0
}

var EthDeployContract = func(clientId uint32, chainIdPtr *byte, chainIdSize uint32, bin string, abiPtr *byte, abiSize uint32, privKeyPtr *byte, privKeySize uint32, addressPtr *byte, methodsSizePtr *uint32, contractIdPtr *uint32, transactionIdPtr *uint32) (error errno.Error) {
	return 0
}

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
	if ok == false {
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
