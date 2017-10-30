package main

import (
	"fmt"
	"io/ioutil"
)

func loadDemo(name string) (r []byte) {

	fmt.Println("\n------- load demo [", name, "] -------\n")

	r, _ = ioutil.ReadFile(`../php/demo/` + name + `.bin`)
	return r
}
