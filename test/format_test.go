package mpp_test

import (
	"testing"

	"github.com/zhengkai/mpp"
)

func TestFormat(t *testing.T) {

	nameList := make(map[string]bool)

	for i := 0; i < 256; i++ {

		b := byte(i)
		s := mpp.GetFormat(b).String()

		if s == `unknown` {
			t.Error(`wrong format`)
		}

		f := mpp.Format(i).String()

		if i == 0 || (i > 5 && i < 192) || i >= 224 {

			if f != `Unknown` {
				t.Error(`wrong format name`)
			}
			continue
		}

		if _, ok := nameList[f]; ok {
			t.Errorf(`duplicate format name "%s"`, f)
		}
		nameList[f] = true
	}

	formatNum := 37
	if len(nameList) != formatNum {
		t.Errorf(`formats number is not match %d / %d`, len(nameList), formatNum)
	}
}
