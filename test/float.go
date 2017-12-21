package main

import (
	"fmt"
	"math"

	"github.com/zhengkai/mpp"
)

func testFloat() {

	v := loadDemo(`float-pi`)

	x, _ := mpp.GetFloat(v)

	if math.Pi == x {
		fmt.Println(`float pi test pass`)
		return
	}

	fmt.Println(`float pi test not pass`)
}
