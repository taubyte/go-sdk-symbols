package i2mvSym

import "github.com/taubyte/go-sdk/errno"

func MockNew(testId int32) {
	MemoryViewNew = func(bufPtr *byte, size, readCloser uint32, idPtr *uint32) (error errno.Error) {
		if testId < 0 {
			return 1
		}

		*idPtr = uint32(testId)

		return 0
	}
}

func MockRead(testId uint32) {
	MemoryViewRead = func(id, offset, count uint32, bufPtr *byte, nPtr *uint32) (error errno.Error) {
		if testId != id {
			return 1
		}

		*nPtr = count

		return 0
	}
}

func MockClose(testClosable bool) {
	MemoryViewClose = func(id uint32) (error errno.Error) {
		if testClosable {
			return 0
		}

		return 1
	}
}

func MockSize(testSize uint32) {
	MemoryViewSize = func(id uint32, sizePtr *uint32) (error errno.Error) {
		if testSize == 0 {
			return 1
		}

		*sizePtr = testSize
		return 0
	}
}
