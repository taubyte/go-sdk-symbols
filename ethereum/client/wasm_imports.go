//go:build wasi || wasm
// +build wasi wasm

package ethereumSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export ethBlockByNumber
func EthBlockByNumber(clientId uint32, size uint32, bufPtr *byte, blockIdPtr *uint64) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethCurrentBlockNumber
func EthCurrentBlockNumber(clientId uint32, blockNumberPtr *uint64) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethBlockNumberFromIdSize
func EthBlockNumberFromIdSize(clientId uint32, blockIdPtr *uint64, lenPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethBlockNumberFromId
func EthBlockNumberFromId(clientId uint32, blockIdPtr *uint64, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethCurrentChainIdSize
func EthCurrentChainIdSize(clientId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethCurrentChainId
func EthCurrentChainId(clientId uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethNew
func EthNew(clientIdPtr *uint32, url string) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethNonceFromPrivateKey
func EthNonceFromPrivateKey(clientId uint32, hexKey string, blockNumberLen uint32, blockNumberPtr *byte, noncePtr *uint64) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethGetTransactionFromBlockByHash
func EthGetTransactionFromBlockByHash(clientId uint32, blockIdPtr *uint64, idPtr *uint32, hashPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethGetTransactionsFromBlockSize
func EthGetTransactionsFromBlockSize(clientId uint32, blockIdPtr *uint64, sizePtr *uint32, arrSizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethGetTransactionsFromBlock
func EthGetTransactionsFromBlock(clientId uint32, blockIdPtr *uint64, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethGetTransactionMethodSize
func EthGetTransactionMethodSize(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, method string, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethGetTransactionMethodBytes
func EthGetTransactionMethodBytes(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, method string, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethTransactionRawSignaturesSize
func EthTransactionRawSignaturesSize(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, vSigSizePtr *uint32, rSigSizePtr *uint32, sSigSizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethTransactionRawSignatures
func EthTransactionRawSignatures(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, vSigBufPtr *byte, rSigBufPtr *byte, sSigBufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethSendTransaction
func EthSendTransaction(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethGetTransactionMethodUint64
func EthGetTransactionMethodUint64(clientId uint32, blockIdPtr *uint64, contractId uint32, transactionId uint32, method string, numPtr *uint64) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethNewContractSize
func EthNewContractSize(clientId uint32, abiPtr *byte, abiSize uint32, address string, methodsSizePtr *uint32, contractPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethNewContract
func EthNewContract(clientId uint32, contractId uint32, methodsPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethCallContractSize
func EthCallContractSize(clientId uint32, contractId uint32, method string, inputsPtr *byte, inputsSize uint32, outputSizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethCallContract
func EthCallContract(clientId uint32, contractId uint32, method string, outputPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethGetContractMethodSize
func EthGetContractMethodSize(clientId uint32, contractId uint32, method string, inputSizePtr *uint32, outputSizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethGetContractMethod
func EthGetContractMethod(clientId uint32, contractId uint32, method string, inputPtr *byte, outputPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethTransactContract
func EthTransactContract(clientId uint32, contractId uint32, chainIdPtr *byte, chainIdSize uint32, method string, privKey string, inputPtr *byte, inputSize uint32, transactionIdPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethCloseClient
func EthCloseClient(clientId uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethDeployContractSize
func EthDeployContractSize(clientId uint32, chainIdPtr *byte, chainIdSize uint32, bin string, abiPtr *byte, abiSize uint32, privKey string, addressPtr *byte, methodsSizePtr *uint32, contractIdPtr *uint32, transactionIdPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethSignMessage
func EthSignMessage(message string, privKey string, signaturePtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ethVerifySignature
func EthVerifySignature(message string, signaturePtr *byte, privKey string, verifiedPtr *uint32) (error errno.Error)
