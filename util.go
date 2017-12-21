package mpp

func isPack(it Format) bool {
	switch it {
	case FormatFixArray,
		FormatArray16,
		FormatArray32,
		FormatFixMap,
		FormatMap16,
		FormatMap32:

		return true
	}
	return false
}

func skip(v []byte, j int64) ([]byte, error) {

	if j < 1 {
		return v, nil
	}

	var i int64

	for {
		i++
		if i > j {
			break
		}

		dataLen := getByteLen(v)

		if dataLen < 1 {
			return nil, IncompleteError
		}

		v = v[dataLen:]
	}

	return v, nil
}
