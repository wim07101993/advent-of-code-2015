package day03

import "testing"

func TestCalcNumberOfDeliveredHouses1(t *testing.T) {
	x := calcNumberOfDeliveredHouses(">")
	if x != 2 {
		t.Errorf("The number of delivered houses is 2, not %d", x)
	}
}

func TestCalcNumberOfDeliveredHouses2(t *testing.T) {
	x := calcNumberOfDeliveredHouses("^>v<")
	if x != 4 {
		t.Errorf("The number of delivered houses is 4, not %d", x)
	}
}

func TestCalcNumberOfDeliveredHouses3(t *testing.T) {
	x := calcNumberOfDeliveredHouses("^v^v^v^v^v")
	if x != 2 {
		t.Errorf("The number of delivered houses is 2, not %d", x)
	}
}
