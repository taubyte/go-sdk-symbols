//go:build !wasi && !wasm
// +build !wasi,!wasm

package mocks

import (
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

func DataIdKey(testId uint32, resId uint32, key string, data map[string][]byte, dataPtr *byte) (error errno.Error) {
	if testId != resId {
		return 1
	}

	_data := unsafe.Slice(dataPtr, len(data[key]))
	copy(_data, data[key])
	return 0
}

// Writes data relative to the key to the dataPtr passed to the function
func DataIdKeyFunction(testId uint32, data map[string][]byte) func(uint32, string, *byte) (error errno.Error) {
	return func(resolverId uint32, name string, dataPtr *byte) (error errno.Error) {
		return DataIdKey(testId, resolverId, name, data, dataPtr)
	}
}

func DataString(testString string, stringPtr *byte) (error errno.Error) {
	data := unsafe.Slice(stringPtr, len(testString))
	copy(data, testString)
	return 0
}

func DataStringFunction(testString string) func(*byte) (error errno.Error) {
	return func(stringPtr *byte) (error errno.Error) {
		return DataString(testString, stringPtr)
	}
}

func DataIdString(testId uint32, id uint32, testString string, bufPtr *byte) errno.Error {
	if id != testId {
		return 1
	}

	return DataString(testString, bufPtr)
}

func DataIdStringFunction(testId uint32, testString string) idDataFunction {
	return func(id uint32, bufPtr *byte) (error errno.Error) {
		return DataIdString(testId, id, testString, bufPtr)
	}
}

// Not actually safe, simply ignoring the shared size.
func DataIdStringFunctionSafe(testId uint32, testString string) idDataFunctionSafe {
	return func(id uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
		return DataIdString(testId, id, testString, bufPtr)
	}
}

func DataByte(testByte []byte, bufPtr *byte) (error errno.Error) {
	data := unsafe.Slice(bufPtr, len(testByte))
	copy(data, testByte)
	return 0
}

func DataIdByte(testId uint32, id uint32, testByte []byte, bufPtr *byte) errno.Error {
	if id != testId {
		return 1
	}

	return DataByte(testByte, bufPtr)
}

func DataIdByteFunction(testId uint32, testByte []byte) idDataFunction {
	return func(id uint32, bufPtr *byte) (error errno.Error) {
		return DataIdByte(testId, id, testByte, bufPtr)
	}
}
