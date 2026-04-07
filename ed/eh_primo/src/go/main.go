package main

import "fmt"

func eh_primo(x int, div int) bool {
	if x < 2{
		return false
	}
	if x == 2{
		return true
	}
	if x % 2 == 0{
		return false
	}
	if div * div > x{
		return true
	}

	return eh_primo(x, div+1)
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
