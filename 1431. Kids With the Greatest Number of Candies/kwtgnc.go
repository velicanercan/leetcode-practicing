package kwtgnc

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandie := 0
	for _, candy := range candies {
		if candy > maxCandie {
			maxCandie = candy
		}
	}

	res := make([]bool, len(candies))
	for i, candie := range candies {
		if candie+extraCandies >= maxCandie {
			res[i] = true
		}
	}
	return res
}
