package main
import "fmt"

func pode_colocar(matriz [][]rune, l int, c int, n int, digito rune) bool{

    for i := 0; i < n; i++ {
        if matriz[l][i] == digito{
            return false
        }

        if matriz[i][c] == digito {
            return false
        }   
    }

    sub := 2
    if n == 9 {
        sub = 3
    }

    inicioL := (l / sub) *sub
    inicioC := (c / sub) *sub

    for i := 0; i < sub; i++ {
        for j := 0; j < sub; j++ {
          
        if matriz[inicioL + i][inicioC + j] == digito{
            return false
           }
        }  
    }

    return true
}

func resolver(matriz[][]rune, index int, n int) bool{

    if index == n * n{
        return true
    }

    l := index / n
    c := index % n

    if matriz[l][c] != '.'{
        return resolver(matriz, index + 1, n)
    }

    for d := 1; d <= n; d++ {
        digito := rune('0' + d)

        if pode_colocar(matriz, l, c, n, digito){
            matriz[l][c] = digito

            if resolver(matriz, index + 1, n) {
                return true
            }
            matriz[l][c] = '.'
        }
        
    }
    return false
}

func main() {
   var n int

   if _, err := fmt.Scan(&n); err != nil{
        return
   }

   matriz := make([][]rune, n)
   for i := 0; i < n; i++ {
    var linha string
    if _, err := fmt.Scan(&linha); err != nil{
        return
   }
   matriz[i] = []rune(linha)
   }
   if resolver(matriz, 0, n){
    for i := 0; i < n; i++ {
        fmt.Println(string(matriz[i]))
        
    }
   }
}