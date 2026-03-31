package main

import "fmt"

func imprimir(fila []int, pos int){
	fmt.Print("[ ")
	for i := 0; i < len(fila); i++{
		if i == pos {
			fmt.Printf("%d> ", fila[i])
			continue
		}
		fmt.Printf("%d ", fila[i])
	}
	fmt.Println("]")
}

func main(){
	var n, e int
	fmt.Scan(&n, &e)

	fila := make([]int, 0, n)
	for i := 1; i <= n; i++{
		fila = append(fila, i)
	}

	pos := e - 1

	for len(fila) > 1{
		imprimir(fila, pos)
		morto := (pos + 1) % len(fila)

		for i := morto; i < len(fila)-1; i++{
			fila[i] = fila[i+1]
		}
		fila = fila[:len(fila)-1]

		pos = morto % len(fila)
	}

	imprimir(fila, pos)
}