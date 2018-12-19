package mpp

func MapEach(
	b []byte,
	cb func(i int, k []byte, kt Type, v []byte, vt Type) (isContinue bool),
	key ...string,
) (err error) {

	b, err = byPath(b, key)
	if err != nil {
		return
	}

	f := GetFormat(b[0])

	count, metaLen, pErr := getCount(f, b)

	if pErr != nil || f.Type() != Map {
		return ErrNotMap
	}

	b = b[metaLen:]

	var i int

	for {

		k := b

		if len(k) == 0 {
			err = ErrInvalid
			return
		}

		kt := GetFormat(k[0]).Type()

		trim := GetByteLen(b)
		if len(b) < trim+1 {
			err = ErrInvalid
			return
		}
		b = b[trim:]

		vt := GetFormat(b[0]).Type()

		isContinue := cb(i, k, kt, b, vt)
		if !isContinue {
			break
		}

		i++
		if i >= count {
			break
		}

		trim = GetByteLen(b)
		if len(b) < trim+1 {
			err = ErrInvalid
			return
		}
		b = b[trim:]
	}

	return
}
