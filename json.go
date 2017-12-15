package mpp

import (
	"bytes"
	"fmt"
	"strconv"
)

func ToJSON(v []byte) *bytes.Buffer {

	buf := bytes.NewBuffer([]byte{})

	ArrayEach(v, func(i int64, v []byte, t Type, err error) bool {

		fmt.Println(`#`, i, t)

		if i > 0 {
			buf.WriteByte(',')
		}

		switch t {

		case Array:
			buf.WriteByte('[')
			ToJSON(v).WriteTo(buf)
			buf.WriteByte(']')

		case Object:
			buf.WriteByte('{')
			buf.WriteByte('}')

		case String:
			buf.WriteByte('"')
			s, _, _ := getStr(v)

			fmt.Println(`str`, s)

			buf.WriteString(s)
			buf.WriteByte('"')

		case Integer:

			num, _, _ := getInt(v)

			buf.WriteString(strconv.FormatInt(num, 10))

		default:
			buf.WriteString(`unknown`)

		}

		// fmt.Println(i, t, mpp.GetByteLen(v), err)
		return true
	})

	return buf
}
