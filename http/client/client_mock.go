//go:build !wasi && !wasm
// +build !wasi,!wasm

package httpClientSym

import "github.com/taubyte/go-sdk/errno"

func MockNewClient(testId uint32) {
	NewHttpClient = func(clientId *uint32) (error errno.Error) {
		*clientId = testId
		return 0
	}
}

func MockNewRequest(testClientId uint32, testRequestId uint32) {
	NewHttpRequest = func(clientId uint32, requestIdPtr *uint32) (error errno.Error) {
		if clientId != testClientId {
			return 1
		}

		*requestIdPtr = testRequestId
		return 0
	}
}

func MockNewRequestURL(testClientId uint32, testRequestId uint32, testURL string) {
	SetHttpRequestURL = func(clientId, requestId uint32, url string) (error errno.Error) {
		if clientId != testClientId || requestId != testRequestId {
			return 1
		}

		if testURL != "" && url != testURL {
			return 1
		}

		return 0
	}
}
