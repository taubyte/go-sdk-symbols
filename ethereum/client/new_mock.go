//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
	"math/big"
)

type MockData struct {
	Client                uint32
	BlockByNumber         uint64
	CurrentBlockNumber    uint64
	CurrentChainId        *big.Int
	BlockTransaction      uint32
	BlockTransactions     []uint32
	BlockNumber           *big.Int
	BlockNonce            uint64
	Contract              map[string]MockContractMethod
	ContractId            uint32
	ContractTransactionId uint32
	ContractAddress       string
	MethodTransactionId   uint32
	TransactionU64        uint64
	TransactionBytes      []byte
	VSig                  *big.Int
	RSig                  *big.Int
	SSig                  *big.Int
}

func (m *MockData) Mock() {
	contract := MockContract{
		Methods:              m.Contract,
		ContractSizeClientId: m.Client,
		ContractDataClientId: m.Client,
		MethodSizeClientId:   m.Client,
		MethodDataClientId:   m.Client,
		CallSizeClientId:     m.Client,
		CallDataClientId:     m.Client,
	}

	MockClientNew(int32(m.Client))
	MockBlockByNumber(m.Client, m.BlockByNumber)
	MockBlockNumber(m.Client, m.Client, m.BlockNumber)
	MockBlockTransaction(m.Client, m.BlockTransaction)
	MockBlockTransactions(m.Client, m.Client, m.BlockTransactions, false)
	MockBlockNonce(m.Client, m.BlockNonce)
	MockCurrentBlockNumber(m.Client, m.CurrentBlockNumber)
	MockCurrentChainId(m.Client, m.Client, m.CurrentChainId)
	MockDeployContract(contract, m.ContractAddress, m.ContractTransactionId, m.ContractId, false, false)
	MockNewBoundContract(contract, m.ContractId, false, false)
	MockCall(contract, false, false)
	MockTransactContract(m.Client, m.ContractTransactionId)
	MockU64method(m.Client, m.TransactionU64)
	MockBytesMethod(m.Client, m.Client, m.TransactionBytes)
	MockRawSignatures(m.Client, m.Client, m.VSig, m.RSig, m.SSig)
	MockSendTransaction(m.Client)
}
