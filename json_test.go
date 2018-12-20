package mpp_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/zhengkai/mpp"
)

var (
	jsonID = [...]int{
		2,
		5,
		8,
	}
)

func TestInvalidData(t *testing.T) {
	mpp.ToJSON(nil)

	b := make([]byte, 256)
	for i := 0; i < 256; i++ {
		b[i] = byte(i)
	}

	// try all 3 bit array, make sure no panic
	for _, x := range b {
		for _, y := range b {
			for _, z := range b {
				_ = mpp.ToJSON([]byte{x, y, z})
			}
		}
	}
}

func TestToJSON(t *testing.T) {

	for _, i := range jsonID {

		fileJSON := fmt.Sprintf(`testdata/toJSON/%d.json`, i)
		fileMsgp := fmt.Sprintf(`testdata/toJSON/%d.bin`, i)

		json, err := ioutil.ReadFile(fileJSON)
		if err != nil || len(json) < 100 {
			t.Errorf(`read test data fail, file = %s`, fileJSON)
		}

		msgp, err := ioutil.ReadFile(fileMsgp)
		if err != nil || len(msgp) < 100 {
			t.Errorf(`read test data fail, file = %s`, fileMsgp)
		}

		test := mpp.ToJSON(msgp).Bytes()

		if string(test) != string(json) {
			t.Errorf(`ToJSON fail, sample %d`, i)
		}
	}
}

func Benchmark_ToJSON(b *testing.B) {

	b.StopTimer()

	fileJSON := fmt.Sprintf(`testdata/toJSON/%d.json`, 8)
	json, _ := ioutil.ReadFile(fileJSON)
	jsonStr := string(json)

	fileMsgp := fmt.Sprintf(`testdata/toJSON/%d.bin`, 8)
	v, _ := ioutil.ReadFile(fileMsgp)

	b.StartTimer()

	if len(v) < 1 {
		b.Errorf(`ToJSON output empty`)
	}

	if mpp.ToJSON(v).String() != jsonStr {
		b.Errorf(`ToJSON fail when benchmark`)
	}
}
