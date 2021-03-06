package mpp

import (
	"encoding/binary"
	"math"
)

// GetFloat by path key
func GetFloat(v []byte, key ...string) (f float64, err error) {

	v, err = byPath(v, key)
	if err != nil {
		return
	}

	format := GetFormat(v[0])

	switch format {

	case FormatFloat32:

		if len(v) < 5 {
			err = ErrInvalid
			return
		}
		bits := binary.BigEndian.Uint32(v[1:5])
		f = float64(math.Float32frombits(bits))

	case FormatFloat64:

		if len(v) < 9 {
			err = ErrInvalid
			return
		}
		bits := binary.BigEndian.Uint64(v[1:9])
		f = math.Float64frombits(bits)

	default:

		err = ErrNotFloat
	}

	return
}
