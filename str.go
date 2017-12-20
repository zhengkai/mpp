package mpp

import "bytes"

func GetStr(v []byte) (s string, err error) {
	s, _, err = getStr(v, false)
	return
}

func getSlashedStr(v []byte) (s string, err error) {
	s, _, err = getStr(v, true)
	return
}

func getStr(v []byte, isSlash bool) (s string, end int64, err error) {

	_, t, metaLen, ext, parseErr := parseMeta(v)

	if t != String || parseErr != nil {
		err = NotStringError
		return
	}

	end = metaLen + ext

	v = v[metaLen:end]

	if isSlash {
		v = bytes.Replace(v, []byte{'\\'}, []byte{'\\', '\\'}, -1)
		v = bytes.Replace(v, []byte{'"'}, []byte{'\\', '"'}, -1)
	}

	s = string(v)

	return
}
