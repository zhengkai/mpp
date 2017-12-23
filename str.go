package mpp

func GetStr(v []byte) (s string, err error) {

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

	s = string(r)
	return
}

func GetBin(v []byte) (r []byte, err error) {

	var tmp []byte
	tmp, _, _, err = getBin(v)

	if err != nil {
		return
	}

	copy(tmp, v)
	return
}

func GetUnsafeBin(v []byte) (r []byte, err error) {

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
