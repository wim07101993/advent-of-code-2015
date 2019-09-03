package day02

import "testing"

func TestParsePresent234(t *testing.T) {
	input := "2x3x4"

	p, err := ParsePresent(input)
	if err != nil {
		t.Error(err)
	}
	if p.Length != 2 {
		t.Errorf("The height is not parsed correctly (%d != 2)", p.Length)
	}
	if p.Width != 3 {
		t.Errorf("The width is not parsed correctly (%d != 3)", p.Width)
	}
	if p.Height != 4 {
		t.Errorf("The width is not parsed correctly (%d != 4)", p.Height)
	}
}

func TestParsePresent1110(t *testing.T) {
	input := "1x1x10"

	p, err := ParsePresent(input)
	if err != nil {
		t.Error(err)
	}
	if p.Length != 1 {
		t.Errorf("The height is not parsed correctly (%d != 1)", p.Length)
	}
	if p.Width != 1 {
		t.Errorf("The width is not parsed correctly (%d != 1)", p.Width)
	}
	if p.Height != 10 {
		t.Errorf("The width is not parsed correctly (%d != 10)", p.Height)
	}
}

func TestParseInput(t *testing.T) {
	input := []string{
		"2x3x4",
		"1x1x10",
	}

	ps, err := ParseInput(input)
	if err != nil {
		t.Error(err)
	}
	if len(ps) != 2 {
		t.Errorf("Expected two parsed presents, but got %d", len(ps))
	}
}

func TestCalcWrappingPaperArea234(t *testing.T) {
	p, _ := ParsePresent("2x3x4")
	a := p.CalcWrappingPaperArea()
	if a != 58 {
		t.Errorf("Present 2x3x4 should have a total wrapping paper area of 58, not %d", a)
	}
}

func TestCalcWrappingPaperArea1110(t *testing.T) {
	p, _ := ParsePresent("1x1x10")
	a := p.CalcWrappingPaperArea()
	if a != 43 {
		t.Errorf("Present 1x1x10 should have a total wrapping paper area of 43, not %d", a)
	}
}

func TestCalcRibbonLength234(t *testing.T) {
	p, _ := ParsePresent("2x3x4")
	l := p.CalcRibbonLength()
	if l != 34 {
		t.Errorf("Present 2x3x4 should have a total robbon length of 34, not %d", l)
	}
}

func TestCalcRibbonLength1110(t *testing.T) {
	p, _ := ParsePresent("1x1x10")
	l := p.CalcRibbonLength()
	if l != 14 {
		t.Errorf("Present 2x3x4 should have a total robbon length of 14, not %d", l)
	}
}

func TestSolver(t *testing.T) {
	input := func() []string {
		return []string{
			"2x3x4",
			"1x1x10",
		}
	}

	solver := New(input)
	a := solver.SolvePart1()
	if a != "101" {
		t.Errorf("The total wrapping paper area should be 101, not %s", a)
	}
	l := solver.SolvePart2()
	if l != "48" {
		t.Errorf("The total length ribbon should be 48, not %s", l)
	}
}
