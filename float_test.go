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

func TestGetFloat64(t *testing.T) {

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

	t.Logf(`test float64 done`)
}

func TestGetFloat32(t *testing.T) {

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

	t.Logf(`test float32 done`)
}
