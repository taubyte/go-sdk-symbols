//go:build !wasi && !wasm
// +build !wasi,!wasm

package pubsubSym

import (
	eventSym "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk-symbols/mocks"
	"github.com/taubyte/go-sdk/common"
)

type MockData struct {
	EventId   uint32
	Channel   string
	EventData []byte
}

func (_m MockData) Mock() *MockData {
	m := &_m

	eventSym.MockEventType(m.EventId, common.EventTypePubsub)
	MockEventData(m.EventId, m.EventData)
	MockEventChannel(m.EventId, m.Channel)

	return m
}

func MockEventData(testEventId uint32, testData []byte) {
	GetMessageData = mocks.DataIdByteFunction(testEventId, testData)
	GetMessageDataSize = mocks.SizeIdByteFunction(testEventId, testData)
}

func MockEventChannel(testEventId uint32, testChannel string) {
	GetMessageChannel = mocks.DataIdStringFunction(testEventId, testChannel)
	GetMessageChannelSize = mocks.SizeIdStringFunction(testEventId, testChannel)
}
