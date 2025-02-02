package daily

// s := "abc"
// shifts := []int{3, 5, 9}
//ans = "rpl"

func ShiftingLetters(s string, shifts []int) string {
	byteArray := []byte(s)

	sum := 0

	for i := len(shifts) - 1; i >= 0; i-- {
		sum += shifts[i]
		shifts[i] = sum
	}

	for i := range byteArray {
		byteArray[i] += byte(shifts[i] % 26)

		if byteArray[i] > 'z' {
			byteArray[i] -= 'a'
		}
	}

	return string(byteArray)
}
