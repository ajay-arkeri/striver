package recursion

import (
	"fmt"
)

// Print name N times recursively
// TC --> O(N), SC --> O(N)(internal stack space)
func PrintName(counter int, n int) {
	if counter > n {
		return
	}
	fmt.Println("Ajay")
	PrintName(counter+1, n)
}

// Print linearly from 1 to N
func PrintNumbers(n, counter int) {
	if counter > n {
		return
	}
	fmt.Println(counter)
	PrintNumbers(n, counter+1)
}

// Print N to 1
func PrintNumbers_Reverse(n int) {
	if n < 1 {
		return
	}

	fmt.Println(n)
	PrintNumbers_Reverse(n - 1)
}

// Print 1 to N using backtracking
func PrintNumbers_BackTracking(n int) {
	if n <= 0 {
		return
	}
	PrintNumbers_BackTracking(n - 1)
	fmt.Println(n)
}

// Print N to 1 using backtracking
func PrintNumbers_BackTracking_Reverse(i, n int) {
	if i > n {
		return
	}
	PrintNumbers_BackTracking_Reverse(i+1, n)
	fmt.Println(i)

}

// Summation of first N numbers
func Sum_FirstN_Numbers_Parameterized(sum, n int) {
	if n < 1 {
		fmt.Println(sum)
		return
	}
	Sum_FirstN_Numbers_Parameterized(sum+n, n-1)

}

// Summation of first N numbers
func Sum_FirstN_Numbers_functional(n int) int {
	if n <= 0 {
		return 0
	}

	return n + Sum_FirstN_Numbers_functional(n-1)
}

// return factorial of n
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func ReverseArray(arr []int) []int {
	n := len(arr)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func ReverseArray_Using_Recursion_Parameterized(arr []int, index int) {
	//base condition
	if index >= len(arr)/2 {
		fmt.Println(arr)
		return
	}

	arr[index], arr[len(arr)-1-index] = arr[len(arr)-1-index], arr[index]
	ReverseArray_Using_Recursion_Parameterized(arr, index+1)
}

//check if string is palindrome or not using recursion