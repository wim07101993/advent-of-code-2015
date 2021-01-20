package day11

type Password []byte

func NewPassword(seed string) Password {
	p := make(Password, 8)
	for i := 0; i < len(p); i++ {
		p[i] = seed[i]
	}
	return p
}

func(p Password) Next() {
	p.Increment()
	for !p.IsValid() {
		p.Increment()
	}
}

func (p Password) Increment() {
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] == 'z' {
			p[i] = 'a'
		} else {
			p[i] = p[i] + 1
			break
		}
	}
}

func (p Password) IsValid() bool {
	return p.IncludesOneIncreasingStraightOfAtLeastThreeLetters() &&
		p.DoesNotContainiol() &&
		p.ContainsAtLeastTwoDifferentNonOverlappingPairsOfLetters()
}

func (p Password) IncludesOneIncreasingStraightOfAtLeastThreeLetters() bool {
	for i := 2; i < len(p); i++ {
		x := p[i]
		if p[i-1] == x-1 && p[i-2] == x-2 {
			return true
		}
	}
	return false
}

func (p Password) DoesNotContainiol() bool {
	for _, c := range p {
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}
	}
	return true
}

func (p Password) ContainsAtLeastTwoDifferentNonOverlappingPairsOfLetters() bool {
	foundFirstPair := false
	for i := 1; i < len(p); i++ {
		if p[i] == p[i-1] {
			if foundFirstPair {
				return true
			}
			foundFirstPair = true
			i++
		}
	}
	return false
}

func (p Password) String() string {
	return string(p)
}
