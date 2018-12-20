package mpp

func ArrayEach(
	b []byte,
	cb func(i int, v []byte, t Type) (isContinue bool),
	key ...string,
) (err error) {

	b, err = byPath(b, key)
	if err != nil {
		return
	}

	f := GetFormat(b[0])

	count, metaLen, err := getCount(f, b)

	if err != nil || f.Type() != Array {
		return ErrNotArray
	}

	if len(b) < metaLen+1 {
		err = ErrInvalid
		return
	}

	b = b[metaLen:]

	var i int

	for {

		isContinue := cb(i, b, GetFormat(b[0]).Type())
		if !isContinue {
			break
		}

		i++
		if i >= count {
			break
		}

		l := GetByteLen(b)
		if len(b) < l+1 {
			err = ErrInvalid
			return
		}

		b = b[l:]
	}

	return
}
