package basicmaths

import (
	"fmt"
	"math"
	"sort"
)

func Print_Digits(n int) {
	var k int
	for n > 0 {
		digit := n % 10
		fmt.Println(digit)
		n = n / 10
		k = k*10 + digit
	}
	fmt.Println(k)
}

// TC --> O(log10(N))  since division is by 10, if by 2 then O(log2(N))
func Count_Digits(n int) int {
	var i int
	for n > 0 {
		n = n / 10
		i++
	}
	return i
}

func Count_Digits_Logarithmic(n int) int {
	return int(math.Log10(float64(n)) + 1)
}

func ReverseNumber(n int) int {
	var reverse int
	for n > 0 {
		reverse = reverse*10 + (n % 10)
		n = n / 10
	}
	return reverse
}

func Check_Palindrome(n int) bool {
	num := n
	var reverse int

	for num > 0 {
		reverse = reverse*10 + (num % 10)
		num = num / 10
	}

	if reverse == n {
		return true
	} else {
		return false
	}
}

func Check_ArmstrongNumber(n int) bool {
	sum := 0
	num := n
	for num > 0 {
		sum += int(math.Pow(float64(num%10), float64(3)))
		num = num / 10
	}

	fmt.Println(sum)
	if sum == n {
		return true
	} else {
		return false
	}
}

func PrintAllDivisors(n int) {
	// for i := 1; i <= n; i++ {
	// 	if n%i == 0 {
	// 		fmt.Print(i, " ")
	// 	}
	// }  O(N) solution.
	list := []int{}
	for i := 1; i*i <= n; i++ { //int(math.Sqrt(float64(n)))
		if n%i == 0 {
			list = append(list, i)
			if (n / i) != i {
				list = append(list, n/i)
			}
		}
	}
	sort.Ints(list)
	fmt.Println(list)

	//TC --> O(sqrt(N)) + O(nlogn)
}

func CheckPrimeNumber(n int) bool {
	// counter := 0
	// for i := 1; i <= n; i++ {
	// 	if n%i == 0 {
	// 		counter++
	// 	}
	// }

	// if counter == 2 {
	// 	return true
	// } else {
	// 	return false
	// }  TC --> O(N)

	counter := 0

	for i := 1; i*i < n; i++ {
		if n%i == 0 {
			counter++
			if (n / i) != i {
				counter++
			}
		}
	}

	if counter == 2 {
		return true
	} else {
		return false
	}

	//TC --> O(sqrt(N))
}

func Find_GCD_HCF(a int, b int) int {
	gcd := 0

	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	for i := 1; i <= min(a, b); i++ {
		if a%i == 0 && b%i == 0 {
			gcd = max(i, gcd)
		}
	}

	return gcd
}

func GCD_Euclidean(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	if a == 0 {
		return b
	} else {
		return a
	}
}
