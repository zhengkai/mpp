package mpp

func GetBool(v []byte) (b bool, err error) {

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
