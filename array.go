package mpp

func ArrayEach(in []byte, cb func(i int64, v []byte, t Type) (isContinue bool)) (err error) {

	_, t, metaLen, ext, parseErr := parseMeta(in)
	if t != Array || parseErr != nil {
		return NotArrayError
	}

	in = in[metaLen:]

	var i int64

	for {

		_, t, _, _, parseErr := parseMeta(in)
		if parseErr != nil {
			err = parseErr
			return
		}

		isContinue := cb(i, in, t)
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
