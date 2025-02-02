package daily

// s := "abc"
// shifts := [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}

func ShiftingLetters2_Brute(s string, shifts [][]int) string {
	shiftArray := make([]int, len(s))

	for _, shift := range shifts {
		start := shift[0]
		end := shift[1]

		for i := start; i <= end; i++ {
			if shift[2] == 0 {
				shiftArray[i]--
			} else {
				shiftArray[i]++
			}
		}
	}

	byteArray := []byte(s)

	for i := range byteArray {
		byteArray[i] += byte(shiftArray[i] % 26)

		if byteArray[i] > 'z' {
			byteArray[i] -= 26
		}

		if byteArray[i] < 'a' {
			byteArray[i] += 26
		}
	}

	return string(byteArray)
}

func ShiftingLetters2_Optimal(s string, shifts [][]int) string {
	diff := make([]int, len(s)+1)

	for _, shift := range shifts {
		start := shift[0]
		end := shift[1]

		if shift[2] == 0 {
			diff[start] -= 1
			diff[end+1] += 1
		} else {
			diff[start] += 1
			diff[end+1] -= 1
		}
	}

	byteArray := []byte(s)

	shiftSum := 0

	for i := range byteArray {
		shiftSum += diff[i]

		byteArray[i] += byte(shiftSum % 26)

		if byteArray[i] > 'z' {
			byteArray[i] -= 26
		}

		if byteArray[i] < 'a' {
			byteArray[i] += 26
		}
	}

	return string(byteArray)
}
