package mpp_test

import (
	"math"
	"testing"

	"github.com/vmihailenco/msgpack"
	"github.com/zhengkai/mpp"
)

var (
	floatBound = [...]float64{
		math.E,
		math.Ln2,
		-math.E,
		-math.Ln2,
	}
)

func TestGetFloat(t *testing.T) {

	// float64

	for _, f := range floatBound {

		b, _ := msgpack.Marshal(f)
		tf, err := mpp.GetFloat(b)

		if err != nil {
			t.Errorf(`test float64 %v throw err %s`, f, err)
		}

		if tf != f {
			t.Errorf(`test float64 %v fail`, f)
		}
	}

	// float32

	for _, f64 := range floatBound {

		f := float32(f64)

		b, _ := msgpack.Marshal(f)

		if len(b) != 5 {
			t.Error(`can not encode float32`)
		}

		tf64, err := mpp.GetFloat(b)

		tf := float32(tf64)

		if err != nil {
			t.Errorf(`test float32 %v throw err %s`, f, err)
		}

		if tf != f {
			t.Errorf(`test float32 %v fail`, f)
		}
	}

	// path

	b, _ := msgpack.Marshal([...]float32{1.1, 3.14, 0.618})

	f, err := mpp.GetFloat(b, `2`)
	if float32(f) != 0.618 || err != nil {
		t.Error(`fail with path`)
	}

	// error

	var zero float64

	f, err = mpp.GetFloat(b, `3`)
	if f != zero || err != mpp.ErrKeyPathNotFound {
		t.Error(`fail with wrong path`)
	}

	f, err = mpp.GetFloat(b)
	if f != zero || err != mpp.ErrNotFloat {
		t.Error(`fail with not float`)
	}
}
