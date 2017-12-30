package mpp_test

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/zhengkai/mpp"
)

func TestStr(t *testing.T) {

	var s string
	var b []byte
	var r []byte
	var err error

	// utf-8 str

	demo := `UTF-8 汉字测试`
	b, _ = msgpack.Marshal(demo)

	s, err = mpp.GetStr(b)

	if s != demo || err != nil {
		t.Error(`getStr utf-8 fail`)
	}

	s, err = mpp.GetStr(b, `foo`)
	if s != `` || err != mpp.ErrKeyPathNotFound {
		t.Error(`path not found`)
	}

	// utf-8 malformed

	malformed := make([]byte, len(b))
	copy(malformed, b)

	malformed[len(malformed)-3] = byte(97)

	s, err = mpp.GetStr(malformed)

	if s != `` || err != mpp.ErrMalformedUtf8 {
		t.Error(`malformed data should be error, but not`)
	}

	// bin

	bin := md5.Sum([]byte(`YES RPG 3`))

	b, _ = msgpack.Marshal(bin)

	r, err = mpp.GetBin(b)
	if err != nil {
		t.Error(`get bin fail`)
	}

	if len(r) != len(bin) {
		fmt.Println(r)
		t.Errorf(`bin len not match %d / %d`, len(r), len(bin))
	}

	for idx, bt := range r {
		if bt != bin[idx] {
			t.Error(`bin not match`)
		}
	}

	// bin safe

	b[len(b)-1] = byte(255)

	isSame := true
	for idx, bt := range r {
		if bt != bin[idx] {
			isSame = false
		}
	}

	// bin path

	r, err = mpp.GetBin(b, `foo`)
	if r != nil || err != mpp.ErrKeyPathNotFound {
		t.Error(`path not found`)
	}

	/*
		fmt.Println()
		fmt.Printf("  r = %x\n", r)
		fmt.Printf("  b = %x\n", b[len(b)-16:])
		fmt.Printf("bin = %x\n", bin)
		fmt.Println()
	*/

	if !isSame {
		t.Error(`bin not safe`)
	}

	// unsafe bin

	b, _ = msgpack.Marshal(bin)

	r, err = mpp.GetUnsafeBin(b)
	if err != nil {
		t.Error(`get bin fail`)
	}

	if len(r) != len(bin) {
		fmt.Println(r)
		t.Errorf(`bin len not match %d / %d`, len(r), len(bin))
	}

	for idx, bt := range r {
		if bt != bin[idx] {
			t.Error(`bin not match`)
		}
	}

	// unsafe bin not safe

	b[len(b)-1] = byte(255)

	isSame = true
	for idx, bt := range r {
		if bt != bin[idx] {
			isSame = false
		}
	}

	if isSame {
		t.Error(`unsafe bin is safe`)
	}

	// unsafe bin path

	r, err = mpp.GetUnsafeBin(b, `foo`)
	if r != nil || err != mpp.ErrKeyPathNotFound {
		t.Error(`path not found`)
	}

	// type error

	bInt, _ := msgpack.Marshal(1514701937)

	s, err = mpp.GetStr(bInt)
	if s != `` || err != mpp.ErrNotStr {
		t.Error(`getStr type error fail`)
	}

	b, err = mpp.GetBin(bInt)
	if b != nil || err != mpp.ErrNotBin {
		t.Error(`getBin type error fail`)
	}

	b, err = mpp.GetUnsafeBin(bInt)
	if b != nil || err != mpp.ErrNotBin {
		t.Error(`getUnsafeBin error fail`)
	}

	// diffect length

	lenList := [...]int{
		1,
		31,
		32,
		33,
		65535,
		65536,
		65537,
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	char := []byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)
	charLen := len(char)

	for _, i := range lenList {

		data := make([]byte, i)

		for i, _ := range data {
			data[i] = char[rnd.Intn(charLen)]
		}

		s := string(data)
		b, _ := msgpack.Marshal(s)

		r, err := mpp.GetStr(b)
		if r != s || err != nil {
			t.Errorf(`str fail when len = %d`, len(s))
		}
	}
}
