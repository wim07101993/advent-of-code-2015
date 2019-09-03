package day04

import "testing"

func TestHash(t *testing.T) {
	h := hash([]byte("abcdef609043"))
	if h != "000001dbbfa3a5c83a2d506429c7b00e" {
		t.Errorf("The hash is supposed to be 000001dbbfa3a5c83a2d506429c7b00e, not %s", h)
	}
}
