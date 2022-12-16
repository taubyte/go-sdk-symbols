//go:build !wasi && !wasm
// +build !wasi,!wasm

package dnsSym

import (
	"github.com/taubyte/go-sdk-symbols/mocks"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

func MockTxT(testId uint32, data map[string][]byte) {
	DnsLookupTxTSize = mocks.SizeIdKeyFunction(testId, data)
	DnsLookupTxT = mocks.DataIdKeyFunction(testId, data)
}

func MockAddress(testId uint32, data map[string][]byte) {
	DnsLookupAddressSize = mocks.SizeIdKeyFunction(testId, data)
	DnsLookupAddress = mocks.DataIdKeyFunction(testId, data)
}

func MockCname(testId uint32, data map[string][]byte) {
	DnsLookupCNAMESize = mocks.SizeIdKeyFunction(testId, data)
	DnsLookupCNAME = mocks.DataIdKeyFunction(testId, data)
}

func MockMx(testId uint32, data map[string][]byte) {
	DnsLookupMXSize = mocks.SizeIdKeyFunction(testId, data)
	DnsLookupMX = mocks.DataIdKeyFunction(testId, data)
}

func MockReroute(testId uint32, expectedAddress string, expectedNet string) {
	DnsRerouteResolver = func(resolverId uint32, address, net string) (error errno.Error) {
		if testId != resolverId {
			return 1
		}

		// Only verify if given
		if expectedAddress != "" && address != expectedAddress {
			return 1
		}
		if expectedNet != "" && net != expectedNet {
			return 1
		}

		return 0
	}
}

func MockReset(testId uint32) {
	DnsResetResolver = func(resolverId uint32) (error errno.Error) {
		if testId != resolverId {
			return 1
		}

		return 0
	}
}

func Convert(slice []string) []byte {
	bytes := []byte{}
	codec.Convert(slice).To(&bytes)
	return bytes
}
