package main

import (
	"fmt"

	"github.com/zhengkai/mpp"
)

func testStr() {
	var list = [...]int{
		1,
		31,
		32,
		33,
		65535,
		65536,
		65537,
	}

	for _, i := range list {

		file := fmt.Sprintf(`str/len-%d`, i)

		v := loadDemo(file)

		if len(v) < 1 {
			fmt.Println(`load file`, file, `fail`)
			return
		}

		s, err := mpp.GetStr(v)

		if err != nil {
			fmt.Println(`error when load str`, i, err)
			return
		}

		if len(s) != i {
			fmt.Println(`str len not match`, i)
			return
		}
	}

	fmt.Println(`all str test pass`)
}
