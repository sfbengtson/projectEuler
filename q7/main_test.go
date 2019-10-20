package main

import "testing"

func TestIsPrime(t *testing.T) {
	var a bool
	a = isPrime(2)
	if a != true {
		t.Error("expected 2 to be prime!")
	}

	a = isPrime(5)
	if a != true {
		t.Error("Expected 5 to be prime!")
	}

	a = isPrime(9)
	if a != false {
		t.Error("Expected 9 to not be prime!")
	}
}

func TestFindNthPrime(t *testing.T) {
	var a int
	a = findNthPrime(6)
	if a != 13 {
		t.Error("Expected 13 to be the 6th prime, got", a)
	}
}
