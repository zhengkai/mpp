package mpp

const (
	FormatUnknown = Format(iota)
	FormatFixInt
	FormatFixMap
	FormatFixArray
	FormatFixStr
	FormatNegativeFixInt

	FormatNil      = Format(0xc0)
	FormatNa       = Format(0xc1)
	FormatFalse    = Format(0xc2)
	FormatTrue     = Format(0xc3)
	FormatBin8     = Format(0xc4)
	FormatBin16    = Format(0xc5)
	FormatBin32    = Format(0xc6)
	FormatExt8     = Format(0xc7)
	FormatExt16    = Format(0xc8)
	FormatExt32    = Format(0xc9)
	FormatFloat32  = Format(0xca)
	FormatFloat64  = Format(0xcb)
	FormatUint8    = Format(0xcc)
	FormatUint16   = Format(0xcd)
	FormatUint32   = Format(0xce)
	FormatUint64   = Format(0xcf)
	FormatInt8     = Format(0xd0)
	FormatInt16    = Format(0xd1)
	FormatInt32    = Format(0xd2)
	FormatInt64    = Format(0xd3)
	FormatFixExt1  = Format(0xd4)
	FormatFixExt2  = Format(0xd5)
	FormatFixExt4  = Format(0xd6)
	FormatFixExt8  = Format(0xd7)
	FormatFixExt16 = Format(0xd8)
	FormatStr8     = Format(0xd9)
	FormatStr16    = Format(0xda)
	FormatStr32    = Format(0xdb)
	FormatArray16  = Format(0xdc)
	FormatArray32  = Format(0xdd)
	FormatMap16    = Format(0xde)
	FormatMap32    = Format(0xdf)
)

// Format https://github.com/msgpack/msgpack/blob/master/spec.md#formats
type Format uint8

func (f Format) String() (s string) {
	switch f {
	case FormatFixInt:
		s = `FixInt`
	case FormatFixMap:
		s = `FixMap`
	case FormatFixArray:
		s = `FixArray`
	case FormatFixStr:
		s = `FixStr`
	case FormatNegativeFixInt:
		s = `NegativeFixInt`
	case FormatNil:
		s = `Nil`
	case FormatNa:
		s = `Na`
	case FormatFalse:
		s = `False`
	case FormatTrue:
		s = `True`
	case FormatBin8:
		s = `Bin8`
	case FormatBin16:
		s = `Bin16`
	case FormatBin32:
		s = `Bin32`
	case FormatExt8:
		s = `Ext8`
	case FormatExt16:
		s = `Ext16`
	case FormatExt32:
		s = `Ext32`
	case FormatFloat32:
		s = `Float32`
	case FormatFloat64:
		s = `Float64`
	case FormatUint8:
		s = `Uint8`
	case FormatUint16:
		s = `Uint16`
	case FormatUint32:
		s = `Uint32`
	case FormatUint64:
		s = `Uint64`
	case FormatInt8:
		s = `Int8`
	case FormatInt16:
		s = `Int16`
	case FormatInt32:
		s = `Int32`
	case FormatInt64:
		s = `Int64`
	case FormatFixExt1:
		s = `FixExt1`
	case FormatFixExt2:
		s = `FixExt2`
	case FormatFixExt4:
		s = `FixExt4`
	case FormatFixExt8:
		s = `FixExt8`
	case FormatFixExt16:
		s = `FixExt16`
	case FormatStr8:
		s = `Str8`
	case FormatStr16:
		s = `Str16`
	case FormatStr32:
		s = `Str32`
	case FormatArray16:
		s = `Array16`
	case FormatArray32:
		s = `Array32`
	case FormatMap16:
		s = `Map16`
	case FormatMap32:
		s = `Map32`
	default:
		s = `Unknown`
	}
	return
}

func (f Format) Type() (t Type) {

	switch f {

	case FormatFixStr,
		FormatStr8,
		FormatStr16,
		FormatStr32:

		t = Str

	case FormatBin8,
		FormatBin16,
		FormatBin32:

		t = Bin

	case FormatFixInt,
		FormatNegativeFixInt,
		FormatInt8,
		FormatInt16,
		FormatInt32,
		FormatInt64,
		FormatUint8,
		FormatUint16,
		FormatUint32,
		FormatUint64:

		t = Int

	case FormatFloat32,
		FormatFloat64:

		t = Float

	case FormatFixArray,
		FormatArray16,
		FormatArray32:

		t = Array

	case FormatFixMap,
		FormatMap16,
		FormatMap32:

		t = Map

	case FormatTrue,
		FormatFalse:

		t = Bool

	case FormatNil:

		t = Nil

	case FormatExt8,
		FormatExt16,
		FormatExt32,
		FormatFixExt1,
		FormatFixExt2,
		FormatFixExt4,
		FormatFixExt8,
		FormatFixExt16:

		t = Ext

	default:

		t = Unknown
	}

	return
}

// GetFormat .
func GetFormat(v byte) Format {

	switch {

	case v <= 0x7f:
		return FormatFixInt

	case v <= 0x8f:
		return FormatFixMap

	case v <= 0x9f:
		return FormatFixArray

	case v <= 0xbf:
		return FormatFixStr

	case v >= 0xe0:
		return FormatNegativeFixInt
	}

	return Format(v)
}

func (f Format) metaLen() (len int) {

	switch f {

	case FormatFixArray,
		FormatFixMap,
		FormatFixInt,
		FormatNegativeFixInt,
		FormatFixStr,
		FormatNil,
		FormatNa,
		FormatFalse,
		FormatTrue:

		len = 1

	case FormatBin8,
		FormatUint8,
		FormatInt8,
		FormatStr8,
		FormatFixExt1,
		FormatFixExt2,
		FormatFixExt4,
		FormatFixExt8,
		FormatFixExt16:

		len = 2

	case FormatExt8,
		FormatArray16,
		FormatMap16,
		FormatBin16,
		FormatUint16,
		FormatInt16,
		FormatStr16:

		len = 3

	case FormatExt16:

		len = 4

	case FormatArray32,
		FormatMap32,
		FormatBin32,
		FormatFloat32,
		FormatUint32,
		FormatStr32,
		FormatInt32:

		len = 5

	case FormatExt32:

		len = 6

	case FormatFloat64,
		FormatUint64,
		FormatInt64:

		len = 9
	}

	return
}
