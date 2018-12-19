package mpp

import (
	"bytes"
	"encoding/json"
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

		toJSONArray(v, buf)

	case Map:

		toJSONMap(v, buf)

	case Str:

		s, _ := GetStr(v)
		j, _ := json.Marshal(s)
		buf.Write(j)

	case Int:

		num, _ := GetInt(v)
		j, _ := json.Marshal(num)
		buf.Write(j)

	case Float:

		num, _ := GetFloat(v)
		j, _ := json.Marshal(num)
		buf.Write(j)

	case Bool:

		b, _ := GetBool(v)
		j, _ := json.Marshal(b)
		buf.Write(j)

	case Nil:

		buf.WriteString(`null`)

	default:
		buf.WriteString(`unknown`)

	}

	return buf
}

func toJSONStr(v []byte, buf *bytes.Buffer, t Type) {

	switch t {

	case Str,
		Int,
		Bool,
		Nil:

		nb := ToJson(v)

		if t == Str {

			nb.WriteTo(buf)

		} else {

			buf.WriteByte('"')
			nb.WriteTo(buf)
			buf.WriteByte('"')
		}

	default:
		buf.WriteString(`"unknown"`)
	}
}

func toJSONArray(v []byte, buf *bytes.Buffer) {

	buf.WriteByte('[')

	err := ArrayEach(v, func(i int, v []byte, t Type) bool {

		if i > 0 {
			buf.WriteByte(',')
		}

		ToJson(v).WriteTo(buf)

		return true
	})

	if err != nil {
		// fmt.Println(`json array error`, err)
	}

	buf.WriteByte(']')
}

func toJSONMap(v []byte, buf *bytes.Buffer) {

	buf.WriteByte('{')
	MapEach(v, func(i int, k []byte, kt Type, v []byte, t Type) bool {

		if i > 0 {
			buf.WriteByte(',')
		}

		// ToJson(k).WriteTo(buf)

		toJSONStr(k, buf, kt)
		buf.WriteByte(':')
		ToJson(v).WriteTo(buf)

		return true
	})
	buf.WriteByte('}')
}
