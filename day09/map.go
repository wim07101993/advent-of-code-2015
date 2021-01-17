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
	minDistance := 0
	var minPath History
	for i := range m {
		d, h := walkShortest(History{}, 0, m[i])
		if minDistance > d || minDistance == 0 {
			minDistance = d
			minPath = h
		}
	}
	return minDistance, minPath
}

func (m Map) LongestDistance() (int, History) {
	maxDistance := 0
	var maxPath History
	for i := range m {
		d, h := walkLongest(History{}, 0, m[i])
		if d > maxDistance {
			maxDistance = d
			maxPath = h
		}
	}
	return maxDistance, maxPath
}

func walkShortest(history History, currentDistance int, point *Point) (int, History) {
	history = append(history, point)
	minDistance := 0
	minPath := history
	for np, d := range point.Connections {
		if history.Contains(np) {
			continue
		}

		distance, h := walkShortest(history, currentDistance, np)
		distance += d
		if minDistance > distance || minDistance == 0 {
			minDistance = distance
			minPath = h
		}
	}

	return minDistance, minPath
}

func walkLongest(history History, currentDistance int, point *Point) (int,
	History) {
	history = append(history, point)
	maxDistance := 0
	maxPath := history
	for np, d := range point.Connections {
		if history.Contains(np) {
			continue
		}

		distance, h := walkLongest(history, currentDistance, np)
		distance += d
		if distance > maxDistance {
			maxDistance = distance
			maxPath = h
		}
	}

	return maxDistance, maxPath
}
