package mpp

import (
	"fmt"
)

func Debug(s []byte) {
	for _, b := range s {

		s := ` `
		i := int(b)
		if i >= 32 && i <= 127 {
			s = string(b)
		}

		fmt.Printf("%s %02x %08b %d\n", s, b, b, i)
	}
	fmt.Println()
}

func getFormat() {
}

func JSONtoMSGP(s []byte) []byte {
	// var buff bytes.Buffer
	return s
}

func MSGPtoJSON(s []byte) []byte {
	l := len(s)
	i := 0
	for i < l {
		v := s[i]
		t, _, j := GetInType([]byte{v})
		fmt.Printf("v hex = %02x, v = %08b, t = %08b, i = %02x\n", v, v, t, j)

		i++
	}
	return s
}
