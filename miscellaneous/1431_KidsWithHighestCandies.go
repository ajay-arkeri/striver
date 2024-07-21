package miscellaneous

func KidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	ans := make([]bool, n)

	highest := candies[0]

	for i := 1; i < n; i++ {
		highest = max(candies[i], highest)
	}

	for i := 0; i < n; i++ {
		ans[i] = candies[i]+extraCandies >= highest
	}

	return ans

}