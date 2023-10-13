package main

import "fmt"

func isNegative(nb int) {
	if nb < 0 {
		fmt.Print("T")
		fmt.Println("")
	} else {
		fmt.Print("F")
		fmt.Println("")
	}
}

func main() {

	isNegative(-5)
	isNegative(0)
	isNegative(10)
}
