package mpp

func MapEach(in []byte, cb func(i int64, k []byte, kt Type, v []byte, vt Type) (isContinue bool)) (err error) {

	f := GetFormat(in[0])

	count, metaLen, pErr := getCount(f, in)

	if pErr != nil || f.Type() != Map {
		return NotMapError
	}

	in = in[metaLen:]

	var i int64

	for {

		k := in

		kt := GetFormat(k[0]).Type()

		in = in[getByteLen(in):]

		vt := GetFormat(in[0]).Type()

		isContinue := cb(i, k, kt, in, vt)
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
