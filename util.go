package mpp

import "fmt"

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

		it, metaLen, iPack := GetInType(v)
		t := getType(it)

		switch t {

		case String:

			_, end, err := getStr(v)
			if err != nil {
				return nil, WrongFormatError
			}
			v = v[end:]

		case Integer,
			Float,
			Boolean,
			Null:

			v = v[getMetaLen(it):]

		case Object:

			fmt.Println(`obj`, metaLen, iPack)

		case Array:

			fmt.Println(`array`, metaLen, iPack)

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
