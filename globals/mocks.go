//go:build !wasi && !wasm
// +build !wasi,!wasm

package globalSym

import (
	"path"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

type MockData struct {
	Data map[string][]byte
}

func getPath(application, function uint32, name string) string {
	return path.Join(getPathPrefix(application, function), name)
}

func getPathPrefix(application, function uint32) string {
	if application == 1 && function == 1 {
		return "/" + path.Join("application", "function")
	}
	if application == 1 {
		return "/" + path.Join("application")
	}
	if function == 1 {
		return "/" + path.Join("function")
	}

	return "/"
}

func (_m MockData) Mock() *MockData {
	m := &_m

	if m.Data == nil {
		m.Data = make(map[string][]byte)
	}

	GetGlobalValue = func(name string, application, function uint32, valuePtr *byte) errno.Error {
		path := getPath(application, function, name)

		value, ok := m.Data[path]
		if !ok {
			return 1
		}

		data := unsafe.Slice(valuePtr, len(value))
		copy(data, value)
		return 0
	}

	GetGlobalValueSize = func(name string, application, function uint32, valueSizePtr *uint32) errno.Error {
		path := getPath(application, function, name)
		value, ok := m.Data[path]
		if !ok {
			return 1
		}

		*valueSizePtr = uint32(len(value))
		return 0
	}

	GetOrCreateGlobalValueSize = func(name string, application, function uint32, valueSizePtr *uint32) errno.Error {
		path := getPath(application, function, name)
		v, ok := m.Data[path]
		if !ok {
			m.Data[path] = []byte{}
			return 0
		}

		*valueSizePtr = uint32(len(v))

		return 0
	}

	PutGlobalValue = func(name string, application, function uint32, value []byte) errno.Error {
		path := getPath(application, function, name)

		m.Data[path] = value
		return 0
	}

	return m
}
