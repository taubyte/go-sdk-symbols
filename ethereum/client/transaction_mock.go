//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
	"fmt"
	"math/big"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func MockU64method(testClientId uint32, testValue uint64) {
	EthGetTransactionMethodUint64 = func(clientId uint32, blockIdPtr *uint64, contractId, transactionId uint32, method string, numPtr *uint64) (error errno.Error) {
		if clientId != testClientId {
			return 1
		}

		*numPtr = testValue
		return 0
	}
}

func MockBytesMethod(testSizeClientId, testBytesClientId uint32, testBytes []byte) {
	EthGetTransactionMethodSize = func(clientId uint32, blockIdPtr *uint64, contractId, transactionId uint32, method string, sizePtr *uint32) (error errno.Error) {
		if clientId != testSizeClientId {
			return 1
		}

		*sizePtr = uint32(len(testBytes))
		return 0
	}

	EthGetTransactionMethodBytes = func(clientId uint32, blockIdPtr *uint64, contractId, transactionId uint32, method string, bufPtr *byte) (error errno.Error) {
		if clientId != testBytesClientId {
			return 1
		}

		data := unsafe.Slice(bufPtr, len(testBytes))
		copy(data, testBytes)
		return 0
	}
}

func MockRawSignatures(testSizeClientId, testBytesClientId uint32, v, r, s *big.Int) error {
	if v == nil || r == nil || s == nil {
		return fmt.Errorf("v,r,s cannot be nil")
	}
	vBytes, rBytes, sBytes := v.Bytes(), r.Bytes(), s.Bytes()

	EthTransactionRawSignaturesSize = func(clientId uint32, blockIdPtr *uint64, contractId, transactionId uint32, vSigSizePtr, rSigSizePtr, sSigSizePtr *uint32) (error errno.Error) {
		if clientId != testSizeClientId {
			return 1
		}

		*vSigSizePtr, *rSigSizePtr, *sSigSizePtr = uint32(len(vBytes)), uint32(len(rBytes)), uint32(len(sBytes))
		return 0
	}

	EthTransactionRawSignatures = func(clientId uint32, blockIdPtr *uint64, contractId, transactionId uint32, vSigBufPtr, rSigBufPtr, sSigBufPtr *byte) (error errno.Error) {
		if clientId != testBytesClientId {
			return 1
		}

		data := unsafe.Slice(vSigBufPtr, len(vBytes))
		copy(data, vBytes)

		data = unsafe.Slice(rSigBufPtr, len(rBytes))
		copy(data, rBytes)

		data = unsafe.Slice(sSigBufPtr, len(sBytes))
		copy(data, sBytes)
		return 0
	}

	return nil
}

func MockSendTransaction(testClientId uint32) {
	EthSendTransaction = func(clientId uint32, blockIdPtr *uint64, contractId, transactionId uint32) (error errno.Error) {
		if clientId != testClientId {
			return 1
		}

		return 0
	}
}
