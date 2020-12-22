package main

import (
	"testing"
)

func TestCheckNumArgs(t *testing.T) {

	// Assign CheckNumArgs to f
	f := CheckNumArgs

	// Test valid options
	if got := f(2, false); got != nil {
		t.Errorf("Want: %v, Got: %v", nil, got)
	}

	if got := f(1, true); got != nil {
		t.Errorf("Want: %v, Got: %v", nil, got)
	}

	// Check if fail on too few args
	if got := f(0, true); got == nil {
		t.Errorf("Got: %v, Want: %v", got, "Too few arguments")
	}

	if got := f(1, false); got == nil {
		t.Errorf("Got: %v, Want: %v", got, "Too few arguments")
	}

	// Check if fail on too many args
	if got := f(3, false); got == nil {
		t.Errorf("Got: %v, Want: %v", got, "Too many arguments")
	}

	if got := f(2, true); got == nil {
		t.Errorf("Got: %v, Want: %v", got, "Too many arguments")
	}
}
