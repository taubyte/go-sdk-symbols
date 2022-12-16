//go:build !wasi && !wasm
// +build !wasi,!wasm

package pubsubSym

import (
	"unsafe"

	eventSym "github.com/taubyte/go-sdk-symbols/event"
	"github.com/taubyte/go-sdk-symbols/mocks"
	"github.com/taubyte/go-sdk/common"
	"github.com/taubyte/go-sdk/errno"
)

type MockData struct {
	EventId       uint32
	Channel       string
	WebSocketURL  string
	PublishedData []byte
	EventData     []byte
	Subscriptions []string
}

func (_m MockData) Mock() *MockData {
	m := &_m

	eventSym.MockEventType(m.EventId, common.EventTypePubsub)
	MockWebSocketUrl(m.Channel, m.WebSocketURL)
	MockEventData(m.EventId, m.EventData)
	MockEventChannel(m.EventId, m.Channel)

	if m.Subscriptions == nil {
		m.Subscriptions = make([]string, 0)
	}

	m.Subscribe()
	m.Publish()
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

func (m *MockData) Subscribe() {
	SetSubscriptionChannel = func(channel string) (error errno.Error) {
		if channel != m.Channel {
			return 1
		}

		m.Subscriptions = append(m.Subscriptions, channel)
		return 0
	}
}

func (m *MockData) Publish() {
	PublishToChannel = func(channel string, data *byte, dataSize uint32) (error errno.Error) {
		if channel != m.Channel {
			return 1
		}

		m.PublishedData = unsafe.Slice(data, dataSize)
		return 0
	}
}

func MockWebSocketUrl(testChannel, url string) {
	GetWebSocketURLSize = func(channel string, sizePtr *uint32) (error errno.Error) {
		if channel != testChannel {
			return 1
		}

		return mocks.SizeString(url, sizePtr)
	}

	GetWebSocketURL = func(channel string, socketURLPtr *byte) (error errno.Error) {
		if channel != testChannel {
			return 1
		}

		return mocks.DataString(url, socketURLPtr)
	}
}
