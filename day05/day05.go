package day05

import (
	"strconv"
	"strings"
)

type SolverDay5 struct {
	inputProvider func() []string
	input         []string
	solution1     int
	solution2     int
}

func New(inputProvider func() []string) *SolverDay5 {
	return &SolverDay5{
		inputProvider: inputProvider,
	}
}

func (s *SolverDay5) SolvePart1() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	good := removeNaughty(s.input, isGood1)
	return strconv.Itoa(len(good))
}

func (s *SolverDay5) SolvePart2() string {
	if s.input == nil {
		s.input = s.inputProvider()
	}

	good := removeNaughty(s.input, isGood2)
	return strconv.Itoa(len(good))
}

func removeNaughty(ss []string, filter func(string) bool) []string {
	ret := []string{}
	for _, s := range ss {
		if filter(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func isGood1(s string) bool {
	not := []string{"ab", "cd", "pq", "xy"}

	return containsNVowels(s, 3) && containsXThatAppearsNTimesInARow(s, 2) && doesntContain(s, not)
}

func isGood2(s string) bool {
	return containsPairOfLastTwoLetters(s) && containsPairWithLetterInBetween(s)
}

func containsNVowels(s string, n int) bool {
	x := 0
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			x++
			if x >= n {
				return true
			}
		}
	}
	return false
}

func containsXThatAppearsNTimesInARow(s string, n int) bool {
	streak := 1
	var check rune
	for _, c := range s {
		if check == c {
			streak++
			if streak >= n {
				return true
			}
		} else {
			check = c
			streak = 1
		}
	}
	return false
}

func doesntContain(s string, ss []string) bool {
	for _, x := range ss {
		if strings.Contains(s, x) {
			return false
		}
	}
	return true
}

func containsPairOfLastTwoLetters(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		pair := s[i : i+2]
		c := s[i+2:]
		if strings.Contains(c, pair){
			return true
		}
	}
	return false
}

func containsPairWithLetterInBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
