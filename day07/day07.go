package day07

import (
	"strconv"
	"strings"
)

type SolverDay7 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
	wires         map[string]uint16
}

func New(inputProvider func() []string) *SolverDay7 {
	return &SolverDay7{
		inputProvider: inputProvider,
		wires:         make(map[string]uint16),
	}
}

func (s *SolverDay7) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func (s *SolverDay7) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	return ""
}

func InterpreteInstruction(i string, wires map[string]uint16) {
	ss := strings.Split(i, " ")
	var val uint16

	if ss[0] == "NOT" {
		val = SelectVar(ss[1], wires)
	} else if ss[1] == "->" {
		val = SelectVar(ss[0], wires)
	} else {
		first := SelectVar(ss[0], wires)
		second := SelectVar(ss[2], wires)

		switch ss[1] {
		case "AND":
			val = first & second
		case "OR":
			val = first | second
		case "LSHIFT":
			val = first << second
		case "RSHIFT":
			val = first >> second
		default:
			return
		}
	}

	name := ss[len(ss)-1]
	wires[name] = val
}

func SelectVar(v string, wires map[string]uint16) uint16 {
	i, err := strconv.Atoi(v)
	if err != nil {
		return wires[v]
	}
	return uint16(i)
}
