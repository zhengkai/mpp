package mpp

import (
	"bytes"
	"encoding/binary"
)

func GetInt(v []byte) (i int64, err error) {
	i, _, err = getInt(v)
	return
}

func getInt(v []byte) (i int64, metaLen int64, err error) {

	it, metaLen, iPack := GetInType(v)

	switch it {

	case InTypeFixInt:
		i = int64(iPack)

	case InTypeInt8:
		i = int64(int8(v[1]))

	case InTypeInt16:

		var i32 int32
		buf := bytes.NewBuffer([]byte{0, 0, v[1], v[2]})
		binary.Read(buf, binary.BigEndian, &i32)
		i = int64(i32)

	case InTypeInt32:

		var i32 int32
		buf := bytes.NewBuffer(v[1:metaLen])
		binary.Read(buf, binary.BigEndian, &i32)
		i = int64(i32)

	case InTypeInt64:

		buf := bytes.NewBuffer(v[1:metaLen])
		binary.Read(buf, binary.BigEndian, &i)

	case InTypeUint8:
		i = int64(uint8(v[1]))

	case InTypeUint16,
		InTypeUint32,
		InTypeUint64:

		i = int64(binary.BigEndian.Uint64(v[1:metaLen]))

	default:
		err = NotIntegerError
		return
	}

	return
}
