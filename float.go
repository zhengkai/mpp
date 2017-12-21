package mpp

import (
	"encoding/binary"
	"math"
)

func GetFloat(v []byte) (f float64, err error) {

	it, t, metaLen, _, parseErr := parseMeta(v)
	if t != Float || parseErr != nil {
		err = NotIntegerError
		return
	}

	switch it {

	case FormatFloat32:
		bits := binary.BigEndian.Uint32(v[1:metaLen])
		f = float64(math.Float32frombits(bits))

	case FormatFloat64:
		bits := binary.BigEndian.Uint64(v[1:metaLen])
		f = math.Float64frombits(bits)
	}

	return
}
