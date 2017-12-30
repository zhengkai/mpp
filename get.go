package mpp

import (
	"strconv"
)

func byPath(b []byte, key []string) (r []byte, err error) {
	if len(key) == 0 {
		return b, nil
	}
	r, _, err = Get(b, key...)
	return
}

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

		count, metaLen, _ := getCount(f, v)
		v = v[metaLen:]

		findKey, key = key[0], key[1:]

		var i int64

		for {

			if i >= count {
				err = ErrKeyPathNotFound
				return
			}
			i++

			if len(v) == 0 {
				err = ErrInvalid
				return
			}

			var subErr error

			kv, end, _, subErr := getBin(v)
			if subErr != nil {
				err = ErrInvalid
				return
			}

			k := string(kv)

			v = v[end:]

			if k == findKey {
				return Get(v, key...)
			}

			v, subErr = skip(v, 1)
			if subErr != nil {
				err = ErrInvalid
				return
			}
		}

	case Array:

		findKey, key = key[0], key[1:]

		var i int64
		var tI int
		tI, err = strconv.Atoi(findKey)
		if err != nil || tI < 0 {
			err = ErrKeyPathNotFound
			return
		}
		i = int64(tI)

		count, metaLen, _ := getCount(f, v)

		v = v[metaLen:]

		if i >= count {
			err = ErrKeyPathNotFound
			return
		}

		var subErr error
		v, subErr = skip(v, i)

		if subErr != nil {
			err = ErrInvalid
			return
		}

		return Get(v, key...)

	default:

		err = ErrKeyPathNotFound
	}

	return
}
