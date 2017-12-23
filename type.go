package mpp

import (
	"errors"
)

const (
	NotExist = Type(iota)
	Str
	Bin
	Ext
	Int
	Float
	Map
	Array
	Bool
	Nil
	Unknown
)

var (
	KeyPathNotFoundError = errors.New("Key path not found")
	WrongFormatError     = errors.New("Wrong format")
	NotStrError          = errors.New("Not a string")
	NotBinError          = errors.New("Not a binary")
	NotIntError          = errors.New("Not a integer")
	NotFloatError        = errors.New("Not a float")
	NotBoolError         = errors.New("Not a boolean")
	NotArrayError        = errors.New("Not a array")
	NotMapError          = errors.New("Not a map")
	NotFixedDataError    = errors.New("Not a fixed data")
	IncompleteError      = errors.New("Not complete yet")
	IllegalMapKeyError   = errors.New("Iillegal map key")
	CanNotCountError     = errors.New("this format can not be count")
	FormatError          = errors.New("Unknown Format")

	TypeName = map[Type]string{
		Ext:     `Ext`,
		Str:     `Str`,
		Bin:     `Bin`,
		Int:     `Int`,
		Float:   `Float`,
		Map:     `Map`,
		Array:   `Array`,
		Bool:    `Bool`,
		Nil:     `Nil`,
		Unknown: `Unknown`,
	}
)

type Type uint8
type ExtType int8

func (t Type) String() string {
	return TypeName[t]
}
