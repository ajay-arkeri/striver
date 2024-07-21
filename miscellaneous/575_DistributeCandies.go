package miscellaneous

import "sort"

func DistributeCandies1(candyType []int) int {
    mp:=make(map[int]int)

    for i:=0;i<len(candyType);i++{
        mp[candyType[i]]++
    }
    
    differentCandies:= len(mp)
    
    return min(len(candyType)/2,differentCandies)
}

func DistributeCandies2(candyType []int) int {
	n := len(candyType)

	sort.Ints(candyType)
	count := 1
	for i := 1; i < n; i++ {
		if candyType[i] != candyType[i-1] {
			count++
		}
	}

	return min(count, n/2)
}

