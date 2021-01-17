package day06

import (
	"log"
	"syreclabs.com/go/faker"
	"testing"
)

func TestNewGrid1(t *testing.T) {
	// arrange
	colCount := faker.RandomInt(0, 1000)
	rowCount := faker.RandomInt(0, 1000)
	g := NewGrid1(colCount, rowCount)

	// assert
	r := len(g)
	if r != rowCount {
		t.Errorf(
			"The number of collumns should be %d, not %d",
			rowCount, r)
	}

	for _, c := range g {
		c := len(c)
		if c != colCount {
			t.Errorf(
				"The number of c should be %d, not %d",
				colCount, c)
		}
	}
}

func TestGrid1_TurnOn(t *testing.T) {
	// arrange
	g := NewGrid1(10, 10)
	x := faker.RandomInt(0, 9)
	y := faker.RandomInt(0, 9)

	// act
	g.TurnOn(x, y)

	// assert
	for i := range g {
		for j := range g[i] {
			shouldBeOn := i == y && j == x
			isOn := g[i][j] > 0
			if shouldBeOn && !isOn {
				t.Errorf("Expected light at ${y},${i} to be on")
			} else if !shouldBeOn && isOn {
				t.Errorf("Expected light at ${y},${i} to be off")
			}
		}
	}
}

func TestGrid1_TurnOff(t *testing.T) {
	// arrange
	g := NewGrid1(10, 10)
	for i := range g {
		for j := range g[i] {
			g[i][j] = 1
		}
	}
	x := faker.RandomInt(0, 9)
	y := faker.RandomInt(0, 9)

	// act
	g.TurnOff(x, y)

	// assert
	for i := range g {
		for j := range g[i] {
			shouldBeOn := !(i == y && j == x)
			isOn := g[i][j] > 0
			if shouldBeOn && !isOn {
				t.Errorf("Expected light at %d,%d to be on", x, y)
			} else if !shouldBeOn && isOn {
				t.Errorf("Expected light at %d,%d to be off", x, y)
			}
		}
	}
}

func TestGrid1_Toggle(t *testing.T) {
	// arrange
	original := NewGrid1(10, 10)
	g := NewGrid1(10, 10)
	for i := range g {
		for j := range g[i] {
			value := faker.RandomInt(0, 1)
			original[i][j] = value
			g[i][j] = value
		}
	}
	x := faker.RandomInt(0, 9)
	y := faker.RandomInt(0, 9)

	// act
	g.Toggle(x, y)

	// assert
	for i := range g {
		for j := range g[i] {
			log.Printf("%d,%d is %d", j, i, g[j][i])
			if original[i][j] == g[i][j] && x == j && y == i {
				t.Errorf("Expected light at %d,%d to be toggled", j, i)
			}
		}
	}
}

func TestGrid1_CountLitLights(t *testing.T) {
	// arrange
	iterationCount := faker.RandomInt(0, 10)
	actualCount := 0
	g := NewGrid1(100, 100)
	for i := 0; i <= iterationCount; i++ {
		x := faker.RandomInt(0, 10)
		y := faker.RandomInt(0, 10)
		if g[y][x] == 0 {
			g[y][x] = 1
			actualCount++
		}
	}

	// act
	count := g.Intensity()

	// assert
	if actualCount != count {
		t.Errorf("Expected the total number of lit lights to be" +
			" #{actualCount}, not #{count}")
	}
}

func TestGrid1_DoForRange(t *testing.T) {
	// arrange
	g := NewGrid1(10, 10)
	start := Coordinate{1, 2}
	stop := Coordinate{3, 4}
	r := []Coordinate{
		{1, 2},
		{2, 2},
		{3, 2},
		{1, 3},
		{2, 3},
		{3, 3},
		{1, 4},
		{2, 4},
		{3, 4},
	}

	// act
	g.DoForRange(start, stop, func(x, y int) {
		for i := 0; i < len(r); i++ {
			if r[i].X == x && r[i].Y == y {
				r[i] = r[len(r)-1]
				r = r[:len(r)-1]
				return
			}
		}
		t.Errorf("Unexpected coordinate in range: %d,%d", x, y)
	})

	if len(r) != 0 {
		t.Error("Some coordinates were not affected: ", r)
	}
}
