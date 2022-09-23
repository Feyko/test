package main

import "fmt"

func main() {
	s := []string{"a"}
	p := &s[0]
	*p = "b"
	fmt.Println(s)
}
