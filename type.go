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

	InTypeFixInt = InType(iota)
	InTypeFixMap
	InTypeFixArray
	InTypeFixStr
	InTypeNegativeFixInt

	InTypeNil      = InType(0xc0)
	InTypeNa       = InType(0xc1)
	InTypeFalse    = InType(0xc2)
	InTypeTrue     = InType(0xc3)
	InTypeBin8     = InType(0xc4)
	InTypeBin16    = InType(0xc5)
	InTypeBin32    = InType(0xc6)
	InTypeExt8     = InType(0xc7)
	InTypeExt16    = InType(0xc8)
	InTypeExt32    = InType(0xc9)
	InTypeFloat32  = InType(0xca)
	InTypeFloat64  = InType(0xcb)
	InTypeUint8    = InType(0xcc)
	InTypeUint16   = InType(0xcd)
	InTypeUint32   = InType(0xce)
	InTypeUint64   = InType(0xcf)
	InTypeInt8     = InType(0xd0)
	InTypeInt16    = InType(0xd1)
	InTypeInt32    = InType(0xd2)
	InTypeInt64    = InType(0xd3)
	InTypeFixExt1  = InType(0xd4)
	InTypeFixExt2  = InType(0xd5)
	InTypeFixExt4  = InType(0xd6)
	InTypeFixExt8  = InType(0xd7)
	InTypeFixExt16 = InType(0xd8)
	InTypeStr8     = InType(0xd9)
	InTypeStr16    = InType(0xda)
	InTypeStr32    = InType(0xdb)
	InTypeArray16  = InType(0xdc)
	InTypeArray32  = InType(0xdd)
	InTypeMap16    = InType(0xde)
	InTypeMap32    = InType(0xdf)

	InTypeDevUnknown = InTypeNa
)

var (
	KeyPathNotFoundError = errors.New("Key path not found")
	WrongFormatError     = errors.New("Wrong format")
	NotStringError       = errors.New("Not a string")
	NotIntegerError      = errors.New("Not a integer")
	NotArrayError        = errors.New("Not a array")
	NotMapError          = errors.New("Not a map")
	NotFixedDataError    = errors.New("Not a fixed data")
	IncompleteError      = errors.New("Not complete yet")
	IllegalMapKeyError   = errors.New("Iillegal map key")
	InTypeError          = errors.New("Unknown InType")

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
)

type Type uint8
type InType uint8

func DebugGetType(it InType) (t Type) {
	return getType(it)
}

func getType(it InType) (t Type) {

	t = Unknown

	switch it {

	case InTypeFixStr,
		InTypeStr8,
		InTypeStr16,
		InTypeStr32:

		t = String

	case InTypeFixInt,
		InTypeInt8,
		InTypeInt16,
		InTypeInt32,
		InTypeInt64,
		InTypeUint8,
		InTypeUint16,
		InTypeUint32,
		InTypeUint64:

		t = Integer

	case InTypeFixArray,
		InTypeArray16,
		InTypeArray32:

		t = Array

	case InTypeFixMap,
		InTypeMap16,
		InTypeMap32:

		t = Map

	case InTypeTrue,
		InTypeFalse:

		t = Boolean

	case InTypeNil:

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

		fmt.Println(TypeName[t], t)

		panic(`unknown type`)
	}

	return

}

func getMetaLen(v InType) (len int64) {
	switch v {

	case
		InTypeFixArray,
		InTypeFixMap,
		InTypeFixInt,
		InTypeNil,
		InTypeNa,
		InTypeFalse,
		InTypeTrue:

		len = 1

	case
		InTypeBin8,
		InTypeExt8,
		InTypeUint8,
		InTypeInt8,
		InTypeFixExt8,
		InTypeStr8:

		len = 2

	case
		InTypeArray16,
		InTypeMap16,
		InTypeBin16,
		InTypeExt16,
		InTypeUint16,
		InTypeInt16,
		InTypeStr16,
		InTypeFixExt16:

		len = 3

	case
		InTypeArray32,
		InTypeMap32,
		InTypeBin32,
		InTypeExt32,
		InTypeFloat32,
		InTypeUint32,
		InTypeStr32,
		InTypeInt32:

		len = 5

	case
		InTypeFloat64,
		InTypeUint64,
		InTypeInt64:

		len = 9

	default:

		panic(`incomplete type`)
	}

	return
}

/*
func GetInType(v []byte) (t InType, metaLen int64, iPack uint32) {

	metaLen = getMetaLen(v)
	if metaLen > 0 {
		t = in
		return
	}

	return InTypeDevUnknown, 0, 0
}
*/

func getFixedMeta(b InType) (it InType, t Type, metaLen int64, ext int64) {

	if b <= 0x7f {
		return InTypeFixInt, Integer, 1, int64(b & 0x7f)
	}

	if b <= 0x8f {
		return InTypeFixMap, Map, 1, int64(b & 0x0f)
	}

	if b <= 0x9f {
		return InTypeFixArray, Array, 1, int64(b & 0x0f)
	}

	if b <= 0xbf {
		return InTypeFixStr, String, 1, int64(b & 0x1f)
	}

	return
}

func parseMeta(v []byte) (it InType, t Type, metaLen int64, ext int64, err error) {

	// TODO: combine GetInType / getMetaLen / getType and special get len func

	first := InType(v[0])

	if first <= 0xbf {
		it, t, metaLen, ext = getFixedMeta(first)
		return
	}

	it = first
	t = getType(it)
	metaLen = getMetaLen(it)

	switch it {

	case InTypeStr8:
		ext = int64(uint8(v[1]))

	case InTypeStr16:
		ext = int64(binary.BigEndian.Uint16(v[1:metaLen]))

	case InTypeStr32:
		ext = int64(binary.BigEndian.Uint32(v[1:metaLen]))

	case InTypeArray16:
		ext = int64(binary.BigEndian.Uint16(v[1:metaLen]))

	case InTypeArray32:
		ext = int64(binary.BigEndian.Uint32(v[1:metaLen]))

	case InTypeMap16:
		ext = int64(binary.BigEndian.Uint16(v[1:metaLen]))

	case InTypeMap32:
		ext = int64(binary.BigEndian.Uint32(v[1:metaLen]))

	default:
		err = InTypeError
	}

	return
}
