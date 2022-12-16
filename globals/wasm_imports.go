//go:build wasi || wasm
// +build wasi wasm

package globalSym

import "github.com/taubyte/go-sdk/errno"

//go:wasm-module taubyte/sdk
//export getGlobalValueSize
func GetGlobalValueSize(
	name string,
	application, function uint32,
	valueSizePtr *uint32,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getGlobalValue
func GetGlobalValue(
	name string,
	application, function uint32,
	valuePtr *byte,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getOrCreateGlobalValueSize
func GetOrCreateGlobalValueSize(
	name string,
	application, function uint32,
	valueSizePtr *uint32,
) (error errno.Error)

//go:wasm-module taubyte/sdk
//export putGlobalValue
func PutGlobalValue(
	name string,
	application, function uint32,
	value []byte,
) (error errno.Error)
