package mpp

func isPack(it InType) bool {
	switch it {
	case InTypeFixArray,
		InTypeArray16,
		InTypeArray32,
		InTypeFixMap,
		InTypeMap16,
		InTypeMap32:

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
