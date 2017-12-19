package main

import (
	"fmt"
	"reflect"

	"github.com/zhengkai/mpp"
)

func testMap() {

	v := loadDemo(`string`)

	result := make(map[string]string)

	assert := make(map[string]string)
	assert[`abc`] = `def`
	assert[`a1`] = `a2`
	assert[`foo`] = `bar`

	err := mpp.MapEach(v, func(i int64, k []byte, kt mpp.Type, v []byte, vt mpp.Type) bool {

		ks, _ := mpp.GetStr(k)
		vs, _ := mpp.GetStr(v)

		result[ks] = vs

		return true
	})

	if err != nil {
		fmt.Println(`map each fail`)
		return
	}

	if !reflect.DeepEqual(result, assert) {
		fmt.Println(`map each compare fail`)
		return
	}

	fmt.Println(`map each test pass`)
}
