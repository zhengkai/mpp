package mpp

import (
	"encoding/binary"
	"math"
)

func GetFloat(v []byte, key ...string) (f float64, err error) {

	if len(key) > 0 {
		v, _, err = Get(v, key...)
		if err != nil {
			return
		}
	}

	format := GetFormat(v[0])

	switch format {

	case FormatFloat32:

		bits := binary.BigEndian.Uint32(v[1:5])
		f = float64(math.Float32frombits(bits))

	case FormatFloat64:

		bits := binary.BigEndian.Uint64(v[1:9])
		f = math.Float64frombits(bits)

	default:

		err = NotFloatError
	}

	return
}
