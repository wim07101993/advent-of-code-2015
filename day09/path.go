package day09

import "strings"

type Path []*Node

func (p Path) String() string {
	if len(p) == 0 {
		return "no path"
	}
	var b strings.Builder

	b.WriteString(p[0].Name)
	for i := 1; i < len(p); i++ {
		b.WriteString(" -> ")
		b.WriteString(p[i].Name)
	}

	d := p.Distance()
	b.WriteString(": ")
	b.WriteString(d.String())
	return b.String()
}

func (p Path) Distance() Distance {
	var distance Distance = 0

	for i := 1; i < len(p); i++ {
		n1 := p[i-1]
		n2 := p[i]

		for n, d := range n2.Connections {
			if n.Name == n1.Name {
				distance += d
			}
		}
	}

	return distance
}

func (p Path) Contains(value *Node) bool {
	for _, n := range p {
		if n == value {
			return true
		}
	}
	return false
}
