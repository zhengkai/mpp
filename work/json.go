package main

import (
	"fmt"

	"github.com/zhengkai/mpp"
)

func testJSON() {

	i := 0
	for {
		i++
		if i > 10 {
			break
		}

		file := fmt.Sprintf(`json/%d`, i)

		fileJson := file + `.json`

		v := loadDemo(file)

		t1 := mpp.ToJson(v).String()

		// fmt.Println(`final json =`, b.String())

		t2 := string(loadDemoData(fileJson))

		if t1 != t2 {
			fmt.Println(`json not match`, i, len(t1), len(t2))
			fmt.Println(t1)
			fmt.Println(t2)
			return
		}

		fmt.Println(`json match`, i)
	}
}
