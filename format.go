package mpp

const (
	FormatFixInt = Format(iota)
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

var (
	FormatName = map[Format]string{
		FormatFixInt:         `FixInt`,
		FormatFixMap:         `FixMap`,
		FormatFixArray:       `FixArray`,
		FormatFixStr:         `FixStr`,
		FormatNegativeFixInt: `NegativeFixInt`,
		FormatNil:            `Nil`,
		FormatNa:             `Na`,
		FormatFalse:          `False`,
		FormatTrue:           `True`,
		FormatBin8:           `Bin8`,
		FormatBin16:          `Bin16`,
		FormatBin32:          `Bin32`,
		FormatExt8:           `Ext8`,
		FormatExt16:          `Ext16`,
		FormatExt32:          `Ext32`,
		FormatFloat32:        `Float32`,
		FormatFloat64:        `Float64`,
		FormatUint8:          `Uint8`,
		FormatUint16:         `Uint16`,
		FormatUint32:         `Uint32`,
		FormatUint64:         `Uint64`,
		FormatInt8:           `Int8`,
		FormatInt16:          `Int16`,
		FormatInt32:          `Int32`,
		FormatInt64:          `Int64`,
		FormatFixExt1:        `FixExt1`,
		FormatFixExt2:        `FixExt2`,
		FormatFixExt4:        `FixExt4`,
		FormatFixExt8:        `FixExt8`,
		FormatFixExt16:       `FixExt16`,
		FormatStr8:           `Str8`,
		FormatStr16:          `Str16`,
		FormatStr32:          `Str32`,
		FormatArray16:        `Array16`,
		FormatArray32:        `Array32`,
		FormatMap16:          `Map16`,
		FormatMap32:          `Map32`,
	}
)

type Format uint8

func (f Format) String() string {
	return FormatName[f]
}

func (f Format) Type() (t Type) {

	t = Unknown

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
	}

	return
}

func GetFormat(v byte) Format {

	if v <= 0x7f {
		return FormatFixInt
	}

	if v <= 0x8f {
		return FormatFixMap
	}

	if v <= 0x9f {
		return FormatFixArray
	}

	if v <= 0xbf {
		return FormatFixStr
	}

	if v >= 0xe0 {
		return FormatNegativeFixInt
	}

	return Format(v)
}

func (f Format) MetaLen() (len int64) {

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

	default:

		len = 0
		// panic(`incomplete type ` + f.String())
	}

	return
}
