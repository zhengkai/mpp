package mpp_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/zhengkai/mpp"
)

var (
	jsonId = [...]int{
		2,
		5,
		8,
	}
)

func TestToJson(t *testing.T) {

	for _, i := range jsonId {

		fileJson := fmt.Sprintf(`testdata/toJson/%d.json`, i)
		fileMsgp := fmt.Sprintf(`testdata/toJson/%d.bin`, i)

		json, err := ioutil.ReadFile(fileJson)
		if err != nil || len(json) < 100 {
			t.Errorf(`read test data fail, file = %s`, fileJson)
		}

		msgp, err := ioutil.ReadFile(fileMsgp)
		if err != nil || len(msgp) < 100 {
			t.Errorf(`read test data fail, file = %s`, fileMsgp)
		}

		test := mpp.ToJson(msgp).Bytes()

		if string(test) != string(json) {
			t.Errorf(`ToJson fail, sample %d`, i)
		}
	}
}
