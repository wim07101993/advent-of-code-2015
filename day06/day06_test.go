package day06

import "testing"

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
