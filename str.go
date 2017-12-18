package mpp

func GetStr(v []byte) (s string, err error) {
	s, _, err = getStr(v)
	return
}

func getStr(v []byte) (s string, end int64, err error) {

	_, t, metaLen, ext, parseErr := parseMeta(v)

	if t != String || parseErr != nil {
		err = NotStringError
		return
	}

	end = metaLen + ext

	s = string(v[metaLen:end])

	return
}
