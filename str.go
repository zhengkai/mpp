package mpp

import (
	"errors"
	"unicode/utf8"
)

var (
	InvalidUtf8Error = errors.New(`invalid utf8`)
)

func GetStr(v []byte, key ...string) (s string, err error) {

	if len(key) > 0 {
		v, _, err = Get(v, key...)
		if err != nil {
			return
		}
	}

	var r []byte
	var t Type
	r, _, t, err = getBin(v)
	if t != Str {
		err = NotStrError
		return
	}

	if err != nil {
		return
	}

	if !utf8.Valid(r) {
		err = InvalidUtf8Error
		return
	}

	s = string(r)
	return
}

func GetBin(v []byte, key ...string) (r []byte, err error) {

	if len(key) > 0 {
		v, _, err = Get(v, key...)
		if err != nil {
			return
		}
	}

	var tmp []byte
	tmp, _, _, err = getBin(v)

	if err != nil {
		return
	}

	copy(tmp, v)
	return
}

func GetUnsafeBin(v []byte, key ...string) (r []byte, err error) {

	if len(key) > 0 {
		v, _, err = Get(v, key...)
		if err != nil {
			return
		}
	}

	r, _, _, err = getBin(v)

	return
}

func getBin(v []byte) (r []byte, end int64, t Type, err error) {

	f := GetFormat(v[0])
	t = f.Type()
	if t != Str && t != Bin {
		err = NotBinError
		return
	}

	count, metaLen, _ := getCount(f, v)

	end = metaLen + count

	r = v[metaLen:end]

	return
}
