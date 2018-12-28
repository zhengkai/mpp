package mpp_test

import (
	"testing"

	"github.com/vmihailenco/msgpack"
	"github.com/zhengkai/mpp"
)

type boolDemo struct {
	True  bool   `msgpack:"t"`
	False bool   `msgpack:"f"`
	No    string `msgpack:"n"`
}

func TestGetBool(t *testing.T) {

	b, _ := msgpack.Marshal(boolDemo{
		True: true,
		No:   `no`,
	})

	var r bool
	var err error

	r, err = mpp.GetBool(b, `t`)
	if r != true || err != nil {
		t.Error(`get bool fail when true(t)`)
	}

	r, err = mpp.GetBool(b, `f`)
	if r != false || err != nil {
		t.Error(`get bool fail when false(f)`)
	}

	r, err = mpp.GetBool(b, `n`)
	if r != false || err != mpp.ErrNotBool {
		t.Error(`get bool fail when str(n)`)
	}

	r, err = mpp.GetBool(b, `not-exists`)
	if r != false || err != mpp.ErrKeyPathNotFound {
		t.Error(`get bool fail when wrong path`, err)
	}

	b, _ = msgpack.Marshal(123)
	r, err = mpp.GetBool(b, `not-exists`)
	if r != false || err != mpp.ErrKeyPathNotFound {
		t.Error(`get bool fail when wrong path`)
	}

	var tp mpp.Type
	b, _ = msgpack.Marshal(nil)
	_, tp, err = mpp.Get(b)
	if tp != mpp.Nil || err != nil {
		t.Error(`get bil fail`)
	}
}
