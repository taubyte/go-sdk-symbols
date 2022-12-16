//go:build wasi || wasm
// +build wasi wasm

package dnsSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export dnsNewResolver
func DnsNewResolver(resolverIdPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsRerouteResolver
func DnsRerouteResolver(resolverIdPtr uint32, address string, net string) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsResetResolver
func DnsResetResolver(resolverIdPtr uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsLookupTxTSize
func DnsLookupTxTSize(resolverId uint32, name string, sizeptr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsLookupTxT
func DnsLookupTxT(resolverId uint32, name string, txtPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsLookupAddressSize
func DnsLookupAddressSize(resolverId uint32, name string, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsLookupAddress
func DnsLookupAddress(resolverId uint32, name string, addrPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsLookupCNAMESize
func DnsLookupCNAMESize(resolverId uint32, name string, sizeptr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsLookupCNAME
func DnsLookupCNAME(resolverId uint32, name string, cnamePtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsLookupMXSize
func DnsLookupMXSize(resolverId uint32, name string, sizeptr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export dnsLookupMX
func DnsLookupMX(resolverId uint32, name string, recordPtr *byte) (error errno.Error)
