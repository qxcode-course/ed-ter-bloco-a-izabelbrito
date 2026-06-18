package main
import "fmt"

func sufixos(s string){
    if len(s) == 0{
        return
    }

    sufixos(s[1:])

    fmt.Println(s)
}
func main() {
    var palavra string

    fmt.Scan(&palavra)

    sufixos(palavra)
}