package mpp

// GetBool by path key
func GetBool(v []byte, key ...string) (b bool, err error) {

	v, err = byPath(v, key)
	if err != nil {
		return
	}

	f := GetFormat(v[0])

	switch f {

	case FormatTrue:

		b = true

	case FormatFalse:

		b = false

	default:

		err = ErrNotBool
	}

	return
}
