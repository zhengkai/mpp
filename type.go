package mpp

import (
	"errors"
)

const (
	NotExist = Type(iota)
	String
	Integer
	Float
	Object
	Array
	Boolean
	Null
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
	IncompleteError      = errors.New("Not complete yet")
)

type Type uint8
type InType uint8

func GetType(v []byte) {
}

func DebugGetType(it InType) (t Type) {
	return getType(it)
}

func getType(it InType) (t Type) {

	t = Unknown

	switch it {

	case InTypeFixInt,
		InTypeUint8,
		InTypeUint16,
		InTypeUint32:

		t = Integer

	case InTypeFixArray,
		InTypeArray16,
		InTypeArray32:

		t = Array

	case InTypeFixMap,
		InTypeMap16,
		InTypeMap32:

		t = Object

	case
		InTypeFixStr,
		InTypeStr8,
		InTypeStr16,
		InTypeStr32:

		t = String
	}

	return
}

func getLen(v InType) (len int64) {
	switch v {

	case
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
		InTypeFixExt8:

		len = 2

	case
		InTypeBin16,
		InTypeExt16,
		InTypeUint16,
		InTypeInt16,
		InTypeFixExt16:

		len = 3

	case
		InTypeBin32,
		InTypeExt32,
		InTypeFloat32,
		InTypeUint32,
		InTypeInt32:

		len = 5

	case
		InTypeFloat64,
		InTypeUint64,
		InTypeInt64:

		len = 9
	}

	return
}

func GetInType(v []byte) (t InType, len int64, iPack uint32) {

	in := InType(v[0])
	if in <= 0x7f {
		return InTypeFixInt, 1, uint32(in & 0x7f)
	}

	if in <= 0x8f {
		return InTypeFixMap, 1, uint32(in & 0x0f)
	}

	if in <= 0x9f {
		return InTypeFixArray, 1, uint32(in & 0x0f)
	}

	if in <= 0xbf {
		return InTypeFixStr, 1, uint32(in & 0x1f)
	}

	l := getLen(in)
	if l > 0 {
		return in, l, 0
	}

	return InTypeDevUnknown, 0, 0
}
