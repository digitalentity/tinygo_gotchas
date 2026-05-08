package main

import (
	"github.com/ungerik/go3d/vec3"
)

// This example demonstrates how TinyGo handles heap allocations
// and how to monitor memory usage.

func PerformMath() vec3.T {
	a := vec3.T{1, 2, 3}
	b := vec3.T{4, 5, 6} // TinyGo allocates this on heap

	c := b.Scale(0.5)
	return vec3.Cross(&a, c)
}

func main() {
	PerformMath()
}
