package main
import "fmt"

func contar_bloquinhos(n int) int{

    if n == 1{
        return 20
    }

    return contar_bloquinhos(n-1) + 8
}
func main() {
    var n int

    if _, err := fmt.Scan(&n); err != nil{
        return
    }

    fmt.Println(contar_bloquinhos(n))
}