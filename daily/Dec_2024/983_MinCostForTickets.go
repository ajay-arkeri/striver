package daily

func MincostTickets(days []int, costs []int) int {
	mp := make(map[int]bool)
	for _, val := range days {
		mp[val] = true
	}
	memo := make(map[int]int)
	var recursiveHelper func(int) int

	recursiveHelper = func(currentDay int) int {
		if currentDay > days[len(days)-1] {
			return 0
		}

		if val, ok := memo[currentDay]; ok {
			return val
		}

		if _, ok := mp[currentDay]; !ok {
			memo[currentDay] = recursiveHelper(currentDay + 1)
			return memo[currentDay]
		}

		oneDay := recursiveHelper(currentDay+1) + costs[0]
		sevenDay := recursiveHelper(currentDay+7) + costs[1]
		thirtyDay := recursiveHelper(currentDay+30) + costs[2]

		memo[currentDay] = min(oneDay, min(sevenDay, thirtyDay))
		return memo[currentDay]

	}

	return recursiveHelper(days[0])
}
