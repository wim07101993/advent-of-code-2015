package day09

import (
	"fmt"
	"syreclabs.com/go/faker"
	"testing"
)

func TestParsePointToPoint(t *testing.T) {
	// arrange
	nameP1 := faker.Address().Country()
	nameP2 := faker.Address().Country()
	distance := faker.RandomInt(0, 900)

	// act
	ptp := ParsePointToPoint(fmt.Sprintf("%s to %s = %d",
		nameP1, nameP2, distance))

	// assert
	if ptp.p1 != nameP1 {
		t.Errorf("Expected point1 to be %s, not %s", nameP1, ptp.p1)
	}
	if ptp.p2 != nameP2 {
		t.Errorf("Expected point2 to be %s, not %s", nameP2, ptp.p2)
	}
	if ptp.distance != distance {
		t.Errorf("Expected distance to be %d, not %d",
			distance, ptp.distance)
	}
}
