//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
	"fmt"
	"math/big"
	"math/rand"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

func MockBlockTransaction(testClientId, testTransactionId uint32) {
	EthGetTransactionFromBlockByHash = func(clientId uint32, blockIdPtr *uint64, idPtr *uint32, hashPtr *byte) (error errno.Error) {
		if clientId != testClientId {
			return 1
		}

		*idPtr = testTransactionId
		return 0
	}
}

func MockBlockTransactions(testSizeClientId, testDataClientId uint32, testTransactionIds []uint32, writeError bool) error {
	var encoded []byte
	if writeError == true {
		encoded = make([]byte, 8)
		rand.Read(encoded)
	} else {
		err := codec.Convert(testTransactionIds).To(&encoded)
		if err != nil {
			return err
		}
	}

	EthGetTransactionsFromBlockSize = func(clientId uint32, blockIdPtr *uint64, sizePtr, arraySize *uint32) (error errno.Error) {
		if clientId != testSizeClientId {
			return 1
		}

		*sizePtr = uint32(len(encoded))
		*arraySize = uint32(len(testTransactionIds))
		return 0
	}

	if writeError == true {
		encoded = []byte("Hello worlds")
	}

	EthGetTransactionsFromBlock = MockClientUint64IdBuf(testDataClientId, encoded)

	return nil
}

func MockClientUint64IdBuf(testClientId uint32, data []byte) func(clientId uint32, idPtr *uint64, bufPtr *byte) (error errno.Error) {
	return func(clientId uint32, idPtr *uint64, bufPtr *byte) (error errno.Error) {
		if testClientId != clientId {
			return 1
		}

		d := unsafe.Slice(bufPtr, len(data))
		copy(d, data)

		return 0
	}
}

func MockBlockNumber(testSizeClientId, testDataClientId uint32, blockNumber *big.Int) error {
	if blockNumber == nil {
		return fmt.Errorf("Block Number cannot be nil")
	}

	blockNumberBytes := blockNumber.Bytes()

	EthBlockNumberFromIdSize = func(clientId uint32, blockIdPtr *uint64, sizePtr *uint32) (error errno.Error) {
		if testSizeClientId != clientId {
			return 1
		}

		*sizePtr = uint32(len(blockNumberBytes))
		return 0
	}

	EthBlockNumberFromId = MockClientUint64IdBuf(testDataClientId, blockNumberBytes)

	return nil
}

func MockBlockNonce(testClientId uint32, testNonce uint64) {
	EthNonceFromPrivateKey = func(clientId uint32, hexKey string, blockNumberLen uint32, blockNumberPtr *byte, noncePtr *uint64) (error errno.Error) {
		if clientId != testClientId {
			return 1
		}

		*noncePtr = testNonce
		return 0
	}
}
