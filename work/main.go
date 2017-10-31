package main

import (
	"encoding/binary"
	"fmt"

	"github.com/zhengkai/mpp"
	// "github.com/zhengkai/mpp"
)

func x(a []byte) {
	a = a[uint8(12):]

	fmt.Println(a)
}

func main() {

	r := []byte{}

	f := binary.BigEndian.Uint16([]byte{0, 0, 0, 0, 0, 0, 0, 1})
	fmt.Println(f)

	// mpp.Debug([]byte(`ka there is now cow level`))
	r = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(r[1:3], r[3:5])

	r = loadDemo(`string`)
	mpp.Debug(r)

	x, _, _ := mpp.Get(r, `foo`)
	fmt.Println(`get`)

	s, _ := mpp.GetString(x)
	fmt.Println(s)

	r = loadDemo(`arrays`)
	mpp.Debug(r)

	x, _, _ = mpp.Get(r, `2`)
	fmt.Println(`get`, len(x))

	s, _ = mpp.GetString(x)
	fmt.Println(s)

	return

	/*
		r = loadDemo(`string`)
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
	*/

}
