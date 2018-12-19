package mpp_test

import (
	"fmt"
	"testing"

	"github.com/vmihailenco/msgpack"
	"github.com/zhengkai/mpp"
)

var (
	intBound = [...]int64{
		0,
		1,
		2,
		126,
		127,
		128,
		255,
		256,
		65535,
		65536,
		4294967295,
		4294967296,
		9223372036854775807,
		-1,
		-2,
		-32,
		-33,
		-128,
		-129,
		-256,
		-257,
		-65536,
		-65537,
		-4294967296,
		-4294967297,
		-9223372036854775808,
	}
)

func ExampleGetInt() {
	v := []byte{0xce, 0x00, 0x01, 0xe2, 0x40}

	i, _ := mpp.GetInt(v)
	fmt.Println(i)
	// Output:
	// 123456
}

func TestGetInt(t *testing.T) {

	// solo

	for _, i := range intBound {

		b, _ := msgpack.Marshal(i)
		// fmt.Println(i, b)
		ti, err := mpp.GetInt(b)

		if err != nil {
			t.Errorf(`test int %d throw err %s`, i, err)
		}

		if ti != i {
			t.Errorf(`test int %d fail`, i)
		}

		// bound check
		k := len(b)
		for j := 0; j < k; j++ {
			_, err := mpp.GetInt(b[:j])
			if err != mpp.ErrInvalid {
				t.Error(`no error when data broken`)
			}
		}
	}

	// err

	b, _ := msgpack.Marshal(`not int`)
	i, err := mpp.GetInt(b)

	if i > 0 {
		t.Error(`int not zero when err`)
	}
	if err != mpp.ErrNotInt {
		t.Error(`not throw NotIntError`)
	}

	// path

	b, _ = msgpack.Marshal(2017)

	i, err = mpp.GetInt(b, `0`)

	if i != 0 || err != mpp.ErrKeyPathNotFound {
		t.Error(`path not found`)
	}
}
