package main

import (
	"fmt"
	"reflect"
)

type Struct struct {
	a int
}

func main() {
	var v any = new(Struct)
	wrap(v)
}

func wrap(v any) {
	test(v)
}

func test(p any) {
	v := reflect.ValueOf(p).Elem()
	t := reflect.TypeOf(p)
	fmt.Println(v, t)
	fmt.Print(v.Kind())
	return
}
