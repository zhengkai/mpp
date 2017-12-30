package mpp_test

import (
	"testing"

	"github.com/vmihailenco/msgpack"
	"github.com/zhengkai/mpp"
)

type mapList []demoObj

func TestMapEach(t *testing.T) {

	v := mapList{
		{String: `foo`, Int: 2017},
		{String: `bar`, Int: 2018},
	}

	b, _ := msgpack.Marshal(v)

	var s string
	mpp.MapEach(
		b,
		func(i int64, k []byte, kt mpp.Type, v []byte, vt mpp.Type) bool {

			if vt == mpp.Str {
				s, _ = mpp.GetStr(v)
				return false
			}

			return true
		},
		`1`,
	)

	if s != `bar` {
		t.Error(`foreach map fail`)
	}

	var err error

	err = mpp.MapEach(
		b,
		func(i int64, k []byte, kt mpp.Type, v []byte, vt mpp.Type) bool {

			if vt == mpp.Str {
				s, _ = mpp.GetStr(v)
				return false
			}

			return true
		},
		`2`,
	)

	if err != mpp.ErrKeyPathNotFound {
		t.Error(`should be fail but not`)
	}

	err = mpp.MapEach(
		b,
		func(i int64, k []byte, kt mpp.Type, v []byte, vt mpp.Type) bool {

			if vt == mpp.Str {
				s, _ = mpp.GetStr(v)
				return false
			}

			return true
		},
	)

	if err != mpp.ErrNotMap {
		t.Error(`should be fail but not`)
	}
}
