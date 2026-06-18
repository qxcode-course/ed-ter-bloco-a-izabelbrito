package main
import (
    "fmt"
    "strings"
)

func main() {
   var T, R int

   if _, err := fmt.Scan(&T, &R); err != nil{
            return
   }

   vetor := make([]string, T)

   for i := 0; i < T; i++ {
    fmt.Scan(&vetor[i]) 
   }

   R = R % T

   corte := T - R

   var rotacionar []string
   rotacionar = append(rotacionar, vetor[corte:]...)
   rotacionar = append(rotacionar, vetor[:corte]...)

   fmt.Printf("[ %s ]\n", strings.Join(rotacionar, " "))
}