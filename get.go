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

// Get value by key
func Get(v []byte, key ...string) (r []byte, t Type, err error) {

	tier := len(key)

	f := GetFormat(v[0])
	t = f.Type()

	if tier < 1 {
		r = v
		return
	}

	switch t {

	case Map:

		r, t, err = getMap(v, key, f)

	case Array:

		r, t, err = getArray(v, key, f)

	default:

		err = ErrKeyPathNotFound
	}

	return
}

func getMap(v []byte, key []string, f Format) (r []byte, t Type, err error) {

	var findKey string

	count, metaLen, err := getCount(f, v)
	if err != nil {
		return
	}
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
}

func getArray(v []byte, key []string, f Format) (r []byte, t Type, err error) {

	var findKey string
	findKey, key = key[0], key[1:]

	var i int64
	var tI int
	tI, err = strconv.Atoi(findKey)
	if err != nil || tI < 0 {
		err = ErrKeyPathNotFound
		return
	}
	i = int64(tI)

	count, metaLen, err := getCount(f, v)
	if err != nil {
		return
	}

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
}
