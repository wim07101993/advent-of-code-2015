package day09

import (
	"syreclabs.com/go/faker"
	"testing"
)

func TestNewPoint(t *testing.T) {
	// arrange
	name := faker.Name().Name()

	// act
	p := NewPoint(name)

	// assert
	if name != p.Name {
		t.Errorf("Expected name to be %s, not %s", name, p.Name)
	}
	if p.Connections == nil {
		t.Error("Expected connections not to be nil")
	}
}

func TestParseMap(t *testing.T) {
	// arrange
	input := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	london := NewPoint("London")
	dublin := NewPoint("Dublin")
	belfast := NewPoint("Belfast")
	london.Connections = map[*Point]int{dublin: 464, belfast: 518}
	dublin.Connections = map[*Point]int{london: 464, belfast: 141}
	belfast.Connections = map[*Point]int{london: 518, dublin: 141}

	// act
	m := ParseMap(input)

	// assert
	if len(m) != 3 {
		t.Errorf("Expcted the length to be 3, not %d", len(m))
	}
	pointEquals(t, m, 0, "London", "Dublin", "Belfast", 464, 518)
	pointEquals(t, m, 1, "Dublin", "London", "Belfast", 464, 141)
	pointEquals(t, m, 2, "Belfast", "London", "Dublin", 518, 141)
}

func pointEquals(t *testing.T, m Map, i int, name, p1, p2 string, d1, d2 int) {
	if m[i].Name != name {
		t.Errorf("Expected name to be %s, not %s", name, m[i].Name)
	}
	for k, v := range m[i].Connections {
		if k.Name == p1 {
			if v != d1 {
				t.Errorf("Expected distance to be %d, not %d", d1, v)
			}
		} else if k.Name == p2 {
			if v != d2 {
				t.Errorf("Expected distance to be %d, not %d", d2, v)
			}
		} else {
			t.Errorf("Unexpected name %s", k.Name)
		}
	}
}

func TestMap_ShortestDistance(t *testing.T) {
	// arrange
	input := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	london := NewPoint("London")
	dublin := NewPoint("Dublin")
	belfast := NewPoint("Belfast")
	london.Connections = map[*Point]int{dublin: 464, belfast: 518}
	dublin.Connections = map[*Point]int{london: 464, belfast: 141}
	belfast.Connections = map[*Point]int{london: 518, dublin: 141}

	// act
	d, h := ParseMap(input).ShortestDistance()

	//
	if d != 605 {
		t.Errorf("Expected distance 605, not %d", d)
	}
	if h.String() != "London -> Dublin -> Belfast" {
		t.Errorf("Expected history London -> Dublin -> Belfast, not %s", h)
	}
}