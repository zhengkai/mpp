package mpp

import (
	"fmt"
)

func Debug(s []byte) {
	for _, b := range s {
		fmt.Printf("%s %x %#08b\n", string(b), b, b)
	}
}
