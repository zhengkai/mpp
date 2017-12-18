package mpp

import (
	"bytes"
	"strconv"
)

func ToJson(v []byte) *bytes.Buffer {

	buf := bytes.NewBuffer([]byte{})

	_, t, err := Get(v)

	if err != nil {
		buf.WriteString(`!ERROR!`)
		buf.WriteString(err.Error())
		return buf
	}

	switch t {

	case Array:

		toJsonArray(v, buf)

	case Map:
		buf.WriteByte('{')
		buf.WriteByte('}')

	case String:
		buf.WriteByte('"')
		s, _, _ := getStr(v)
		buf.WriteString(s)
		buf.WriteByte('"')

	case Integer:

		num, _, _ := getInt(v)

		buf.WriteString(strconv.FormatInt(num, 10))

	case Boolean,
		Nil:

		s := ``
		switch InType(v[0]) {
		case InTypeFalse:
			s = `false`
		case InTypeTrue:
			s = `true`
		case InTypeNil:
			s = `null`
		}
		buf.WriteString(s)

	default:
		buf.WriteString(`unknown`)

	}

	return buf
}

func toJsonArray(v []byte, buf *bytes.Buffer) {

	buf.WriteByte('[')
	ArrayEach(v, func(i int64, v []byte, t Type, err error) bool {

		if i > 0 {
			buf.WriteByte(',')
		}

		ToJson(v).WriteTo(buf)

		// fmt.Println(i, t, mpp.GetByteLen(v), err)
		return true
	})
	buf.WriteByte(']')
}
