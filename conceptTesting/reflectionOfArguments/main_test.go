package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := &post{ArgumentHeight: 10, ArgumenAngle: 10.5}
		checkSwitch(p)
	}
}
