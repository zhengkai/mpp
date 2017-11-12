package mpp

import (
	"encoding/binary"
)

func getArrayLen(v []byte) (arrayLen int64, metaLen int64, err error) {

	it, metaLen, iPack := GetInType(v)

	switch it {

	case InTypeFixArray:
		arrayLen = int64(iPack)

	case InTypeArray16:
		arrayLen = int64(binary.BigEndian.Uint16(v[1:metaLen]))

	case InTypeArray32:
		arrayLen = int64(binary.BigEndian.Uint32(v[1:metaLen]))

	default:
		err = NotArrayError
	}

	return
}
