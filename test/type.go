package main

import (
	"fmt"

	"github.com/zhengkai/mpp"
)

func testType() {

	it := mpp.Format(0xcd)

	// t := it.Type()

	fmt.Println(it.String())
	// fmt.Println(t.Name())
}
