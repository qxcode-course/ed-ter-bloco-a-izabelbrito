package main
import "fmt"

func main() {

  var n int

  fmt.Scan(&n)

  descasados := make(map[int]int)
  casais := 0

  for i := 0; i < n; i++{

    var animal int

    fmt.Scan(&animal)

    par := -animal

    if descasados[par] > 0{
        casais++
        descasados[par]--
    }else{
        descasados[animal]++
    }
  }

fmt.Println(casais)

}
