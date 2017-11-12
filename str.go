package mpp

import (
	"encoding/binary"
)

func GetStr(v []byte) (s string, err error) {
	s, _, err = getStr(v)
	return
}

func getStr(v []byte) (s string, end int64, err error) {

	var strLen int64
	var metaLen int64
	strLen, metaLen, err = getStrLen(v)

	if err != nil {
		return
	}

	end = metaLen + strLen

	s = string(v[metaLen:end])

	return
}

func getStrLen(v []byte) (strLen int64, metaLen int64, err error) {

	it, metaLen, iPack := GetInType(v)

	switch it {

	case InTypeFixStr:
		strLen = int64(iPack)

	case InTypeStr8:
		strLen = int64(uint8(v[1]))

	case InTypeStr16:
		strLen = int64(binary.BigEndian.Uint16(v[1:metaLen]))

	case InTypeStr32:
		strLen = int64(binary.BigEndian.Uint32(v[1:metaLen]))

	default:
		err = NotStringError
	}

	return
}
