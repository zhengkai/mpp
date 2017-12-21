package mpp

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	NotExist = Type(iota)
	String
	Integer
	Float
	Map
	Array
	Boolean
	Nil
	Unknown

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

	FormatDevUnknown = FormatNa
)

var (
	KeyPathNotFoundError = errors.New("Key path not found")
	WrongFormatError     = errors.New("Wrong format")
	NotStringError       = errors.New("Not a string")
	NotIntegerError      = errors.New("Not a integer")
	NotFloatError        = errors.New("Not a float")
	NotArrayError        = errors.New("Not a array")
	NotMapError          = errors.New("Not a map")
	NotFixedDataError    = errors.New("Not a fixed data")
	IncompleteError      = errors.New("Not complete yet")
	IllegalMapKeyError   = errors.New("Iillegal map key")
	FormatError          = errors.New("Unknown Format")

	TypeName = map[Type]string{
		String:  `String`,
		Integer: `Integer`,
		Float:   `Float`,
		Map:     `Map`,
		Array:   `Array`,
		Boolean: `Boolean`,
		Nil:     `Nil`,
		Unknown: `Unknown`,
	}

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

func (t Type) String() string {
	return TypeName[t]
}

func (f Format) String() string {
	return FormatName[f]
}

type Type uint8
type Format uint8

func (f Format) Type() (t Type) {

	t = Unknown

	switch f {

	case FormatFixStr,
		FormatStr8,
		FormatStr16,
		FormatStr32:

		t = String

	case FormatFixInt,
		FormatInt8,
		FormatInt16,
		FormatInt32,
		FormatInt64,
		FormatUint8,
		FormatUint16,
		FormatUint32,
		FormatUint64:

		t = Integer

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

		t = Boolean

	case FormatNil:

		t = Nil
	}

	return
}

func GetByteLen(v []byte) (byteLen int64) {
	return getByteLen(v)
}

func getByteLen(v []byte) (byteLen int64) {

	_, t, metaLen, ext, _ := parseMeta(v)

	switch t {

	case Integer,
		Boolean,
		Float,
		Nil:

		byteLen = metaLen

	case String:

		byteLen = metaLen + ext

	case Map:

		limit := ext * 2
		var i int64
		byteLen = metaLen
		for {
			i++
			if i > limit {
				break
			}
			byteLen += getByteLen(v[byteLen:])
		}

	case Array:

		var i int64
		byteLen = metaLen
		for {
			i++
			if i > ext {
				break
			}
			byteLen += getByteLen(v[byteLen:])
		}

	default:

		panic(`unknown type ` + t.String())
	}

	return
}

func (v Format) MetaLen() (len int64) {

	switch v {

	case
		FormatFixArray,
		FormatFixMap,
		FormatFixInt,
		FormatFixStr,
		FormatNil,
		FormatNa,
		FormatFalse,
		FormatTrue:

		len = 1

	case
		FormatBin8,
		FormatExt8,
		FormatUint8,
		FormatInt8,
		FormatFixExt8,
		FormatStr8:

		len = 2

	case
		FormatArray16,
		FormatMap16,
		FormatBin16,
		FormatExt16,
		FormatUint16,
		FormatInt16,
		FormatStr16,
		FormatFixExt16:

		len = 3

	case
		FormatArray32,
		FormatMap32,
		FormatBin32,
		FormatExt32,
		FormatFloat32,
		FormatUint32,
		FormatStr32,
		FormatInt32:

		len = 5

	case
		FormatFloat64,
		FormatUint64,
		FormatInt64:

		len = 9

	default:

		panic(`incomplete type` + v.String())
	}

	return
}

func getFixedMeta(b Format) (it Format, ext int64) {

	if b <= 0x7f {
		return FormatFixInt, int64(b & 0x7f)
	}

	if b <= 0x8f {
		return FormatFixMap, int64(b & 0x0f)
	}

	if b <= 0x9f {
		return FormatFixArray, int64(b & 0x0f)
	}

	if b <= 0xbf {
		return FormatFixStr, int64(b & 0x1f)
	}

	return
}

func parseMeta(v []byte) (it Format, t Type, metaLen int64, ext int64, err error) {

	first := Format(v[0])

	if first <= 0xbf {
		it, ext = getFixedMeta(first)
		t = it.Type()
		metaLen = it.MetaLen()
		return
	}

	it = first
	t = it.Type()
	metaLen = it.MetaLen()

	switch it {

	case FormatStr8:
		ext = int64(uint8(v[1]))

	case FormatStr16:
		ext = int64(binary.BigEndian.Uint16(v[1:it.MetaLen()]))

	case FormatStr32:
		ext = int64(binary.BigEndian.Uint32(v[1:it.MetaLen()]))

	case FormatArray16:
		ext = int64(binary.BigEndian.Uint16(v[1:it.MetaLen()]))

	case FormatArray32:
		ext = int64(binary.BigEndian.Uint32(v[1:it.MetaLen()]))

	case FormatMap16:
		ext = int64(binary.BigEndian.Uint16(v[1:it.MetaLen()]))

	case FormatMap32:
		ext = int64(binary.BigEndian.Uint32(v[1:it.MetaLen()]))

	case FormatNil,
		FormatTrue,
		FormatFalse,
		FormatInt16,
		FormatInt32,
		FormatInt64,
		FormatInt8,
		FormatUint16,
		FormatUint32,
		FormatUint64,
		FormatUint8,
		FormatFloat32,
		FormatFloat64:

		ext = 0

	default:
		fmt.Printf("unknown type 0x%x %s\n", it, FormatName[it])
		err = FormatError
	}

	return
}
