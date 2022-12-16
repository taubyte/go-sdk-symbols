//go:build !wasi && !wasm
// +build !wasi,!wasm

package mocks

import (
	"github.com/taubyte/go-sdk/errno"
)

func SizeIdKey(testId uint32, id uint32, key string, data map[string][]byte, sizePtr *uint32) (error errno.Error) {
	if testId != id {
		return 1
	}

	_, ok := data[key]
	if ok == false {
		return 1
	}

	*sizePtr = uint32(len(data[key]))
	return 0
}

// Writes data size relative to the key to the sizePtr passed to the function
func SizeIdKeyFunction(testId uint32, data map[string][]byte) idKeyPtrFunction {
	return func(resolverId uint32, name string, sizePtr *uint32) (error errno.Error) {
		return SizeIdKey(testId, resolverId, name, data, sizePtr)
	}
}

func SizeString(testString string, sizePtr *uint32) errno.Error {
	*sizePtr = uint32(len(testString))
	return 0
}

func SizeStringFunction(testString string) func(*uint32) errno.Error {
	return func(sizePtr *uint32) errno.Error {
		return SizeString(testString, sizePtr)
	}
}

func SizeIdString(testId uint32, id uint32, testString string, sizePtr *uint32) errno.Error {
	if id != testId {
		return 1
	}

	return SizeString(testString, sizePtr)
}

func SizeIdStringFunction(testId uint32, testString string) idSizeFunction {
	return func(id uint32, sizePtr *uint32) (error errno.Error) {
		return SizeIdString(testId, id, testString, sizePtr)
	}
}

func SizeByte(testByte []byte, sizePtr *uint32) errno.Error {
	*sizePtr = uint32(len(testByte))
	return 0
}

func SizeIdByte(testId uint32, id uint32, testByte []byte, sizePtr *uint32) errno.Error {
	if id != testId {
		return 1
	}

	return SizeByte(testByte, sizePtr)
}

func SizeIdByteFunction(testId uint32, testByte []byte) idSizeFunction {
	return func(id uint32, sizePtr *uint32) (error errno.Error) {
		return SizeIdByte(testId, id, testByte, sizePtr)
	}
}
