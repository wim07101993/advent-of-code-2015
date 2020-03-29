package day09

import "testing"

func TestCalculateShortestDistance(t *testing.T) {
	input := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}

	m := NewMap()
	m.Parse(input)

	p := m.CalculateShortestPath()
	if p.Distance() != 605 {
		t.Errorf("The correct solution is 605, no %d, %s", p.Distance(), p)
	}
}
