package day09

import "strings"

type Node struct {
	Name        string
	Connections map[*Node]Distance
}

func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString(n.Name)
	builder.WriteString(": [")

	for k, v := range n.Connections {
		builder.WriteString(" ")
		builder.WriteString(k.Name)
		builder.WriteString(" => ")
		builder.WriteString(v.String())
		builder.WriteString(",")
	}

	builder.WriteString("]")

	return builder.String()
}
