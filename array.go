package mpp

/*
func ArrayEach(in []byte, cb func(i int64, v []byte, t Type, err error) (isGoon bool)) (err error) {

	_, t, metaLen, ext, parseErr := parseMeta(in)
	if t != Array || parseErr != nil {
		return NotArrayError
	}

	in = in[metaLen:]

	var i int64

	for {

		it, t, metaLen, ext, parseErr := parseMeta(in)

		isGoon := cb(i, in, t, parseErr)
		if !isGoon {
			break
		}

		i++
		if i > ext {
			break
		}

		in = in[metaLen+ext:]
	}

	return
}
*/
