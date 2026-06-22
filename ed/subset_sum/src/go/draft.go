package main
import "fmt"

func tem_soma(vetor []int, alvo int, i int, soma int) bool{

    if soma == alvo{
        return true
    }

    if soma > alvo{
        return false
    }

    if i == len(vetor){
        return false
    }

    if tem_soma(vetor, alvo, i + 1, soma + vetor[i]){
        return true
    }

    if tem_soma(vetor, alvo, i + 1, soma){
        return true
    }
    
    return false
}
func main() {
    var n, alvo int

    if _, err := fmt.Scan(&n, &alvo); err != nil{
        return
    }

    vetor := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i])   
    }

    if tem_soma(vetor, alvo, 0, 0){
        fmt.Println("true")
    }else{
        fmt.Println("false")
    }
}
