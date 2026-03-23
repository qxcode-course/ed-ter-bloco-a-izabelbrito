package main

import "fmt"

func main(){

    var n int

    fmt.Scan(&n)
    
    solteiros := make(map[int]int)
    casais := 0

    for i := 0; i < n; i++{
        var animal int

        fmt.Scan(&animal)

        par := -animal

        if solteiros[par] > 0{
            casais++
            solteiros[par]--
        }else{
            solteiros[animal]++
        }
    }

    fmt.Println(casais)
}
