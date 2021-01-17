package day09

type Point struct {
	Connections map[*Point]int
	Name        string
}

type Map []*Point

func NewPoint(name string) *Point {
	return &Point{
		Name:        name,
		Connections: map[*Point]int{},
	}
}

func ParseMap(ss []string) Map {
	m := map[string]*Point{}

	for _, s := range ss {
		ptp := ParsePointToPoint(s)

		p1, ok := m[ptp.p1]
		if !ok {
			p1 = NewPoint(ptp.p1)
			m[ptp.p1] = p1
		}
		p2, ok := m[ptp.p2]
		if !ok {
			p2 = NewPoint(ptp.p2)
			m[ptp.p2] = p2
		}

		p1.Connections[p2] = ptp.distance
		p2.Connections[p1] = ptp.distance
	}

	mp := make(Map, 0, len(m))
	for _, p := range m {
		mp = append(mp, p)
	}
	return mp
}

func (m Map) ShortestDistance() (int, History) {
	distance := 0
	var minPath History
	for i := range m {
		d, h := walk(History{}, 0, m[i])
		if distance > d || distance == 0 {
			distance = d
			minPath = h
		}
	}
	return distance, minPath
}

func walk(history History, currentDistance int, point *Point) (int, History) {
	minDistance := 0
	history = append(history, point)
	minPath := history
	for np, d := range point.Connections {
		if history.Contains(np) {
			continue
		}

		total, h := walk(history, currentDistance, np)
		total += d
		if minDistance > total || minDistance == 0 {
			minDistance = total
			minPath = h
		}
	}

	return minDistance, minPath
}
