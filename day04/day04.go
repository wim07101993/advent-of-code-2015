package day04

import (
	"crypto/md5"
	"fmt"
)

type SolverDay4 struct {
	input     []byte
	solution1 int
	solution2 int
}

func New(inputProvider func() []byte) *SolverDay4 {
	return &SolverDay4{
		input: inputProvider(),
	}
}

func (s *SolverDay4) SolvePart1() string {
	return ""
}

func (s *SolverDay4) SolvePart2() string {
	return ""
}

func findHash() int {
	return 0
}

func hash(input []byte) string {
	h := md5.New()
	h.Write(input)

	return fmt.Sprintf("%x", h.Sum(nil))
}
