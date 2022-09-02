package main

import (
	"testing"
)

var calls = []string{"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB", "CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"}

var r string

func BenchmarkDispatchMap(b *testing.B) {
	s = ""
	for i := 0; i < b.N; i++ {
		for _, call := range calls {
			m[call]()
		}
	}
	r = s
}

func BenchmarkSwitchInline(b *testing.B) {
	s = ""
	for i := 0; i < b.N; i++ {
		for _, call := range calls {
			SwitchInline(call)
		}
	}
	r = s
}

func BenchmarkSwitchCall(b *testing.B) {
	s = ""
	for i := 0; i < b.N; i++ {
		for _, call := range calls {
			Switch(call)
		}
	}
	r = s
}

func BenchmarkSwitchReturn(b *testing.B) {
	s = ""
	for i := 0; i < b.N; i++ {
		for _, call := range calls {
			SwitchR(call)()
		}
	}
	r = s
}
