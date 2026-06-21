package main
import "fmt"

func contar_pontos(n, m int) int{
    if m == 1{
        return 1
    }

    acrescimo := (n - 2) * (m - 1) + 1
    return contar_pontos(n, m - 1) + acrescimo
}
func main() {
    var n, m int

    if _, err:= fmt.Scan(&n, &m); err != nil{
        return
    }

    fmt.Println(contar_pontos(n, m))
}