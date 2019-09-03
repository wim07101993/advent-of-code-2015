package day04

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	h := hash([]byte("abcdef609043"))
	sh := fmt.Sprintf("%x", h)
	if sh != "000001dbbfa3a5c83a2d506429c7b00e" {
		t.Errorf("The hash is supposed to be 000001dbbfa3a5c83a2d506429c7b00e, not %s", h)
	}
}

func TestCheckHash(t *testing.T) {
	h := hash([]byte("abcdef609043"))
	if !checkHash(h) {
		t.Error("The provide hash is a good hash")
	}

	h = hash([]byte("not a good hash"))
	if checkHash(h) {
		t.Error("The provided hash is not a good hash")
	}
}

func TestFindHash(t *testing.T) {
	s := New(func() []byte {
		return []byte("abcdef")
	}, 4)

	i := s.findHash()
	if i != 609043 {
		t.Errorf("The integer to append should be 609043, not %d", i)
	}
}
