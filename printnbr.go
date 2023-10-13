package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	} else if n == 0 {
		z01.PrintRune('0')
		return
	}

	numCount := 0
	nbr := n
	for nbr > 0 {
		nbr /= 10
		numCount++
	}

	digits := make([]int, numCount)
	i := 0
	for n > 0 {
		digit := n % 10
		digits[i] = digit
		n /= 10
		i++
	}

	for j := i - 1; j >= 0; j-- {
		z01.PrintRune(rune(digits[j]) + '0')
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
