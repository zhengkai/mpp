package mpp

import "encoding/binary"

func skip(v []byte, j int) ([]byte, error) {

	var i int
	for {
		i++
		if i > j {
			break
		}

		dataLen := GetByteLen(v)

		if dataLen < 1 || len(v) < int(dataLen+1) {
			return nil, ErrInvalid
		}

		v = v[dataLen:]
	}

	return v, nil
}

func GetByteLen(v []byte) (byteLen int) {

	if len(v) == 0 {
		return 0
	}

	f := GetFormat(v[0])
	t := f.Type()

	switch t {

	case Int,
		Float:

		byteLen = f.metaLen()

	case Bool,
		Nil:

		byteLen = 1

	case Str,
		Bin,
		Ext:

		count, metaLen, _ := getCount(f, v)
		byteLen = metaLen + count

	case Map,
		Array:

		limit, metaLen, _ := getCount(f, v)
		if t == Map {
			limit *= 2
		}
		var i int
		byteLen = metaLen
		for {
			i++
			if i > limit {
				break
			}

			if len(v) < byteLen {
				return 0
			}

			byteLen += GetByteLen(v[byteLen:])
		}

	default:

		return 0
		// panic(`unknown type ` + f.Type().String())
	}

	return
}

func getCount(f Format, v []byte) (count int, metaLen int, err error) {

	switch f {

	case FormatFixArray:

		metaLen = 1
		count = int(v[0] & 0x0f)

	case FormatFixMap:

		metaLen = 1
		count = int(v[0] & 0x0f)

	case FormatFixStr:

		metaLen = 1
		count = int(v[0] & 0x1f)

	case FormatStr8,
		FormatBin8,
		FormatExt8:

		metaLen = 2
		if len(v) < metaLen {
			err = ErrInvalid
			return
		}
		count = int(uint8(v[1]))

	case FormatStr16,
		FormatBin16,
		FormatExt16,
		FormatMap16,
		FormatArray16:

		metaLen = 3
		if len(v) < metaLen {
			err = ErrInvalid
			return
		}
		count = int(binary.BigEndian.Uint16(v[1:3]))

	case FormatStr32,
		FormatBin32,
		FormatExt32,
		FormatMap32,
		FormatArray32:

		metaLen = 5
		if len(v) < metaLen {
			err = ErrInvalid
			return
		}
		count = int(binary.BigEndian.Uint32(v[1:5]))

	default:

		err = ErrCanNotCount
	}

	return
}
