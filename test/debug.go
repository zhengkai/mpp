package main

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
