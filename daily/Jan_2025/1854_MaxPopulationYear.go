package daily

import "sort"

// logs := [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}

func MaximumPopulation(logs [][]int) int {
	mp := make(map[int]int)
	years := []int{}

	for _, log := range logs {
		startYear := log[0]
		endYear := log[1]

		if _, ok := mp[startYear]; !ok {
			years = append(years, startYear)
		}

		if _, ok := mp[endYear]; !ok {
			years = append(years, endYear)
		}

		mp[startYear]++
		mp[endYear]--
	}

	sort.Ints(years)

	maxPopulation := 0
	currentPopulation := 0
	maxPopulationYear := years[0]

	for _, year := range years {
		currentPopulation += mp[year]

		if currentPopulation > maxPopulation {
			maxPopulation = currentPopulation
			maxPopulationYear = year
		}

	}

	return maxPopulationYear
}
