package mpp

func MapEach(
	b []byte,
	cb func(i int64, k []byte, kt Type, v []byte, vt Type) (isContinue bool),
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

	var i int64

	for {

		k := b

		kt := GetFormat(k[0]).Type()

		b = b[GetByteLen(b):]

		vt := GetFormat(b[0]).Type()

		isContinue := cb(i, k, kt, b, vt)
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
