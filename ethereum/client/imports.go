//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
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

var EthNonceFromPrivateKey = func(clientId uint32, hexKey string, blockNumberLen uint32, blockNumberPtr *byte, noncePtr *uint64) (error errno.Error) {
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

var EthTransactContract = func(clientId uint32, contractId uint32, chainIdPtr *byte, chainIdSize uint32, method string, privKey string, inputPtr *byte, inputSize uint32, transactionIdPtr *uint32) (error errno.Error) {
	return 0
}

var EthCloseClient = func(clientId uint32) (error errno.Error) {
	return 0
}

var EthDeployContractSize = func(clientId uint32, chainIdPtr *byte, chainIdSize uint32, bin string, abiPtr *byte, abiSize uint32, privKey string, addressPtr *byte, methodsSizePtr *uint32, contractIdPtr *uint32, transactionIdPtr *uint32) (error errno.Error) {
	return 0
}

var EthSignMessage = func(message string, privKey string, signaturePtr *byte) (error errno.Error) {
	return 0
}

var EthVerifySignature = func(message string, signaturePtr *byte, privKey string, verifiedPtr *uint32) (error errno.Error) {
	return 0
}
