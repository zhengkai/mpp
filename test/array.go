package main

import (
	"fmt"

	"github.com/zhengkai/mpp"
)

func testArray() {

	arg := [][]string{
		[]string{`a`, `array-a`, `2`},
		[]string{`a`, `array-b`, `0`},
		[]string{`b`, `array-b`, `1`},
		[]string{`c`, `array-b`, `3`},
		[]string{`x`, `array-b`, `2`, `0`},
		[]string{`y`, `array-b`, `2`, `1`},
		[]string{`a`, `array-c`, `0`},
		[]string{`b`, `array-c`, `1`},
		[]string{`c`, `array-c`, `3`},
	}

	for _, row := range arg {
		if !assertArray(row...) {
			return
		}
	}

	fmt.Println(`array test pass`)
}

func assertArray(arg ...string) bool {

	assert, file, key := arg[0], arg[1], arg[2:]

	s := getArrayStr(file, key...)
	if s != assert {
		fmt.Println(`array`, file, key, `fail`, s, assert)
		return false
	}
	return true
}

func getArrayStr(file string, key ...string) (s string) {

	v := loadDemo(file)

	x, _, err := mpp.Get(v, key...)

	if err != nil {
		fmt.Println(`error =`, err)
		return
	}

	s, _ = mpp.GetStr(x)
	return
}
