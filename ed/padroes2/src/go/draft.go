package main
import "fmt"

func num_pecas(n int) int {

    if n == 1{
       
        return 3
    }

    return num_pecas((n - 1)) + (2 * n + 1)
}
func main() {
    var n int

    if _, err := fmt.Scan(&n); err != nil {
        return
        
    }
    fmt.Println(num_pecas(n))
}