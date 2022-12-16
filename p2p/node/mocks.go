//go:build !wasi && !wasm
// +build !wasi,!wasm

package nodeSym

import (
	"path"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
)

type MockData struct {
	Protocol       string
	Command        string
	CommandId      uint32
	SendResponse   []byte
	SentData       []byte
	ListenHash     string
	ListenProtocol string
}

func (_m MockData) Mock() *MockData {
	m := &_m

	MockNew(m.Protocol, m.Command, m.CommandId)
	m.MockSend()
	MockListen(m.ListenHash, m.ListenProtocol)
	return m
}

func MockNew(testProtocol, testCommand string, testId uint32) {
	NewCommand = func(protocol, command string, id *uint32) (error errno.Error) {
		if protocol != testProtocol || testCommand != command {
			return 1
		}

		*id = testId
		return 0
	}
}

func (m *MockData) MockSend() {
	SendCommand = func(id uint32, data *byte, dataSize uint32, responseSize *uint32) (error errno.Error) {
		if id != m.CommandId {
			return 1
		}

		_data := unsafe.Slice(data, dataSize)
		m.SentData = _data

		*responseSize = uint32(len(m.SendResponse))
		return 0
	}

	ReadCommandResponse = func(id uint32, data *byte, dataSize uint32) (error errno.Error) {
		if id != m.CommandId {
			return 1
		}

		_data := unsafe.Slice(data, dataSize)
		copy(_data, m.SendResponse)
		return 0
	}
}

func MockListen(listenHash, testProtocol string) {
	ListenToProtocolSize = func(protocol string, responseSize *uint32) (error errno.Error) {
		if protocol != testProtocol {
			return 1
		}

		_protocol := "/" + path.Join(listenHash, protocol)
		*responseSize = uint32(len(_protocol))
		return 0
	}

	ListenToProtocol = func(protocol string, response *byte, responseSize uint32) (error errno.Error) {
		if protocol != testProtocol {
			return 1
		}

		_protocol := "/" + path.Join(listenHash, protocol)
		data := unsafe.Slice(response, responseSize)
		copy(data, []byte(_protocol))
		return 0
	}

}
