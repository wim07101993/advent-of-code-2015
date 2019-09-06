package day08

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type SolverDay8 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay8 {
	return &SolverDay8{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay8) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	total := 0
	totalCompiled := 0
	for _, str := range s.input {
		total += len(str)
		c, err := CompileString(str)
		if err != nil {
			fmt.Println(err)
		} else {
			totalCompiled += len(c)
		}
	}
	return strconv.Itoa(total - totalCompiled)
}

func (s *SolverDay8) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	total := 0
	totalDecompiled := 0
	for _, str := range s.input {
		total += len(str)
		c := DecompileString(str)
		totalDecompiled += len(c)
	}
	return strconv.Itoa(totalDecompiled - total)
}

func CompileString(s string) (string, error) {
	if !strings.Contains(s, "\\") {
		s = strings.Replace(s, "\"", "", -1)
		return s, nil
	}

	s = s[1 : len(s)-1]
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "\\\\", "\\")

	r := regexp.MustCompile(`\\x[0-9A-f]{2}`)
	ms := r.FindAllString(s, -1)
	for _, m := range ms {
		hs, err := hex.DecodeString(m[2:])
		if err != nil {
			return "", err
		}
		s = strings.Replace(s, m, string(hs), -1)
	}

	return s, nil
}

func DecompileString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")

	return s
}
