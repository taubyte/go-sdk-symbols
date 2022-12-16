//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
	"fmt"
	"math/big"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func MockClientNew(testClientId int32) {
	EthNew = func(clientIdPtr *uint32, url string) (error errno.Error) {
		if testClientId < 0 {
			return 1
		}

		*clientIdPtr = uint32(testClientId)

		return 0
	}
}

func MockBlockByNumber(testClientId uint32, testBlockId uint64) {
	EthBlockByNumber = func(clientId, size uint32, bufPtr *byte, blockIdPtr *uint64) (error errno.Error) {
		if clientId != testClientId {
			return 1
		}

		*blockIdPtr = testBlockId

		return 0
	}
}

func MockCurrentBlockNumber(testClientId uint32, testBlockNumber uint64) {
	EthCurrentBlockNumber = func(clientId uint32, blockNumberPtr *uint64) (error errno.Error) {
		if clientId != testClientId {
			return 1
		}

		*blockNumberPtr = testBlockNumber

		return 0
	}
}

func MockCurrentChainId(testSizeClientId, testDataClientId uint32, testChainId *big.Int) error {
	if testChainId == nil {
		return fmt.Errorf("testChainId cannot be nil")
	}

	chainIdBytes := testChainId.Bytes()
	EthCurrentChainIdSize = func(clientId uint32, sizePtr *uint32) (error errno.Error) {
		if clientId != testSizeClientId {
			return 1
		}

		*sizePtr = uint32(len(chainIdBytes))

		return 0
	}

	EthCurrentChainId = func(clientId uint32, bufPtr *byte) (error errno.Error) {
		if clientId != testDataClientId {
			return 1
		}

		data := unsafe.Slice(bufPtr, len(chainIdBytes))
		copy(data, chainIdBytes)

		return 0
	}

	return nil
}

func MockClientClose(testClientId uint32) {
	EthCloseClient = func(clientId uint32) (error errno.Error) {
		if clientId != testClientId {
			return 1
		}

		return 0
	}
}
