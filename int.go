package mpp

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func GetInt(v []byte) (i int64, err error) {

	it, t, metaLen, ext, parseErr := parseMeta(v)
	if t != Integer || parseErr != nil {
		err = NotIntegerError
		return
	}

	switch it {

	case FormatFixInt:
		i = int64(ext)

	case FormatInt8:
		i = int64(int8(v[1]))

	case FormatInt16:

		var i32 int32
		buf := bytes.NewBuffer([]byte{0, 0, v[1], v[2]})
		binary.Read(buf, binary.BigEndian, &i32)
		i = int64(i32)

	case FormatInt32:

		var i32 int32
		buf := bytes.NewBuffer(v[1:metaLen])
		binary.Read(buf, binary.BigEndian, &i32)
		i = int64(i32)

	case FormatInt64:

		buf := bytes.NewBuffer(v[1:metaLen])
		binary.Read(buf, binary.BigEndian, &i)

	case FormatUint8:
		i = int64(uint8(v[1]))

	case FormatUint16:
		i = int64(binary.BigEndian.Uint16(v[1:metaLen]))

	case FormatUint32:
		i = int64(binary.BigEndian.Uint32(v[1:metaLen]))

	case FormatUint64:
		i = int64(binary.BigEndian.Uint64(v[1:metaLen]))

	default:
		err = NotIntegerError
		fmt.Println(`test`, it)
		return
	}

	return
}
