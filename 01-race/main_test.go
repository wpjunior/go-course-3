package main

import (
	"testing"
)

func TestIncrParallell(t *testing.T) {
	result := incrParallell()
	if result != 4 {
		t.Error("The result must be 4")
	}
}
