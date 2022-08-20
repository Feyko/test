package simd

import (
	"fmt"
	"unsafe"
)

type Object struct {
	Vec [4]float32
	R   [4]float32
}

func (o Object) GetVec() unsafe.Pointer {
	return unsafe.Pointer(&o.Vec)
}

func (o Object) GetResult() unsafe.Pointer {
	return unsafe.Pointer(&o.R)
}

//go:noescape
func _MultiplyBytes(vec1, vec2 unsafe.Pointer) (ret unsafe.Pointer)

func MultiplyAndAdd(someObj Object) {
	fmt.Println(_MultiplyBytes(someObj.GetVec(), unsafe.Pointer(&[16]byte{2, 2, 2, 2})))
}
