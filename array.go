package mpp

func ArrayEach(
	b []byte,
	cb func(i int64, v []byte, t Type) (isContinue bool),
	key ...string,
) (err error) {

	b, err = byPath(b, key)
	if err != nil {
		return
	}

	f := GetFormat(b[0])

	count, metaLen, pErr := getCount(f, b)

	if pErr != nil || f.Type() != Array {
		return ErrNotArray
	}

	b = b[metaLen:]

	var i int64

	for {

		isContinue := cb(i, b, GetFormat(b[0]).Type())
		if !isContinue {
			break
		}

		i++
		if i >= count {
			break
		}

		b = b[GetByteLen(b):]
	}

	return
}
