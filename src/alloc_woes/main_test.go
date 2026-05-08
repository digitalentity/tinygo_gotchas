package main

import (
	"testing"
)

func TestPerformMathAllocations(t *testing.T) {
	allocs := testing.AllocsPerRun(100, func() {
		PerformMath()
	})
	if allocs > 0 {
		t.Errorf("PerformMath allocated %f times, expected 0", allocs)
	}
}
