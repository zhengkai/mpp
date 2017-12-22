package mpp

import (
	"strconv"
)

func Get(v []byte, key ...string) (r []byte, t Type, err error) {

	tier := len(key)

	f := GetFormat(v[0])
	t = f.Type()

	if tier < 1 {
		r = v
		return
	}

	var findKey string

	switch t {

	case Map:

		var i int64
		v = v[f.MetaLen():]

		findKey, key = key[0], key[1:]

		count, _ := getCount(f, v)

		for {
			if i > count {
				err = KeyPathNotFoundError
				break
			}
			i++

			if len(v) == 0 {
				err = WrongFormatError
			}

			var subErr error

			k, end, subErr := getStr(v, false)

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

		count, _ := getCount(f, v)

		v = v[f.MetaLen():]

		if i > count {
			err = KeyPathNotFoundError
			return
		}

		var subErr error
		v, subErr = skip(v, i)

		if subErr != nil {
			err = WrongFormatError
			return
		}

		return Get(v, key...)

	default:
		err = KeyPathNotFoundError
	}

	r = v
	return
}
