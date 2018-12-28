package mpp_test

import (
	"testing"

	"github.com/vmihailenco/msgpack"
	"github.com/zhengkai/mpp"
)

var (
	sl = strList{
		`one`,
		`two`,
		`three`,
	}

	demoV = demoObj{
		String: `abc`,
		Int:    123,
		Array:  sl,
	}
)

type strList []string

type demoObj struct {
	String string  `msgpack:"str"`
	Int    int64   `msgpack:"int"`
	Array  strList `msgpack:"list"`
}

func TestGet(t *testing.T) {

	b, _ := msgpack.Marshal(demoV)

	if len(b) < 1 {
		t.Error(`make demo msgpack fail`)
	}

	v3, mt, _ := mpp.Get(b, `list`, `2`)

	if `Str` != mt.String() {
		t.Error(`get by path, return wrong type`)
		return
	}

	a3, _ := msgpack.Marshal(`three`)

	for i := range a3 {
		if a3[i] != v3[i] {
			t.Error(`get byte array error`)
			return
		}
	}

	s3, _ := mpp.GetStr(v3)

	if `three` != s3 {
		t.Error(`get by path fail`)
	}
}

func TestGetEmtpyMap(t *testing.T) {

	demo := make(map[string]int)

	b, _ := msgpack.Marshal(demo)

	_, _, err := mpp.Get(b, `nothing`)
	if err != mpp.ErrKeyPathNotFound {
		t.Error(`should be fail but not`)
	}

	for i := range b {
		broken := b[:i:i]
		mpp.Get(broken, `nothing`)
	}

	demo[`one`] = 1
	demo[`two`] = 2

	b, _ = msgpack.Marshal(demo)

	_, _, err = mpp.Get(b, `nothing`)
	if err != mpp.ErrKeyPathNotFound {
		t.Error(`should be fail but not`)
	}

	for i := range b {
		broken := b[:i:i]
		mpp.Get(broken, `two`)
	}
}

func TestGetFail(t *testing.T) {

	b, _ := msgpack.Marshal(demoV)

	path := [][]string{
		{`list`, `15`},
		{`list`, `-1`},
		{`list`, `no`},
		{`no`},
		{`str`, `0`},
		{`int`, `foo`},
		{`int`, `foo`, `bar`},
	}

	for _, row := range path {
		_, _, err := mpp.Get(b, row...)
		if err != mpp.ErrKeyPathNotFound && err != mpp.ErrInvalid {
			t.Error(`get by error path, but not fail, path = `)
			t.Error(row)
			t.Error(err)
		}
	}
}

func TestGetMap(test *testing.T) {

	b, _ := msgpack.Marshal(demoV)

	v, t, err := mpp.Get(b, `int`)
	if t != mpp.Int || err != nil {
		test.Error(`fail when get from map`)
	}

	i, err := mpp.GetInt(v)
	if i != demoV.Int || err != nil {
		test.Error(`fail when get int from map`)
	}

	l := len(b)
	for i := 1; i < l; i++ {
		mpp.Get(b[:i], `int`)
	}
}
