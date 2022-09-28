package main

import "fmt"

func main() {
	var s *string
	var a any = s
	fmt.Println(a == nil)
}
