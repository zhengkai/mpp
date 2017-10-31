package mpp

import (
	"bytes"
	"encoding/binary"
)

func GetInt(v []byte) (i int64, err error) {
	i, _, err = getInt(v)
	return
}

func getInt(v []byte) (i int64, end int64, err error) {

	it, _, iPack := GetInType(v)

	switch it {

	case InTypeFixInt:
		end = 1
		i = int64(iPack)

	case InTypeInt8:
		end = 2
		i = int64(int8(v[1]))

	case InTypeInt16:

		end = 3
		var i32 int32
		buf := bytes.NewBuffer([]byte{0, 0, v[1], v[2]})
		binary.Read(buf, binary.BigEndian, &i32)
		i = int64(i32)

	case InTypeInt32:

		end = 5
		var i32 int32
		buf := bytes.NewBuffer(v[1:end])
		binary.Read(buf, binary.BigEndian, &i32)
		i = int64(i32)

	case InTypeInt64:

		end = 9
		buf := bytes.NewBuffer(v[1:end])
		binary.Read(buf, binary.BigEndian, &i)

	case InTypeUint8:
		end = 2
		i = int64(uint8(v[1]))

	case InTypeUint16:
		end = 3
		i = int64(binary.BigEndian.Uint16(v[1:end]))

	case InTypeUint32:
		end = 5
		i = int64(binary.BigEndian.Uint32(v[1:end]))

	case InTypeUint64:
		end = 9
		i = int64(binary.BigEndian.Uint64(v[1:end]))

	default:
		err = NotStringError
		return
	}

	return
}
