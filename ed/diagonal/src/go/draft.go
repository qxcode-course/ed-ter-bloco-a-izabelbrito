package main
import "fmt"

func palavra_diagonal(s string, k int){
    if len(s) == 0{
        return
    }

    for i := 0; i < k; i++ {
        fmt.Print(" ")
        
    }

    fmt.Printf("%c\n", s[0])

    palavra_diagonal(s[1:], k + 1)
}
func main() {
    var palavra string

    fmt.Scan(&palavra)
    
    palavra_diagonal(palavra, 0)
}