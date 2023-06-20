package patterns

import "fmt"

func Pattern1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern3(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j <= i+1; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func Pattern4(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func Pattern5(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern6(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j < n-i+1; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func Pattern7(n int) {
	for i := 0; i < n; i++ {
		//space
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}

		//stars
		for j := 0; j < (2*i)+1; j++ {
			fmt.Print("*")
		}

		//space
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func Pattern8(n int) {
	for i := 0; i < n; i++ {
		//space
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		//stars

		for j := 0; j < (2*n)-(2*i+1); j++ {
			fmt.Print("*")
		}

		//space
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()

	}
}

func Pattern9(n int) {
	Pattern7(n)
	Pattern8(n)
}

func Pattern10(n int) {
	for i := 1; i <= (2*n - 1); i++ {
		stars := i

		if i > n {
			stars = (2 * n) - i
		}

		for j := 1; j <= stars; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern11(n int) {
	for i := 1; i <= n; i++ {
		start := 1

		if i%2 == 0 {
			start = 0
		}

		for j := 1; j <= i; j++ {
			fmt.Print(start)
			start = 1 - start
		}

		fmt.Println()
	}

}

func Pattern12(n int) {
	for i := 1; i <= n; i++ {
		//print
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		//space
		for j := 1; j <= (2*n)-(2*i); j++ {
			fmt.Print(" ")
		}
		//reverse_print
		for j := i; j > 0; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func Pattern13(n int) {
	k := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(k, " ")
			k++
		}
		fmt.Println()
	}
}

func Pattern14(n int) {

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", 'A'+j)
		}
		fmt.Println()
	}
}

func Pattern15(n int) {

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf("%c", 'A'+j)
		}
		fmt.Println()
	}
}

func Pattern16(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", 'A'+i)
		}
		fmt.Println()
	}
}

func Pattern17(n int) {
	for i := 0; i < n; i++ {
		letter := 'A'
		//space
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}

		//letters
		for j := 0; j < (2*i)+1; j++ {
			fmt.Printf("%c", letter)
			if j >= ((2*i + 1) / 2) {
				letter -= 1
			} else {
				letter += 1
			}
		}

		//space
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func Pattern18(n int) {
	letter := 'A' + n - 1

	for i := 0; i < n; i++ {
		char := letter
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", char)
			char += 1
		}
		letter -= 1
		fmt.Println()
	}
}

func pattern19_1(n int) {
	for i := 0; i < n; i++ {
		//star
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		//space
		for j := 0; j < 2*i; j++ {
			fmt.Print(" ")
		}
		//star
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func pattern19_2(n int) {
	for i := 0; i < n; i++ {
		//star
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		//space
		for j := 0; j < 2*(n-i-1); j++ {
			fmt.Print(" ")
		}
		//star
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern19(n int) {
	pattern19_1(n)
	pattern19_2(n)
}

func Pattern20(n int) {
	for i := 1; i <= (2*n - 1); i++ {
		row := i
		space := 2 * (n - i)

		if i > n {
			row = 2*n - i
			space = 2 * (i - n)
		}
		//star
		for j := 1; j <= row; j++ {
			fmt.Print("*")
		}

		//space
		for j := 1; j <= space; j++ {
			fmt.Print(" ")
		}

		//star
		for j := 1; j <= row; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern21(n int) {
	for i := 1; i <= n; i++ {

		for j := 1; j <= n; j++ {

			if i == 1 || i == n || j == 1 || j == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Pattern22(n int) {
	for i := 1; i <= 2*n-1; i++ {
		for j := 1; j <= 2*n-1; j++ {
			top := i
			left := j
			right := 2*n - j
			down := 2*n - i
			num := n - min(min(top, down), min(left, right))
			fmt.Print(num)
		}
		fmt.Println()
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
