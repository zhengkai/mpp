package mpp

func GetBool(v []byte, key ...string) (b bool, err error) {

	if len(key) > 0 {
		v, _, err = Get(v, key...)
		if err != nil {
			return
		}
	}

	f := GetFormat(v[0])

	switch f {

	case FormatTrue:

		b = true

	case FormatFalse:

		b = false

	default:

		err = NotBoolError

		panic(`unknown type ` + f.String() + ` ` + f.Type().String())
	}

	return
}
