//go:build wasi || wasm
// +build wasi wasm

package ipfsClientSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export newIpfsClient
func NewIPFSClient(clientId *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ipfsNewContent
func IpfsNewContent(clientId uint32, contentIdPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ipfsOpenFile
func IpfsOpenFile(clientId uint32, contentIdPtr *uint32, cid *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ipfsCloseFile
func IpfsCloseFile(clientId uint32, contentId uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ipfsFileCid
func IpfsFileCid(clientId uint32, contentId uint32, cidPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ipfsReadFile
func IpfsReadFile(clientId uint32, contentId uint32, buf *byte, bufLen uint32, countPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ipfsWriteFile
func IpfsWriteFile(clientId uint32, contentId uint32, buf *byte, bufLen uint32, writePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ipfsPushFile
func IpfsPushFile(clientId uint32, contentId uint32, cidPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export ipfsSeekFile
func IpfsSeekFile(clientId uint32, contentId uint32, offset int64, whence int, offsetPtr *int) (error errno.Error)
