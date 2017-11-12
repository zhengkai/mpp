package main

import (
	"fmt"

	"github.com/zhengkai/mpp"
	// "github.com/zhengkai/mpp"
)

func main() {

	testStr()
	testInt()
}

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

		fmt.Println(`str`, i, `pass`)
	}
}

func testInt() {

	var list = [...]int64{
		1,
		2,
		3,
		4,
		5,
		15,
		16,
		17,
		127,
		128,
		129,
		255,
		256,
		257,
		65535,
		65536,
		65537,
		4294967295,
		4294967296,
		4294967297,
		9223372036854775806,
		9223372036854775807,
	}

	for _, i := range list {
		file := fmt.Sprintf(`int/n%d`, i)

		v := loadDemo(file)

		if len(v) < 1 {
			fmt.Println(`load file`, file, `fail`)
			return
		}

		j, err := mpp.GetInt(v)
		if err != nil {
			fmt.Println(`error when load int`, i, err)
			return
		}

		if j != i {
			fmt.Println(`int not match`, i)
			return
		}

		fmt.Println(`int`, i, `pass`)
	}
}
