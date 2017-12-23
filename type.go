package mpp

import (
	"errors"
)

const (
	Unknown = Type(iota)
	Str
	Bin
	Ext
	Int
	Float
	Map
	Array
	Bool
	Nil
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

	TypeName = map[Type]string{}
)

type Type uint8
type ExtType int8

func (t Type) String() (s string) {
	switch t {
	case Ext:
		s = `Ext`
	case Str:
		s = `Str`
	case Bin:
		s = `Bin`
	case Int:
		s = `Int`
	case Float:
		s = `Float`
	case Map:
		s = `Map`
	case Array:
		s = `Array`
	case Bool:
		s = `Bool`
	case Nil:
		s = `Nil`
	case Unknown:
		s = `Unknown`
	}
	return
}
