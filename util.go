package mpp

import "encoding/binary"

func skip(v []byte, j int64) ([]byte, error) {

	if j < 1 {
		return v, nil
	}

	var i int64

	for {
		i++
		if i > j {
			break
		}

		dataLen := GetByteLen(v)

		if dataLen < 1 {
			return nil, ErrIncomplete
		}

		v = v[dataLen:]
	}

	return v, nil
}

func GetByteLen(v []byte) (byteLen int64) {

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
		var i int64
		byteLen = metaLen
		for {
			i++
			if i > limit {
				break
			}
			byteLen += GetByteLen(v[byteLen:])
		}

	default:

		panic(`unknown type ` + f.Type().String())
	}

	return
}

func getCount(f Format, v []byte) (count int64, metaLen int64, err error) {

	switch f {

	case FormatFixArray:

		count = int64(v[0] & 0x0f)
		metaLen = 1

	case FormatFixMap:

		count = int64(v[0] & 0x0f)
		metaLen = 1

	case FormatFixStr:

		count = int64(v[0] & 0x1f)
		metaLen = 1

	case FormatStr8,
		FormatBin8,
		FormatExt8:

		count = int64(uint8(v[1]))
		metaLen = 2

	case FormatStr16,
		FormatBin16,
		FormatExt16,
		FormatMap16,
		FormatArray16:

		metaLen = 3
		count = int64(binary.BigEndian.Uint16(v[1:3]))

	case FormatStr32,
		FormatBin32,
		FormatExt32,
		FormatMap32,
		FormatArray32:

		metaLen = 5
		count = int64(binary.BigEndian.Uint32(v[1:5]))

	default:

		err = ErrCanNotCount
	}

	return
}
