package mpp

func Get(v []byte, key ...string) (r []byte, t Type, err error) {

	tier := len(key)

	for {

		it, _ := GetInType(v[0])

		t = getType(it)

		if tier == 0 {
			break
		}

		if t != Object && t != Array {
			return nil, NotExist, KeyPathNotFoundError
		}
	}

	return
}
