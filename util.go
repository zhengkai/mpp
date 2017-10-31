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

func skip(v []byte, j int) ([]byte, error) {

	if j < 1 {
		return v, nil
	}

	var i = 0

	for {

		it, _, _ := GetInType(v)
		t := getType(it)

		switch t {

		case String:
			_, end, err := getString(v)
			if err != nil {
				return nil, WrongFormatError
			}
			v = v[end:]

		case Number:

			return nil, IncompleteError

		default:
			return nil, IncompleteError
		}

		i++
		if i >= j {
			break
		}
	}

	return v, nil
}
