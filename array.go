package mpp

func ArrayEach(in []byte, cb func(i int64, v []byte, t Type) (isContinue bool)) (err error) {

	f := GetFormat(in[0])

	count, pErr := getCount(f, in)

	if pErr != nil || f.Type() != Array {
		return NotArrayError
	}

	in = in[f.MetaLen():]

	var i int64

	for {

		isContinue := cb(i, in, GetFormat(in[0]).Type())
		if !isContinue {
			break
		}

		i++
		if i >= count {
			break
		}

		in = in[getByteLen(in):]
	}

	return
}
