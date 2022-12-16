//go:build !wasi && !wasm
// +build !wasi,!wasm

package ipfsClientSym

import (
	"io"
	"os"
	"unsafe"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-sdk/errno"
)

type mockContent struct {
	Cid  cid.Cid
	File *os.File
}

type MockData struct {
	ClientId     uint32
	ContentIdMap map[uint32]mockContent
	Files        map[cid.Cid]*os.File
}

func (_m MockData) Mock() *MockData {
	m := &_m

	if m.ContentIdMap == nil {
		m.ContentIdMap = make(map[uint32]mockContent)
	}

	if m.Files == nil {
		m.Files = make(map[cid.Cid]*os.File)
	}

	m.NewClient()
	m.Create()
	m.Open()
	m.Close()
	m.Read()
	m.Push()
	m.Seek()
	m.Write()
	m.Cid()
	return m
}

func (m *MockData) Create() {
	IpfsNewContent = func(clientId uint32, contentIdPtr *uint32) (error errno.Error) {
		if clientId != m.ClientId {
			return 1
		}

		*contentIdPtr = uint32(len(m.ContentIdMap))
		file, err := os.CreateTemp("", "file-*")
		if err != nil {
			return 1
		}

		m.ContentIdMap[*contentIdPtr] = mockContent{
			Cid:  cid.Cid{},
			File: file,
		}

		return 0
	}
}

func (m *MockData) Open() {
	IpfsOpenFile = func(clientId uint32, contentIdPtr *uint32, cidPtr *byte, cidSize uint32) (error errno.Error) {
		if clientId != m.ClientId {
			return 1
		}

		cidData := unsafe.Slice(cidPtr, cidSize)
		_cid, err := cid.Parse(cidData)
		if err != nil {
			return 1
		}

		file, ok := m.Files[_cid]
		if ok == false {
			return 1
		}

		*contentIdPtr = uint32(len(m.ContentIdMap))
		m.ContentIdMap[*contentIdPtr] = mockContent{
			Cid:  _cid,
			File: file,
		}

		return 0
	}
}

func (m *MockData) Close() {
	IpfsCloseFile = func(clientId, contentId uint32) (error errno.Error) {
		if clientId != m.ClientId {
			return 1
		}

		content, ok := m.ContentIdMap[contentId]
		if ok == false {
			return 1
		}

		err := content.File.Close()
		if err != nil {
			return 1
		}

		return 0
	}
}

func (m *MockData) Read() {
	IpfsReadFile = func(clientId, contentId uint32, buf *byte, bufLen uint32, countPtr *uint32) (error errno.Error) {
		if clientId != m.ClientId {
			return 1
		}

		content, ok := m.ContentIdMap[contentId]
		if ok == false {
			return 1
		}

		_buf := unsafe.Slice(buf, bufLen)
		n, err := content.File.Read(_buf)
		if err != nil && err != io.EOF {
			return errno.ErrorHttpReadBody
		}

		*countPtr = uint32(n)

		if err == io.EOF {
			return errno.ErrorEOF
		}

		return 0
	}
}

func (m *MockData) NewClient() {
	NewIPFSClient = func(clientId *uint32) (error errno.Error) {
		*clientId = m.ClientId
		return 0
	}
}

func (m *MockData) Write() {
	IpfsWriteFile = func(clientId, contentId uint32, buf *byte, bufLen uint32, countPtr *uint32) (error errno.Error) {
		if clientId != m.ClientId {
			return 1
		}

		content, ok := m.ContentIdMap[contentId]
		if ok == false {
			return 1
		}

		_buf := unsafe.Slice(buf, bufLen)
		n, err := content.File.Write(_buf)
		if err != nil && err != io.EOF {
			return errno.ErrorHttpReadBody
		}

		*countPtr = uint32(n)

		if err == io.EOF {
			return errno.ErrorEOF
		}

		return 0
	}
}

func (m *MockData) Push() {
	IpfsPushFile = func(clientId, contentId uint32, cidPtr *byte) (error errno.Error) {
		if clientId != m.ClientId {
			return 1
		}

		content, ok := m.ContentIdMap[contentId]
		if ok == false {
			return 1
		}

		_, err := content.File.Seek(0, 0)
		if err != nil {
			return 1
		}

		data, err := io.ReadAll(content.File)
		if err != nil {
			return 1
		}

		// https://github.com/ipfs/go-cid
		pref := cid.Prefix{
			Version:  1,
			Codec:    85, // mc.Raw,  	   mc "github.com/multiformats/go-multicodec"
			MhType:   18, // mh.SHA2_256,  mh "github.com/multiformats/go-multihash"
			MhLength: -1, // default length
		}

		// And then feed it some data
		content.Cid, err = pref.Sum(data)
		if err != nil {
			return 1
		}

		cidData := content.Cid.Bytes()
		_cidPtr := unsafe.Slice(cidPtr, 64)
		copy(_cidPtr, cidData)

		file, err := os.CreateTemp("", "file-*")
		if err != nil {
			return 1
		}

		_, err = file.Write(data)
		if err != nil {
			return 1
		}
		file.Seek(0, 0)
		m.Files[content.Cid] = file
		return 0
	}
}

func (m *MockData) Seek() {
	IpfsSeekFile = func(clientId, contentId uint32, offset int64, whence int, offsetPtr *int) (error errno.Error) {
		if clientId != m.ClientId {
			return 1
		}

		content, ok := m.ContentIdMap[contentId]
		if ok == false {
			return 1
		}

		n, err := content.File.Seek(offset, whence)
		if err != nil {
			return 1
		}

		*offsetPtr = int(n)
		return 0
	}
}

func (m *MockData) Cid() {
	IpfsFileCid = func(clientId, contentId uint32, cidPtr *byte) (error errno.Error) {
		if clientId != m.ClientId {
			return 1
		}

		content, ok := m.ContentIdMap[contentId]
		if ok == false {
			return 1
		}

		cidData := content.Cid.Bytes()
		_cidPtr := unsafe.Slice(cidPtr, 64)
		copy(_cidPtr, cidData)
		return 0
	}
}
