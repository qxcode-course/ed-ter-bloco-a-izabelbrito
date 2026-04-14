package main
import "fmt"
func main() {

    var total_album int
    var total_baruel int
    var figura int

    fmt.Scan(&total_album,&total_baruel)

	var contagem [55]int
	for i := 0; i < total_baruel; i++ {
		fmt.Scan(&figura)
		contagem[figura]++
	}
    fmt.Println("Hello, World!")
}
