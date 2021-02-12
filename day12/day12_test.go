package day12

import "testing"

func TestCalculateSum(t *testing.T) {
	testCalculateSum(t, "[1,2,3]", 6)
	testCalculateSum(t, "{\"a\":2,\"b\":4}", 6)
	testCalculateSum(t, "[[[3]]]", 3)
	testCalculateSum(t, "{\"a\":{\"b\":4},\"c\":-1}", 3)
	testCalculateSum(t, "{\"5\":{\"b\":4},\"c\":-1}", 3)
	testCalculateSum(t, "[]", 0)
	testCalculateSum(t, "{}", 0)
	testCalculateSum(t, "[10,2,3]", 15)
}

func testCalculateSum(t *testing.T, input string, expected int) {
	s := CalculateSum(input)
	if s != expected {
		t.Errorf("%s should result in %d, not %d", input, expected, s)
	}
}
