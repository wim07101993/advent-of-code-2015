package day09

import "sync"

type SolverDay9 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay9 {
	return &SolverDay9{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay9) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	m := NewMap()
	m.Parse(s.input)

	d := m.CalculateShortestPath().Distance()
	return d.String()
}

func (s *SolverDay9) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func WaitForRunners(runners []chan []Path) []Path {
	ret := make([]Path, 0)
	merge := make(chan []Path)

	var wg sync.WaitGroup
	wg.Add(len(runners))

	for _, e := range runners {
		go func(c <-chan []Path, merge chan<- []Path) {
			merge <- <-c
			wg.Done()
		}(e, merge)
	}

	go func() {
		wg.Wait()
		close(merge)
	}()

	for p := range merge {
		if p != nil {
			ret = append(ret, p...)
		}
	}

	return ret
}

// func (m *Map) CalculateShortestDistance() Path {
// 	chs := make([]chan Path, 0)

// 	for _, node := range m.Nodes {
// 		ch := make(chan Path)
// 		chs = append(chs, ch)
// 		go m.CalculateShortestDistanceForNode([]*Node{node}, node, ch)
// 	}

// 	return WaitForCalculations(chs)
// }

// func (m *Map) CalculateShortestDistanceForNode(path Path, current *Node, c chan<- Path) {

// 	if len(path) != 1 && path[0] == current {
// 		log.Printf("%s == %s\r\n%s\r\n", path[0].Name, current.Name, path)
// 		c <- path
// 		close(c)
// 		return
// 	}

// 	chs := make([]chan Path, 0)
// 	for node := range current.Connections {
// 		beenHere := false
// 		for _, old := range path {
// 			if node.Name == old.Name {
// 				beenHere = true
// 			}
// 		}
// 		if beenHere {
// 			continue
// 		}

// 		ch := make(chan Path)
// 		chs = append(chs, ch)

// 		newPath := append(path, node)
// 		log.Printf("Calculate: Path=%s\r\n", path)
// 		go m.CalculateShortestDistanceForNode(newPath, node, ch)
// 	}

// 	log.Printf("Wait for routines to finish %s\r\n", current)
// 	c <- WaitForCalculations(chs)
// }

// func WaitForCalculations(chs []chan Path) Path {
// 	merge := make(chan Path)
// 	var wg sync.WaitGroup
// 	wg.Add(len(chs))

// 	for _, ch := range chs {
// 		go func(c <-chan Path, merge chan<- Path) {
// 			p := <-c
// 			log.Printf("Received Path: %s", p)
// 			merge <- p
// 			wg.Done()
// 		}(ch, merge)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(merge)
// 	}()

// 	var minDistancePath Path
// 	var minDistance Distance = 0
// 	maxNodes := 0
// 	for path := range merge {
// 		log.Printf("Path: %s; Len: %d", path, len(path))
// 		if len(path) < maxNodes {
// 			continue
// 		}

// 		log.Printf("Distance: %d", len(path))
// 		distance := path.Distance()
// 		if minDistance == 0 || distance < minDistance {
// 			minDistance = distance
// 			minDistancePath = path
// 		}
// 	}

// 	return minDistancePath
// }
