package mpp

import (
	"encoding/binary"
	"fmt"
)

func getArrayLen(v []byte) (r []byte, end int64, err error) {

	it, metaLen, iPack := GetInType(v)

	var offset int64
	var arrayLen int64

	if metaLen > 0 {
		arrayLen = int64(binary.BigEndian.Uint16(v[1:metaLen]))
	} else {
		arrayLen = int64(iPack)
	}

	start := offset + 1

	fmt.Println(arrayLen, it, start)

	return
}
