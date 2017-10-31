package mpp

import "encoding/binary"

func GetStr(v []byte) (s string, err error) {
	s, _, err = getStr(v)
	return
}

func getStr(v []byte) (s string, end int64, err error) {
	it, _, iPack := GetInType(v)

	var offset int64
	var strLen int64

	switch it {

	case InTypeFixStr:
		offset = 0
		strLen = int64(iPack)

	case InTypeStr8:
		offset = 1
		strLen = int64(uint8(v[1]))

	case InTypeStr16:
		offset = 2
		strLen = int64(binary.BigEndian.Uint16(v[1:3]))

	case InTypeStr32:
		offset = 4
		strLen = int64(binary.BigEndian.Uint32(v[1:5]))

	default:
		err = NotStringError
		return
	}

	start := offset + 1
	end = start + strLen

	s = string(v[start:end])

	return
}
