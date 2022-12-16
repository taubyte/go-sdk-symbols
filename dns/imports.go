//go:build !wasi && !wasm
// +build !wasi,!wasm

package dnsSym

import (
	"github.com/taubyte/go-sdk/errno"
)

var DnsNewResolver = func(resolverIdPtr *uint32) (error errno.Error) {
	return 0
}

var DnsRerouteResolver = func(resolverIdPtr uint32, address string, net string) (error errno.Error) {
	return 0
}

var DnsResetResolver = func(resolverIdPtr uint32) (error errno.Error) {
	return 0
}

var DnsLookupTxTSize = func(resolverId uint32, name string, sizePtr *uint32) (error errno.Error) {
	return 0
}

var DnsLookupTxT = func(resolverId uint32, name string, txtPtr *byte) (error errno.Error) {
	return 0
}

var DnsLookupAddressSize = func(resolverId uint32, name string, sizePtr *uint32) (error errno.Error) {
	return 0
}

var DnsLookupAddress = func(resolverId uint32, name string, addrPtr *byte) (error errno.Error) {
	return 0
}

var DnsLookupCNAMESize = func(resolverId uint32, name string, sizePtr *uint32) (error errno.Error) {
	return 0
}

var DnsLookupCNAME = func(resolverId uint32, name string, cnamePtr *byte) (error errno.Error) {
	return 0
}

var DnsLookupMXSize = func(resolverId uint32, name string, sizePtr *uint32) (error errno.Error) {
	return 0
}

var DnsLookupMX = func(resolverId uint32, name string, mxPtr *byte) (error errno.Error) {
	return 0
}
