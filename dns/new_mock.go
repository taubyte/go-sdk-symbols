//go:build !wasi && !wasm
// +build !wasi,!wasm

package dnsSym

import (
	"github.com/taubyte/go-sdk/errno"
)

func MockNew(testId uint32) {
	DnsNewResolver = func(resolverIdPtr *uint32) (error errno.Error) {
		*resolverIdPtr = testId

		return 0
	}
}

type MockData struct {
	Id uint32
	AddressData,
	CnameData,
	MxData,
	TxtData map[string][]byte
}

func (m MockData) Mock() {
	MockAddress(m.Id, m.AddressData)
	MockCname(m.Id, m.CnameData)
	MockMx(m.Id, m.MxData)
	MockNew(m.Id)
	MockTxT(m.Id, m.TxtData)
	MockReroute(m.Id, "", "")
	MockReset(m.Id)
}
