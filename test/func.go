package main

import (
	"io/ioutil"
)

func loadDemo(name string) (r []byte) {

	// fmt.Println("\n------- load demo [", name, "] -------\n")

	r, _ = ioutil.ReadFile(`../php/demo/` + name + `.bin`)
	return r
}

func loadDemoData(name string) (r []byte) {
	r, _ = ioutil.ReadFile(`../php/demo/` + name)
	return r
}
