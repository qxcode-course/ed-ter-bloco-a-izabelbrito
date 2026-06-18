package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
    var tamanho int

    if _, err := fmt.Scan((&tamanho)); err != nil{
        return
    }

    frequencias := make(map[string]int)

    for i := 0; i < tamanho; i++ {
        var palavra string
        fmt.Scan(&palavra)
        
        frequencias[palavra]++
    }
   var tamanho_consultas int
  
   fmt.Scan(&tamanho_consultas)

   var resultados []string

   for i := 0; i < tamanho_consultas; i++ {
    var consulta string
    fmt.Scan(&consulta)

    ocorrencias := frequencias[consulta]

    resultados = append(resultados, strconv.Itoa(ocorrencias))
    
   }

   fmt.Println(strings.Join(resultados, " "))
}