package day09

import (
	"log"
	"regexp"
	"strconv"
)

type Map struct {
	Regex *regexp.Regexp
	Nodes map[string]*Node
}

func NewMap() *Map {
	return &Map{
		Regex: regexp.MustCompile(`(?P<source>\w*) to (?P<destination>\w*) = (?P<distance>\d*)`),
		Nodes: make(map[string]*Node),
	}
}

func (m *Map) Parse(input []string) {
	for _, s := range input {
		m.ParseNode(s)
	}
}

func (m *Map) ParseNode(s string) {
	ms := m.Regex.FindStringSubmatch(s)
	if len(ms) != 4 {
		log.Printf("Could not parse %s: Regex failed\r\n", s)
	}

	d, err := strconv.Atoi(ms[3])
	if err != nil {
		log.Printf("Could not parse %s: %s\r\n", s, err)
	}

	name1 := ms[1]
	name2 := ms[2]
	distance := Distance(d)

	node1, ok := m.Nodes[name1]
	if !ok {
		node1 = &Node{
			Name:        name1,
			Connections: make(map[*Node]Distance),
		}
		m.Nodes[name1] = node1
	}

	node2, ok := m.Nodes[name2]
	if !ok {
		node2 = &Node{
			Name:        name2,
			Connections: make(map[*Node]Distance),
		}
		m.Nodes[name2] = node2
	}

	node1.Connections[node2] = distance
	node2.Connections[node1] = distance
}

func (m *Map) CalculateShortestPath() Path {
	paths := m.CalculatePaths()

	var minDistancePath Path
	var minDistance Distance = 0
	maxNodes := 0
	for _, p := range paths {
		if len(p) < maxNodes {
			continue
		}

		distance := p.Distance()
		log.Printf("Path: %s; Len: %d; Distance: %d", p, len(p), distance)
		if len(p) > maxNodes {
			maxNodes = len(p)
			minDistance = distance
			minDistancePath = p
			continue
		}

		if minDistance == 0 || distance < minDistance {
			minDistance = distance
			minDistancePath = p
		}
	}

	return minDistancePath
}

func (m *Map) CalculatePaths() []Path {
	ps := make([]Path, 0)

	for _, n := range m.Nodes {
		p := make(Path, 0)
		log.Printf("Calculating from start: %s", n)
		found := m.CalculatePathFromStartingPoint(n, p)
		if found != nil {
			ps = append(ps, found...)
		}
	}

	return ps
}

func (m *Map) CalculatePathFromStartingPoint(at *Node, p Path) []Path {
	if len(p) > 0 && p[0] == at {
		log.Printf("Reached end: %s", p)
		return []Path{p}
	}

	if p.Contains(at) {
		log.Printf("Been here: %s (%s)", at.Name, p)
		return nil
	}

	p = append(p, at)
	ps := make([]Path, 0)

	for n := range at.Connections {
		log.Printf("Moving to next: %s, %s", p, n.Name)
		found := m.CalculatePathFromStartingPoint(n, p)
		if found != nil {
			ps = append(ps, found...)
		}
	}

	return ps
}
