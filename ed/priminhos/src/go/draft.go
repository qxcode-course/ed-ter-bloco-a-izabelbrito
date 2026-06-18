package main
import( 
       "fmt"
       "strconv"
       "strings"
)
func eh_primo(x int, div int) bool{
    if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%div == 0 {
		return false
	}
	if div*div > x {
		return true
	}
	return eh_primo(x, div+1)

}
func guarda_primos(n int, atual int, primos []int) []int{
    if len(primos) == n{
        return primos
    }

    if eh_primo(atual, 2){
        primos = append(primos, atual)
    }

    return guarda_primos(n, atual + 1, primos)
}

func main() {
    var n int
    fmt.Scan(&n)
    
    resultado := guarda_primos(n, 2, []int{})

    var sb strings.Builder
    sb.WriteString("[")

    for i, p := range resultado{
        if i > 0{
            sb.WriteString(", ")
        }

        sb.WriteString(strconv.Itoa(p))
    }

    sb.WriteString("]")
    fmt.Println(sb.String())
    
}