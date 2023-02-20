//go:build wasi || wasm
// +build wasi wasm

package pubsubSym

import (
	"github.com/taubyte/go-sdk/errno"
)

//go:wasm-module taubyte/sdk
//export setSubscriptionChannel
func SetSubscriptionChannel(channel string) (error errno.Error)

//go:wasm-module taubyte/sdk
//export publishToChannel
func PublishToChannel(channel string, data *byte, dataSize uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getWebSocketURLSize
func GetWebSocketURLSize(channel string, sizePtr *uint32) (error errno.Error)

//go:wasm-module taubyte/sdk
//export getWebSocketURL
func GetWebSocketURL(channel string, socketURLPtr *byte) (error errno.Error)
