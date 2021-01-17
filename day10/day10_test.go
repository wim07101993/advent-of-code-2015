package day10

import "testing"

func TestIncrease(t *testing.T) {
	bs := Increase([]byte("211"))

	s := string(bs)
	if s != "1221" {
		t.Errorf("Expected 1221, got %s", s)
	}
}

func TestSequence(t *testing.T) {
	bs := Sequence([]byte("1"), 5)
	s := string(bs)
	if s != "312211" {
		t.Errorf("Expected 312211, got %s", s)
	}
}
