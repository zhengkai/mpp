package mpp

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
