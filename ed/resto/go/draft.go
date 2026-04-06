package main

import "fmt"

func resolver(n int) {
	if n == 0 {
		return
	}

	q := n / 2
	r := n % 2

	resolver(q)

	fmt.Printf("%d %d\n", q, r)
}

func main() {
	var n int
	fmt.Scan(&n)

	if n > 0 {
		resolver(n)
	}
}