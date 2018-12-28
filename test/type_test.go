package mpp_test

import (
	"testing"

	"github.com/zhengkai/mpp"
)

func TestType(t *testing.T) {

	for i := 0; i < 256; i++ {

		f := mpp.Format(i)
		typ := f.Type()
		st := typ.String()

		isUnknown := false
		if f.String() == `Unknown` || f == mpp.FormatNa {
			isUnknown = true
		}

		if (typ == mpp.Unknown) != isUnknown || (st == `Unknown`) != isUnknown {
			t.Errorf(`format "%s (0x%x, %d)" type error`, f.String(), i, i)
		}
	}
}
