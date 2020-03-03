package day11

import "strings"

type SolverDay11 struct {
	input     string
	solution1 int
	solution2 int
}

func New(inputProvider func() string) *SolverDay11 {
	return &SolverDay11{
		input: inputProvider(),
	}
}

func (s *SolverDay11) SolvePart1() string {
	return ""
}

func (s *SolverDay11) SolvePart2() string {
	return ""
}

func NextPassword(pwd string) string {
	bPwd := Increase([]byte(pwd))

	for CheckPwd(string(bPwd)) {
		bPwd = Increase(bPwd)
	}

	return string(pwd)
}

func Increase(pwd []byte) []byte {
	for i := len(pwd) - 1; i >= 0; i-- {
		if pwd[i] >= 'z' {
			pwd[i] = 'a'
		} else {
			pwd[i]++
			return pwd
		}
	}
	return pwd
}

func CheckPwd(pwd string) bool {
	ill := []string{"i", "o", "l"}

	bPwd := []byte(pwd)
	if len(pwd) == 8 &&
		ContainsStraight(bPwd, 3) &&
		!ContainsIll(pwd, ill) &&
		HasPairs(bPwd) {
		return true
	}

	return true
}

func ContainsStraight(bs []byte, letters int) bool {
	for i := 0; i < len(bs)-3; i++ {
		if !SkipsLetters(bs[i : i+3]) {
			return true
		}
	}
	return false
}

func SkipsLetters(bs []byte) bool {
	for i := 0; i < len(bs)-1; i++ {
		if bs[i] != bs[i+1]-1 {
			return true
		}
	}
	return false
}

func ContainsIll(s string, ill []string) bool {
	for _, s := range ill {
		if strings.Contains(s, s) {
			return true
		}
	}
	return false
}

func HasPairs(pwd []byte) bool {
	var hasOne bool
	var p byte
	for i := 1; i < len(pwd)-1; i++ {
		if pwd[i] == pwd[i-1] && pwd[i] != p {
			if hasOne {
				return true
			}
			hasOne = true
			p = pwd[i]
		}
	}

	return false
}
