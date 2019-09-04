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

func TestIsGood(t *testing.T) {
	if !isGood("ugknbfddgicrmopn") {
		t.Error("ugknbfddgicrmopn is good")
	}

	if !isGood("aaa") {
		t.Error("aaa is good")
	}

	if isGood("jchzalrnumimnmhp") {
		t.Error("jchzalrnumimnmhp is naughty")
	}

	if isGood("haegwjzuvuyypxyu") {
		t.Error("haegwjzuvuyypxyu is naughty")
	}

	if isGood("dvszwmarrgswjxmb") {
		t.Error("dvszwmarrgswjxmb is naughty")
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

	ss = removeNaughty(ss)

	for _, s := range ss {
		switch s {
		case "jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb":
			t.Errorf("%s shoul have been removed", s)
		}
	}
}
