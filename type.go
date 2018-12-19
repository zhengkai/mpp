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
	ErrIncomplete      = errors.New("Not complete yet")
	ErrCanNotCount     = errors.New("this format can not be count")
)

// Type https://github.com/msgpack/msgpack/blob/master/spec.md#type-system
type Type uint8

// ExtType https://github.com/msgpack/msgpack/blob/master/spec.md#extension-types
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
