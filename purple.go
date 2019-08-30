package main

import (
	"io/ioutil"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/parser"
)

func main() {
	bs, _ := ioutil.ReadFile("test.js")
	ps, err := parser.ParseFile(nil, "", bs, 0)

	if err != nil {
		panic(err)
	}

	vm, _, _ := otto.Run(ps)
	out, _ := vm.Call("mapper", nil, 5)

	value, _ := out.ToInteger()
	println(value)
}
