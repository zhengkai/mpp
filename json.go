package mpp

import (
	"bytes"
	"fmt"
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

		toJsonMap(v, buf)

	case String:
		buf.WriteByte('"')
		s, _ := getSlashedStr(v)
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

func toJsonStr(v []byte, buf *bytes.Buffer, t Type) {

	switch t {
	case String:
		buf.WriteByte('"')
		s, _ := getSlashedStr(v)
		buf.WriteString(s)
		buf.WriteByte('"')

	case Integer:

		num, _, _ := getInt(v)

		buf.WriteRune('"')
		buf.WriteString(strconv.FormatInt(num, 10))
		buf.WriteRune('"')

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
		buf.WriteRune('"')
		buf.WriteString(s)
		buf.WriteRune('"')

	default:
		buf.WriteString(`"unknown"`)
	}
}

func toJsonArray(v []byte, buf *bytes.Buffer) {

	buf.WriteByte('[')

	err := ArrayEach(v, func(i int64, v []byte, t Type) bool {

		if i > 0 {
			buf.WriteByte(',')
		}

		ToJson(v).WriteTo(buf)

		return true
	})

	if err != nil {
		fmt.Println(`json array error`, err)
	}

	buf.WriteByte(']')
}

func toJsonMap(v []byte, buf *bytes.Buffer) {

	buf.WriteByte('{')
	MapEach(v, func(i int64, k []byte, kt Type, v []byte, t Type) bool {

		if i > 0 {
			buf.WriteByte(',')
		}

		// ToJson(k).WriteTo(buf)

		toJsonStr(k, buf, kt)
		buf.WriteByte(':')
		ToJson(v).WriteTo(buf)

		return true
	})
	buf.WriteByte('}')
}
