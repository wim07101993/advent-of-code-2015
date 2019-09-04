package day05

import "testing"

func TestDoesntContain(t *testing.T) {
	ss := []string{"ab", "cd", "pq", "xy"}

	if doesntContain("haegwjzuvuyypxyu", ss) {
		t.Error("haegwjzuvuyypxyu does contain `ab`, `cd`, `pq`, or `xy`")
	}

	if !doesntContain("ugknbfddgicrmopn", ss) {
		t.Error("ugknbfddgicrmopn doesn't contain `ab`, `cd`, `pq`, or `xy`")
	}
}

func TestContainsXThatAppearsNTimesInARow(t *testing.T) {
	if containsXThatAppearsNTimesInARow("jchzalrnumimnmhp", 2) {
		t.Error("ugknbfddgicrmopn doesn't contain any letter two times in a row")
	}

	if !containsXThatAppearsNTimesInARow("ugknbfddgicrmopn", 2) {
		t.Error("ugknbfddgicrmopn contains two times the letter d in a row")
	}
}

func TestContainsNVowels(t *testing.T) {
	if containsNVowels("dvszwmarrgswjxmb", 3) {
		t.Error("dvszwmarrgswjxmb contains only one vowel")
	}

	if !containsNVowels("ugknbfddgicrmopn", 3) {
		t.Error("ugknbfddgicrmopn contains at least 3 vowels")
	}
}

func TestContainsPairOfLastTwoLetters(t *testing.T) {
	if !containsPairOfLastTwoLetters("xyxy") {
		t.Error("xyxy does contain it's last two letters somewhere else")
	}

	if !containsPairOfLastTwoLetters("qjhvhtzxzqqjkmpb") {
		t.Error("qjhvhtzxzqqjkmpb does contain it's last two letters somewhere else")
	}

	if !containsPairOfLastTwoLetters("aabcdefgaa") {
		t.Error("aabcdefgaa does contain it's last two letters somewhere else")
	}

	if containsPairOfLastTwoLetters("aa") {
		t.Error("aa overlaps")
	}

	if containsPairOfLastTwoLetters("aaa") {
		t.Error("aaa overlaps")
	}
}

func TestContainsPairWithLetterInBetween(t *testing.T) {
	if !containsPairWithLetterInBetween("xyx") {
		t.Error("xyx does repeat the letter x")
	}

	if !containsPairWithLetterInBetween("abcdefeghi") {
		t.Error("abcdefeghi repeats the letter e")
	}

	if !containsPairWithLetterInBetween("aaa") {
		t.Error("aaa repeats the letter a")
	}

	if containsPairWithLetterInBetween("uurcxstgmygtbstg") {
		t.Error("uurcxstgmygtbstg does not repeat any letter")
	}
}

func TestIsGood1(t *testing.T) {
	if !isGood1("ugknbfddgicrmopn") {
		t.Error("ugknbfddgicrmopn is good")
	}

	if !isGood1("aaa") {
		t.Error("aaa is good")
	}

	if isGood1("jchzalrnumimnmhp") {
		t.Error("jchzalrnumimnmhp is naughty")
	}

	if isGood1("haegwjzuvuyypxyu") {
		t.Error("haegwjzuvuyypxyu is naughty")
	}

	if isGood1("dvszwmarrgswjxmb") {
		t.Error("dvszwmarrgswjxmb is naughty")
	}
}

func TestIsGood2(t *testing.T) {
	if !isGood2("qjhvhtzxzqqjkmpb") {
		t.Error("qjhvhtzxzqqjkmpb is good")
	}

	if !isGood2("xxyxx") {
		t.Error("xxyxx is good")
	}

	if isGood2("uurcxstgmygtbstg") {
		t.Error("uurcxstgmygtbstg is naughty")
	}

	if isGood2("ieodomkazucvgmuy") {
		t.Error("ieodomkazucvgmuy is naughty")
	}
}

func TestRemoveNaughty(t *testing.T) {
	ss := []string{
		"ugknbfddgicrmopn",
		"aaa",
		"jchzalrnumimnmhp",
		"haegwjzuvuyypxyu",
		"dvszwmarrgswjxmb",
	}

	ss = removeNaughty(ss, isGood1)

	for _, s := range ss {
		switch s {
		case "jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb":
			t.Errorf("%s shoul have been removed", s)
		}
	}
}
