package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fila := make([]int, n)
	for i := range fila {
		fmt.Scan(&fila[i])
	}

	var m int
	fmt.Scan(&m)

	saiu := make(map[int]bool)
	for i := 0; i < m; i++ {
		var id int
		fmt.Scan(&id)
		saiu[id] = true
	}

	for _, id := range fila {
		if !saiu[id] {
			fmt.Printf("%d ", id)
		}
	}
	fmt.Println()
}