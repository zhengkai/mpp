package mpp

func MapEach(in []byte, cb func(i int64, k []byte, kt Type, v []byte, vt Type) (isContinue bool)) (err error) {

	_, t, metaLen, ext, parseErr := parseMeta(in)
	if t != Map || parseErr != nil {
		return NotMapError
	}

	in = in[metaLen:]

	var i int64

	for {

		_, kt, _, _, parseErr := parseMeta(in)
		if parseErr != nil {
			err = parseErr
			return
		}

		k := in

		in = in[getByteLen(in):]

		_, vt, _, _, parseErr := parseMeta(in)
		if parseErr != nil {
			err = parseErr
			return
		}

		isContinue := cb(i, k, kt, in, vt)
		if !isContinue {
			break
		}

		i++
		if i >= ext {
			break
		}

		in = in[getByteLen(in):]
	}

	return
}
