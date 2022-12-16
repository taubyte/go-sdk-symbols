//go:build wasi || wasm
// +build wasi wasm

package p2pEventSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export getP2PEventData
func GetP2PEventData(eventId uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getP2PEventDataSize
func GetP2PEventDataSize(eventId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getP2PEventCommand
func GetP2PEventCommand(eventId uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getP2PEventCommandSize
func GetP2PEventCommandSize(eventId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getP2PEventFrom
func GetP2PEventFrom(eventId uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getP2PEventFromSize
func GetP2PEventFromSize(eventId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getP2PEventTo
func GetP2PEventTo(eventId uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getP2PEventToSize
func GetP2PEventToSize(eventId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getP2PEventProtocol
func GetP2PEventProtocol(eventId uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getP2PEventProtocolSize
func GetP2PEventProtocolSize(eventId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export writeP2PResponse
func WriteP2PResponse(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error)
