package main

import (
	"strconv"
	"testing"
)

func TestFactorial(t *testing.T) {
	if factorial(0) != 1 {
		t.Error("factorial(0) should be 1 but was " + strconv.Itoa(factorial(0)))
		// note that the other test statements below are still run as well!
		// this is a handy feature!
	}
	if factorial(1) != 1 {
		t.Error("factorial(1) should be 1 but was " + strconv.Itoa(factorial(1)))
	}
	if factorial(5) != 120 {
		t.Error("factorial(5) should be 120 but was " + strconv.Itoa(factorial(5)))
	}
}
