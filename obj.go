package mpp

import "encoding/binary"

func getObjLen(v []byte) (objLen int64, metaLen int64, err error) {

	it, metaLen, iPack := GetInType(v)

	switch it {

	case InTypeFixMap:
		objLen = int64(iPack)

	case InTypeMap16:
		objLen = int64(binary.BigEndian.Uint16(v[1:metaLen]))

	case InTypeMap32:
		objLen = int64(binary.BigEndian.Uint32(v[1:metaLen]))

	default:
		err = NotObjectError
	}

	return
}
