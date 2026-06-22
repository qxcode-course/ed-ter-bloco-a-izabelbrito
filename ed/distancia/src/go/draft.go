package main
import "fmt"

func pode_inserir(vetor []rune, index int, limite int, digito rune) bool{

    inicio := index - limite
    
    if inicio < 0{
        inicio = 0
    }

    for i := inicio; i < index; i++ {
        if vetor[i] == digito {
            return false
        }
    }

    fim:= index + limite

    if fim >= len(vetor){
        fim = len(vetor) - 1
    }

    for i := index + 1; i <= fim; i++ {
        if vetor[i] == digito{
            return false
        }   
    }

    return true
}

func preencher(vetor []rune, index int, limite int) bool{
    if index == len(vetor){
        return true
    }

    if vetor[index] != '.'{
        return preencher(vetor, index + 1, limite)
    }

    for d := 0; d <= limite; d++{
        digito := rune('0' + d)

        if pode_inserir(vetor, index, limite, digito){
            vetor[index] = digito

            if preencher(vetor, index + 1, limite){
                return true
            }

            vetor[index] = '.'
        }
    }

    return false
}

func main() {
   var sequencia string
   var limite int

   if _, err := fmt.Scan(&sequencia, &limite); err != nil{
        return
   }

   vetor := []rune(sequencia)

   if preencher(vetor, 0, limite){
    fmt.Println(string(vetor))
   }
}