//go:build wasi || wasm
// +build wasi wasm

package eventSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export getEventType
func GetEventType(eventId uint32, typeid *uint32)

//go:wasm-module taubyte/sdk
//export getHttpEventMethodSize
func GetHttpEventMethodSize(eventId uint32, size *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventMethod
func GetHttpEventMethod(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export eventHttpWrite
func EventHttpWrite(eventId uint32, bufPtr *byte, bufSize uint32, n *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export eventHttpRetCode
func EventHttpRetCode(eventId uint32, code uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export eventHttpHeaderAdd
func EventHttpHeaderAdd(eventId uint32, key string, val string) (error errno.Error)

//go:wasm-module taubyte/sdk
//export eventHttpRedirect
func EventHttpRedirect(eventId uint32, url string, code uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export readHttpEventBody
func ReadHttpEventBody(eventId uint32, buf *byte, bufSize uint32, countPtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventHostSize
func GetHttpEventHostSize(eventId uint32, size *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventHost
func GetHttpEventHost(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventPathSize
func GetHttpEventPathSize(eventId uint32, size *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventPath
func GetHttpEventPath(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventUserAgentSize
func GetHttpEventUserAgentSize(eventId uint32, size *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventUserAgent
func GetHttpEventUserAgent(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventQueryValueByNameSize
func GetHttpEventQueryValueByNameSize(eventId uint32, size *uint32, key string) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventQueryValueByName
func GetHttpEventQueryValueByName(eventId uint32, key string, bufPtr *byte, bufSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventRequestQueryKeys
func GetHttpEventRequestQueryKeys(eventId uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventRequestQueryKeysSize
func GetHttpEventRequestQueryKeysSize(eventId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventHeadersSize
func GetHttpEventHeaderByNameSize(eventId uint32, size *uint32, key string) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventHeaders
func GetHttpEventHeaderByName(eventId uint32, key string, bufPtr *byte, bufSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventHeaderVarNameById
func GetHttpEventRequestHeaderKeys(eventId uint32, bufPtr *byte) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getHttpEventHeaderVarNameByIdSize
func GetHttpEventRequestHeaderKeysSize(eventId uint32, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export closeHttpEventBody
func CloseHttpEventBody(eventId uint32) (error errno.Error)
