package main

import (
	"fmt"

	"github.com/zhengkai/mpp"
)

func testInt() {

	var list = [...]int64{
		0,
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

	for _, arrow := range [...]int64{1, -1} {

		for _, i := range list {

			i *= arrow
			file := fmt.Sprintf(`int/n%d`, i)

			v := loadDemo(file)

			f := mpp.GetFormat(v[0])

			if len(v) < 1 {
				fmt.Println(`load file`, file, `fail`, f)
				return
			}

			j, err := mpp.GetInt(v)
			if err != nil {
				fmt.Println(`error when load int`, f, i, err)
				return
			}

			if j != i {
				fmt.Println(`int not match`, f, i)
				return
			}

		}
	}

	fmt.Println(`all int test pass`)
}
