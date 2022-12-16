//go:build !wasi && !wasm
// +build !wasi,!wasm

package p2pEventSym

import (
	"unsafe"

	"github.com/ipfs/go-cid"
	eventSym "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk-symbols/mocks"
	"github.com/taubyte/go-sdk/common"
	"github.com/taubyte/go-sdk/errno"
)

type MockData struct {
	ClientId   uint32
	Command    string
	Data       []byte
	From       cid.Cid
	To         cid.Cid
	Protocol   string
	DataToSend []byte
}

func (_m MockData) Mock() *MockData {
	m := &_m

	eventSym.MockEventType(0, common.EventTypeP2P)
	MockCommand(m.ClientId, m.Command)
	MockEventData(m.ClientId, m.Data)
	MockFrom(m.ClientId, m.From)
	MockTo(m.ClientId, m.To)
	MockProtocol(m.ClientId, m.Protocol)
	m.MockWrite()

	return m
}

func MockCommand(testId uint32, testCommand string) {
	GetP2PEventCommandSize = mocks.SizeIdStringFunction(testId, testCommand)
	GetP2PEventCommand = mocks.DataIdStringFunction(testId, testCommand)
}

func MockEventData(testId uint32, testData []byte) {
	GetP2PEventDataSize = mocks.SizeIdByteFunction(testId, testData)
	GetP2PEventData = mocks.DataIdByteFunction(testId, testData)
}

func MockFrom(testId uint32, testFrom cid.Cid) {
	GetP2PEventFromSize = mocks.SizeIdByteFunction(testId, testFrom.Bytes())
	GetP2PEventFrom = mocks.DataIdByteFunction(testId, testFrom.Bytes())
}

func MockTo(testId uint32, testTo cid.Cid) {
	GetP2PEventToSize = mocks.SizeIdByteFunction(testId, testTo.Bytes())
	GetP2PEventTo = mocks.DataIdByteFunction(testId, testTo.Bytes())
}

func MockProtocol(testId uint32, testProtocol string) {
	GetP2PEventProtocolSize = mocks.SizeIdStringFunction(testId, testProtocol)
	GetP2PEventProtocol = mocks.DataIdStringFunction(testId, testProtocol)
}

func (m *MockData) MockWrite() {
	WriteP2PResponse = func(eventId uint32, bufPtr *byte, bufSize uint32) (error errno.Error) {
		if eventId != m.ClientId {
			return 1
		}

		data := unsafe.Slice(bufPtr, bufSize)
		m.DataToSend = data
		return 0
	}
}
