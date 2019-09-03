package utils

func Min(is ...int) int {
	if len(is) == 0 {
		return 0
	}

	min := is[0]
	for _, i := range is {
		if i < min {
			min = i
		}
	}

	return min
}

func Max(is ...int) int {
	if len(is) == 0 {
		return 0
	}

	max := is[0]
	for _, i := range is {
		if i > max {
			max = i
		}
	}

	return max
}

func Sum(is ...int) int {
	if len(is) == 0 {
		return 0
	}

	total := 0
	for _, i := range is {
		total += i
	}
	return total
}
