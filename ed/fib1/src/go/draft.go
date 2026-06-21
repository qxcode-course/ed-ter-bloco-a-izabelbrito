package main
import "fmt"

func coelhos(n, k int) int{
    if n == 1 || n == 2{
        return 1
    }

    return coelhos(n - 1, k) + (k * coelhos(n - 2, k))
}
func main() {
    var n, k int

    if _, err := fmt.Scan(&n, &k); err != nil{
        return
    }
    fmt.Println(coelhos(n, k))
}