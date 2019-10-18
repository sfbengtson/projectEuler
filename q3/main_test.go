package main

import (
	"testing"
)

func TestPrimeFactor(t *testing.T) {
	a, b := 5, 0
	b = primeFactor(a)
	if b != 5 {
		t.Error("expected primefactor of 5 to be 5, got", b)
	}

	a = 4
	b = primeFactor(a)

	if primeFactor(a) != 2 {
		t.Error("expected primefactor to be 2, got", b)
	}

	a = 13195
	b = primeFactor(a)

	if primeFactor(a) != 29 {
		t.Error("expected primefactor to be 29, got", b)
	}
}
