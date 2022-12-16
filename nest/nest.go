package nest

type Child struct {
	A string
	C byte
	d struct{}
}

func (c Child) Ayo() {

}

type ChildTwo struct {
	A string
	D struct{}
}

func (c ChildTwo) Ayo(s string) {

}
