package mpp_test

import (
	"testing"

	"github.com/vmihailenco/msgpack"
	"github.com/zhengkai/mpp"
)

func TestArrayEach(t *testing.T) {

	var list [70000]int

	for i, _ := range list {
		list[i] = i
	}

	b, _ := msgpack.Marshal(list)

	var counter int64

	err := mpp.ArrayEach(
		b,
		func(i int64, v []byte, t mpp.Type) bool {

			if i > 100 {
				return false
			}

			counter += i

			return true
		},
	)

	if counter != 5050 || err != nil {
		t.Error(`array fail when sum(1 to 100)`)
	}
}

func TestArrayEachFail(t *testing.T) {

	b, _ := msgpack.Marshal(true)

	var counter int64

	err := mpp.ArrayEach(
		b,
		func(i int64, v []byte, t mpp.Type) bool {

			if i > 100 {
				return false
			}

			counter += i

			return true
		},
	)

	if err == mpp.ErrNotArray {
		return
	} else {
		t.Error(`array each should fail, but not`)
	}

	if counter > 0 { // code not reach here
		t.Error(`test wrong`)
	}
}

func TestArrayEachFailWithPath(t *testing.T) {

	b, _ := msgpack.Marshal(true)

	var counter int64

	err := mpp.ArrayEach(
		b,
		func(i int64, v []byte, t mpp.Type) bool {

			if i > 100 {
				return false
			}

			counter += i

			return true
		},
		`foo`,
		`bar`,
	)

	if err == mpp.ErrKeyPathNotFound {
		return
	} else {
		t.Error(`array each should fail, but not`)
	}

	if counter > 0 { // code not reach here
		t.Error(`test wrong`)
	}
}
