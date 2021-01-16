package day07

import "testing"

func TestProgram(t *testing.T) {
	p := NewProgram([]string{
		"x AND y -> d",
		"456 -> y",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"123 -> x",
		"y RSHIFT 2 -> g",
		"NOT x -> h",
		"NOT y -> i",
	})
	p.Execute()
	if p.Variables["d"] != 72 {
		t.Errorf("The value of d is 72, not %d", p.Variables["d"])
	}
	if p.Variables["e"] != 507 {
		t.Errorf("The value of e is 507, not %d", p.Variables["e"])
	}
	if p.Variables["f"] != 492 {
		t.Errorf("The value of f is 492, not %d", p.Variables["f"])
	}
	if p.Variables["g"] != 114 {
		t.Errorf("The value of g is 114, not %d", p.Variables["g"])
	}
	if p.Variables["h"] != 65412 {
		t.Errorf("The value of h is 65412, not %d", p.Variables["h"])
	}
	if p.Variables["i"] != 65079 {
		t.Errorf("The value of i is 65079, not %d", p.Variables["i"])
	}
	if p.Variables["x"] != 123 {
		t.Errorf("The value of x is 123, not %d", p.Variables["x"])
	}
	if p.Variables["y"] != 456 {
		t.Errorf("The value of y is 456, not %d", p.Variables["x"])
	}
}
