//go:build !wasi && !wasm
// +build !wasi,!wasm

package nodeSym

import (
	"path"
	"unsafe"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

type MockData struct {
	Protocol        string
	Command         string
	CommandId       uint32
	SendResponse    []byte
	SendTo          cid.Cid
	SentData        []byte
	ListenHash      string
	ListenProtocol  string
	DiscoverMax     uint32
	DiscoverTimeout uint32
	DiscoverId      uint32
	Peers           []cid.Cid
}

func (_m MockData) Mock() *MockData {
	m := &_m

	MockNew(m.Protocol, m.Command, m.CommandId)
	m.MockSend()
	MockListen(m.ListenHash, m.ListenProtocol)
	MockDiscover(m.DiscoverMax, m.DiscoverTimeout, m.DiscoverId, m.Peers)
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

	SendCommandTo = func(id uint32, cidBuf, data *byte, dataSize uint32, responseSize *uint32) (error errno.Error) {
		if id != m.CommandId {
			return 1
		}

		_cid := unsafe.Slice(cidBuf, codec.CidBufferSize)
		_, cid, err := cid.CidFromBytes(_cid)
		if err != nil {
			return 1
		}

		if cid.String() != m.SendTo.String() {
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

func MockDiscover(expectedMax, expectedNSTimeout, discoverId uint32, peers []cid.Cid) {
	DiscoverPeersSize = func(max, nsTimeout uint32, id, peersSize *uint32) (error errno.Error) {
		if max != expectedMax {
			return 1
		}

		if nsTimeout != expectedNSTimeout {
			return 1
		}

		peerBytes := make([][]byte, len(peers))
		for idx, p := range peers {
			peerBytes[idx] = p.Bytes()
		}

		var toWrite []byte
		err := codec.Convert(peerBytes).To(toWrite)
		if err != nil {
			return 1
		}

		*id = discoverId
		*peersSize = uint32(len(toWrite))
		return 0
	}

	DiscoverPeers = func(id uint32, peersBuf *byte) (error errno.Error) {
		if id != discoverId {
			return 1
		}

		peerBytes := make([][]byte, len(peers))
		for idx, p := range peers {
			peerBytes[idx] = p.Bytes()
		}

		var toWrite []byte
		err := codec.Convert(peerBytes).To(toWrite)
		if err != nil {
			return 1
		}

		data := unsafe.Slice(peersBuf, uint32(len(toWrite)))
		copy(data, toWrite)

		return 0
	}
}
