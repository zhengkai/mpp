package mpp

import (
	"unicode/utf8"
)

// GetStr by path key
func GetStr(v []byte, key ...string) (s string, err error) {

	v, err = byPath(v, key)
	if err != nil {
		return
	}

	var r []byte
	var t Type
	r, _, t, err = getBin(v)
	if t != Str || err != nil {
		err = ErrNotStr
		return
	}

	if !utf8.Valid(r) {
		err = ErrMalformedUtf8
		return
	}

	s = string(r)
	return
}

// GetBin by path key
func GetBin(v []byte, key ...string) (r []byte, err error) {

	v, err = byPath(v, key)
	if err != nil {
		return
	}

	var tmp []byte
	tmp, _, _, err = getBin(v)

	if err != nil {
		return
	}

	r = make([]byte, len(tmp))

	copy(r, tmp)
	return
}

// GetUnsafeBin by path key
func GetUnsafeBin(v []byte, key ...string) (r []byte, err error) {

	v, err = byPath(v, key)
	if err != nil {
		return
	}

	r, _, _, err = getBin(v)

	return
}

func getBin(v []byte) (r []byte, end int, t Type, err error) {

	f := GetFormat(v[0])
	t = f.Type()
	if t != Str && t != Bin {
		err = ErrNotBin
		return
	}

	count, metaLen, _ := getCount(f, v)

	end = metaLen + count

	if len(v) < end {
		err = ErrInvalid
		return
	}

	r = v[metaLen:end]

	return
}
