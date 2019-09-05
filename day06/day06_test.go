package day06

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	g := NewGrid(5, 5)

	cols := len(g)
	if cols != 5 {
		t.Errorf("The number of collumns should be 5, not %d", cols)
	}

	for _, c := range g {
		rows := len(c)
		if rows != 5 {
			t.Errorf("The number of rows should be 5, not %d", rows)
		}
	}
}

func TestTurnOnOff(t *testing.T) {
	g := NewGrid(1000, 1000)
	g.TurnOn(Coordinate{0, 0}, Coordinate{999, 999})
	for x, col := range g {
		for y := range col {
			if g[x][y] == 0 {
				t.Errorf("All lights should have been lit (%d,%d)", x, y)
			}
		}
	}

	g.TurnOn(Coordinate{0, 0}, Coordinate{999, 999})
	g.TurnOff(Coordinate{499, 499}, Coordinate{500, 500})
	for x, col := range g {
		for y := range col {
			if x >= 499 && y >= 499 && x <= 500 && y <= 500 {
				if g[x][y] > 0 {
					t.Errorf("The four middle lights should not be on (%d,%d)", x, y)
				}
			} else if g[x][y] == 0 {
				t.Errorf("Only the four middle lights should not be on (%d,%d)", x, y)
			}
		}
	}
}

func TestToggle(t *testing.T) {
	g := NewGrid(1000, 1000)
	g.Toggle(Coordinate{0, 0}, Coordinate{999, 0})
	for x, col := range g {
		if col[0] == 0 {
			t.Errorf("The entire first row should not be lit (%d,%d)", x, 0)
		}
	}

	for x, col := range g {
		for y := 1; y < len(col); y++ {
			if col[y] > 0 {
				t.Errorf("Only the first row should not be lit (%d,%d)", x, y)
			}
		}
	}
}

func TestGetOn(t *testing.T) {
	g := NewGrid(1000, 1000)
	g.TurnOn(Coordinate{0, 0}, Coordinate{999, 999})
	on := g.GetOn()
	if on != 1000000 {
		t.Errorf("The number of lit lights should be 1000 000 (%d)", on)
	}

	g.Toggle(Coordinate{0, 0}, Coordinate{999, 0})
	on = g.GetOn()
	if on != 999000 {
		t.Errorf("The number of lit lights should be 999 000 (%d)", on)
	}

	g.TurnOn(Coordinate{0, 0}, Coordinate{999, 999})
	on = g.GetOn()
	if on != 1000000 {
		t.Errorf("The number of lit lights should be 1000 000 (%d)", on)
	}

	g.TurnOff(Coordinate{499, 499}, Coordinate{500, 500})
	on = g.GetOn()
	if on != 999996 {
		t.Errorf("The number of lit lights should be 999 996 (%d)", on)
	}
}

func TestParseInstruction(t *testing.T) {
	i, err := ParseInstruction("turn off 660,55 through 986,197")
	if err != nil {
		t.Error(err)
	}

	if i.Do != Off {
		t.Error("The action should be turn off")
	}
	if i.From.X != 660 && i.From.Y == 55 {
		t.Error("The from coordinate is not correctly parsed")
	}
	if i.To.X != 986 && i.From.Y != 197 {
		t.Error("The to coordinate is not correctly parsed")
	}
}

