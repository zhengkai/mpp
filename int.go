package mpp

import (
	"bytes"
	"encoding/binary"
)

func GetInt(v []byte, key ...string) (i int64, err error) {

	v, err = byPath(v, key)
	if err != nil {
		return
	}

	f := GetFormat(v[0])

	switch f {

	case FormatFixInt:

		i = int64(v[0] & 0x7f)

	case FormatNegativeFixInt:

		i = int64(int8(v[0]))

	case FormatInt8:

		i = int64(int8(v[1]))

	case FormatInt16:

		var i16 int16
		buf := bytes.NewBuffer(v[1:3])
		binary.Read(buf, binary.BigEndian, &i16)
		i = int64(i16)

	case FormatInt32:

		var i32 int32
		buf := bytes.NewBuffer(v[1:5])
		binary.Read(buf, binary.BigEndian, &i32)
		i = int64(i32)

	case FormatInt64:

		buf := bytes.NewBuffer(v[1:9])
		binary.Read(buf, binary.BigEndian, &i)

	case FormatUint8:

		i = int64(uint8(v[1]))

	case FormatUint16:

		i = int64(binary.BigEndian.Uint16(v[1:3]))

	case FormatUint32:

		i = int64(binary.BigEndian.Uint32(v[1:5]))

	case FormatUint64:

		i = int64(binary.BigEndian.Uint64(v[1:9]))

	default:

		err = NotIntError

		// panic(`unknown type ` + f.String() + ` ` + f.Type().String())
	}

	return
}
