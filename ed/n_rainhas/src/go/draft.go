package main

import (
	"fmt"
)

func salvar_rainha(n int) int{
    count := 0

    colunas := make([]bool, n)
    diag1 := make([]bool, 2 * n)
    diag2 := make([]bool, 2 * n)

    var backtrack func(linha int)
    backtrack = func (linha int)  {
        if linha == n{
            count++
            return 
        }

        for col := 0; col < n; col++{
            d1 := linha - col + n
            d2 := linha + col

            if colunas[col] || diag1[d1] || diag2[d2]{
                continue
            }
             
            colunas[col] = true
            diag1[d1] = true
            diag2[d2] = true
            backtrack(linha + 1)

            colunas[col] = false
            diag1[d1] = false
            diag2[d2] = false


        }
    }
    backtrack(0)
    return count
}
func main() {
    var n int

    if _, err := fmt.Scan(&n); err == nil{
        fmt.Println(salvar_rainha(n))
    }
}