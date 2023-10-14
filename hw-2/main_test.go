package main

import (
	"testing"
)

func TestUnpackClassic(t *testing.T) {
	s := "a4bc2d5e"
	e := "aaaabccddddde"
	c := UnpackString(s)
	if c != e {
		t.Fatalf("bad Unpack for %s: got %s expected %s", s, c, e)
	}
}

func TestUnpackWrong(t *testing.T) {
	s := "45"
	e := ""
	c := UnpackString(s)
	if c != e {
		t.Fatalf("bad Unpack for %s: got %s expected %s", s, c, e)
	}
}

func TestUnpackWithoutDecimals(t *testing.T) {
	s := "abcd"
	e := "abcd"
	c := UnpackString(s)
	if c != e {
		t.Fatalf("bad Unpack for %s: got %s expected %s", s, c, e)
	}
}
