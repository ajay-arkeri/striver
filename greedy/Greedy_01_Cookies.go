package greedy

import "sort"

//TC - nlogn + mlogm + m

func AssignCookies(g []int, s []int) int {
	l, r := 0, 0

	sort.Ints(g)
	sort.Ints(s)

	for l < len(s) && r < len(g) {
		if s[l] >= g[r] {
			l++
			r++
		} else {
			l++
		}
	}

	return r + 1
}