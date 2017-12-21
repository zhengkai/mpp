package mpp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

func Get(v []byte, key ...string) (r []byte, t Type, err error) {

	tier := len(key)
	buf := bytes.NewBuffer(v)

	if tier > 0 {
		for {
		}
	}

	return
}

type layer struct {
	len   int
	cur   int
	isMap bool
	skip  bool
}

const (
	searchEmpty = uint(iota)
	searchArray = uint(iota)
	searchMap   = uint(iota)
)

func searchKey(buf *bytes.Buffer) (r []byte, err error) {

	// var layer []int
	// var offset int

	// len := len(v)
	searchTag := searchEmpty

	i := 0
	j := 0

	item := uint32(0)

	var isMap bool

	for {
		if i > j {
			err = errors.New(`length error`)
			break
		}
		if searchTag == searchEmpty {
			isMap, item, err = getSearchType(FormatMap32, item, buf)
			if err != nil {
				return
			}
			if isMap {
				searchTag = searchMap
			} else {
				searchTag = searchArray
			}
			if item == 0 {

				t := FormatArray16

				switch t {

				case FormatArray16,
					FormatMap16:

					read := buf.Next(2)
					item = binary.BigEndian.Uint32([]byte{0, 0, read[0], read[1]})

				case
					FormatArray32,
					FormatMap32:

					read := buf.Next(4)
					item = binary.BigEndian.Uint32(read)

				default:
					err = errors.New(`type check error`)
					return
				}
			}
		}

		fmt.Println(item)

		return
	}
	return
}

func isPack(it Format) bool {
	switch it {
	case FormatFixArray,
		FormatArray16,
		FormatArray32,
		FormatFixMap,
		FormatMap16,
		FormatMap32:

		return true
	}
	return false
}

func getSearchType(t Format, item uint32, buf *bytes.Buffer) (isMap bool, ritem uint32, err error) {

	switch t {

	case FormatFixArray:

		isMap = false

	case FormatArray16:

		isMap = false
		read := buf.Next(2)
		item = binary.BigEndian.Uint32([]byte{0, 0, read[0], read[1]})

	case FormatArray32:

		isMap = false
		read := buf.Next(4)
		item = binary.BigEndian.Uint32(read)

	case FormatFixMap:

		isMap = true

	case FormatMap16:

		isMap = true
		read := buf.Next(2)
		item = binary.BigEndian.Uint32([]byte{0, 0, read[0], read[1]})

	case FormatMap32:

		isMap = true
		read := buf.Next(4)
		item = binary.BigEndian.Uint32(read)

	default:
		err = errors.New(`not array or map`)
	}

	return
}
