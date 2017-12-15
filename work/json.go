package main

import (
	"fmt"

	"github.com/zhengkai/mpp"
)

func testJSON() {

	file := `json/4`
	v := loadDemo(file)

	b := mpp.ToJSON(v)

	fmt.Println(`final json =`, b.String())
}
