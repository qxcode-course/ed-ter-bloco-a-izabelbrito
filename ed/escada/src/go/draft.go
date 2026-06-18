package main
import "fmt"

func maneiras_de_subir(n int) int{
    if n == 1{
        return 1
    }

    if n == 2 {
        return 1
    }

    if n == 3 {
        return 2
    }

    return maneiras_de_subir(n -1) + maneiras_de_subir(n - 3)

}

func main() {
    var n int

    fmt.Scan(&n)

    fmt.Println(maneiras_de_subir(n))
}