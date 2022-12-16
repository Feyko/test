package main

import "test/nest"

type Parent struct {
	A string
}

type SomeUnion interface {
	Parent | nest.Child
}

type Ayo interface {
	Ayo()
}

func main() {
	f(Parent{}.Ayo())
	p := Parent{}
}

func f(a Ayo) {
	a.Ayo()
}
