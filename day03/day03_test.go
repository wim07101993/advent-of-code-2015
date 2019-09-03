package day03

import "testing"

func TestIncreaseCoordinateUp(t *testing.T) {
	c := coordinate{}
	increaseCoordinate('^', &c)
	if c.X != 0 {
		t.Error("The coordinate moved horizontally while it shouldn't have.")
	}
	if c.Y != 1 {
		t.Error("The coordinate did not move up however is should have.")
	}
}

func TestIncreaseCoordinateDown(t *testing.T) {
	c := coordinate{X: 1, Y: 3}
	increaseCoordinate('v', &c)
	if c.X != 1 {
		t.Error("The coordinate moved horizontally while it shouldn't have.")
	}
	if c.Y != 2 {
		t.Error("The coordinate did not move down however is should have.")
	}
}

func TestIncreaseCoordinateLeft(t *testing.T) {
	c := coordinate{X: 10, Y: 5}
	increaseCoordinate('<', &c)
	if c.X != 9 {
		t.Error("The coordinate did not move left however is should have.")
	}
	if c.Y != 5 {
		t.Error("THe coordinate moved vertically while it shouldn't have.")
	}
}

func TestIncreaseCoordinateRight(t *testing.T) {
	c := coordinate{}
	increaseCoordinate('>', &c)
	if c.X != 1 {
		t.Error("The coordinate did not move right however is should have.")
	}
	if c.Y != 0 {
		t.Error("The coordinate did not move vertical however is should have.")
	}
}

func TestCalcDeliveredHouses1(t *testing.T) {
	x := calcDeliveredHouses(">")
	if x != 2 {
		t.Errorf("The number of delivered houses is 2, not %d", x)
	}
}

func TestCalcDeliveredHouses2(t *testing.T) {
	x := calcDeliveredHouses("^>v<")
	if x != 4 {
		t.Errorf("The number of delivered houses is 4, not %d", x)
	}
}

func TestCalcDeliveredHouses3(t *testing.T) {
	x := calcDeliveredHouses("^v^v^v^v^v")
	if x != 2 {
		t.Errorf("The number of delivered houses is 2, not %d", x)
	}
}

func TestCalcDeliveredHousesWithRobot1(t *testing.T) {
	x := calcDeliveredHousesWithRobot("^v")
	if x != 3 {
		t.Errorf("The number of delivered houses is 3, not %d", x)
	}
}

func TestCalcDeliveredHousesWithRobot2(t *testing.T) {
	x := calcDeliveredHousesWithRobot("^>v<")
	if x != 3 {
		t.Errorf("The number of delivered houses is 3, not %d", x)
	}
}

func TestCalcDeliveredHousesWithRobot3(t *testing.T) {
	x := calcDeliveredHousesWithRobot("^v^v^v^v^v")
	if x != 11 {
		t.Errorf("The number of delivered houses is 11, not %d", x)
	}
}