func TestTurnOff2(t *testing.T) {
	g := NewGrid(5, 5)
	g.TurnOn(Coordinate{X: 0, Y: 0}, Coordinate{X: 4, Y: 4})
	g.TurnOn(Coordinate{X: 0, Y: 0}, Coordinate{X: 4, Y: 4})

	g.TurnOff2(Coordinate{X: 1, Y: 2}, Coordinate{X: 3, Y: 4})
	if g[0][0] != 2 || g[1][0] != 2 || g[2][0] != 2 || g[3][0] != 2 || g[4][0] != 2 ||
		g[0][1] != 2 || g[1][1] != 2 || g[2][1] != 2 || g[3][1] != 2 || g[4][1] != 2 ||
		g[0][2] != 2 || g[1][2] != 1 || g[2][2] != 1 || g[3][2] != 1 || g[4][2] != 2 ||
		g[0][3] != 2 || g[1][3] != 1 || g[2][3] != 1 || g[3][3] != 1 || g[4][3] != 2 ||
		g[0][4] != 2 || g[1][4] != 1 || g[2][4] != 1 || g[3][4] != 1 || g[4][4] != 2 {
		t.Error("The turn off did not happen correctly")
	}

	g.TurnOff2(Coordinate{X: 0, Y: 1}, Coordinate{X: 2, Y: 3})
	if g[0][0] != 2 || g[1][0] != 2 || g[2][0] != 2 || g[3][0] != 2 || g[4][0] != 2 ||
		g[0][1] != 1 || g[1][1] != 1 || g[2][1] != 1 || g[3][1] != 2 || g[4][1] != 2 ||
		g[0][2] != 1 || g[1][2] != 0 || g[2][2] != 0 || g[3][2] != 1 || g[4][2] != 2 ||
		g[0][3] != 1 || g[1][3] != 0 || g[2][3] != 0 || g[3][3] != 1 || g[4][3] != 2 ||
		g[0][4] != 2 || g[1][4] != 1 || g[2][4] != 1 || g[3][4] != 1 || g[4][4] != 2 {
		t.Error("The turn off did not happen correctly")
	}

	g.TurnOff2(Coordinate{X: 0, Y: 1}, Coordinate{X: 2, Y: 3})
	if g[0][0] != 2 || g[1][0] != 2 || g[2][0] != 2 || g[3][0] != 2 || g[4][0] != 2 ||
		g[0][1] != 0 || g[1][1] != 0 || g[2][1] != 0 || g[3][1] != 2 || g[4][1] != 2 ||
		g[0][2] != 0 || g[1][2] != 0 || g[2][2] != 0 || g[3][2] != 1 || g[4][2] != 2 ||
		g[0][3] != 0 || g[1][3] != 0 || g[2][3] != 0 || g[3][3] != 1 || g[4][3] != 2 ||
		g[0][4] != 2 || g[1][4] != 1 || g[2][4] != 1 || g[3][4] != 1 || g[4][4] != 2 {
		t.Errorf("The turn off did not happen correctly\r\n%v", g)
	}
}

func TestToggle2(t *testing.T) {
	g := NewGrid(5, 5)
	g.Toggle2(Coordinate{X: 1, Y: 2}, Coordinate{X: 3, Y: 4})
	if g[0][0] != 0 || g[1][0] != 0 || g[2][0] != 0 || g[3][0] != 0 || g[4][0] != 0 ||
		g[0][1] != 0 || g[1][1] != 0 || g[2][1] != 0 || g[3][1] != 0 || g[4][1] != 0 ||
		g[0][2] != 0 || g[1][2] != 2 || g[2][2] != 2 || g[3][2] != 2 || g[4][2] != 0 ||
		g[0][3] != 0 || g[1][3] != 2 || g[2][3] != 2 || g[3][3] != 2 || g[4][3] != 0 ||
		g[0][4] != 0 || g[1][4] != 2 || g[2][4] != 2 || g[3][4] != 2 || g[4][4] != 0 {
		t.Error("The toggle did not happen correctly")
	}

	g.Toggle2(Coordinate{X: 0, Y: 1}, Coordinate{X: 2, Y: 3})
	if g[0][0] != 0 || g[1][0] != 0 || g[2][0] != 0 || g[3][0] != 0 || g[4][0] != 0 ||
		g[0][1] != 2 || g[1][1] != 2 || g[2][1] != 2 || g[3][1] != 0 || g[4][1] != 0 ||
		g[0][2] != 2 || g[1][2] != 4 || g[2][2] != 4 || g[3][2] != 2 || g[4][2] != 0 ||
		g[0][3] != 2 || g[1][3] != 4 || g[2][3] != 4 || g[3][3] != 2 || g[4][3] != 0 ||
		g[0][4] != 0 || g[1][4] != 2 || g[2][4] != 2 || g[3][4] != 2 || g[4][4] != 0 {
		t.Error("The toggle did not happen correctly")
	}
}

func TestGetTotalBightNess(t *testing.T) {
	g := NewGrid(5, 5)
	g.Toggle2(Coordinate{X: 1, Y: 2}, Coordinate{X: 3, Y: 4})
	g.Toggle2(Coordinate{X: 0, Y: 1}, Coordinate{X: 2, Y: 3})

	b := g.GetTotalBrightness()
	if b != 36 {
		t.Errorf("The total brightness should be 34, not %d", b)
	}

	g = NewGrid(1000, 1000)
	g.TurnOn(Coordinate{X: 0, Y: 0}, Coordinate{X: 0, Y: 0})
	b = g.GetTotalBrightness()
	if b != 1 {
		t.Errorf("The total brightness should be 1, not %d", b)
	}

	g = NewGrid(1000, 1000)
	g.Toggle2(Coordinate{X: 0, Y: 0}, Coordinate{X: 999, Y: 999})
	b = g.GetTotalBrightness()
	if b != 2000000 {
		t.Errorf("The total brightness should be 2000000, not %d", b)
	}
}
