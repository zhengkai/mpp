package mpp

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func GetInt(v []byte) (i int64, err error) {
	i, _, err = getInt(v)
	return
}

func getInt(v []byte) (i int64, metaLen int64, err error) {

	it, t, metaLen, ext, parseErr := parseMeta(v)
	if t != Integer || parseErr != nil {
		fmt.Println(`test`, it)
		err = NotIntegerError
		return
	}

	switch it {

	case InTypeFixInt:
		i = int64(ext)

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

	case InTypeUint16:
		i = int64(binary.BigEndian.Uint16(v[1:metaLen]))

	case InTypeUint32:
		i = int64(binary.BigEndian.Uint32(v[1:metaLen]))

	case InTypeUint64:
		i = int64(binary.BigEndian.Uint64(v[1:metaLen]))

	default:
		err = NotIntegerError
		fmt.Println(`test`, it)
		return
	}

	return
}
