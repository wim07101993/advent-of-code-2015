package day09

import (
	"fmt"
	"strings"
)

type History []*Point

func (h History) Contains(point *Point) bool {
	for _, p := range h {
		if p.Name == point.Name {
			return true
		}
	}
	return false
}

func (h History) Last() *Point {
	return h[len(h)-1]
}

func (h History) String() string {
	if len(h) == 0 {
		return "no history"
	}
	builder := strings.Builder{}
	builder.WriteString(h[0].Name)
	if len(h) == 1 {
		return builder.String()
	}

	for _, p := range h[1:] {
		builder.WriteString(fmt.Sprintf(" -> %s", p.Name))
	}
	return builder.String()
}
