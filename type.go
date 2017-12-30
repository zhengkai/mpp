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
	ErrInvalid         = errors.New("Invalid or malformed data")
	ErrKeyPathNotFound = errors.New("Key path not found")
	ErrMalformedUtf8   = errors.New("Malformed UTF-8 characters, possibly incorrectly encoded")
	ErrNotArray        = errors.New("Not an array")
	ErrNotBin          = errors.New("Not a binary")
	ErrNotBool         = errors.New("Not a boolean")
	ErrNotFloat        = errors.New("Not a float")
	ErrNotInt          = errors.New("Not a integer")
	ErrNotMap          = errors.New("Not a map")
	ErrNotStr          = errors.New("Not a string")

	NotMapError        = errors.New("Not a map")
	NotFixedDataError  = errors.New("Not a fixed data")
	IncompleteError    = errors.New("Not complete yet")
	IllegalMapKeyError = errors.New("Iillegal map key")
	CanNotCountError   = errors.New("this format can not be count")
	FormatError        = errors.New("Unknown Format")
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
	default:
		s = `Unknown`
	}
	return
}
