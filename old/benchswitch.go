package main

var s string

var m = map[string]func(){
	"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA": A,
	"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB": B,
	"CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC": C,
}

func Switch(s string) {
	switch s {
	case "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA":
		A()
	case "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB":
		B()
	case "CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC":
		C()
	}
}

func SwitchInline(s string) {
	switch s {
	case "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA":
		s += "A"
	case "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB":
		s += "B"
	case "CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC":
		s += "C"
	}
}

func SwitchR(s string) func() {
	switch s {
	case "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA":
		return A
	case "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB":
		return B
	case "CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC":
		return C
	}
	return nil
}

func A() {
	s += "A"
}

func B() {
	s += "B"
}

func C() {
	s += "C"
}

type T struct {
}

func (t T) A() {
	s += "A"
}

func (t T) B() {
	s += "B"
}

func (t T) C() {
	s += "C"
}

func (t T) A_dispatch() {
	A()
}

func (t T) B_dispatch() {
	B()
}

func (t T) C_dispatch() {
	C()
}

func (t T) A_resolve() func() {
	return A
}

func (t T) B_resolve() func() {
	return B
}

func (t T) C_resolve() func() {
	return C
}
