//go:build !wasi && !wasm
// +build !wasi,!wasm

package eventSym

import (
	"bytes"
	"io"
	"os"
	"unsafe"

	"github.com/taubyte/go-sdk-symbols/mocks"
	"github.com/taubyte/go-sdk/common"
	"github.com/taubyte/go-sdk/errno"
)

type MockData struct {
	EventId    uint32
	EventType  common.EventType
	Body       []byte
	Headers    map[string]string
	Queries    map[string]string
	Host       string
	Method     string
	Path       string
	UserAgent  string
	ReturnBody []byte
	ReturnCode int

	RedirectedTo    string
	RedirectionCode uint32
}

func (_m MockData) Mock() *MockData {
	m := &_m

	MockEventType(m.EventId, m.EventType)

	if m.Headers == nil {
		m.Headers = make(map[string]string)
	}

	if m.Queries == nil {
		m.Queries = make(map[string]string)
	}

	MockHeaders(m.EventId, m.Headers)
	MockQueries(m.EventId, m.Queries)

	mockBodyReader := bytes.NewReader(m.Body)
	MockBody(m.EventId, mockBodyReader)

	MockHost(m.EventId, m.Host)
	MockMethod(m.EventId, m.Method)
	MockPath(m.EventId, m.Path)
	MockUserAgent(m.EventId, m.UserAgent)

	m.MockRedirect()
	m.MockReturnCode()
	m.MockReturnBody()

	return m
}

func MockHost(testId uint32, testHost string) {
	GetHttpEventHostSize = mocks.SizeIdStringFunction(testId, testHost)
	GetHttpEventHost = mocks.DataIdStringFunctionSafe(testId, testHost)
}

func MockMethod(testId uint32, testMethod string) {
	GetHttpEventMethodSize = mocks.SizeIdStringFunction(testId, testMethod)
	GetHttpEventMethod = mocks.DataIdStringFunctionSafe(testId, testMethod)
}

func MockPath(testId uint32, testPath string) {
	GetHttpEventPathSize = mocks.SizeIdStringFunction(testId, testPath)
	GetHttpEventPath = mocks.DataIdStringFunctionSafe(testId, testPath)
}

func MockUserAgent(testId uint32, testUserAgent string) {
	GetHttpEventUserAgentSize = mocks.SizeIdStringFunction(testId, testUserAgent)
	GetHttpEventUserAgent = mocks.DataIdStringFunctionSafe(testId, testUserAgent)
}

func (m *MockData) MockRedirect() {
	EventHttpRedirect = func(eventId uint32, url string, code uint32) (error errno.Error) {
		if eventId != m.EventId {
			return 1
		}

		m.RedirectionCode = code
		m.RedirectedTo = url

		return 0
	}
}

func (m *MockData) MockReturnCode() {
	EventHttpRetCode = func(eventId, code uint32) (error errno.Error) {
		if eventId != m.EventId {
			return 1
		}

		m.ReturnCode = int(code)
		return 0
	}
}

func (m *MockData) MockReturnBody() {
	EventHttpWrite = func(eventId uint32, bufPtr *byte, bufSize uint32, n *uint32) (error errno.Error) {
		if eventId != m.EventId {
			return 1
		}

		data := unsafe.Slice(bufPtr, bufSize)
		copy(m.ReturnBody, data)

		f, err := os.CreateTemp("", "*")
		if err != nil {
			return errno.ErrorCap
		}
		defer f.Close()

		wroteN, err := f.Write(data)
		defer func() {
			*n = uint32(wroteN)
		}()
		if err != nil {
			return errno.ErrorHttpWrite
		}

		_, err = f.Seek(0, 0)
		if err != nil {
			return errno.ErrorHttpWrite
		}

		m.ReturnBody, err = io.ReadAll(f)
		if err != nil {
			return errno.ErrorHttpWrite
		}

		return 0
	}
}
