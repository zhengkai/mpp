package mpp

import (
	"strconv"
)

func Get(v []byte, key ...string) (r []byte, t Type, err error) {

	tier := len(key)

	it, al, _ := GetInType(v)
	t = getType(it)
	if tier < 1 {
		r = v
		return
	}

	var findKey string

	switch t {
	case Object:

		var i int64
		j, _, _ := getObjLen(v)
		v = v[al:]

		findKey, key = key[0], key[1:]

		for {
			if i > j {
				err = KeyPathNotFoundError
				break
			}
			i++

			if len(v) == 0 {
				err = WrongFormatError
			}

			var subErr error

			k, end, subErr := getStr(v)

			if subErr != nil {
				err = WrongFormatError
				return
			}

			v = v[end:]

			if k == findKey {
				return Get(v, key...)
			}

			v, subErr = skip(v, 1)
			if subErr != nil {
				err = WrongFormatError
			}
		}

	case Array:

		findKey, key = key[0], key[1:]

		var i int64
		var tI int
		tI, err = strconv.Atoi(findKey)
		if err != nil {
			err = KeyPathNotFoundError
			return
		}
		i = int64(tI)

		if i > 0 {

			j, _, _ := getArrayLen(v)
			v = v[al:]
			if i > j {
				err = KeyPathNotFoundError
				return
			}

			var subErr error
			v, subErr = skip(v, i)

			if subErr != nil {
				err = WrongFormatError
				return
			}
		} else {

			v = v[al:]
		}

		return Get(v, key...)

	default:
		err = KeyPathNotFoundError
	}

	r = v
	return
}
