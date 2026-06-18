package main
import "fmt"

func eh_primo( x int, div int) bool{
    if x < 2 {
        return false
    }

    if x == 2 {
        return true
    }

    if x % div == 0 {
        return false
    }

    if div * div > x {
        return true
    }

    return eh_primo(x, div + 1)
}

func enesimo_primo(n int, atual int, cont int) int{
    if eh_primo(atual, 2){
        cont++

        if cont == n{
            return atual
        }
        
    }

    return enesimo_primo(n, atual + 1, cont)

}
func main() {

    var n int
    fmt.Scan(&n)

    resultado := enesimo_primo(n, 2, 0)

    fmt.Println(resultado)

    
}
