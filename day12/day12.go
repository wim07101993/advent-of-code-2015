package day12

import (
	"strconv"
	"strings"
)

type SolverDay12 struct {
	inputProvider func() string
	input         string
	solution1     int
	solution2     int
}

func New(inputProvider func() string) *SolverDay12 {
	return &SolverDay12{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay12) SolvePart1() string {
	if s.input == "" {
		s.input = s.inputProvider()
	}

	return strconv.Itoa(CalculateSum(s.input))
}

func (s *SolverDay12) SolvePart2() string {
	if s.input == "" {
		s.input = s.inputProvider()
	}

	return ""
}

func CalculateSum(s string) int {
	sum := 0
	isInQuotes := false
	builder := strings.Builder{}

	for _, c := range s {
		if c == '"' {
			isInQuotes = !isInQuotes
		} else if !isInQuotes {
			if c == '-' || (c >= 48 && c <= 57) {
				builder.WriteRune(c)
			} else {
				i, _ := strconv.Atoi(builder.String())
				sum += i
				builder.Reset()
			}
		}
	}

	return sum
}
