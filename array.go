package mpp

func ArrayEach(in []byte, cb func(i int64, v []byte, t Type, err error) (isContinue bool)) (err error) {

	_, t, metaLen, ext, parseErr := parseMeta(in)
	if t != Array || parseErr != nil {
		return NotArrayError
	}

	in = in[metaLen:]

	var i int64

	for {

		_, t, _, _, parseErr := parseMeta(in)

		isContinue := cb(i, in, t, parseErr)
		if !isContinue || parseErr != nil {
			break
		}

		i++
		if i >= ext {
			break
		}

		l := getByteLen(in)

		in = in[l:]
	}

	return
}
