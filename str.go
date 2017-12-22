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

	f := GetFormat(v[0])

	count, pErr := getCount(f, v)

	if pErr != nil || f.Type() != String {
		err = NotStringError
		return
	}

	metaLen := f.MetaLen()

	end = metaLen + count

	v = v[metaLen:end]

	if isSlash {
		v = bytes.Replace(v, []byte{'\\'}, []byte{'\\', '\\'}, -1)
		v = bytes.Replace(v, []byte{'"'}, []byte{'\\', '"'}, -1)
	}

	s = string(v)

	return
}
