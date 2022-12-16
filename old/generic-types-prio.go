package main

type A[T any] struct {
}

func main() {
	a := A[string]{}
	b := types[int](a)
	b = b
	return
}

func types[U any, T any](a A[T]) A[U] {
	return A[U]{}
}
