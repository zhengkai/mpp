package main

import (
	"fmt"

	"github.com/zhengkai/mpp"
)

func main() {
	// mpp.Debug([]byte(`ka there is now cow level`))

	r := loadDemo(`string`)
	mpp.Debug(r)

	r = loadDemo(`int128`)
	mpp.Debug(r)

	r = loadDemo(`int-1`)
	mpp.Debug(r)

	r = loadDemo(`int1`)
	mpp.Debug(r)

	r = loadDemo(`array1`)
	mpp.Debug(r)

	r = loadDemo(`int109`)
	mpp.Debug(r)

	fmt.Println("\n---- msgp to json ----\n")

	rt, t, _ := mpp.Get(r)
	fmt.Println(t, rt)

	fmt.Println(`type =`, mpp.DebugGetType(mpp.InTypeArray16))

	rt, t, err := mpp.Get(r, `no`)
	fmt.Println(err)

}
