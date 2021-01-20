package day11

import "testing"

func TestPassword_Next(t *testing.T) {
	passwordNext(t, "abcdefgh", "abcdffaa")
	passwordNext(t, "ghijklmn", "ghjaabcc")
}

func TestPassword_Increment(t *testing.T) {
	passwordIncrement(t, "aaaaaaaa", "aaaaaaab")
	passwordIncrement(t, "aaaaaaaz", "aaaaaaba")
}

func TestPassword_IsValid(t *testing.T) {
	if p := NewPassword("hijklmmn"); p.IsValid() {
		t.Errorf("%s is not valid (contains i and l)", p)
	}
	if p := NewPassword("abbceffg"); p.IsValid() {
		t.Errorf("%s is not valid (does not contain a straight", p)
	}
	if p := NewPassword("abbcegjk"); p.IsValid() {
		t.Errorf("%s is not valid (does not contain two pairs)", p)
	}
	if p := NewPassword("abbccdek"); !p.IsValid() {
		t.Errorf("%s is valid", p)
	}
}

func TestPassword_IncludesOneIncreasingStraightOfAtLeastThreeLetters(t *testing.T) {
	if p := NewPassword("aaaaaaaa");
		p.IncludesOneIncreasingStraightOfAtLeastThreeLetters() {
		t.Errorf("%s does not include a straight", p)
	}
	if p := NewPassword("aabcaaaa");
		!p.IncludesOneIncreasingStraightOfAtLeastThreeLetters() {
		t.Errorf("%s does include a straight", p)
	}
	if p := NewPassword("aaacbaaa");
		p.IncludesOneIncreasingStraightOfAtLeastThreeLetters() {
		t.Errorf("%s does not include a straight", p)
	}
}

func TestPassword_DoesNotContainiol(t *testing.T) {
	if p := NewPassword("abcdefgh");
		!p.DoesNotContainiol() {
		t.Errorf("%s does not contain i,o or l", p)
	}
	if p := NewPassword("jklisljk");
		p.DoesNotContainiol() {
		t.Errorf("%s does contain i,o or l", p)
	}
}

func TestPassword_DoesContainAtLeastTwoDifferentNonOverlappingPairsOfLetters(t *testing.T) {
	if p := NewPassword("abcdefgh");
		p.ContainsAtLeastTwoDifferentNonOverlappingPairsOfLetters() {
		t.Errorf("%s does not contain pairs", p)
	}
	if p := NewPassword("kkkejbie");
		p.ContainsAtLeastTwoDifferentNonOverlappingPairsOfLetters() {
		t.Errorf("%s does not contain two pairs", p)
	}
	if p := NewPassword("aaejllow");
		!p.ContainsAtLeastTwoDifferentNonOverlappingPairsOfLetters() {
		t.Errorf("%s does contain two pairs", p)
	}
}

func TestPassword_String(t *testing.T) {
	passwordStringTester(t, "abcdefgh")
	passwordStringTester(t, "skikoxny")
}

func passwordStringTester(t *testing.T, password string) {
	p := NewPassword(password)
	s := p.String()
	if s != password {
		t.Errorf("expected %s, not %s", password, s)
	}
}

func passwordNext(t *testing.T, password, expected string) {
	p := NewPassword(password)
	p.Next()
	if p.String() != expected {
		t.Errorf("expected password to increment to %s, not %s",
			expected, p.String())
	}
}

func passwordIncrement(t *testing.T, password, expected string) {
	p := NewPassword(password)
	p.Increment()
	if p.String() != expected {
		t.Errorf("expected password to increment to %s, not %s",
			expected, p.String())
	}
}
